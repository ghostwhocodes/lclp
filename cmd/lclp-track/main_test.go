//go:build linux

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/oklog/ulid/v2"
)

func makeInstalledBaseDir(t *testing.T) string {
	t.Helper()

	baseDir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(baseDir, "workflow"), 0o755); err != nil {
		t.Fatalf("create workflow dir: %v", err)
	}
	return baseDir
}

func TestRunTrackerLifecycle(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	addOutput := captureStdout(t, func() {
		err := run([]string{
			"add",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--phase", "proposal_capture",
			"--action", "create",
			"--status", "partial",
			"--prompt", "process/roadmaps/proposals/CHANGE_PROPOSAL_CAPTURE_PROMPT.md",
			"--artifact", "workflow/roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md",
			"--note", "initial draft created",
		})
		if err != nil {
			t.Fatalf("add returned error: %v", err)
		}
	})

	added := decodeEvent(t, addOutput)
	assertValidULID(t, added.ID)
	if added.Action != "create" {
		t.Fatalf("expected add action to be create, got %q", added.Action)
	}

	replaceOutput := captureStdout(t, func() {
		err := run([]string{
			"replace",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--id", added.ID,
			"--status", "ready",
			"--note", "corrected operator mistake",
		})
		if err != nil {
			t.Fatalf("replace returned error: %v", err)
		}
	})

	replaced := decodeEvent(t, replaceOutput)
	assertValidULID(t, replaced.ID)
	if replaced.ID == added.ID {
		t.Fatalf("expected replacement event ID to differ from original, got %q", replaced.ID)
	}
	if replaced.Action != "create" {
		t.Fatalf("expected replacement to preserve action create, got %q", replaced.Action)
	}
	if replaced.Replaces == nil || *replaced.Replaces != added.ID {
		t.Fatalf("expected replacement to target ID %q, got %#v", added.ID, replaced.Replaces)
	}

	lastOutput := captureStdout(t, func() {
		err := run([]string{
			"last",
			"--base-dir", baseDir,
			"--slug", "P-0001",
		})
		if err != nil {
			t.Fatalf("last returned error: %v", err)
		}
	})

	last := decodeEvent(t, lastOutput)
	if last.ID != replaced.ID {
		t.Fatalf("expected last event ID %q, got %q", replaced.ID, last.ID)
	}

	getOutput := captureStdout(t, func() {
		err := run([]string{
			"get",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--status", "ready",
		})
		if err != nil {
			t.Fatalf("get returned error: %v", err)
		}
	})

	var filtered []event
	if err := json.Unmarshal(getOutput, &filtered); err != nil {
		t.Fatalf("unmarshal get output: %v", err)
	}
	if len(filtered) != 1 || filtered[0].ID != replaced.ID {
		t.Fatalf("expected one ready event with ID %q, got %+v", replaced.ID, filtered)
	}

	events, path, err := loadEvents(baseDir, "P-0001")
	if err != nil {
		t.Fatalf("loadEvents returned error: %v", err)
	}
	if len(events) != 2 {
		t.Fatalf("expected 2 stored events, got %d", len(events))
	}
	if filepath.Base(path) != "P-0001.audit.log.jsonl" {
		t.Fatalf("unexpected audit path %q", path)
	}
}

func TestReplaceCanOverrideActionAndGetDefaultsToEffectiveEvents(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	added := decodeEvent(t, captureStdout(t, func() {
		err := run([]string{
			"add",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--phase", "proposal_capture",
			"--action", "cretae",
			"--status", "partial",
		})
		if err != nil {
			t.Fatalf("add returned error: %v", err)
		}
	}))

	replaced := decodeEvent(t, captureStdout(t, func() {
		err := run([]string{
			"replace",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--id", added.ID,
			"--action", "create",
			"--status", "ready",
		})
		if err != nil {
			t.Fatalf("replace returned error: %v", err)
		}
	}))

	if replaced.Action != "create" {
		t.Fatalf("expected replacement to override action to create, got %q", replaced.Action)
	}

	createOutput := captureStdout(t, func() {
		err := run([]string{
			"get",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--action", "create",
		})
		if err != nil {
			t.Fatalf("get create returned error: %v", err)
		}
	})
	createEvents := decodeEventList(t, createOutput)
	if len(createEvents) != 1 || createEvents[0].ID != replaced.ID {
		t.Fatalf("expected effective get to return replacement ID %q, got %+v", replaced.ID, createEvents)
	}

	staleOutput := captureStdout(t, func() {
		err := run([]string{
			"get",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--action", "cretae",
		})
		if err != nil {
			t.Fatalf("get stale action returned error: %v", err)
		}
	})
	if staleEvents := decodeEventList(t, staleOutput); len(staleEvents) != 0 {
		t.Fatalf("expected effective get to hide superseded stale action, got %+v", staleEvents)
	}

	rawOutput := captureStdout(t, func() {
		err := run([]string{
			"get",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--all",
			"--action", "cretae",
		})
		if err != nil {
			t.Fatalf("get raw history returned error: %v", err)
		}
	})
	rawEvents := decodeEventList(t, rawOutput)
	if len(rawEvents) != 1 || rawEvents[0].ID != added.ID {
		t.Fatalf("expected raw get to return original mistyped action ID %q, got %+v", added.ID, rawEvents)
	}

	supersededIDOutput := captureStdout(t, func() {
		err := run([]string{
			"get",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--id", added.ID,
		})
		if err != nil {
			t.Fatalf("get by superseded id returned error: %v", err)
		}
	})
	if supersededIDEvents := decodeEventList(t, supersededIDOutput); len(supersededIDEvents) != 0 {
		t.Fatalf("expected effective get by superseded id to be empty, got %+v", supersededIDEvents)
	}

	rawIDOutput := captureStdout(t, func() {
		err := run([]string{
			"get",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--all",
			"--id", added.ID,
		})
		if err != nil {
			t.Fatalf("get raw by id returned error: %v", err)
		}
	})
	rawIDEvents := decodeEventList(t, rawIDOutput)
	if len(rawIDEvents) != 1 || rawIDEvents[0].ID != added.ID {
		t.Fatalf("expected raw get by id to return original event ID %q, got %+v", added.ID, rawIDEvents)
	}
}

