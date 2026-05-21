# Checkpoint Plan: P-0001 / Milestone 1 - Freeze The Split-Gate Wording

> Curated source-repo example only. In a live installed repo, the equivalent
> artifact would live under `workflow/...` rather than `examples/...`.

## Summary

Tighten the gate language and root plan structure without adding workflow
engine-style complexity.

## Parent Plan Context

- parent root plan: `examples/fix-forward/roadmaps/plans/P-0001/IMPLEMENTATION_PLAN_P-0001.md`
- focus milestone: `Milestone 1`

## Authorized Scope

- this checkpoint is authorized to complete: `make milestone 1 wording and plan-structure changes only`
- auto-advance after completion: `no`

## Scope Boundary

- in scope: `plan-set index, gate wording, shared reason set`
- out of scope: `automation or lint enforcement`

## Workstreams

### Workstream 1: Add The Slug Landing Page

- purpose: `make the slug plan set easier to inspect`
- scope: `PLANSET_INDEX.md shape and references`
- dependencies: `proposal P-0001`
- completion signal: `the root slug directory has one obvious landing page`

### Workstream 2: Align Gate Language

- purpose: `keep proposal, planning, and coding-start gates distinct`
- scope: `core flow docs and planning prompts`
- dependencies: `Workstream 1`
- completion signal: `the docs use one coherent gate vocabulary`

## Validation Focus

- the landing page must reduce ambiguity rather than add bureaucracy
- the gate language must not reintroduce a hard pre-planning durable-doc gate

## Risks And Watchpoints

- avoid adding fields that no gate or handoff actually needs

## Expected Summary Outputs

- execution summary log: `P-0001_M1_EXECUTION_SUMMARY.md`
- review summary log: `P-0001_M1_REVIEW_SUMMARY.md`
- proposal rollup log: `P-0001_PROPOSAL_ROLLUP.md`
- expected rollup inputs:
  - `decisions made`
  - `docs likely needing update`
  - `unresolved deviations`

## Related Docs

- `examples/fix-forward/roadmaps/plans/P-0001/PLANSET_INDEX.md`
- `examples/fix-forward/roadmaps/plans/P-0001/IMPLEMENTATION_PLAN_P-0001.md`
