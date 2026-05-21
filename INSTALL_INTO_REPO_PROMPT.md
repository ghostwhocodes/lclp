# Install LCLP Into Another Repo

Supported install shape:

- source repo checkout contains exported prompt-pack material under
  `bootstrap/`, `process/`, and `templates/workflow/`
- target repo receives a live `workflow/` tree
- installed `workflow/` content is runtime state, not disposable template input
- the main prompt-pack install step is a manual copy of those three exported
  directories into an empty destination
- the optional tracker is a separate manual build step from `cmd/lclp-track/`
- repo-local source material such as `docs/`, `examples/`, `scripts/`, and
  `test/` is not part of the installable prompt pack

You are copying the Layered Context Loop Process from a source repository into
the current repository.

Copy these directories into a user-chosen base path inside the target repo:

- `process/`
- `bootstrap/`
- `templates/workflow/` as the destination `workflow/`

Ask before copying anything:

- ask for the install base path unless it is already supplied
- ask for the source path to the LCLP repo unless it is already supplied

Rules:

- do not invent a different directory layout
- copy the directories exactly as they exist in the source LCLP repo
- preserve filenames and subdirectories
- treat the source repo's `templates/workflow/` as the starter skeleton that
  becomes the live runtime tree after it is copied into the target repo as
  `workflow/`
- treat all `process/...`, `bootstrap/...`, and `workflow/...` references
  inside the installed prompt pack as relative to the chosen base path
- refuse installs back into the source LCLP repo itself
- refuse installs into a linked git worktree of the source LCLP repo
- if any of `<base path>/process/`, `<base path>/bootstrap/`, or
  `<base path>/workflow/` already exist, show the collisions and stop
- do not merge into an existing live `workflow/` tree
- do not overwrite existing install content in place
- if the target already has a `workflow/` tree with live proposals, plans,
  logs, or other runtime artifacts, treat that as live user state and stop
- if the target has a broken placeholder at `<base path>/workflow` such as a
  file or symlink rather than a directory, report it as a collision instead of
  partially installing around it
- if any existing component of `<base path>` is a symlink, stop instead of
  installing through it
- if any existing component of `<base path>` is not a directory, stop instead
  of partially installing around it
- do not copy source-repo reference material such as `docs/`, `examples/`,
  `scripts/`, `test/`, or maintainer-only release/governance docs into the
  target repo
- after copying, briefly summarize what was installed and where

Optional tracker:

- install the prompt pack first, then handle the tracker as a separate
  explicit step
- if the user wants the optional `lclp-track` binary and Go is available on a
  supported Linux host, build `cmd/lclp-track/` into
  `<base path>/bin/lclp-track`
- if the tracker is installed and the user wants to run it from CI or from
  outside the target repo, suggest setting `LCLP_BASE_DIR=<base path>` once or
  passing `--base-dir <base path>` per command
- run the tracker build from the source LCLP repo root so Go can see the
  `cmd/lclp-track/` module context
- `cmd/lclp-track/` is exported source material for the optional tracker, but
  it is not part of the main prompt-pack copy step
- refuse tracker installs back into the source LCLP repo itself
- refuse tracker installs into a linked git worktree of the source LCLP repo
- if any existing component of `<base path>/bin` is a symlink, stop instead of
  building through it
- if any existing component of `<base path>/bin` is not a directory, stop
  instead of partially building around it
- if you install the tracker, also ensure `<base path>/bin/.gitignore` keeps
  `lclp-track` out of normal git status
- do not report tracker install success on unsupported hosts
- do not overwrite an existing tracker binary in place

Install mapping:

- source `process/` -> `<base path>/process/`
- source `bootstrap/` -> `<base path>/bootstrap/`
- source `templates/workflow/` -> `<base path>/workflow/`

Path interpretation after install:

- `process/...` means `<base path>/process/...`
- `bootstrap/...` means `<base path>/bootstrap/...`
- `workflow/...` means `<base path>/workflow/...`
- runtime artifacts such as proposals, plans, checkpoints, and durable product
  or technical docs live under `<base path>/workflow/`
- transient tracker lockfiles under
  `<base path>/workflow/logs/audit/.locks/` should be ignored in the installed
  repo
- `process/` remains the definition layer
- `bootstrap/` remains the retrofit and install guidance layer
- source-repo docs, examples, tests, helper scripts, and maintainer-only
  governance material remain outside the installed target repo

Example for base path `ai/`:

- `process/` -> `ai/process/`
- `bootstrap/` -> `ai/bootstrap/`
- `workflow/` -> `ai/workflow/`
- `process/specs/product/` -> `ai/process/specs/product/`
- `process/specs/technical/` -> `ai/process/specs/technical/`
- `process/roadmaps/proposals/` -> `ai/process/roadmaps/proposals/`
- `process/roadmaps/plans/` -> `ai/process/roadmaps/plans/`
- `process/logs/checkpoints/` -> `ai/process/logs/checkpoints/`
- `workflow/bootstrap/` -> `ai/workflow/bootstrap/`
- `workflow/specs/product/` -> `ai/workflow/specs/product/`
- `workflow/specs/technical/` -> `ai/workflow/specs/technical/`
- `workflow/roadmaps/proposals/` -> `ai/workflow/roadmaps/proposals/`
- `workflow/roadmaps/plans/` -> `ai/workflow/roadmaps/plans/`
- `workflow/logs/checkpoints/` -> `ai/workflow/logs/checkpoints/`

Process:

1. Ask for the install base path if needed.
2. Ask for the source path if needed.
3. Inspect the destination for collisions.
4. If collisions exist, report them and stop.
5. Create the chosen base path if needed, then copy `process/`,
   `bootstrap/`, and `templates/workflow/` into it so the installed
   destinations are exactly `<base path>/process/`, `<base path>/bootstrap/`,
   and `<base path>/workflow/`.
6. If requested and supported, run a separate tracker build into
   `<base path>/bin/lclp-track`.
7. Confirm the installed destinations.

When reporting completion, include:

- the base path used
- the source path used
- the top-level installed directories
- whether the optional tracker was installed, skipped, or unsupported
- any collisions that blocked the install