func TestReplaceCanClearPromptAndArtifacts(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	added := decodeEvent(t, captureStdout(t, func() {
		err := run([]string{
			"add",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--phase", "proposal_capture",
			"--action", "create",
			"--status", "partial",
			"--prompt", "process/roadmaps/proposals/CHANGE_PROPOSAL_CAPTURE_PROMPT.md",
			"--artifact", "workflow/roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md",
		})
		if err != nil {
			t.Fatalf("seed add returned error: %v", err)
		}
	}))

	replaceOutput := captureStdout(t, func() {
		err := run([]string{
			"replace",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--id", added.ID,
			"--status", "ready",
			"--clear-prompt",
			"--clear-artifacts",
		})
		if err != nil {
			t.Fatalf("replace returned error: %v", err)
		}
	})

	replaced := decodeEvent(t, replaceOutput)
	if replaced.Replaces == nil || *replaced.Replaces != added.ID {
		t.Fatalf("expected replacement to target ID %q, got %#v", added.ID, replaced.Replaces)
	}
	if replaced.Prompt != "" {
		t.Fatalf("expected cleared prompt, got %q", replaced.Prompt)
	}
	if len(replaced.Artifacts) != 0 {
		t.Fatalf("expected cleared artifacts, got %+v", replaced.Artifacts)
	}
	if bytes.Contains(replaceOutput, []byte(`"prompt"`)) {
		t.Fatalf("expected cleared prompt to be omitted from output, got %s", replaceOutput)
	}
	if bytes.Contains(replaceOutput, []byte(`"artifacts"`)) {
		t.Fatalf("expected cleared artifacts to be omitted from output, got %s", replaceOutput)
	}
}

func TestReplacePreservesPriorCorrectionsWhenReplacingSameIDTwice(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	added := decodeEvent(t, captureStdout(t, func() {
		err := run([]string{
			"add",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--phase", "proposal_capture",
			"--action", "create",
			"--status", "partial",
			"--prompt", "process/roadmaps/proposals/CHANGE_PROPOSAL_CAPTURE_PROMPT.md",
			"--artifact", "workflow/roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md",
		})
		if err != nil {
			t.Fatalf("seed add returned error: %v", err)
		}
	}))

	firstReplace := decodeEvent(t, captureStdout(t, func() {
		err := run([]string{
			"replace",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--id", added.ID,
			"--clear-prompt",
			"--clear-artifacts",
		})
		if err != nil {
			t.Fatalf("first replace returned error: %v", err)
		}
	}))

	secondReplace := decodeEvent(t, captureStdout(t, func() {
		err := run([]string{
			"replace",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--id", added.ID,
			"--status", "ready",
		})
		if err != nil {
			t.Fatalf("second replace returned error: %v", err)
		}
	}))

	if secondReplace.Replaces == nil || *secondReplace.Replaces != firstReplace.ID {
		t.Fatalf("expected second replacement to supersede entry %q, got %#v", firstReplace.ID, secondReplace.Replaces)
	}
	if secondReplace.Prompt != "" {
		t.Fatalf("expected second replacement to keep cleared prompt, got %q", secondReplace.Prompt)
	}
	if len(secondReplace.Artifacts) != 0 {
		t.Fatalf("expected second replacement to keep cleared artifacts, got %+v", secondReplace.Artifacts)
	}
	if secondReplace.Status != "ready" {
		t.Fatalf("expected second replacement status ready, got %q", secondReplace.Status)
	}
	if secondReplace.Action != "create" {
		t.Fatalf("expected second replacement to keep action create, got %q", secondReplace.Action)
	}
}

