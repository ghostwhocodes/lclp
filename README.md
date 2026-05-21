# Layered Context Loop Process

`LCLP` is a lightweight human-plus-LLM process for shaping, planning,
executing, and reviewing software changes across durable and temporary context
layers.

This repository contains:

- the exported process definition, templates, and prompts in [process/](./process/README.md)
- the exported bootstrap and retrofit guidance in [bootstrap/](./bootstrap/README.md)
- the exported runtime skeleton in [templates/workflow/](./templates/workflow/README.md)
- repo-local reference docs in [docs/](./docs/TRACKER_MINIMALISM_CASE_STUDY.md)
- repo-local worked examples in [examples/](./examples/README.md)
- optional released tracker source in [`cmd/lclp-track/`](./cmd/lclp-track/main.go)

## Start Here

Choose the entrypoint that matches your situation:

- I want the AI to choose the next valid phase for one change:
  [process/RUN_LCLP_FOR_CHANGE.md](./process/RUN_LCLP_FOR_CHANGE.md)
- I want to run the phases manually:
  [process/README.md](./process/README.md)
- I need manual copy/install guidance for another repository, or I want to
  build the optional tracker:
  [INSTALL_INTO_REPO_PROMPT.md](./INSTALL_INTO_REPO_PROMPT.md)
- I need to bootstrap durable context in an existing repository:
  [bootstrap/README.md](./bootstrap/README.md)
- I want to see a complete worked example first:
  [examples/README.md](./examples/README.md)

## What LCLP Separates

LCLP keeps durable context distinct from temporary change work.

Durable layers:

- product specs
- technical specs
- decision logs

Temporary layers:

- change proposals
- plan sets
- checkpoint plans
- execution and review summaries

That split is the point of the process: durable docs stay stable, while
day-to-day change work can move quickly without overwriting long-lived context.

## Repo Layout

There are two different scopes in this project:

- private development repo: this repo, including maintainer-only material such
  as examples, docs, tests, and any private governance material kept outside
  the install boundary
- installed target repo: the unrelated user repo that receives copied LCLP
  material

Only part of this repo should be copied into a target repo.

Exported/installable boundary:

- `bootstrap/`: copied into the target repo as-is
- `process/`: copied into the target repo as-is
- `templates/workflow/`: copied into the target repo as `workflow/`
- `cmd/lclp-track/`: optional released source/tooling, built separately only if wanted

Repo-local source material:

- `docs/`: maintainer and design-reference notes; not copied into target repos
- `examples/`: curated reference artifacts; not copied into target repos
- `scripts/`: source-repo helpers; not copied into target repos
- `test/`: source-repo validation; not copied into target repos

The main runtime artifact families are:

1. `workflow/bootstrap/`
2. `workflow/specs/product/`
3. `workflow/specs/technical/`
4. `workflow/roadmaps/proposals/`
5. `workflow/roadmaps/plans/`
6. `workflow/logs/checkpoints/`

Boundary:

- `process/` defines how the process works
- this repo's `templates/workflow/` directory is the starter tree
- an installed copy's `workflow/` directory stores the artifacts created while running it
- `docs/`, `examples/`, `scripts/`, and `test/` stay in this source repo

## Common Paths

### 1. Run One Change In A Repo That Already Has Context

If the repository already has enough durable product and technical context,
use this order:

1. allocate or confirm a slug with
   [`process/roadmaps/proposals/ALLOCATE_PROPOSAL_SLUG_PROMPT.md`](./process/roadmaps/proposals/ALLOCATE_PROPOSAL_SLUG_PROMPT.md)
   if the user has not already supplied one
2. draft the proposal with
   [`process/roadmaps/proposals/CHANGE_PROPOSAL_CAPTURE_PROMPT.md`](./process/roadmaps/proposals/CHANGE_PROPOSAL_CAPTURE_PROMPT.md)
3. split the proposal into a plan set with
   [`process/roadmaps/proposals/PROPOSAL_TO_PLANSET_CAPTURE_PROMPT.md`](./process/roadmaps/proposals/PROPOSAL_TO_PLANSET_CAPTURE_PROMPT.md)
4. write the root implementation plan with
   [`process/roadmaps/plans/MULTI_MILESTONE_PLAN_CAPTURE_PROMPT.md`](./process/roadmaps/plans/MULTI_MILESTONE_PLAN_CAPTURE_PROMPT.md)
5. add product and technical plans only when needed with
   [`PRODUCT_PLAN_CAPTURE_PROMPT.md`](./process/roadmaps/plans/PRODUCT_PLAN_CAPTURE_PROMPT.md)
   and
   [`TECHNICAL_PLAN_CAPTURE_PROMPT.md`](./process/roadmaps/plans/TECHNICAL_PLAN_CAPTURE_PROMPT.md)
6. run the coding start gate with
   [`process/roadmaps/plans/CODING_START_READINESS_CHECK_PROMPT.md`](./process/roadmaps/plans/CODING_START_READINESS_CHECK_PROMPT.md)
7. create a checkpoint plan with
   [`process/logs/checkpoints/CHECKPOINT_PLAN_CAPTURE_PROMPT.md`](./process/logs/checkpoints/CHECKPOINT_PLAN_CAPTURE_PROMPT.md)
   unless the parent plan explicitly allows direct execution
