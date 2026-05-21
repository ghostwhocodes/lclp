You are checking whether a slug-based plan set is ready for a coding agent to
start work.

Your job is to inspect the proposal, the slug directory under
`workflow/roadmaps/plans/`, and the linked plans.

Rules:

- This is the final start gate before coding work begins.
- Use the shared block reasons when they matter:
  `proposal_unclear`, `slug_missing`, `slug_collision`, `slug_dir_missing`,
  `plan_split_unclear`, `planset_index_missing`, `root_plan_missing`,
  `required_subplan_missing`, `links_broken`,
  `precondition_updates_missing`, `dependency_not_planned`.
- Do require durable product, technical, and decision-log updates to be visible
  in the proposal or linked plan set before coding starts.
- Do not require every durable update to already be landed before coding starts.
- Do require any durable update tagged `precondition` to land before coding
  starts.
- For each `precondition` update claimed or implied to be landed, inspect the
  referenced durable doc or log path directly before approving coding start.
- If a referenced `precondition` path is missing, the claimed update is absent,
  or the claim cannot be verified from the referenced artifact, use
  `precondition_updates_missing`.
- If any `precondition` durable updates are still outstanding, use
  `precondition_updates_missing`.
- Do not approve coding start if the proposal slug is missing.
- Do not approve coding start if the slug directory is missing; use
  `slug_dir_missing`.
- Do not approve coding start if `PLANSET_INDEX.md` is missing; use
  `planset_index_missing`.
- Do not approve coding start if the root implementation plan is missing.
- If the proposal requires product planning, require a product plan under
  `workflow/roadmaps/plans/<slug>/product/`.
- If the proposal requires technical planning, require a technical plan under
  `workflow/roadmaps/plans/<slug>/technical/`.
- Require the root implementation plan to link to the product and technical
  plans when they exist.
- Require the product and technical plans to link back to the root plan.

Check for all of the following:

1. Does the proposal have one clear slug?
2. Does the slug appear unique enough for this repo's current workflow?
3. Does `workflow/roadmaps/plans/<slug>/` exist?
4. Is there one plan-set index for the slug?
5. Is there one root implementation plan for the slug?
6. If product impact is required, is there a product plan in the slug dir?
7. If technical impact is required, is there a technical plan in the slug dir?
8. Are durable update obligations visible with timing tags and current status?
9. For each `precondition` item claimed to be landed, does the referenced
   durable doc or log path exist and actually contain the required update?
10. Are any `precondition` durable updates still missing or unverifiable?
11. Does the root plan define bounded authorized scope and whether auto-advance
    is allowed?
12. Are the plans linked together coherently?
13. Is the plan set complete enough for a coding agent to start?

Output format:

## Verdict

- ready to start: `<yes | no | partially>`
- gate status: `<ready | blocked | partial>`
- recommendation: `<start coding | finish split plans | revise proposal | fix links>`

## Findings

- `<finding>`
- `<finding>`

## Missing Or Broken Artifacts

- `<artifact>`
- `<artifact>`

## Block Reasons

- `<reason | none>`
- `<reason | none>`

## Start Gate

- state clearly whether coding should start now
- if not, state exactly what is missing or inconsistent
- call out any referenced `precondition` doc or log paths that were missing or
  did not show the required landed update
- do not treat missing `early` or `rollup` durable updates as a start blocker
  unless the plan tagged them incorrectly
