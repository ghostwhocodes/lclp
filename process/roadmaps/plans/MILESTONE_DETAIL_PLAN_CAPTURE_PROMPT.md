You are helping expand one milestone from a parent multi-milestone
implementation plan into a more detailed milestone-specific plan.

Your job is to interview the user as needed and then produce a detailed plan
for one milestone only.

Save a separate milestone plan to:

- `workflow/roadmaps/plans/<slug>/MILESTONE_PLAN_<slug>_M<n>.md`

Rules:

- This plan is subordinate to the parent multi-milestone implementation plan.
- Do not rewrite the whole parent plan.
- Expand only the selected milestone.
- Save milestone subplans in the root slug directory, not in `product/` or
  `technical/`, unless the user explicitly asks for a different layout.
- Use this prompt only when the milestone is complex enough to justify extra
  detail.
- If the milestone is already simple and clear, say that a separate milestone
  plan is probably unnecessary.
- Keep the plan concrete enough to guide execution, but do not turn it into a
  run log.
- Keep scope bounded to the milestone.
- Carry forward the parent's authorized-scope rule into this milestone plan.
- Do not treat milestone completion as permission to auto-advance into adjacent
  work unless the parent plan explicitly allows it.
- If the milestone is already simple and clear, do not invent a milestone plan.
  Instead, say so explicitly and state that execution may proceed directly
  against the parent milestone while still writing execution and review summary
  logs under `workflow/logs/checkpoints/`.

What you are trying to capture:

- the detailed internal structure of one milestone
- the bounded authorized scope for that milestone
- workstreams or sub-steps inside that milestone
- milestone-specific risks, dependencies, and validation
- the most sensible execution order inside the milestone

Process:

1. Ask for the parent implementation plan and the specific milestone to expand
   if they have not been supplied.
2. Ask a short first round of clarifying questions.
3. Ask no more than 8 questions in one round.
4. Prioritize questions that affect scope, sequencing, risk, and validation for
   that milestone.
5. Ask only when a blocker affects a core requirement, would violate a stated
   constraint or invariant, would require an irreversible external choice, or
   would force a decision between materially different tradeoff paths.
6. Otherwise, make a reasonable planning decision, record it, and keep moving.
7. If the milestone is too small to justify a separate plan, say so clearly and
   name the direct execution path plus required summary artifacts.
8. Otherwise, draft the milestone-specific plan.

Use this output structure exactly when a separate milestone plan is warranted:

# Milestone Plan: <proposal name> / Milestone <n> - <title>

## Summary

## Parent Context

## Authorized Scope

- this milestone plan is authorized to complete: `<bounded milestone scope>`
- auto-advance beyond this milestone when current scope is done: `<no by default | yes if explicitly approved>`
- stop and ask for approval only when: `<approval-needed decision categories>`

## Scope Boundary

## Work Breakdown

For each work item, use this exact sub-structure:

### Work Item <n>: <title>

- purpose: `<why this work item exists>`
- scope: `<what it covers>`
- dependencies: `<other work items or prerequisites>`
- likely artifacts touched: `<artifact types or areas>`
- completion signal: `<what indicates this work item is done>`

## Validation And Review Focus

## Risks And Watchpoints

## Out Of Scope

## Related Docs

When drafting:

- keep the plan subordinate to the parent milestone structure
- make the milestone more executable without exploding it into trivia
- keep clear what belongs inside this milestone and what does not
- assume the milestone plan will sit flat under `workflow/roadmaps/plans/<slug>/`
- keep the authorized scope explicit so downstream checkpoint execution stays bounded
- say explicitly if no separate milestone plan is warranted

If no separate milestone plan is warranted, use this output structure exactly:

# Milestone Plan Decision: <proposal name> / Milestone <n> - <title>

## Decision

- separate milestone plan warranted: `no`
- execute directly against: `<parent implementation plan> / Milestone <n>`
- execution summary log: `<filename>`
- review summary log: `<filename>`
- reason: `<why the parent milestone is already specific enough>`
- scope reminder: `<bounded scope that still applies>`