8. execute with
   [`process/logs/checkpoints/EXECUTE_CHECKPOINT_PLAN_PROMPT.md`](./process/logs/checkpoints/EXECUTE_CHECKPOINT_PLAN_PROMPT.md)
9. review and fix forward with
   [`process/logs/checkpoints/REVIEW_AND_FIX_FORWARD_PROMPT.md`](./process/logs/checkpoints/REVIEW_AND_FIX_FORWARD_PROMPT.md)
10. close out with
    [`process/logs/checkpoints/PROPOSAL_ROLLUP_PROMPT.md`](./process/logs/checkpoints/PROPOSAL_ROLLUP_PROMPT.md)

Example artifact paths for slug `P-0001`:

- proposal:
  `workflow/roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md`
- plan set landing page:
  `workflow/roadmaps/plans/P-0001/PLANSET_INDEX.md`
- root plan:
  `workflow/roadmaps/plans/P-0001/IMPLEMENTATION_PLAN_P-0001.md`
- optional product plan:
  `workflow/roadmaps/plans/P-0001/product/PRODUCT_PLAN_P-0001.md`
- optional technical plan:
  `workflow/roadmaps/plans/P-0001/technical/TECHNICAL_PLAN_P-0001.md`
- checkpoint outputs:
  `workflow/logs/checkpoints/P-0001_M1_CHECKPOINT_PLAN.md`,
  `workflow/logs/checkpoints/P-0001_M1_EXECUTION_SUMMARY.md`,
  `workflow/logs/checkpoints/P-0001_M1_REVIEW_SUMMARY.md`

If you want one AI prompt to route the phase for you, start with
[process/RUN_LCLP_FOR_CHANGE.md](./process/RUN_LCLP_FOR_CHANGE.md) instead of
walking the list manually.

### 2. Install LCLP Into Another Repository

Install by manually copying the shipped source material.

Main prompt-pack install:

1. Choose a repo-relative base path in the target repo, such as `ai/` or `.`.
2. Confirm the target repo is not this source repo and is not one of this
   repo's linked git worktrees.
3. Stop if `<base path>/bootstrap/`, `<base path>/process/`, or
   `<base path>/workflow/` already exist.
4. Stop if any existing component of `<base path>` is a symlink or a
   non-directory.
5. Create `<base path>/` if it does not already exist.
6. Copy `bootstrap/` exactly as it exists here.
7. Copy `process/` exactly as it exists here.
8. Copy `templates/workflow/` into the target as `workflow/`.

Install contract:

- that three-directory copy is the entire main prompt-pack install
- do not merge into an existing live `workflow/` tree
- do not copy source-repo-only material such as `docs/`, `examples/`,
  `scripts/`, `test/`, or maintainer-only governance/release notes
- after install, every prompt reference remains relative to the chosen base
  path inside the target repo

Example:

- chosen base path: `ai/`
- installed directories:
  `ai/process/`, `ai/bootstrap/`, and `ai/workflow/`
- a prompt reference like `process/roadmaps/plans/...` means
  `ai/process/roadmaps/plans/...` inside the target repo

Optional tracker install:

- install the prompt pack first
- build `cmd/lclp-track/` separately on Linux only
- place the binary at `<base path>/bin/lclp-track`
- keep `lclp-track` out of normal git status, typically with
  `<base path>/bin/.gitignore`
- if you want to call the tracker from CI or outside the target repo, set
  `LCLP_BASE_DIR=<base path>` once or pass `--base-dir <base path>` per
  command
- do not build into this source repo or one of its linked worktrees
- do not overwrite an existing tracker binary in place

If the target repository does not yet have durable product and technical
artifacts, continue with [bootstrap/README.md](./bootstrap/README.md) after the
copy.

If you want AI-assisted install guidance for that same manual-copy flow, use
[INSTALL_INTO_REPO_PROMPT.md](./INSTALL_INTO_REPO_PROMPT.md).

### 3. Review The Intended Shape Before Using It

The worked example under
[examples/README.md](./examples/README.md) includes two curated source-repo
reference sets. Start with the happy-path example for the default clean
completion flow, then use the fix-forward example for a partial milestone and
review-rejected recovery path. These examples remain source-repo reference
material only and are not copied into installed target repositories. Their
artifact paths are intentionally example-local so the sets can be read in
place without pretending to be live installed `workflow/` trees.

## Optional Audit CLI

This repository includes a small Go CLI for append-only audit logging:

- source: [`cmd/lclp-track/`](./cmd/lclp-track/main.go)
- design note: [docs/TRACKER_MINIMALISM_CASE_STUDY.md](./docs/TRACKER_MINIMALISM_CASE_STUDY.md)
- output: `workflow/logs/audit/<slug>.audit.log.jsonl`
- runtime notes:
  [templates/workflow/logs/audit/README.md](./templates/workflow/logs/audit/README.md)
- runtime locks: `workflow/logs/audit/.locks/<slug>.lock`
- platform support: Linux only; non-Linux builds compile, but tracker commands
  return an unsupported-platform error because audit locking uses `flock`

