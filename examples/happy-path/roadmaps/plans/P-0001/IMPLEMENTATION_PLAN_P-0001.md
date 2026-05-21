# Implementation Plan: Plan-Set Index And Coding-Start Gate Cleanup

> Artifact type: `temporary planning artifact`
>
> Curated source-repo example only. In a live installed repo, the equivalent
> artifact would live under `workflow/...` rather than `examples/...`.

## Slug

- slug: `P-0001`
- root dir: `examples/happy-path/roadmaps/plans/P-0001/`

## Plan Paths

- plan-set index: `examples/happy-path/roadmaps/plans/P-0001/PLANSET_INDEX.md`
- root plan: `examples/happy-path/roadmaps/plans/P-0001/IMPLEMENTATION_PLAN_P-0001.md`
- product plan: `examples/happy-path/roadmaps/plans/P-0001/product/PRODUCT_PLAN_P-0001.md`
- technical plan: `examples/happy-path/roadmaps/plans/P-0001/technical/TECHNICAL_PLAN_P-0001.md`

## Summary

Land a clearer split between proposal shaping and coding-start gating by adding
one slug landing page, aligning the prompts, and improving checkpoint-to-rollup
handoff.

## Proposal Source

- proposal path: `examples/happy-path/roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md`

## Planning Inputs

- proposal split already identifies both product and technical planning work
- current ambiguity sits at the plan-set handoff, not the proposal itself

## Planning Gate

- status: `ready`
- block reasons: `none`
- note: `coding may start because the required subplans exist, the durable update obligations are visible, no precondition items remain open, and the current milestone scope is bounded`

## Authorized Scope

- this plan is authorized to complete: `proposal-to-plan, start-gate, checkpoint, and rollup contract cleanup for P-0001`
- auto-advance to the next milestone when current scope is done: `no by default`
- stop and ask for approval only when: `a change would alter user-visible workflow materially, violate a stated constraint, require an irreversible external choice, or force a materially different tradeoff path`

## Durable Update Obligations

- product docs:
  - doc: `process/FLOW.md`
  - why it matters: `operators need the visibility-before-coding rule stated clearly`
  - timing: `early`
  - status: `identified`
- technical docs:
  - doc: `process/roadmaps/plans/IMPLEMENTATION_PLAN_TEMPLATE.md`
  - why it matters: `the plan contract must carry scope and durable update obligations`
  - timing: `early`
  - status: `identified`
- decision logs:
  - doc: `technical decision log`
  - why it matters: `record the rationale for bounded execution scope and no implicit auto-advance`
  - timing: `rollup`
  - status: `identified`

## Milestone Plan

### Milestone 1: Freeze The Split-Gate Wording

- purpose: `make proposal, planning, and coding-start language coherent`
- scope: `prompt and template wording, plan-set index shape, and gate taxonomy`
- dependencies: `proposal P-0001`
- likely artifacts touched: `proposal template, implementation plan template, readiness prompts`
- exit criteria: `all core planning artifacts use the same gate language and shared reason set`
- detailed milestone plan likely needed: `no`

### Milestone 2: Improve Checkpoint Handoff

- purpose: `make execution summaries feed rollup directly`
- scope: `checkpoint execution, review, and rollup structures`
- dependencies: `Milestone 1`
- likely artifacts touched: `checkpoint prompts and process docs`
- exit criteria: `checkpoint summaries carry decisions, doc follow-ups, and unresolved deviations clearly enough for rollup`
- detailed milestone plan likely needed: `maybe`

## Linked Subplans

- product plan linkage: `defines operator-facing language and process teaching changes`
- technical plan linkage: `defines prompt/template contract changes and gate fields`

## Sequencing Notes

- make the durable update obligations visible and timing-tagged before coding
  starts; this example has no `precondition` items, so land the `early` items
  during milestone 1 and the decision-log item in rollup
- milestone completion does not authorize silent scope expansion into the next
  milestone

## Risks And Watchpoints

- the landing page must remain light enough to avoid process bloat
- the shared taxonomy must stay diagnostic rather than bureaucratic

## Related Docs

- `examples/happy-path/roadmaps/plans/P-0001/PLANSET_INDEX.md`
- `examples/happy-path/roadmaps/plans/P-0001/product/PRODUCT_PLAN_P-0001.md`
- `examples/happy-path/roadmaps/plans/P-0001/technical/TECHNICAL_PLAN_P-0001.md`
