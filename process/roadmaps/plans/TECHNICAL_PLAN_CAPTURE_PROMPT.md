You are helping create or refine the technical subplan for one proposal slug.

Your job is to produce the technical plan that belongs under:

`workflow/roadmaps/plans/<slug>/technical/`

Save the resulting document to:

- `workflow/roadmaps/plans/<slug>/technical/TECHNICAL_PLAN_<slug>.md`

Rules:

- Use this only when the proposal has technical impact.
- Keep it technical-design facing rather than milestone-log shaped.
- Do not replace the root implementation plan.
- Use the same timing tags as the root plan: `precondition`, `early`, or
  `rollup`.
- Do not add a separate boolean start gate here; if this work must land before
  coding, express that with `precondition` timing and a short rationale.
- Link this plan back to the root implementation plan.
- Link this plan to the slug plan-set index.
- Capture user-supplied rules, screenshots, codebase constraints, or other
  technical planning inputs here when they belong with technical planning.

Use this output structure exactly:

# Technical Plan: <proposal name>

## Slug

## Summary

## Scope

## Timing And Dependency

## Requested Technical Changes

## Inputs, Rules, And Resources

## Deliverables

## Validation

## Links
