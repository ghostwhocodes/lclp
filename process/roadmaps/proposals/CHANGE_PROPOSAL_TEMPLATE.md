# Change Proposal: <proposal name>

> Artifact type: `temporary change-shaping artifact`

## Proposal Slug

- slug: `<ABC-123 | P-20260418-153045-1K7F>`
- slug source: `<external issue key | user supplied | generated fallback>`

## Summary

One short paragraph describing the change being proposed and why it matters
now.

## Problem

What concrete problem, limitation, risk, or opportunity is driving this
proposal?

- `<problem>`
- `<problem>`

## Current Context

What current repository, implementation, or workflow context is relevant to
this proposal?

- `<current context>`
- `<current context>`

## Proposed Change

Describe the intended change at proposal level, without turning this into a
milestone plan.

- `<proposal point>`
- `<proposal point>`

## Scope Boundary

What is in scope and out of scope for this proposal?

- in scope:
  - `<in-scope item>`
  - `<in-scope item>`
- out of scope:
  - `<out-of-scope item>`
  - `<out-of-scope item>`

## Core Requirements

State the minimum non-negotiable outcomes and red lines for this proposal.
Keep this section short and high-signal.

- must achieve:
  - `<core outcome>`
  - `<core outcome>`
- must not violate:
  - `<constraint or invariant>`
  - `<constraint or invariant>`
- approval-needed decisions:
  - `<decision that needs explicit user approval if encountered>`
  - `<decision that needs explicit user approval if encountered>`

## Cross-Layer Impact

Which higher layers does this proposal affect?

- product impact: `<none | yes>`
- durable technical-doc impact: `<none | yes>`
- subsystem-design docs likely needed: `<no | yes | maybe>`
- decision-log updates required: `<list or none>`

## Required Plan Split

List the slug-based planning outputs this proposal requires.

Each required plan should live under `workflow/roadmaps/plans/<slug>/`.

- root plan dir: `workflow/roadmaps/plans/<slug>/`
- product plan required: `<yes | no>`
- product plan why: `<short reason | n/a>`
- technical plan required: `<yes | no>`
- technical plan why: `<short reason | n/a>`
- product plan dir: `workflow/roadmaps/plans/<slug>/product/`
- technical plan dir: `workflow/roadmaps/plans/<slug>/technical/`
- expected root plan files:
  - `PLANSET_INDEX.md`
  - `IMPLEMENTATION_PLAN_<slug>.md`
- expected product plan files: `<PRODUCT_PLAN_<slug>.md | none>`
- expected technical plan files: `<TECHNICAL_PLAN_<slug>.md | none>`

## Durable Update Visibility

List the durable product, technical, and decision-log updates this change is
expected to require.
These updates must be visible before coding starts even when they do not all
need to land first. Any update tagged `precondition` must land before coding
starts.
Status here is planning metadata only. A claimed `landed` item does not prove
readiness by itself; the coding-start check must inspect the referenced durable
doc or log path directly.

- product updates:
  - doc: `<path | none>`
  - why it matters: `<short reason>`
  - timing: `<precondition | early | rollup | not applicable>`
  - status: `<identified | drafted | landed | deferred>`
- technical updates:
  - doc: `<path | none>`
  - why it matters: `<short reason>`
  - timing: `<precondition | early | rollup | not applicable>`
  - status: `<identified | drafted | landed | deferred>`
- decision-log updates:
  - doc: `<path | none>`
  - why it matters: `<short reason>`
  - timing: `<precondition | early | rollup | not applicable>`
  - status: `<identified | drafted | landed | deferred>`

## Constraints And Guardrails

- `<constraint>`
- `<constraint>`

## Risks And Tradeoffs

- `<risk or tradeoff>`
- `<risk or tradeoff>`

## Open Questions

- `<question>`
- `<question>`

## Related Docs

- `<product docs>`
- `<technical design docs>`
- `<decision logs>`
