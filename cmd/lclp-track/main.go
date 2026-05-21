package main

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

var (
	allowedPhases = []string{
		"bootstrap",
		"proposal_capture",
		"plan_split",
		"root_plan",
		"product_plan",
		"technical_plan",
		"coding_start_gate",
		"milestone_detail",
		"checkpoint_plan",
		"execution",
		"review_fix_forward",
		"proposal_rollup",
	}
	slugPattern = regexp.MustCompile(`^[A-Z0-9._-]+$`)
)

const (
	baseDirEnvVar   = "LCLP_BASE_DIR"
	baseDirFlagHelp = "install root containing workflow/; defaults to LCLP_BASE_DIR, then the current git worktree root when that root already has workflow/"
)

type event struct {
	ID        string   `json:"id"`
	TS        string   `json:"ts"`
	Slug      string   `json:"slug"`
	Phase     string   `json:"phase"`
	Action    string   `json:"action"`
	Status    string   `json:"status"`
	Prompt    string   `json:"prompt,omitempty"`
	Artifacts []string `json:"artifacts,omitempty"`
	Note      string   `json:"note,omitempty"`
	Replaces  *string  `json:"replaces,omitempty"`
}

type stringSlice []string

func (s *stringSlice) String() string {
	return strings.Join(*s, ",")
}

