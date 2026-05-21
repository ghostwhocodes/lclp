# Layered Context Loop

## Purpose

This document explains the Layered Context Loop, a lightweight manual planning
and delivery process used with an LLM in this repository's workflow.

Templates and prompts for the process live under:

- [README.md](./README.md)
- [FLOW.md](./FLOW.md)

Live generated artifacts for the process live under:

- the installed `workflow/` directory beside `process/`

The process is intentionally lightweight, iterative, and human-directed.
It is designed to:

- keep high-level intent stable
- keep durable technical context close to how coding agents actually work
- avoid duplicated technical-doc artifacts
- separate product context, technical context, change shaping, and execution
- reduce drift by making durable-doc updates explicit in proposals and plans
- keep durable-doc drift visible before coding without forcing every update to
  land first
- keep each change grouped under one slug-specific plan directory
- keep new artifacts justified by gates, ambiguity reduction, or handoff value

## Core Principle

Different layers change in different ways.

- product context is mostly evergreen
- technical design should stay current and practical
- proposals define bounded change contracts
- implementation plans define milestone sequencing
- checkpoint and attempt artifacts are temporary execution aids

The process should respect those differences instead of forcing one document
shape across every layer.

## Artifact Model

### 1. Product Brief

The `Product Brief` is a high-level durable overview of the product.
It should stay broad and change rarely.

### 2. Feature Index

The `Feature Index` is the evergreen record of product capabilities being added
over time.
It should accumulate features gradually without becoming a roadmap.

### 3. Feature Spec

A `Feature Spec` is optional.
Use it only when a feature is large, risky, or important enough that a short
feature-index entry is not sufficient.

### 4. Product Decision Log

Product-level reasoning belongs in the product decision log when the tradeoffs
should survive beyond a single conversation or proposal.

### 5. System Overview

`System Overview vN` is the main durable technical-design reference.

Use it to capture:

- the current durable system shape
- major boundaries and flows
- system invariants
- operating context
- compatibility and migration stance

Version it only when versioning materially helps.

### 6. Subsystem Design

`Subsystem Design` docs hold durable technical detail for one subsystem,
contract surface, or implementation area.

Use them for:

- schemas and state shapes
- APIs, events, commands, and file formats
- naming conventions
- subsystem responsibilities
- low-level technical constraints and guardrails

These docs are evergreen references, not plans.

### 7. Technical Change

A `Technical Change` is an optional bounded durable technical-change contract.

Use it when a change needs:

- explicit guardrails
- frozen decisions
- a compatibility or migration stance
- a durable statement of which technical docs need updating

Do not use it as a milestone plan.

### 8. Technical Decision Log

The `Technical Decision Log` is the single durable decision log for the
technical-design layer.

Use it for both system-level and subsystem-level technical decisions.
Do not split decisions into parallel technical-decision logs.

### 9. Change Proposals

The `workflow/roadmaps/proposals/` layer holds temporary day-to-day change
proposals.

These proposals may cross product and technical-design boundaries, but they
must not replace durable product or technical docs.

Each proposal must have one unique slug.
That slug is the anchor for all planning artifacts related to the change.
Prefer a user-supplied external issue key such as `ABC-123` when one exists.
If no external key exists, generate a fallback slug such as
`P-20260418-153045-1K7F`.
Do not allocate proposal slugs from a shared mutable counter file.

If a proposal affects product or technical design materially, it must say which
separate plan artifacts need to exist under the slug directory.
It should also make the expected durable updates visible before coding starts,
including whether each one is a `precondition`, `early`, or `rollup` update.
Any update tagged `precondition` must land before coding starts.

### 10. Implementation Plans

`Implementation Plan` converts a proposal into milestone sequencing.

It owns:

- milestone boundaries
- dependencies
- exit criteria
- likely touched areas
- validation expectations
- durable update obligations and their timing tags
- the bounded authorized scope for the current plan
- whether milestone completion permits auto-advance

Each proposal should map into a slug directory shaped like:

- `workflow/roadmaps/plans/<slug>/`
- `workflow/roadmaps/plans/<slug>/product/`
- `workflow/roadmaps/plans/<slug>/technical/`

The root slug directory holds the overall implementation plan and any flat
subplans that coordinate the whole change.

The `product/` and `technical/` subdirectories hold separated plan artifacts
and supporting resources for those areas when needed.

`PLANSET_INDEX.md` should act as the human-readable landing page for the slug
plan set.
It is a planning artifact, not the final coding-start approval record.

Proposals do not own milestone ladders.

### 11. Checkpoint Plans And Execution Logs

The `workflow/logs/checkpoints/` layer holds execution-facing artifacts such
as:

- saved internal checkpoint plans
- execution summary logs
- review and fix-forward checkpoint plans
- proposal rollup summaries

These artifacts are temporary and execution-facing, not evergreen context.
For a milestone that planning explicitly marks as too small for a separate
checkpoint plan, execution may run directly against the parent plan milestone,
but the execution and review summaries still belong in this layer.

### 12. Attempt Plans

