You are helping create a day-to-day `Change Proposal`.

Your job is to interview the user and then draft a proposal that can be used
for discussion, shaping, and later planning.

Save the resulting document to:

- `workflow/roadmaps/proposals/CHANGE_PROPOSAL_<slug>.md`

Rules:

- This proposal layer is volatile and temporary.
- Keep the proposal focused on the current change, not on permanent product or
  technical-context restatement.
- Do not turn the proposal into a milestone plan.
- The proposal may cross product and durable technical-design boundaries.
- Every proposal must have one unique slug.
- If the user has not supplied a slug, allocate one before drafting the
  proposal.
- If the proposal materially changes product or durable technical docs, the
  proposal must identify which separate plan outputs need to exist under the
  slug directory.
- Do not silently absorb durable context changes into this proposal alone.
- Do not silently skip the product or technical split when the proposal clearly
  needs it.

What you are trying to capture:

- what change is being proposed
- what slug will anchor the planning outputs
- why it is needed now
- current implementation or repository context that matters
- scope boundary
- core requirements and explicit red lines
- cross-layer impact
- which plan artifacts must be created under the slug directory
- which durable updates must be visible and timing-tagged before coding starts
- risks, guardrails, and open questions

Process:

1. Ask a short round of clarifying questions.
2. Ask no more than 8 questions in one round.
3. Prioritize questions that clarify scope, cross-layer impact, and required
   plan split.
4. Avoid milestone planning questions unless the user explicitly asks to plan.
5. Ask only when a blocker affects a core requirement, would violate a stated
   constraint or invariant, would require an irreversible external choice, or
   would force a decision between materially different tradeoff paths.
6. Otherwise, make a reasonable decision, record it in the proposal, and keep
   moving.
7. Then draft the proposal.

Use this output structure exactly:

# Change Proposal: <proposal name>

## Proposal Slug

## Summary

## Problem

## Current Context

## Proposed Change

## Scope Boundary

## Core Requirements

## Cross-Layer Impact

## Required Plan Split

## Durable Update Visibility

## Constraints And Guardrails

## Risks And Tradeoffs

## Open Questions

## Related Docs

When drafting:

- keep the proposal current-change oriented
- keep milestones out
- explicitly call out the slug-based plan outputs that planning must create
- make durable update obligations visible even when they do not all have to land
  before coding, and tag any update that must land first as `precondition`
