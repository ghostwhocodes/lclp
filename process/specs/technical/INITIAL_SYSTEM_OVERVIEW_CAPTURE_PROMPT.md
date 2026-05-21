You are helping create the first version of a lightweight, technology-aware
but durable system-overview document for a software product.

Your job is to interview the user and then produce `System Overview v1`.

Save the resulting document to:

- `workflow/specs/technical/SYSTEM_OVERVIEW_v1.md`

Rules:

- Keep this durable and implementation-aware, but avoid turning it into a
  low-level subsystem reference.
- Do not choose specific languages, frameworks, clouds, databases, queues, or
  vendors unless the user explicitly wants that or they are already part of the
  stable system context.
- Do not turn this into a milestone plan or implementation sequence.
- This prompt creates the initial system-overview version only.
- Later system states should be represented as later versions of the same
  architecture document type, not as a separate target-document type.
- Ask clarifying questions where ambiguity would materially affect the quality
  of the initial system summary.
- If something is unknown, mark it as `Open` rather than inventing details.

What you are trying to capture:

 - the initial system shape
 - major boundaries or bounded contexts
 - core flows that matter structurally
 - operating context and integration posture
 - non-functional expectations in broad terms
 - compatibility, migration, and shim policy
 - stable areas versus change-sensitive areas

Process:

1. Start with a small first round of clarifying questions.
2. Ask no more than 8 questions in one round.
3. Prioritize questions that most affect the accuracy of the system-overview
   summary or its compatibility stance.
4. Avoid stack-selection questions unless the user explicitly wants them.
5. After answers, ask a short follow-up round only if needed.
6. Then draft `System Overview v1`.
7. Keep it compact and high-signal so humans can actually maintain it.

Use this output structure exactly:

# System Overview v1: <product name>

## Document Status

## Summary

## Current State

## System Invariants

## Boundaries

## Core Flows

## Operating Context

## Quality Attributes

## Compatibility And Migration Stance

## Stable Areas

## Change-Sensitive Areas

## Non-Goals

## Decision Status

## Open Questions

## Related Docs

When drafting:

- prefer short paragraphs and flat bullets
- explicitly separate `Frozen`, `Directional`, and `Open`
- keep the document readable enough to act as a persistent architectural
  summary
- make it clear this is version `v1`
