# Checkpoint Layer

This folder contains prompts and guidance for the checkpoint layer of the
Layered Context Loop.

The checkpoint layer sits between planning and direct coding work.

Use it to:

- create saved internal Codex checkpoint plans
- run or resume work against those checkpoint plans
- run lightweight milestone work directly against the parent plan only when
  planning explicitly said no separate checkpoint plan is warranted
- capture execution summaries
- review against parent plans and create fix-forward checkpoint plans
- produce proposal-level rollup summaries
- preserve decisions, doc follow-ups, and unresolved deviations for rollup

The live generated artifacts for this layer belong under:

- `workflow/logs/checkpoints/` in an installed target repo
- `templates/workflow/logs/checkpoints/` in this source repo

Maintainer-only governance or release-boundary review prompts should live
outside the installed `process/` tree.