func TestReplaceFollowsLatestHistoricalCorrection(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	path, err := auditPath(baseDir, "P-0001")
	if err != nil {
		t.Fatalf("build audit path: %v", err)
	}

	appendSeedEvent(t, path, event{
		ID:        "01HZX3P5Q8K0M0G9H7K0D9Y1AA",
		TS:        "2026-04-22T00:00:00Z",
		Slug:      "P-0001",
		Phase:     "proposal_capture",
		Action:    "create",
		Status:    "partial",
		Prompt:    "process/roadmaps/proposals/CHANGE_PROPOSAL_CAPTURE_PROMPT.md",
		Artifacts: []string{"workflow/roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md"},
	})
	appendSeedEvent(t, path, event{
		ID:       "01HZX3P5Q8K0M0G9H7K0D9Y1AB",
		TS:       "2026-04-22T00:01:00Z",
		Slug:     "P-0001",
		Phase:    "proposal_capture",
		Action:   "create",
		Status:   "partial",
		Note:     "cleared mistaken fields",
		Replaces: stringPtr("01HZX3P5Q8K0M0G9H7K0D9Y1AA"),
	})
	appendSeedEvent(t, path, event{
		ID:       "01HZX3P5Q8K0M0G9H7K0D9Y1AC",
		TS:       "2026-04-22T00:02:00Z",
		Slug:     "P-0001",
		Phase:    "proposal_capture",
		Action:   "create",
		Status:   "ready",
		Note:     "updated status after clearing fields",
		Replaces: stringPtr("01HZX3P5Q8K0M0G9H7K0D9Y1AB"),
	})

	replaced := decodeEvent(t, captureStdout(t, func() {
		err := run([]string{
			"replace",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--id", "01HZX3P5Q8K0M0G9H7K0D9Y1AA",
			"--status", "done",
		})
		if err != nil {
			t.Fatalf("replace returned error: %v", err)
		}
	}))

	if replaced.Replaces == nil || *replaced.Replaces != "01HZX3P5Q8K0M0G9H7K0D9Y1AC" {
		t.Fatalf("expected replacement to supersede latest historical correction %q, got %#v", "01HZX3P5Q8K0M0G9H7K0D9Y1AC", replaced.Replaces)
	}
	if replaced.Prompt != "" {
		t.Fatalf("expected replacement to inherit cleared prompt, got %q", replaced.Prompt)
	}
	if len(replaced.Artifacts) != 0 {
		t.Fatalf("expected replacement to inherit cleared artifacts, got %+v", replaced.Artifacts)
	}
	if replaced.Note != "supersedes entry 01HZX3P5Q8K0M0G9H7K0D9Y1AC" {
		t.Fatalf("expected default note to reference effective entry ID, got %q", replaced.Note)
	}
	if replaced.Action != "create" {
		t.Fatalf("expected replacement to inherit effective logical action, got %q", replaced.Action)
	}
}

func TestReplaceRejectsAmbiguousHistoricalCorrectionBranches(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	path, err := auditPath(baseDir, "P-0001")
	if err != nil {
		t.Fatalf("build audit path: %v", err)
	}

	appendSeedEvent(t, path, event{
		ID:     "01HZX3P5Q8K0M0G9H7K0D9Y1BA",
		TS:     "2026-04-22T00:00:00Z",
		Slug:   "P-0001",
		Phase:  "proposal_capture",
		Action: "create",
		Status: "partial",
		Note:   "original entry",
	})
	appendSeedEvent(t, path, event{
		ID:       "01HZX3P5Q8K0M0G9H7K0D9Y1BB",
		TS:       "2026-04-22T00:00:01Z",
		Slug:     "P-0001",
		Phase:    "proposal_capture",
		Action:   "create",
		Status:   "partial",
		Note:     "branch A correction",
		Replaces: stringPtr("01HZX3P5Q8K0M0G9H7K0D9Y1BA"),
	})
	appendSeedEvent(t, path, event{
		ID:       "01HZX3P5Q8K0M0G9H7K0D9Y1BC",
		TS:       "2026-04-22T00:00:02Z",
		Slug:     "P-0001",
		Phase:    "proposal_capture",
		Action:   "create",
		Status:   "ready",
		Note:     "branch B correction",
		Replaces: stringPtr("01HZX3P5Q8K0M0G9H7K0D9Y1BA"),
	})

	err = run([]string{
		"replace",
		"--base-dir", baseDir,
		"--slug", "P-0001",
		"--id", "01HZX3P5Q8K0M0G9H7K0D9Y1BA",
		"--status", "done",
	})
	if err == nil {
		t.Fatal("expected ambiguous replace lineage to return an error")
	}
	if got := err.Error(); !strings.Contains(got, "replacement lineage for entry \"01HZX3P5Q8K0M0G9H7K0D9Y1BA\" is ambiguous") {
		t.Fatalf("expected ambiguity error, got %q", got)
	}
	if got := err.Error(); !strings.Contains(got, "01HZX3P5Q8K0M0G9H7K0D9Y1BB") || !strings.Contains(got, "01HZX3P5Q8K0M0G9H7K0D9Y1BC") {
		t.Fatalf("expected ambiguity error to list both leaf ids, got %q", got)
	}

	events, _, err := loadEvents(baseDir, "P-0001")
	if err != nil {
		t.Fatalf("loadEvents returned error: %v", err)
	}
	if len(events) != 3 {
		t.Fatalf("expected ambiguous replace not to append a new event, got %d events", len(events))
	}
}

func TestRunRejectsInvalidPhase(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	err := run([]string{
		"add",
		"--base-dir", baseDir,
		"--slug", "P-0001",
		"--phase", "not_a_phase",
		"--action", "create",
		"--status", "partial",
	})
	if err == nil {
		t.Fatal("expected invalid phase to return an error")
	}
}