func (s *stringSlice) Set(value string) error {
	if strings.TrimSpace(value) == "" {
		return errors.New("empty value is not allowed")
	}
	*s = append(*s, value)
	return nil
}

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "lclp-track:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		printUsage()
		return errors.New("missing command")
	}

	switch args[0] {
	case "add":
		return runAdd(args[1:])
	case "replace":
		return runReplace(args[1:])
	case "last":
		return runLast(args[1:])
	case "get":
		return runGet(args[1:])
	case "help", "-h", "--help":
		printUsage()
		return nil
	default:
		printUsage()
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func runAdd(args []string) error {
	fs := flag.NewFlagSet("add", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	baseDir := fs.String("base-dir", "", baseDirFlagHelp)
	slug := fs.String("slug", "", "LCLP slug")
	phase := fs.String("phase", "", "LCLP phase")
	action := fs.String("action", "", "event action")
	status := fs.String("status", "", "event status")
	prompt := fs.String("prompt", "", "prompt or doc path")
	note := fs.String("note", "", "free-form note")
	var artifacts stringSlice
	fs.Var(&artifacts, "artifact", "artifact path to attach; repeat to add more than one")

	if err := fs.Parse(args); err != nil {
		return err
	}
	if err := validateSlug(*slug); err != nil {
		return err
	}
	if err := validatePhase(*phase); err != nil {
		return err
	}
	if strings.TrimSpace(*action) == "" {
		return errors.New("--action is required")
	}
	if strings.TrimSpace(*status) == "" {
		return errors.New("--status is required")
	}
	baseDirExplicit := flagWasProvided(fs, "base-dir")

	return withAuditWriteLockForInput(*baseDir, baseDirExplicit, *slug, func() error {
		_, path, err := loadEventsForInput(*baseDir, baseDirExplicit, *slug)
		if err != nil {
			return err
		}

		nextID, err := newEventID()
		if err != nil {
			return err
		}
		record := event{
			ID:        nextID,
			TS:        time.Now().UTC().Format(time.RFC3339),
			Slug:      *slug,
			Phase:     *phase,
			Action:    strings.TrimSpace(*action),
			Status:    strings.TrimSpace(*status),
			Prompt:    strings.TrimSpace(*prompt),
			Artifacts: compactStrings(artifacts),
			Note:      strings.TrimSpace(*note),
		}

		if err := appendEvent(path, record); err != nil {
			return err
		}
		return writeJSON(os.Stdout, record)
	})
}

func runReplace(args []string) error {
	fs := flag.NewFlagSet("replace", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	baseDir := fs.String("base-dir", "", baseDirFlagHelp)
	slug := fs.String("slug", "", "LCLP slug")
	targetID := fs.String("id", "", "existing event id to supersede")
	phase := fs.String("phase", "", "LCLP phase; defaults to the replaced entry")
	action := fs.String("action", "", "event action; defaults to the replaced entry")
	status := fs.String("status", "", "event status; defaults to the replaced entry")
	prompt := fs.String("prompt", "", "prompt or doc path; defaults to the replaced entry")
	clearPrompt := fs.Bool("clear-prompt", false, "clear the prompt instead of inheriting it")
	note := fs.String("note", "", "free-form note")
	clearArtifacts := fs.Bool("clear-artifacts", false, "clear attached artifacts instead of inheriting them")
	var artifacts stringSlice
	fs.Var(&artifacts, "artifact", "artifact path to attach; repeat to add more than one")

	if err := fs.Parse(args); err != nil {
		return err
	}
	if err := validateSlug(*slug); err != nil {
		return err
	}
	if err := validateEventID(*targetID); err != nil {
		return err
	}

	promptProvided := false
	artifactProvided := false
	fs.Visit(func(f *flag.Flag) {
		switch f.Name {
		case "prompt":
			promptProvided = true
		case "artifact":
			artifactProvided = true
		}
	})
	if *clearPrompt && promptProvided && strings.TrimSpace(*prompt) != "" {
		return errors.New("--clear-prompt cannot be used with a non-empty --prompt")
	}
	if *clearArtifacts && artifactProvided {
		return errors.New("--clear-artifacts cannot be used with --artifact")
	}
	baseDirExplicit := flagWasProvided(fs, "base-dir")

	return withAuditWriteLockForInput(*baseDir, baseDirExplicit, *slug, func() error {
		events, path, err := loadEventsForInput(*baseDir, baseDirExplicit, *slug)
		if err != nil {
			return err
		}
		if len(events) == 0 {
			return fmt.Errorf("no audit log exists yet for slug %q", *slug)
		}

		target, err := resolveEffectiveEvent(events, *targetID)
		if err != nil {
			return err
		}

		effectivePhase := strings.TrimSpace(*phase)
		if effectivePhase == "" {
			effectivePhase = target.Phase
		}
		if err := validatePhase(effectivePhase); err != nil {
			return err
		}

		effectiveAction := strings.TrimSpace(*action)
		if effectiveAction == "" {
			effectiveAction = target.Action
		}
		if effectiveAction == "" {
			return errors.New("--action is required when the replaced entry has no action")
		}

		effectiveStatus := strings.TrimSpace(*status)
		if effectiveStatus == "" {
			effectiveStatus = target.Status
		}
		if effectiveStatus == "" {
			return errors.New("--status is required when the replaced entry has no status")
		}

		effectivePrompt := target.Prompt
		if *clearPrompt {
			effectivePrompt = ""
		} else if promptProvided && strings.TrimSpace(*prompt) != "" {
			effectivePrompt = strings.TrimSpace(*prompt)
		}

		effectiveArtifacts := target.Artifacts
		if *clearArtifacts {
			effectiveArtifacts = nil
		} else if artifactProvided {
			effectiveArtifacts = compactStrings(artifacts)
		}

		nextID, err := newEventID()
		if err != nil {
			return err
		}
		record := event{
			ID:        nextID,
			TS:        time.Now().UTC().Format(time.RFC3339),
			Slug:      *slug,
			Phase:     effectivePhase,
			Action:    effectiveAction,
			Status:    effectiveStatus,
			Prompt:    effectivePrompt,
			Artifacts: effectiveArtifacts,
			Note:      strings.TrimSpace(*note),
			Replaces:  stringPtr(target.ID),
		}

		if record.Note == "" {
			record.Note = fmt.Sprintf("supersedes entry %s", target.ID)
		}

		if err := appendEvent(path, record); err != nil {
			return err
		}
		return writeJSON(os.Stdout, record)
	})
}

func runLast(args []string) error {
	fs := flag.NewFlagSet("last", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	baseDir := fs.String("base-dir", "", baseDirFlagHelp)
	slug := fs.String("slug", "", "LCLP slug")

	if err := fs.Parse(args); err != nil {
		return err
	}
	if err := validateSlug(*slug); err != nil {
		return err
	}
	baseDirExplicit := flagWasProvided(fs, "base-dir")

	return withAuditReadLockForInput(*baseDir, baseDirExplicit, *slug, func() error {
		events, _, err := loadEventsForInput(*baseDir, baseDirExplicit, *slug)
		if err != nil {
			return err
		}
		if len(events) == 0 {
			return fmt.Errorf("no entries found for slug %q", *slug)
		}

		return writeJSON(os.Stdout, events[len(events)-1])
	})
}

func runGet(args []string) error {
	fs := flag.NewFlagSet("get", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)

	baseDir := fs.String("base-dir", "", baseDirFlagHelp)
	slug := fs.String("slug", "", "LCLP slug")
	id := fs.String("id", "", "specific event id to return")
	phase := fs.String("phase", "", "filter by phase")
	action := fs.String("action", "", "filter by action")
	status := fs.String("status", "", "filter by status")
	all := fs.Bool("all", false, "return raw append-only history instead of only effective live events")

	if err := fs.Parse(args); err != nil {
		return err
	}
	if err := validateSlug(*slug); err != nil {
		return err
	}
	if strings.TrimSpace(*id) != "" {
		if err := validateEventID(*id); err != nil {
			return err
		}
	}
	if *phase != "" {
		if err := validatePhase(*phase); err != nil {
			return err
		}
	}
	baseDirExplicit := flagWasProvided(fs, "base-dir")

	return withAuditReadLockForInput(*baseDir, baseDirExplicit, *slug, func() error {
		events, _, err := loadEventsForInput(*baseDir, baseDirExplicit, *slug)
		if err != nil {
			return err
		}

		filterSource := events
		if !*all {
			filterSource = effectiveEvents(events)
		}

		filtered := filterEvents(filterSource, strings.TrimSpace(*id), *phase, *action, *status)
		return writeJSON(os.Stdout, filtered)
	})
}

func loadEvents(baseDir, slug string) ([]event, string, error) {
	return loadEventsForInput(baseDir, true, slug)
}

func loadEventsForInput(baseDir string, explicit bool, slug string) ([]event, string, error) {
	path, err := auditPathForInput(baseDir, explicit, slug)
	if err != nil {
		return nil, "", err
	}
	file, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, path, nil
		}
		return nil, "", err
	}
	defer file.Close()

	var events []event
	scanner := bufio.NewScanner(file)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	lineNo := 0
	for scanner.Scan() {
		lineNo++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var evt event
		if err := json.Unmarshal([]byte(line), &evt); err != nil {
			return nil, "", fmt.Errorf("invalid JSON on line %d of %s: %w", lineNo, path, err)
		}
		if err := validateStoredSlug(path, lineNo, evt.Slug, slug); err != nil {
			return nil, "", err
		}
		events = append(events, evt)
	}
	if err := scanner.Err(); err != nil {
		return nil, "", err
	}

	return events, path, nil
}

func appendEvent(path string, evt event) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o777); err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o666)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	if _, err := file.Write(append(data, '\n')); err != nil {
		return err
	}
	return nil
}

