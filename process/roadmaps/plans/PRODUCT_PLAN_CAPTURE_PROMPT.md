You are helping create or refine the product subplan for one proposal slug.

Your job is to produce the product plan that belongs under:

`workflow/roadmaps/plans/<slug>/product/`

Save the resulting document to:

- `workflow/roadmaps/plans/<slug>/product/PRODUCT_PLAN_<slug>.md`

Rules:

- Use this only when the proposal has product impact.
- Keep it product-facing rather than implementation-facing.
- Do not replace the root implementation plan.
- Use the same timing tags as the root plan: `precondition`, `early`, or
  `rollup`.
- Do not add a separate boolean start gate here; if this work must land before
  coding, express that with `precondition` timing and a short rationale.
- Link this plan back to the root implementation plan.
- Link this plan to the slug plan-set index.
- Capture user-supplied resources here when they belong with product planning.

Use this output structure exactly:

# Product Plan: <proposal name>

## Slug

## Summary

## Scope

## Timing And Dependency

## Requested Product Changes

## Inputs And Resources

## Deliverables

## Validation

## Links
