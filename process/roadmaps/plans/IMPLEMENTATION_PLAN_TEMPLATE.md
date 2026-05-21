# Implementation Plan: <proposal name>

> Artifact type: `temporary planning artifact`

## Slug

- slug: `<slug>`
- root dir: `workflow/roadmaps/plans/<slug>/`

## Plan Paths

- plan-set index: `workflow/roadmaps/plans/<slug>/PLANSET_INDEX.md`
- root plan: `workflow/roadmaps/plans/<slug>/IMPLEMENTATION_PLAN_<slug>.md`
- product plan: `<path | not needed>`
- technical plan: `<path | not needed>`

## Summary

## Proposal Source

- proposal path: `<workflow/roadmaps/proposals/...>`

## Planning Inputs

- `<input>`
- `<input>`

## Planning Gate

- status: `<ready | blocked | partial>`
- block reasons: `<none | reason list>`
- note: `<short note>`

## Authorized Scope

- this plan is authorized to complete: `<describe the bounded scope>`
- auto-advance to the next milestone when current scope is done: `<no by default | yes if explicitly approved>`
- stop and ask for approval only when: `<approval-needed decision categories>`

## Durable Update Obligations

List the durable product, technical, and decision-log updates this change is
expected to drive.
Visibility is required before coding starts.
Landing timing depends on the tagged obligation.
Status here is planning metadata only. A claimed `landed` item does not prove
readiness by itself; the coding-start check must inspect the referenced durable
doc or log path directly.

- product docs:
  - doc: `<path | none>`
  - why it matters: `<short reason>`
  - timing: `<precondition | early | rollup | not applicable>`
  - status: `<identified | drafted | landed | deferred>`
- technical docs:
  - doc: `<path | none>`
  - why it matters: `<short reason>`
  - timing: `<precondition | early | rollup | not applicable>`
  - status: `<identified | drafted | landed | deferred>`
- decision logs:
  - doc: `<path | none>`
  - why it matters: `<short reason>`
  - timing: `<precondition | early | rollup | not applicable>`
  - status: `<identified | drafted | landed | deferred>`

## Milestone Plan

### Milestone <n>: <title>

- purpose: `<why this milestone exists>`
- scope: `<what it covers>`
- dependencies: `<prior milestones or prerequisites>`
- likely artifacts touched: `<artifact types or areas>`
- exit criteria: `<what must be true to consider it done>`
- detailed milestone plan likely needed: `<yes | no | maybe>`

## Linked Subplans

- product plan linkage: `<summary>`
- technical plan linkage: `<summary>`

## Sequencing Notes

- make clear which durable update obligations are `precondition`, `early`, or
  `rollup`
- milestone completion does not authorize silent scope expansion into the next
  milestone

## Risks And Watchpoints

## Related Docs