func TestRunRejectsNonCanonicalSlugCase(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	err := run([]string{
		"add",
		"--base-dir", baseDir,
		"--slug", "p-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	})
	if err == nil {
		t.Fatal("expected lowercase slug to return an error")
	}
	if got := err.Error(); !strings.Contains(got, `invalid slug "p-0001"`) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestLoadEventsRejectsStoredSlugMismatch(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	path, err := auditPath(baseDir, "P-0001")
	if err != nil {
		t.Fatalf("build audit path: %v", err)
	}

	appendSeedEvent(t, path, event{
		ID:     "01HZX3P5Q8K0M0G9H7K0D9Y1CA",
		TS:     "2026-04-22T00:00:00Z",
		Slug:   "p-0001",
		Phase:  "proposal_capture",
		Action: "create",
		Status: "partial",
	})

	_, _, err = loadEvents(baseDir, "P-0001")
	if err == nil {
		t.Fatal("expected loadEvents to reject mismatched stored slug")
	}
	if got := err.Error(); !strings.Contains(got, `has slug "p-0001"; expected "P-0001"`) {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestReplaceRejectsClearPromptWithPrompt(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	added := decodeEvent(t, captureStdout(t, func() {
		err := run([]string{
			"add",
			"--base-dir", baseDir,
			"--slug", "P-0001",
			"--phase", "proposal_capture",
			"--action", "create",
			"--status", "partial",
		})
		if err != nil {
			t.Fatalf("seed add returned error: %v", err)
		}
	}))

	err := run([]string{
		"replace",
		"--base-dir", baseDir,
		"--slug", "P-0001",
		"--id", added.ID,
		"--status", "ready",
		"--clear-prompt",
		"--prompt", "process/roadmaps/proposals/CHANGE_PROPOSAL_CAPTURE_PROMPT.md",
	})
	if err == nil {
		t.Fatal("expected conflicting prompt flags to return an error")
	}
}

func TestReplaceRejectsClearArtifactsWithArtifacts(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	if err := run([]string{
		"add",
		"--base-dir", baseDir,
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	}); err != nil {
		t.Fatalf("seed add returned error: %v", err)
	}

	addedEvents, _, err := loadEvents(baseDir, "P-0001")
	if err != nil {
		t.Fatalf("loadEvents returned error: %v", err)
	}

	err = run([]string{
		"replace",
		"--base-dir", baseDir,
		"--slug", "P-0001",
		"--id", addedEvents[0].ID,
		"--status", "ready",
		"--clear-artifacts",
		"--artifact", "workflow/roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md",
	})
	if err == nil {
		t.Fatal("expected conflicting artifact flags to return an error")
	}
}

func TestAuditUsesCanonicalRepoPathAcrossAliases(t *testing.T) {
	tempDir := t.TempDir()
	baseDir := filepath.Join(tempDir, "repo")
	if err := os.MkdirAll(filepath.Join(baseDir, "workflow"), 0o755); err != nil {
		t.Fatalf("create base dir: %v", err)
	}

	aliasDir := filepath.Join(tempDir, "repo-link")
	if err := os.Symlink(baseDir, aliasDir); err != nil {
		t.Fatalf("create repo alias: %v", err)
	}

	if err := run([]string{
		"add",
		"--base-dir", baseDir,
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	}); err != nil {
		t.Fatalf("first add returned error: %v", err)
	}

	if err := run([]string{
		"add",
		"--base-dir", aliasDir,
		"--slug", "P-0001",
		"--phase", "review_fix_forward",
		"--action", "append",
		"--status", "ready",
	}); err != nil {
		t.Fatalf("second add returned error: %v", err)
	}

	realAuditPath, err := auditPath(baseDir, "P-0001")
	if err != nil {
		t.Fatalf("build real audit path: %v", err)
	}
	aliasAuditPath, err := auditPath(aliasDir, "P-0001")
	if err != nil {
		t.Fatalf("build alias audit path: %v", err)
	}
	if realAuditPath != aliasAuditPath {
		t.Fatalf("expected canonical audit path match, got %q and %q", realAuditPath, aliasAuditPath)
	}

	realLockPath, err := auditLockPath(baseDir, "P-0001")
	if err != nil {
		t.Fatalf("build real lock path: %v", err)
	}
	aliasLockPath, err := auditLockPath(aliasDir, "P-0001")
	if err != nil {
		t.Fatalf("build alias lock path: %v", err)
	}
	if realLockPath != aliasLockPath {
		t.Fatalf("expected canonical lock path match, got %q and %q", realLockPath, aliasLockPath)
	}

	events, _, err := loadEvents(baseDir, "P-0001")
	if err != nil {
		t.Fatalf("loadEvents returned error: %v", err)
	}
	if len(events) != 2 {
		t.Fatalf("expected 2 stored events, got %d", len(events))
	}
	for _, evt := range events {
		assertValidULID(t, evt.ID)
	}
}

func TestAuditLockPathIsRepoScopedAndIgnoresTmpDir(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	expected := filepath.Join(baseDir, "workflow", "logs", "audit", ".locks", "P-0001.lock")

	t.Setenv("TMPDIR", filepath.Join(t.TempDir(), "tmp-a"))
	firstPath, err := auditLockPath(baseDir, "P-0001")
	if err != nil {
		t.Fatalf("build first lock path: %v", err)
	}
	if firstPath != expected {
		t.Fatalf("expected repo-scoped lock path %q, got %q", expected, firstPath)
	}

	t.Setenv("TMPDIR", filepath.Join(t.TempDir(), "tmp-b"))
	secondPath, err := auditLockPath(baseDir, "P-0001")
	if err != nil {
		t.Fatalf("build second lock path: %v", err)
	}
	if secondPath != expected {
		t.Fatalf("expected repo-scoped lock path %q, got %q", expected, secondPath)
	}
}

func TestDefaultBaseDirUsesGitRepoRoot(t *testing.T) {
	tempDir := t.TempDir()
	repoDir := filepath.Join(tempDir, "repo")
	if err := os.MkdirAll(repoDir, 0o755); err != nil {
		t.Fatalf("create repo dir: %v", err)
	}
	initGitRepo(t, repoDir)
	if err := os.MkdirAll(filepath.Join(repoDir, "workflow"), 0o755); err != nil {
		t.Fatalf("create workflow dir: %v", err)
	}

	nestedDir := filepath.Join(repoDir, "subdir", "nested")
	if err := os.MkdirAll(nestedDir, 0o755); err != nil {
		t.Fatalf("create nested dir: %v", err)
	}

	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working dir: %v", err)
	}
	if err := os.Chdir(nestedDir); err != nil {
		t.Fatalf("chdir to nested dir: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(originalWD); err != nil {
			t.Fatalf("restore working dir: %v", err)
		}
	})

	if err := run([]string{
		"add",
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	}); err != nil {
		t.Fatalf("add returned error: %v", err)
	}

	repoAuditPath := filepath.Join(repoDir, "workflow", "logs", "audit", "P-0001.audit.log.jsonl")
	if _, err := os.Stat(repoAuditPath); err != nil {
		t.Fatalf("expected audit log at repo root, got err=%v", err)
	}

	nestedAuditPath := filepath.Join(nestedDir, "workflow", "logs", "audit", "P-0001.audit.log.jsonl")
	if _, err := os.Stat(nestedAuditPath); !os.IsNotExist(err) {
		t.Fatalf("expected no nested audit log, got err=%v", err)
	}
}

func TestDefaultBaseDirRequiresInstalledWorkflowRootInGitRepo(t *testing.T) {
	tempDir := t.TempDir()
	repoDir := filepath.Join(tempDir, "repo")
	if err := os.MkdirAll(repoDir, 0o755); err != nil {
		t.Fatalf("create repo dir: %v", err)
	}
	initGitRepo(t, repoDir)
	if err := os.MkdirAll(filepath.Join(repoDir, "templates", "workflow"), 0o755); err != nil {
		t.Fatalf("create templates workflow dir: %v", err)
	}

	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working dir: %v", err)
	}
	if err := os.Chdir(repoDir); err != nil {
		t.Fatalf("chdir to repo dir: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(originalWD); err != nil {
			t.Fatalf("restore working dir: %v", err)
		}
	})

	err = run([]string{
		"add",
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	})
	if err == nil {
		t.Fatal("expected add to fail without a live workflow root")
	}
	if got := err.Error(); !strings.Contains(got, "default base-dir resolution requires a workflow/ directory at the current git worktree root; otherwise set LCLP_BASE_DIR or pass --base-dir") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestExplicitBaseDirStillUsesRequestedTree(t *testing.T) {
	tempDir := t.TempDir()
	repoDir := filepath.Join(tempDir, "repo")
	if err := os.MkdirAll(repoDir, 0o755); err != nil {
		t.Fatalf("create repo dir: %v", err)
	}
	initGitRepo(t, repoDir)

	nestedDir := filepath.Join(repoDir, "subdir")
	if err := os.MkdirAll(nestedDir, 0o755); err != nil {
		t.Fatalf("create nested dir: %v", err)
	}

	customBaseDir := filepath.Join(repoDir, "custom-base")
	if err := os.MkdirAll(filepath.Join(customBaseDir, "workflow"), 0o755); err != nil {
		t.Fatalf("create custom base dir: %v", err)
	}
	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working dir: %v", err)
	}
	if err := os.Chdir(nestedDir); err != nil {
		t.Fatalf("chdir to nested dir: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(originalWD); err != nil {
			t.Fatalf("restore working dir: %v", err)
		}
	})

	if err := run([]string{
		"add",
		"--base-dir", customBaseDir,
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	}); err != nil {
		t.Fatalf("add returned error: %v", err)
	}

	customAuditPath := filepath.Join(customBaseDir, "workflow", "logs", "audit", "P-0001.audit.log.jsonl")
	if _, err := os.Stat(customAuditPath); err != nil {
		t.Fatalf("expected audit log in explicit base dir, got err=%v", err)
	}

	repoAuditPath := filepath.Join(repoDir, "workflow", "logs", "audit", "P-0001.audit.log.jsonl")
	if _, err := os.Stat(repoAuditPath); !os.IsNotExist(err) {
		t.Fatalf("expected no repo-root audit log for explicit base dir, got err=%v", err)
	}
}

