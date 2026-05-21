You are performing a proposal-level rollup after checkpoint execution and
review cycles.

Your job is to assess the current repository state against a proposal-level
source of truth and summarize:

- what has landed
- what remains outstanding
- what differs materially from what was specified

Inputs may include:

- the proposal in `workflow/roadmaps/proposals/`
- related product and technical-design artifacts in
  `workflow/`
- current code and tests
- checkpoint execution and review summary logs in `workflow/logs/checkpoints/`

Use checkpoint summaries to recover, at minimum:

- decisions made
- docs likely needing update
- unresolved deviations
- cross-layer follow-ups

Proposal completion rule:

- If code landed but durable update obligations remain stale, treat the proposal
  as `partially landed` until those obligations are landed or explicitly
  deferred.

Write the rollup log under:

- `workflow/logs/checkpoints/`

Use this output structure exactly:

# Proposal Rollup: <proposal name>

## Verdict

- status: `<substantially landed | partially landed | blocked>`

## What Landed

- `<landed item>`
- `<landed item>`

## What Is Outstanding

- `<outstanding item>`
- `<outstanding item>`

## What Differs From Specified

- `<difference>`
- `<difference>`

## Cross-Layer State

- product-design layer: `<updated | partially updated | not updated | not applicable>`
- technical-design layer: `<updated | partially updated | not updated | not applicable>`

## Evergreen Follow-Up

- product docs to update:
  - `<doc>`
  - `<doc>`
- technical docs to update:
  - `<doc>`
  - `<doc>`
- decision logs likely needed:
  - `<log>`
  - `<log>`
- timing:
  - product docs: `<precondition | early | rollup | not applicable>`
  - technical docs: `<precondition | early | rollup | not applicable>`
  - decision logs: `<precondition | early | rollup | not applicable>`

## Notes

- `<note>`
