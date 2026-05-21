# Repo Boundary Notes

LCLP has two distinct boundaries that must not be mixed:

1. this private development repo
2. the user's unrelated target repo that receives an installed copy

The exported/installable boundary is:

- `bootstrap/`
- `process/`
- `templates/workflow/`
- `cmd/lclp-track/`

These paths must stay safe to ship to other repositories. They should contain
only the reusable LCLP process, runtime skeleton, or optional tracker source.

Release/install model:

- the private development repo contains maintainer-only material
- users then manually copy only `bootstrap/`, `process/`, and
  `templates/workflow/` from this repo into a user's target repo
- `cmd/lclp-track/` is released as optional source/tooling, but is not part of
  the main prompt-pack copy step

Important install/export rules:

- install-facing docs should describe a manual copy of `bootstrap/`,
  `process/`, and `templates/workflow/` as `workflow/`
- tracker installation is a separate manual build of `cmd/lclp-track/` into a
  target-repo binary
- files inside shipped `process/` or `templates/workflow/` must not depend on
  repo-local-only paths such as `examples/`, `docs/`, or release/test
  scripts
- do not place maintainer-only governance prompts, release-review prompts, or
  source-repo operational notes under the exported/installable paths

Repo-local, non-installed material includes:

- `examples/`
- `docs/`
- `scripts/`
- `test/`

If a file would make an engineer in a target repo ask "why is this here?",
it probably belongs outside the exported/installable boundary.