func TestExplicitBaseDirRequiresInstalledWorkflowRoot(t *testing.T) {
	tempDir := t.TempDir()
	baseDir := filepath.Join(tempDir, "custom-base")
	if err := os.MkdirAll(baseDir, 0o755); err != nil {
		t.Fatalf("create custom base dir: %v", err)
	}

	err := run([]string{
		"add",
		"--base-dir", baseDir,
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	})
	if err == nil {
		t.Fatal("expected add to fail without an installed workflow root")
	}
	if got := err.Error(); !strings.Contains(got, "explicit --base-dir must already contain workflow/") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestExplicitBaseDirRejectsWorkflowSymlink(t *testing.T) {
	tempDir := t.TempDir()
	baseDir := filepath.Join(tempDir, "custom-base")
	externalWorkflowRoot := filepath.Join(tempDir, "external-workflow-root")
	if err := os.MkdirAll(baseDir, 0o755); err != nil {
		t.Fatalf("create custom base dir: %v", err)
	}
	if err := os.MkdirAll(externalWorkflowRoot, 0o755); err != nil {
		t.Fatalf("create external workflow root: %v", err)
	}
	if err := os.Symlink(externalWorkflowRoot, filepath.Join(baseDir, "workflow")); err != nil {
		t.Fatalf("create workflow symlink: %v", err)
	}

	err := run([]string{
		"add",
		"--base-dir", baseDir,
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	})
	if err == nil {
		t.Fatal("expected add to fail with a symlinked workflow root")
	}
	if got := err.Error(); !strings.Contains(got, "not a symlink") {
		t.Fatalf("unexpected error: %v", err)
	}

	externalAuditPath := filepath.Join(externalWorkflowRoot, "logs", "audit", "P-0001.audit.log.jsonl")
	if _, statErr := os.Stat(externalAuditPath); !os.IsNotExist(statErr) {
		t.Fatalf("expected no external audit log through workflow symlink, got err=%v", statErr)
	}
}

func TestExplicitBaseDirRejectsSymlinkedAuditDir(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)
	externalAuditDir := filepath.Join(t.TempDir(), "external-audit")
	if err := os.MkdirAll(filepath.Join(baseDir, "workflow", "logs"), 0o755); err != nil {
		t.Fatalf("create logs dir: %v", err)
	}
	if err := os.MkdirAll(externalAuditDir, 0o755); err != nil {
		t.Fatalf("create external audit dir: %v", err)
	}
	if err := os.Symlink(externalAuditDir, filepath.Join(baseDir, "workflow", "logs", "audit")); err != nil {
		t.Fatalf("create audit symlink: %v", err)
	}

	err := run([]string{
		"add",
		"--base-dir", baseDir,
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	})
	if err == nil {
		t.Fatal("expected add to fail with a symlinked audit dir")
	}
	if got := err.Error(); !strings.Contains(got, "audit storage path must not be a symlink") {
		t.Fatalf("unexpected error: %v", err)
	}

	externalAuditPath := filepath.Join(externalAuditDir, "P-0001.audit.log.jsonl")
	if _, statErr := os.Stat(externalAuditPath); !os.IsNotExist(statErr) {
		t.Fatalf("expected no external audit log through audit symlink, got err=%v", statErr)
	}
}

