You are in planning mode and must create a saved internal checkpoint plan for
Codex execution.

Your job is to take either:

- a milestone-specific plan from `workflow/roadmaps/plans/`, or
- one milestone reference inside a larger implementation plan in
  `workflow/roadmaps/plans/`

and turn it into a concrete Codex-internal checkpoint plan.

This plan must be saved under:

- `workflow/logs/checkpoints/`

Rules:

- Preserve the milestone intent from the parent plan.
- This checkpoint plan should reflect Codex's own execution shaping.
- Do not over-expand it into a line-by-line implementation script.
- Keep it concrete enough to drive execution, but bounded enough to remain
  automatable.
- The checkpoint plan is allowed to be more execution-shaped than the human
  milestone plan, but it must remain subordinate to that milestone.
- The checkpoint must stay inside the parent plan's authorized scope.
- Do not treat completion of this checkpoint as permission to auto-advance into
  the next milestone unless the parent plan explicitly allows it.
- If the milestone is too small to justify a separate checkpoint plan, do not
  invent one. Instead, say so explicitly, direct execution against the parent
  plan milestone, and name the required execution and review summary logs that
  still must be written under `workflow/logs/checkpoints/`.
- Add a new artifact only when it changes a gate, reduces ambiguity, or
  materially improves handoff.

Use this output structure exactly when a separate checkpoint plan is warranted:

# Checkpoint Plan: <proposal or milestone name>

## Summary

## Parent Plan Context

## Authorized Scope

- this checkpoint is authorized to complete: `<bounded checkpoint scope>`
- auto-advance after completion: `<no | yes if explicitly approved in parent plan>`

## Scope Boundary

## Workstreams

For each workstream:

### Workstream <n>: <title>

- purpose: `<why this workstream exists>`
- scope: `<what it covers>`
- dependencies: `<other workstreams or prerequisites>`
- completion signal: `<what indicates it is done>`

## Validation Focus

## Risks And Watchpoints

## Expected Summary Outputs

- execution summary log: `<filename>`
- review summary log: `<filename>`
- proposal rollup log: `<filename if relevant>`
- expected rollup inputs:
  - `decisions made`
  - `docs likely needing update`
  - `unresolved deviations`

## Related Docs

If no separate checkpoint plan is warranted, use this output structure exactly:

# Checkpoint Decision: <proposal or milestone name>

- checkpoint plan needed: `no`
- execute directly against: `<parent plan name> / Milestone <n>`
- execution summary log: `<filename>`
- review summary log: `<filename>`
- proposal rollup log: `<filename if relevant>`
- reason: `<why a separate checkpoint plan is unnecessary>`
- scope reminder: `<bounded scope that still applies>`
