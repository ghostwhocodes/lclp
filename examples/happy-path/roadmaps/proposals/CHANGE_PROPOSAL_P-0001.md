# Change Proposal: Plan-Set Index And Coding-Start Gate Cleanup

> Artifact type: `temporary change-shaping artifact`
>
> Curated source-repo example only. In a live installed repo, the equivalent
> artifact would live under `workflow/...` rather than `examples/...`.

## Proposal Slug

- slug: `P-0001`
- slug source: `example only`

## Summary

Refine the planning handoff so every proposal produces a clearer slug plan set,
with one landing page and a distinct coding-start gate.

## Problem

- Proposal shaping and coding-start gating have been too easy to confuse.
- Slug plan directories need one obvious landing page for humans.

## Current Context

- Proposals already own slugs and plan splits.
- Root, product, and technical plans already exist as separate artifact types.

## Proposed Change

- Add `PLANSET_INDEX.md` as the human-readable landing page for each slug.
- Make coding-start gating inspect the linked plan set instead of the proposal
  alone.
- Make plan timing explicit for evergreen product, technical, and decision-log
  updates.

## Scope Boundary

- in scope:
  - plan-set landing page
  - clearer gate wording
  - plan timing fields
- out of scope:
  - automated plan linting
  - code-generation helpers

## Core Requirements

- must achieve:
  - make durable update obligations visible and timing-tagged before coding
    starts
  - preserve one agent-driven change flow across proposal, planning, execution,
    and review
- must not violate:
  - do not force every durable update to land before coding
  - do not encourage silent scope creep once one milestone is complete
- approval-needed decisions:
  - introducing irreversible external workflow or storage dependencies
  - expanding the change beyond planning and checkpoint handoff semantics

## Cross-Layer Impact

- product impact: `yes`
- durable technical-doc impact: `yes`
- subsystem-design docs likely needed: `no`
- decision-log updates required: `product and technical decision logs`

## Required Plan Split

- root plan dir: `examples/happy-path/roadmaps/plans/P-0001/`
- product plan required: `yes`
- product plan why: `operator-facing workflow language changes`
- technical plan required: `yes`
- technical plan why: `prompt and template contract changes`
- product plan dir: `examples/happy-path/roadmaps/plans/P-0001/product/`
- technical plan dir: `examples/happy-path/roadmaps/plans/P-0001/technical/`
- expected root plan files:
  - `PLANSET_INDEX.md`
  - `IMPLEMENTATION_PLAN_P-0001.md`
- expected product plan files:
  - `PRODUCT_PLAN_P-0001.md`
- expected technical plan files:
  - `TECHNICAL_PLAN_P-0001.md`

## Durable Update Visibility

- product updates:
  - doc: `process/FLOW.md`
  - why it matters: `operators need a visible split between proposal clarity and coding readiness`
  - timing: `early`
  - status: `identified`
- technical updates:
  - doc: `process/roadmaps/plans/IMPLEMENTATION_PLAN_TEMPLATE.md`
  - why it matters: `the root plan must carry durable update obligations and scope boundaries`
  - timing: `early`
  - status: `identified`
- decision-log updates:
  - doc: `technical decision log`
  - why it matters: `record why visibility-before-coding replaced docs-first gating`
  - timing: `rollup`
  - status: `identified`

## Constraints And Guardrails

- keep the process lightweight
- do not collapse evergreen docs into the slug plan set

## Risks And Tradeoffs

- more explicit gating adds one more visible artifact
- the landing page is only worth keeping if it reduces handoff ambiguity

## Open Questions

- none in this example

## Related Docs

- `process/FLOW.md`
- `examples/happy-path/roadmaps/plans/P-0001/PLANSET_INDEX.md`
