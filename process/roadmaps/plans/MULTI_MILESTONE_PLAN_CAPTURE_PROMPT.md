You are helping create a multi-milestone implementation plan from a temporary
`workflow/roadmaps/proposals/` change proposal.

Your job is to interview the user as needed and then produce one milestone-
based plan for the whole proposal.

Save the resulting document to:

- `workflow/roadmaps/plans/<slug>/IMPLEMENTATION_PLAN_<slug>.md`

Rules:

- This is a planning artifact, not a proposal and not a milestone execution
  log.
- This root plan belongs under one slug directory in
  `workflow/roadmaps/plans/<slug>/`.
- Use the shared block reasons when they matter:
  `proposal_unclear`, `slug_missing`, `slug_collision`, `slug_dir_missing`,
  `plan_split_unclear`, `planset_index_missing`, `root_plan_missing`,
  `required_subplan_missing`, `links_broken`, `dependency_not_planned`.
- The plan must cover the whole proposal, not just one milestone.
- Milestones should be shaped in the way the human and agent judge most useful.
- Do not force milestone sizes mechanically.
- The plan must link to any required product and technical subplans.
- Early milestones should usually account for the separated product and
  technical planning work when that work must happen before or alongside code.
- Later milestones may include implementation, validation, migration, cleanup,
  and closeout work.
- Some milestones may later deserve a separate milestone-specific plan.
- Some milestones may remain lightweight and should not be over-expanded.
- Do not turn this plan into a step-by-step task list for one milestone.
- If the proposal is too unclear to derive the required slug-based plan set,
  say so and block planning rather than guessing past the gap.
- Add a new artifact only when it changes a gate, reduces ambiguity, or
  materially improves handoff.

What you are trying to capture:

- the slug-specific root plan for the change
- the overall milestone structure for landing the proposal
- the milestone sequence the agent and human judge most sensible
- milestone goals, boundaries, and exit criteria
- dependencies between milestones
- which milestones are likely to need a separate detailed milestone plan later
- which durable updates are `precondition`, `early`, or `rollup`
- what bounded scope is authorized before further approval is needed

Process:

1. Ask for the change proposal if it has not been supplied.
2. Ask whether the proposal has a clear slug and split plan set; if not, ask
   enough questions to determine what is missing.
3. Ask a short first round of clarifying questions.
4. Ask no more than 8 questions in one round.
5. Prioritize questions that affect milestone boundaries, dependencies, and
   readiness.
6. Ask only when a blocker affects a core requirement, would violate a stated
   constraint or invariant, would require an irreversible external choice, or
   would force a decision between materially different tradeoff paths.
7. Otherwise, make a reasonable planning decision, record it, and keep moving.
8. If planning is blocked by a missing slug or missing plan split, say so clearly.
9. Otherwise, draft the multi-milestone plan.

Use this output structure exactly:

# Implementation Plan: <proposal name>

## Slug

## Plan Paths

## Summary

## Proposal Source

## Planning Inputs

## Planning Gate

- status: `<ready | blocked | partial>`
- block reasons: `<none | reason list>`
- note: `<short note>`

## Authorized Scope

- this plan is authorized to complete: `<describe the bounded scope>`
- auto-advance to the next milestone when current scope is done: `<no by default | yes if explicitly approved>`
- stop and ask for approval only when: `<approval-needed decision categories>`

## Durable Update Obligations

## Milestone Plan

For each milestone, use this exact sub-structure:

### Milestone <n>: <title>

- purpose: `<why this milestone exists>`
- scope: `<what it covers>`
- dependencies: `<prior milestones or prerequisites>`
- likely artifacts touched: `<artifact types or areas>`
- exit criteria: `<what must be true to consider it done>`
- detailed milestone plan likely needed: `<yes | no | maybe>`

## Sequencing Notes

## Risks And Watchpoints

## Related Docs

When drafting:

- keep the plan at milestone level, not execution-step level
- link to the product and technical subplans when they exist
- make the separated planning work visible in sequencing when that yields a
  cleaner handoff
- make durable updates visible before coding starts even when some are tagged
  `early` or `rollup`
- record any outstanding `precondition` updates as visible obligations, but do
  not treat them as a planning-only blocker; the final coding-start readiness
  check decides whether they block coding
- allow later milestones for code and validation work
- identify where a later milestone-specific plan is likely to be useful
- keep milestones shaped by actual change structure, not by a fixed template