Attempt plans are temporary execution aids created for a milestone or a missing
subset of a milestone.
They are deliberately short-lived.

## Working Flow

### Product-Design Flow

1. maintain the `Product Brief`
2. record new capabilities in the `Feature Index`
3. create a `Feature Spec` only when a feature needs more detail
4. record durable product reasoning in the product decision log when needed
5. assess whether the feature implies technical-design updates

### Technical-Design Flow

1. keep `System Overview` current enough to guide work
2. create `Subsystem Design` docs where detailed technical reference is useful
3. record durable technical reasoning in the `Technical Decision Log`
4. create a `Technical Change` only when a bounded technical change needs its
   own durable contract

The goal is one coherent technical-doc layer, not a split between abstract
architecture docs and implementation-spec docs.

### Proposal Flow

1. create or refine a `workflow/roadmaps/proposals/` change proposal
2. assign or confirm a unique proposal slug
3. identify whether it needs product plans, technical plans, or both
4. confirm the proposal is complete enough to split into plans
5. split the proposal into a slug-based plan set under
   `workflow/roadmaps/plans/<slug>/`
6. create the root implementation plan plus any required product and technical
   subplans
7. make durable update obligations visible with `precondition`, `early`, or
   `rollup` timing
8. create or refine `PLANSET_INDEX.md`
9. run the separate planning-side readiness check before coding starts

### Delivery Flow

1. decide whether the change needs product, technical-doc, decision-log,
   proposal, or planning updates
2. make the required plan split explicit in the proposal first
3. make expected durable updates visible and timing-tagged before coding starts
4. create the slug-based plan set
5. confirm the slug plan set is complete enough to start coding, including that
   any `precondition` updates are already landed in the referenced durable docs
   or logs
6. only then start coding work
7. execute with checkpoint and review loops as needed, or use the lightweight
   direct milestone path when planning explicitly says no separate checkpoint
   plan is warranted
8. compare final outcome against the plan and proposal for drift checking
9. decide whether the change is complete enough to roll up

## Why The Layers Differ

Product-design docs:

- are broader
- change more slowly
- describe intent and capabilities

Technical-design docs:

- must stay close to the realities coding agents work with
- should be current-state references wherever possible
- should avoid duplication across parallel doc types

Proposal docs:

- are temporary and change-oriented
- are allowed to be volatile
- should name the slug, plan split, and required separated plans rather than
  replacing those durable artifacts
- should make core requirements and durable update obligations visible early

Checkpoint docs:

- are temporary and execution-facing
- preserve execution shaping across context clears
- should leave behind enough decision and follow-up detail for rollup
- should stay within explicitly authorized scope rather than auto-expanding into
  the next milestone

## Core Requirement Escalation

An LLM following this process should ask for user input only when a blocker
affects a core requirement.

A blocker affects a core requirement when it would:

- materially change user-visible behavior
- violate a stated constraint or invariant
- require an irreversible external choice
- force a decision between materially different tradeoff paths

Otherwise, the LLM should make a reasonable decision, document it, and continue.

## Durable Update Timing

Use these timing tags for durable product, technical, and decision-log updates:

- `precondition`: must land before coding starts
- `early`: may land during execution and should usually be handled in early milestones
- `rollup`: may land after implementation work, before proposal closeout
- `not applicable`: no durable update is needed for that layer

## LLM Guidance

An LLM using this process should follow these rules:

- do not rewrite the `Product Brief` unless the product really changed
- prefer one durable technical doc path per topic
- use `System Overview` for durable system shape
- use `Subsystem Design` for durable detailed technical reference
- use `Technical Change` only when a bounded technical change needs a durable
  contract
- use one `Technical Decision Log` rather than parallel architecture/spec logs
- use `workflow/roadmaps/proposals/` as a temporary shaping layer, not as a
  substitute for durable product or technical-doc artifacts
- assign a clear unique slug to each proposal before planning outputs are
  created
- create planning outputs under `workflow/roadmaps/plans/<slug>/`
- create separate product and technical subplans when the proposal requires
  them
- keep `PLANSET_INDEX.md` current enough to act as the slug landing page
- do not treat `PLANSET_INDEX.md` as the final coding-start approval record
- make expected durable updates visible before coding starts even when some are
  planned for `early` or `rollup`; any update tagged `precondition` must land
  before coding starts, and that readiness check must verify the referenced
  durable docs or logs directly
- do not start coding until the planning-side readiness check says the slug
  plan set is complete, linked coherently, has visible durable update
  obligations, and has no unlanded or unverifiable `precondition` updates
- keep milestones out of proposals and in implementation plans
- keep subsystem design docs out of milestone-planning and run-journal mode
- stay within the current authorized scope and do not auto-advance into the
  next milestone unless the plan explicitly allows it
- if planning explicitly says a milestone does not need a checkpoint plan,
  execute directly against the parent milestone but still write execution and
  review summaries under `workflow/logs/checkpoints/`
- add a new artifact only when it changes a gate, reduces ambiguity, or
  materially improves handoff
