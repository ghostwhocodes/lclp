# LCLP Flow

This is the short canonical flow for the Layered Context Loop Process.

Use it to understand the handoff between proposal shaping, split planning,
coding-start gating, checkpoint execution, and rollup.

If you want one AI-facing supervisor entrypoint instead of manual phase
selection, start with `process/RUN_LCLP_FOR_CHANGE.md`.

## End-To-End Shape

1. create or refine a proposal under `workflow/roadmaps/proposals/`
2. assign or confirm one unique slug
3. make the required plan split and durable update obligations explicit
4. create the slug-based plan set under `workflow/roadmaps/plans/<slug>/`
5. run the separate coding-start readiness check and verify any referenced
   `precondition` durable updates directly in their docs or logs
6. execute work with checkpoint plans or, for explicitly lightweight
   milestones, direct parent-plan execution plus execution summaries and review
   loops
7. roll up the landed outcome back toward evergreen product and technical docs

## Prompt Map

Use these prompt files for the canonical flow:

1. `process/roadmaps/proposals/CHANGE_PROPOSAL_CAPTURE_PROMPT.md`
2. `process/roadmaps/proposals/PROPOSAL_TO_PLANSET_CAPTURE_PROMPT.md`
3. `process/roadmaps/plans/MULTI_MILESTONE_PLAN_CAPTURE_PROMPT.md`
4. `process/roadmaps/plans/PRODUCT_PLAN_CAPTURE_PROMPT.md` if product impact is
   required
5. `process/roadmaps/plans/TECHNICAL_PLAN_CAPTURE_PROMPT.md` if technical
   impact is required
6. `process/roadmaps/plans/CODING_START_READINESS_CHECK_PROMPT.md`
7. `process/logs/checkpoints/CHECKPOINT_PLAN_CAPTURE_PROMPT.md` or the direct
   lightweight milestone path when the plan explicitly allows it
8. `process/logs/checkpoints/EXECUTE_CHECKPOINT_PLAN_PROMPT.md`
9. `process/logs/checkpoints/REVIEW_AND_FIX_FORWARD_PROMPT.md`
10. `process/logs/checkpoints/PROPOSAL_ROLLUP_PROMPT.md`

Prompt ownership at the proposal-to-planning handoff:

- `PROPOSAL_TO_PLANSET_CAPTURE_PROMPT.md` is the split wrapper
- `MULTI_MILESTONE_PLAN_CAPTURE_PROMPT.md` owns the root implementation plan
- `PRODUCT_PLAN_CAPTURE_PROMPT.md` owns the product subplan when needed
- `TECHNICAL_PLAN_CAPTURE_PROMPT.md` owns the technical subplan when needed

Slug rule:

- prefer a user-supplied external issue key such as `ABC-123` when one exists
- otherwise generate a fallback slug such as `P-20260418-153045-1K7F`
- check for collisions against existing proposal artifacts and slug plan
  directories
- do not allocate slugs from a shared mutable counter file

## Gates

### Proposal Gate

Question:

- is the proposal complete enough to split into plans?

The proposal gate is about clarity, not coding readiness.
It is satisfied when the proposal has:

- one clear slug
- a clear plan split
- visible core requirements and red lines
- visible durable update obligations with timing tags
- enough scope and context to derive the root plan and any required subplans

### Planning Gate

Question:

- is the slug plan set complete enough to start coding?

The planning gate is satisfied when:

- the slug directory exists
- the plan-set index exists
- the root implementation plan exists
- required product and technical subplans exist
- durable update obligations are visible in the proposal or plan set
- any durable updates tagged `precondition` are already landed in their
  referenced durable docs or logs
- the root plan defines bounded authorized scope and whether auto-advance is allowed
- the plans link to each other coherently

### Rollup Gate

Question:

- is the change complete enough to roll up?

The rollup gate is satisfied when checkpoint execution and review artifacts are
rich enough to explain:

- what landed
- what differed from plan
- what docs likely need updating
- what decisions should survive into evergreen context
- whether any durable update obligations remain open or explicitly deferred

## Artifact Roles

- proposal: temporary change-shaping artifact
- slug plan set: temporary planning artifacts that define milestone sequencing
  and separated product or technical planning work
- evergreen docs: durable product and technical context outside the slug plan
  set
- checkpoint artifacts: temporary execution-facing artifacts that preserve
  momentum across context clears
- rollup: temporary closeout artifact that bridges execution outcomes back into
  evergreen updates

## Minimal Slug Plan Set

The usual minimum slug plan set is:

```text
workflow/roadmaps/plans/<slug>/
  PLANSET_INDEX.md
  IMPLEMENTATION_PLAN_<slug>.md
  product/
    PRODUCT_PLAN_<slug>.md        # if product impact is required
  technical/
    TECHNICAL_PLAN_<slug>.md      # if technical impact is required
```

`PLANSET_INDEX.md` is the human-readable landing page for the slug.
It is not the final coding-start approval artifact.

## Lightweight Milestone Path

If planning explicitly says a milestone is too small to justify a separate
checkpoint plan:

- execute directly against the parent plan milestone
- stay inside the same authorized scope and auto-advance rules
- still write execution and review summaries under `workflow/logs/checkpoints/`
- do not treat the missing checkpoint plan as a process failure

## Block Reasons

Use these shared reasons across proposal splitting, plan creation, and
coding-start gating when relevant:

- `proposal_unclear`
- `slug_missing`
- `slug_collision`
- `slug_dir_missing`
- `plan_split_unclear`
- `planset_index_missing`
- `root_plan_missing`
- `required_subplan_missing`
- `links_broken`
- `dependency_not_planned`

Use `precondition_updates_missing` only in the final coding-start gate when
required `precondition` updates are still not landed.

## Durable Update Timing

Use these timing tags for product, technical, and decision-log updates:

- `precondition`: must land before coding starts
- `early`: may land during execution and should usually be handled in early milestones
- `rollup`: may land after implementation work, before the proposal is closed out
- `not applicable`: no durable update is needed for that layer

## Anti-Bloat Rule

Add a new artifact only when it does at least one of these:

- changes a gate
- reduces ambiguity
- materially improves handoff

If it does none of those, it is probably process bloat.

## Worked Example

Worked examples live in the source LCLP repo and are intentionally not part of
the installed prompt pack.