func TestExplicitBaseDirRejectsSymlinkedAuditLockDir(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)
	externalLockDir := filepath.Join(t.TempDir(), "external-locks")
	if err := os.MkdirAll(filepath.Join(baseDir, "workflow", "logs", "audit"), 0o755); err != nil {
		t.Fatalf("create audit dir: %v", err)
	}
	if err := os.MkdirAll(externalLockDir, 0o755); err != nil {
		t.Fatalf("create external lock dir: %v", err)
	}
	if err := os.Symlink(externalLockDir, filepath.Join(baseDir, "workflow", "logs", "audit", ".locks")); err != nil {
		t.Fatalf("create lock dir symlink: %v", err)
	}

	err := run([]string{
		"add",
		"--base-dir", baseDir,
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	})
	if err == nil {
		t.Fatal("expected add to fail with a symlinked audit lock dir")
	}
	if got := err.Error(); !strings.Contains(got, "audit storage path must not be a symlink") {
		t.Fatalf("unexpected error: %v", err)
	}

	externalLockPath := filepath.Join(externalLockDir, "P-0001.lock")
	if _, statErr := os.Stat(externalLockPath); !os.IsNotExist(statErr) {
		t.Fatalf("expected no external lock through .locks symlink, got err=%v", statErr)
	}
}

func TestExplicitDotBaseDirUsesCurrentDirectoryOutsideGit(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working dir: %v", err)
	}
	if err := os.Chdir(baseDir); err != nil {
		t.Fatalf("chdir to base dir: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(originalWD); err != nil {
			t.Fatalf("restore working dir: %v", err)
		}
	})

	if err := run([]string{
		"add",
		"--base-dir", ".",
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	}); err != nil {
		t.Fatalf("add returned error: %v", err)
	}

	auditPath := filepath.Join(baseDir, "workflow", "logs", "audit", "P-0001.audit.log.jsonl")
	if _, err := os.Stat(auditPath); err != nil {
		t.Fatalf("expected audit log in current directory, got err=%v", err)
	}
}