func auditLockPath(baseDir, slug string) (string, error) {
	return auditLockPathForInput(baseDir, true, slug)
}

func auditLockPathForInput(baseDir string, explicit bool, slug string) (string, error) {
	canonicalBaseDir, err := resolveBaseDir(baseDir, explicit)
	if err != nil {
		return "", err
	}

	lockPath := filepath.Join(canonicalBaseDir, "workflow", "logs", "audit", ".locks", slug+".lock")
	if err := validateAuditLeafPath(lockPath); err != nil {
		return "", err
	}
	return lockPath, nil
}

func auditPath(baseDir, slug string) (string, error) {
	return auditPathForInput(baseDir, true, slug)
}

func auditPathForInput(baseDir string, explicit bool, slug string) (string, error) {
	canonicalBaseDir, err := resolveBaseDir(baseDir, explicit)
	if err != nil {
		return "", err
	}
	auditPath := filepath.Join(canonicalBaseDir, "workflow", "logs", "audit", slug+".audit.log.jsonl")
	if err := validateAuditLeafPath(auditPath); err != nil {
		return "", err
	}
	return auditPath, nil
}

func resolveBaseDir(baseDir string, explicit bool) (string, error) {
	if explicit {
		return resolveConfiguredBaseDir(baseDir, "explicit --base-dir")
	}

	if envBaseDir, ok := baseDirFromEnv(); ok {
		return resolveConfiguredBaseDir(envBaseDir, baseDirEnvVar)
	}

	return resolveDefaultBaseDir()
}

