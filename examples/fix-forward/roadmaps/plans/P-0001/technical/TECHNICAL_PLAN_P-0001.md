# Technical Plan: Plan-Set Index And Coding-Start Gate Cleanup

> Artifact type: `temporary planning artifact`
>
> Curated source-repo example only. In a live installed repo, the equivalent
> artifact would live under `workflow/...` rather than `examples/...`.

## Slug

- slug: `P-0001`
- directory: `examples/fix-forward/roadmaps/plans/P-0001/technical/`

## Summary

Update prompt and template contracts so the plan-set index, gate taxonomy, and
checkpoint-to-rollup fields are explicit and reusable.

## Scope

- prompt output structures
- plan and checkpoint templates
- gate-reason fields

## Timing And Dependency

- expected landing point: `early`
- timing rationale: `the prompt and template contract changes should land early, with start readiness governed by the root plan's durable update obligations`

## Requested Technical Changes

- add `PLANSET_INDEX.md` to the root slug plan set
- standardize shared block reasons across proposal split, plan creation, and
  coding-start gating
- add checkpoint summary fields that rollup can consume directly

## Inputs, Rules, And Resources

- keep artifact additions lightweight
- prefer one human-readable landing page over a machine state file

## Deliverables

- updated prompt structures
- updated planning templates
- worked example artifacts with consistent fields

## Validation

- all linked artifacts use the same gate statuses and reason names
- checkpoint summaries and rollups reference the same decision and follow-up
  categories

## Links

- plan-set index: `examples/fix-forward/roadmaps/plans/P-0001/PLANSET_INDEX.md`
- root implementation plan: `examples/fix-forward/roadmaps/plans/P-0001/IMPLEMENTATION_PLAN_P-0001.md`
- proposal: `examples/fix-forward/roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md`
- related technical docs: `process/FLOW.md`