func TestEnvBaseDirUsesInstalledTreeOutsideGit(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)
	t.Setenv(baseDirEnvVar, baseDir)

	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working dir: %v", err)
	}
	outsideDir := t.TempDir()
	if err := os.Chdir(outsideDir); err != nil {
		t.Fatalf("chdir to outside dir: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(originalWD); err != nil {
			t.Fatalf("restore working dir: %v", err)
		}
	})

	if err := run([]string{
		"add",
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	}); err != nil {
		t.Fatalf("add returned error: %v", err)
	}

	auditPath := filepath.Join(baseDir, "workflow", "logs", "audit", "P-0001.audit.log.jsonl")
	if _, err := os.Stat(auditPath); err != nil {
		t.Fatalf("expected audit log in env-selected base dir, got err=%v", err)
	}
}

func TestExplicitBaseDirOverridesEnvBaseDir(t *testing.T) {
	envBaseDir := makeInstalledBaseDir(t)
	explicitBaseDir := makeInstalledBaseDir(t)
	t.Setenv(baseDirEnvVar, envBaseDir)

	if err := run([]string{
		"add",
		"--base-dir", explicitBaseDir,
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	}); err != nil {
		t.Fatalf("add returned error: %v", err)
	}

	explicitAuditPath := filepath.Join(explicitBaseDir, "workflow", "logs", "audit", "P-0001.audit.log.jsonl")
	if _, err := os.Stat(explicitAuditPath); err != nil {
		t.Fatalf("expected audit log in explicit base dir, got err=%v", err)
	}

	envAuditPath := filepath.Join(envBaseDir, "workflow", "logs", "audit", "P-0001.audit.log.jsonl")
	if _, err := os.Stat(envAuditPath); !os.IsNotExist(err) {
		t.Fatalf("expected no audit log in env base dir when explicit flag is set, got err=%v", err)
	}
}

func TestEnvBaseDirRequiresInstalledWorkflowRoot(t *testing.T) {
	baseDir := t.TempDir()
	t.Setenv(baseDirEnvVar, baseDir)

	err := run([]string{
		"add",
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	})
	if err == nil {
		t.Fatal("expected add to fail when LCLP_BASE_DIR lacks workflow/")
	}
	if got := err.Error(); !strings.Contains(got, "LCLP_BASE_DIR must already contain workflow/") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDefaultBaseDirOutsideGitRequiresExplicitFlag(t *testing.T) {
	baseDir := t.TempDir()

	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working dir: %v", err)
	}
	if err := os.Chdir(baseDir); err != nil {
		t.Fatalf("chdir to base dir: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(originalWD); err != nil {
			t.Fatalf("restore working dir: %v", err)
		}
	})

	err = run([]string{
		"add",
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	})
	if err == nil {
		t.Fatal("expected add to fail outside a git worktree without --base-dir")
	}
	if got := err.Error(); !strings.Contains(got, "default base-dir resolution requires LCLP_BASE_DIR, or running inside a git worktree and passing --base-dir explicitly") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDefaultBaseDirRejectsWorkflowSymlinkInGitRepo(t *testing.T) {
	tempDir := t.TempDir()
	repoDir := filepath.Join(tempDir, "repo")
	externalWorkflowRoot := filepath.Join(tempDir, "external-workflow-root")
	if err := os.MkdirAll(repoDir, 0o755); err != nil {
		t.Fatalf("create repo dir: %v", err)
	}
	if err := os.MkdirAll(externalWorkflowRoot, 0o755); err != nil {
		t.Fatalf("create external workflow root: %v", err)
	}
	initGitRepo(t, repoDir)
	if err := os.Symlink(externalWorkflowRoot, filepath.Join(repoDir, "workflow")); err != nil {
		t.Fatalf("create workflow symlink: %v", err)
	}

	originalWD, err := os.Getwd()
	if err != nil {
		t.Fatalf("get working dir: %v", err)
	}
	if err := os.Chdir(repoDir); err != nil {
		t.Fatalf("chdir to repo dir: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(originalWD); err != nil {
			t.Fatalf("restore working dir: %v", err)
		}
	})

	err = run([]string{
		"add",
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	})
	if err == nil {
		t.Fatal("expected add to fail with a symlinked workflow root")
	}
	if got := err.Error(); !strings.Contains(got, "workflow/ symlink") {
		t.Fatalf("unexpected error: %v", err)
	}

	externalAuditPath := filepath.Join(externalWorkflowRoot, "logs", "audit", "P-0001.audit.log.jsonl")
	if _, statErr := os.Stat(externalAuditPath); !os.IsNotExist(statErr) {
		t.Fatalf("expected no external audit log through workflow symlink, got err=%v", statErr)
	}
}