func resolveConfiguredBaseDir(baseDir, source string) (string, error) {
	absBaseDir, err := filepath.Abs(baseDir)
	if err != nil {
		return "", err
	}
	canonicalBaseDir, err := filepath.EvalSymlinks(absBaseDir)
	if err != nil {
		return "", err
	}

	err = validateWorkflowDirRoot(canonicalBaseDir)
	if err == nil {
		if err := validateAuditStorageRoot(canonicalBaseDir); err != nil {
			return "", err
		}
		return canonicalBaseDir, nil
	}
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", fmt.Errorf("%s must already contain workflow/: %s", source, canonicalBaseDir)
		}
		if errors.Is(err, errWorkflowDirSymlink) {
			return "", fmt.Errorf("%s must contain a real workflow/ directory, not a symlink: %s", source, canonicalBaseDir)
		}
		if errors.Is(err, errWorkflowDirNotDir) {
			return "", fmt.Errorf("%s must already contain workflow/: %s", source, canonicalBaseDir)
		}
		return "", err
	}

	return "", fmt.Errorf("%s must already contain workflow/: %s", source, canonicalBaseDir)
}

func baseDirFromEnv() (string, bool) {
	value, ok := os.LookupEnv(baseDirEnvVar)
	if !ok {
		return "", false
	}
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return "", false
	}
	return trimmed, true
}

func flagWasProvided(fs *flag.FlagSet, name string) bool {
	provided := false
	fs.Visit(func(f *flag.Flag) {
		if f.Name == name {
			provided = true
		}
	})
	return provided
}

func canonicalizeBaseDir(baseDir string) (string, error) {
	return resolveBaseDir(baseDir, true)
}

var (
	errWorkflowDirSymlink = errors.New("workflow directory must not be a symlink")
	errWorkflowDirNotDir  = errors.New("workflow path is not a directory")
)

func validateWorkflowDirRoot(root string) error {
	workflowDir := filepath.Join(root, "workflow")
	info, err := os.Lstat(workflowDir)
	if err != nil {
		return err
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return errWorkflowDirSymlink
	}
	if !info.IsDir() {
		return errWorkflowDirNotDir
	}
	return nil
}

func validateAuditStorageRoot(root string) error {
	for _, dir := range []string{
		filepath.Join(root, "workflow", "logs"),
		filepath.Join(root, "workflow", "logs", "audit"),
		filepath.Join(root, "workflow", "logs", "audit", ".locks"),
	} {
		if err := validateOptionalDirectoryPath(dir); err != nil {
			return err
		}
	}
	return nil
}

func validateOptionalDirectoryPath(path string) error {
	info, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return fmt.Errorf("audit storage path must not be a symlink: %s", path)
	}
	if !info.IsDir() {
		return fmt.Errorf("audit storage path must be a directory: %s", path)
	}
	return nil
}

func validateAuditLeafPath(path string) error {
	info, err := os.Lstat(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if info.Mode()&os.ModeSymlink != 0 {
		return fmt.Errorf("audit storage path must not be a symlink: %s", path)
	}
	if info.IsDir() {
		return fmt.Errorf("audit storage path must not be a directory: %s", path)
	}
	return nil
}

func resolveDefaultBaseDir() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.CombinedOutput()
	if err != nil {
		detail := strings.TrimSpace(string(output))
		if detail == "" {
			return "", fmt.Errorf("default base-dir resolution requires %s, or running inside a git worktree and passing --base-dir explicitly", baseDirEnvVar)
		}
		return "", fmt.Errorf("default base-dir resolution requires %s, or running inside a git worktree and passing --base-dir explicitly: %s", baseDirEnvVar, detail)
	}

	root := strings.TrimSpace(string(output))
	if root == "" {
		return "", errors.New("default --base-dir resolved to an empty repository root")
	}

	err = validateWorkflowDirRoot(root)
	if err == nil {
		return root, nil
	}
	if errors.Is(err, errWorkflowDirSymlink) {
		return "", fmt.Errorf("default base-dir resolution requires a real workflow/ directory at the current git worktree root, not a workflow/ symlink; otherwise set %s or pass --base-dir", baseDirEnvVar)
	}
	if err != nil && !errors.Is(err, os.ErrNotExist) && !errors.Is(err, errWorkflowDirNotDir) {
		return "", err
	}

	return "", fmt.Errorf("default base-dir resolution requires a workflow/ directory at the current git worktree root; otherwise set %s or pass --base-dir", baseDirEnvVar)
}

