You are taking one proposal and splitting it into the plan set that will live
under its slug directory.

Your job is to read a proposal from `workflow/roadmaps/proposals/` and create
or refine the plan outputs that belong under:

- `workflow/roadmaps/plans/<slug>/`
- `workflow/roadmaps/plans/<slug>/product/`
- `workflow/roadmaps/plans/<slug>/technical/`

This prompt is the split-orchestration wrapper for the planning handoff.
It owns the slug-based split, required-plan decision, slug directory shape, and
`PLANSET_INDEX.md`.
The root implementation plan it creates or refines should follow the structure
defined by `process/roadmaps/plans/MULTI_MILESTONE_PLAN_CAPTURE_PROMPT.md` and
`process/roadmaps/plans/IMPLEMENTATION_PLAN_TEMPLATE.md`.
Any product or technical subplans it creates or refines should follow the
corresponding prompt and template files under `process/roadmaps/plans/`.

Rules:

- Do not start coding.
- Do not treat this as a final coding-start readiness check.
- Use the shared block reasons when they matter:
  `proposal_unclear`, `slug_missing`, `slug_collision`, `slug_dir_missing`,
  `plan_split_unclear`, `planset_index_missing`, `root_plan_missing`,
  `required_subplan_missing`, `links_broken`, `dependency_not_planned`.
- First confirm that the proposal has a clear slug.
- If the slug is missing, stop and say so clearly.
- If the proposal is too unclear to derive the required plan set, stop and say
  so clearly.
- Always create or refine `workflow/roadmaps/plans/<slug>/PLANSET_INDEX.md`.
- Treat `PLANSET_INDEX.md` as the slug landing page and planning snapshot only.
- Do not stamp a final coding-start approval or denial into `PLANSET_INDEX.md`.
- Always create or refine the root implementation plan under the slug dir.
- Create a product plan under the slug dir if the proposal has product impact.
- Create a technical plan under the slug dir if the proposal has technical
  impact.
- Carry the proposal's durable update obligations into the plan set and make
  their timing tags visible.
- If any durable updates are tagged `precondition`, carry them forward as
  visible follow-up obligations; the final coding-start readiness check decides
  whether they block coding.
- If user-supplied resources belong with product or technical planning, place
  them in the corresponding subdirectory.
- The root plan must link to any product and technical subplans.
- Product and technical subplans must link back to the root plan.
- Add a new artifact only when it changes a gate, reduces ambiguity, or
  materially improves handoff.

What you are trying to produce:

- one coherent slug-based plan set
- one human-readable slug landing page
- one root implementation plan
- separated product and technical plans when required
- visible durable update obligations with `precondition`, `early`, or `rollup`
  timing
- no self-certified coding-start verdict inside the plan-set index
- clear linkage between the plans

Process:

1. Read the proposal.
2. Confirm the proposal slug.
3. Confirm whether product and technical subplans are required.
4. Create or refine the plan outputs in the slug directory.
5. State what was created and what still needs follow-up.

Output format:

## Split Result

- slug: `<slug>`
- gate status: `<ready | blocked | partial>`
- block reasons: `<none | reason list>`
- plan-set index: `<path>`
- root plan: `<path>`
- product plan: `<path | not needed>`
- technical plan: `<path | not needed>`

## Findings

- `<finding>`
- `<finding>`

## Remaining Gaps

- `<gap>`
- `<gap>`

## Next Step

- `<next step>`