The tracker source is released separately on purpose. The main prompt-pack
install copies only `bootstrap/`, `process/`, and `templates/workflow/`.

Build it with:

```bash
go build -o lclp-track ./cmd/lclp-track
```

Or build it directly into an already installed prompt-pack repo as
`<base path>/bin/lclp-track` on Linux, and keep that binary out of normal git
status with `<base path>/bin/.gitignore`.

If you build the tracker in this source repository, either pass `--base-dir`
explicitly to point at the target repo root or another live `workflow/` tree,
or set `LCLP_BASE_DIR` once before running commands.

Commands:

- `add`
- `replace`
- `last`
- `get`

Allowed phases:

- `bootstrap`
- `proposal_capture`
- `plan_split`
- `root_plan`
- `product_plan`
- `technical_plan`
- `coding_start_gate`
- `milestone_detail`
- `checkpoint_plan`
- `execution`
- `review_fix_forward`
- `proposal_rollup`

Behavior that matters:

- base-dir resolution order is: `--base-dir`, then `LCLP_BASE_DIR`, then the
  current git worktree root when that root already has a real `workflow/`
  directory
- use `LCLP_BASE_DIR` when you want one stable install root across many calls
  in CI or shell sessions
- pass `--base-dir` explicitly when a script must be fully self-contained,
  including `--base-dir .`
- if neither `--base-dir` nor `LCLP_BASE_DIR` is set and the current git
  worktree root does not have a live `workflow/` directory, the tracker fails
- `add` requires `--slug`, `--phase`, `--action`, and `--status`
- event IDs are opaque ULID strings emitted in command output and stored in the log
- `replace` requires `--slug` and `--id`, and inherits `phase`, `action`,
  `status`, `prompt`, and `artifacts` from the single effective corrective
  leaf in that lineage when those flags are omitted
- if that replacement lineage has multiple live leaves, `replace` fails and you
  must rerun it with the specific leaf event ID you want to supersede
- `replace --clear-prompt` clears a mistaken prompt instead of inheriting it
- `replace --clear-artifacts` clears mistaken artifacts instead of inheriting them
- `last` prints the most recent event for a slug
- `get` prints a JSON array of effective live events by default and can filter
  by `--id`, `--phase`, `--action`, or `--status`
- `get --all` returns raw append-only history instead of only live leaf events
- slug values must match `^[A-Z0-9._-]+$`
- stored event `slug` values must exactly match the requested slug or the log is
  rejected as inconsistent
- logs are append-only; `replace` records a corrective event instead of
  rewriting prior lines

Example:

```bash
export LCLP_BASE_DIR=/path/to/repo

./lclp-track add \
  --slug P-0001 \
  --phase proposal_capture \
  --action create \
  --status partial \
  --prompt process/roadmaps/proposals/CHANGE_PROPOSAL_CAPTURE_PROMPT.md \
  --artifact workflow/roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md \
  --note "initial draft created"

./lclp-track replace \
  --slug P-0001 \
  --id 01K06W6GG1Z5KJ0QMD4X9N9D4H \
  --action create \
  --status ready \
  --note "corrected operator mistake"

./lclp-track last --slug P-0001
./lclp-track get --slug P-0001 --phase proposal_capture
./lclp-track get --slug P-0001 --status ready
./lclp-track get --slug P-0001 --all
```

Typical stored event:

```json
{
  "id": "01K06W6GG1Z5KJ0QMD4X9N9D4H",
  "ts": "2026-04-19T01:40:24Z",
  "slug": "P-0001",
  "phase": "proposal_capture",
  "action": "create",
  "status": "partial",
  "prompt": "process/roadmaps/proposals/CHANGE_PROPOSAL_CAPTURE_PROMPT.md",
  "artifacts": [
    "workflow/roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md"
  ],
  "note": "initial draft created"
}
```

If you want one default install root outside the current git worktree, set
`LCLP_BASE_DIR`:

```bash
export LCLP_BASE_DIR=/path/to/repo
./lclp-track last --slug P-0001
```

If a particular command needs a different root, use `--base-dir`:

```bash
./lclp-track add \
  --base-dir /path/to/repo \
  --slug P-0001 \
  --phase execution \
  --action continue \
  --status partial \
  --artifact workflow/logs/checkpoints/P-0001_M1_EXECUTION_SUMMARY.md \
  --note "checkpoint execution resumed"
```

## Why This Exists

LCLP is meant for repositories where:

- humans and LLMs collaborate continuously
- context is too large to keep in one chat
- product, technical design, and implementation move at different speeds
- temporary execution artifacts should not replace durable docs

## Version

This repository contains `LCLP v1`.

## License

This repository is licensed under the Creative Commons Attribution 4.0
International License.

Exception:

- the Go tool in `cmd/lclp-track/` is licensed under MIT; see
  [`cmd/lclp-track/LICENSE`](./cmd/lclp-track/LICENSE)

Attribution should credit:

- `Layered Context Loop Process (LCLP)`
- `Copyright (c) 2026 Nos Doughty`

See [LICENSE](./LICENSE) for the full license text.