func newEventID() (string, error) {
	id, err := ulid.New(ulid.Timestamp(time.Now().UTC()), rand.Reader)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func findEventByID(events []event, id string) (event, error) {
	for _, evt := range events {
		if evt.ID == id {
			return evt, nil
		}
	}
	return event{}, fmt.Errorf("entry %q not found", id)
}

func buildDescendantsByID(events []event) map[string][]string {
	descendantsByID := make(map[string][]string, len(events))
	for _, evt := range events {
		if evt.Replaces == nil {
			continue
		}
		descendantsByID[*evt.Replaces] = append(descendantsByID[*evt.Replaces], evt.ID)
	}
	return descendantsByID
}

func effectiveEvents(events []event) []event {
	descendantsByID := buildDescendantsByID(events)
	filtered := make([]event, 0, len(events))

	for _, evt := range events {
		if len(descendantsByID[evt.ID]) != 0 {
			continue
		}
		filtered = append(filtered, evt)
	}

	return filtered
}

func filterEvents(events []event, id, phase, action, status string) []event {
	filtered := make([]event, 0, len(events))
	for _, evt := range events {
		if id != "" && evt.ID != id {
			continue
		}
		if phase != "" && evt.Phase != phase {
			continue
		}
		if action != "" && evt.Action != action {
			continue
		}
		if status != "" && evt.Status != status {
			continue
		}
		filtered = append(filtered, evt)
	}

	return filtered
}

func resolveEffectiveEvent(events []event, id string) (event, error) {
	if _, err := findEventByID(events, id); err != nil {
		return event{}, err
	}

	descendantsByID := buildDescendantsByID(events)

	lineage := map[string]struct{}{id: {}}
	queue := []string{id}
	var leafIDs []string

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		children := descendantsByID[current]
		if len(children) == 0 {
			leafIDs = append(leafIDs, current)
			continue
		}

		for _, childID := range children {
			if _, seen := lineage[childID]; seen {
				continue
			}
			lineage[childID] = struct{}{}
			queue = append(queue, childID)
		}
	}

	if len(leafIDs) == 1 {
		return findEventByID(events, leafIDs[0])
	}

	slices.Sort(leafIDs)
	return event{}, fmt.Errorf("replacement lineage for entry %q is ambiguous; rerun replace with one of these leaf ids: %s", id, strings.Join(leafIDs, ", "))
}

func compactStrings(values []string) []string {
	var out []string
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed != "" {
			out = append(out, trimmed)
		}
	}
	return out
}

func validateSlug(slug string) error {
	if strings.TrimSpace(slug) == "" {
		return errors.New("--slug is required")
	}
	if !slugPattern.MatchString(slug) {
		return fmt.Errorf("invalid slug %q; use uppercase letters, numbers, dot, underscore, or hyphen", slug)
	}
	return nil
}

func validateStoredSlug(path string, lineNo int, storedSlug, expectedSlug string) error {
	if storedSlug != expectedSlug {
		return fmt.Errorf("audit log line %d of %s has slug %q; expected %q", lineNo, path, storedSlug, expectedSlug)
	}
	return nil
}

func validatePhase(phase string) error {
	if strings.TrimSpace(phase) == "" {
		return errors.New("--phase is required")
	}
	if !slices.Contains(allowedPhases, phase) {
		return fmt.Errorf("invalid phase %q; allowed phases: %s", phase, strings.Join(allowedPhases, ", "))
	}
	return nil
}

func validateEventID(id string) error {
	if strings.TrimSpace(id) == "" {
		return errors.New("--id is required")
	}
	return nil
}

func writeJSON(file *os.File, value any) error {
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(value)
}

func stringPtr(value string) *string {
	return &value
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `lclp-track manages append-only LCLP audit logs.

Usage:
  lclp-track <command> [flags]

Commands:
  add       append a new event
  replace   append a corrective event that supersedes an earlier entry
  last      print the last event for a slug
  get       print live events for a slug, optionally filtered

Examples:
  LCLP_BASE_DIR=/path/to/repo lclp-track add --slug P-0001 --phase proposal_capture --action create --status partial
  lclp-track add --slug P-0001 --phase proposal_capture --action create --status partial --artifact workflow/roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md --note "initial draft"
  lclp-track replace --slug P-0001 --id 01K06W6GG1Z5KJ0QMD4X9N9D4H --action create --status ready --note "corrected mistaken block"
  lclp-track last --slug P-0001
  lclp-track get --slug P-0001 --phase execution
  lclp-track get --slug P-0001 --all
`)
}