func TestReadLockWaitsForActiveWriter(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	if err := run([]string{
		"add",
		"--base-dir", baseDir,
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	}); err != nil {
		t.Fatalf("seed add returned error: %v", err)
	}

	lockPath, err := auditLockPath(baseDir, "P-0001")
	if err != nil {
		t.Fatalf("build lock path: %v", err)
	}
	lockFile, err := os.OpenFile(lockPath, os.O_CREATE|os.O_RDWR, 0o666)
	if err != nil {
		t.Fatalf("open lock file: %v", err)
	}
	defer lockFile.Close()

	if err := syscall.Flock(int(lockFile.Fd()), syscall.LOCK_EX); err != nil {
		t.Fatalf("acquire writer lock: %v", err)
	}

	done := make(chan error, 1)
	go func() {
		done <- withAuditReadLock(baseDir, "P-0001", func() error {
			events, _, err := loadEvents(baseDir, "P-0001")
			if err != nil {
				return err
			}
			if len(events) != 1 {
				return fmt.Errorf("expected 1 event, got %d", len(events))
			}
			return nil
		})
	}()

	select {
	case err := <-done:
		t.Fatalf("read lock should block behind writer lock, got early result: %v", err)
	case <-time.After(100 * time.Millisecond):
	}

	if err := syscall.Flock(int(lockFile.Fd()), syscall.LOCK_UN); err != nil {
		t.Fatalf("release writer lock: %v", err)
	}

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("read lock returned error: %v", err)
		}
	case <-time.After(1 * time.Second):
		t.Fatal("timed out waiting for read lock to complete")
	}
}

func TestTrackerCreatesHiddenRepoScopedAuditLockFile(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)

	if err := run([]string{
		"add",
		"--base-dir", baseDir,
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	}); err != nil {
		t.Fatalf("add returned error: %v", err)
	}

	lockPath := filepath.Join(baseDir, "workflow", "logs", "audit", ".locks", "P-0001.lock")
	if _, err := os.Stat(lockPath); err != nil {
		t.Fatalf("expected hidden repo-scoped audit lock file, got err=%v", err)
	}

	visibleSidecarPath := filepath.Join(baseDir, "workflow", "logs", "audit", "P-0001.audit.lock")
	if _, err := os.Stat(visibleSidecarPath); !os.IsNotExist(err) {
		t.Fatalf("expected no visible sidecar audit lock file, got err=%v", err)
	}
}

func TestTrackerRespectsUmaskForRuntimeAuditPermissions(t *testing.T) {
	baseDir := makeInstalledBaseDir(t)
	restoreUmaskForTest(t, 0o002)

	if err := run([]string{
		"add",
		"--base-dir", baseDir,
		"--slug", "P-0001",
		"--phase", "proposal_capture",
		"--action", "create",
		"--status", "partial",
	}); err != nil {
		t.Fatalf("add returned error: %v", err)
	}

	assertPerm(t, filepath.Join(baseDir, "workflow", "logs", "audit"), 0o775)
	assertPerm(t, filepath.Join(baseDir, "workflow", "logs", "audit", ".locks"), 0o775)
	assertPerm(t, filepath.Join(baseDir, "workflow", "logs", "audit", "P-0001.audit.log.jsonl"), 0o664)
	assertPerm(t, filepath.Join(baseDir, "workflow", "logs", "audit", ".locks", "P-0001.lock"), 0o664)
}

func restoreUmaskForTest(t *testing.T, mask int) {
	t.Helper()

	previous := syscall.Umask(mask)
	t.Cleanup(func() {
		syscall.Umask(previous)
	})
}

func assertPerm(t *testing.T, path string, want os.FileMode) {
	t.Helper()

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat %s: %v", path, err)
	}
	if got := info.Mode().Perm(); got != want {
		t.Fatalf("expected mode %o for %s, got %o", want, path, got)
	}
}

func captureStdout(t *testing.T, fn func()) []byte {
	t.Helper()

	original := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatalf("create pipe: %v", err)
	}

	os.Stdout = writer
	defer func() {
		os.Stdout = original
	}()

	fn()

	if err := writer.Close(); err != nil {
		t.Fatalf("close writer: %v", err)
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatalf("read stdout: %v", err)
	}
	if err := reader.Close(); err != nil {
		t.Fatalf("close reader: %v", err)
	}

	return bytes.TrimSpace(data)
}

func decodeEvent(t *testing.T, data []byte) event {
	t.Helper()

	var evt event
	if err := json.Unmarshal(data, &evt); err != nil {
		t.Fatalf("unmarshal event: %v", err)
	}
	return evt
}

func decodeEventList(t *testing.T, data []byte) []event {
	t.Helper()

	var events []event
	if err := json.Unmarshal(data, &events); err != nil {
		t.Fatalf("unmarshal event list: %v", err)
	}
	return events
}

func appendSeedEvent(t *testing.T, path string, evt event) {
	t.Helper()

	if err := appendEvent(path, evt); err != nil {
		t.Fatalf("append seed event: %v", err)
	}
}

func assertValidULID(t *testing.T, value string) {
	t.Helper()

	if strings.TrimSpace(value) == "" {
		t.Fatal("expected non-empty ULID")
	}
	if _, err := ulid.Parse(value); err != nil {
		t.Fatalf("expected valid ULID %q: %v", value, err)
	}
}

func initGitRepo(t *testing.T, dir string) {
	t.Helper()

	cmd := exec.Command("git", "init", "--quiet")
	cmd.Dir = dir
	if output, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("git init %q: %v (%s)", dir, err, bytes.TrimSpace(output))
	}
}
