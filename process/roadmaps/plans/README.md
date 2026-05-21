# Planning Layer

This folder holds planning artifacts derived from `workflow/roadmaps/proposals/` change
proposals.

The planning layer is where a temporary day-to-day proposal is turned into a
milestone-based delivery plan.

Important boundaries:

- plans own milestone sequencing
- proposals do not own milestone sequencing
- every plan set belongs to one unique proposal slug
- milestones should be shaped in the way the human and agent judge most useful
- some milestones may be broad enough to justify their own detailed milestone
  plan
- some milestones may remain lightweight and not need a separate milestone
  plan

Typical pattern:

1. start from a `workflow/roadmaps/proposals/` change proposal
2. confirm the proposal slug
3. create the slug directory in `workflow/roadmaps/plans/<slug>/`
4. create or refine `PLANSET_INDEX.md` in that slug directory
5. create the root implementation plan in that slug directory
6. create `product/` and `technical/` subplans if the proposal requires them
7. run the planning-side readiness check before starting coding
8. execute milestones one by one
9. optionally create a more detailed milestone-specific plan when a milestone
   is complex enough to benefit from one

Prompt roles:

- `MULTI_MILESTONE_PLAN_CAPTURE_PROMPT.md` owns the root
  `IMPLEMENTATION_PLAN_<slug>.md`
- `PRODUCT_PLAN_CAPTURE_PROMPT.md` owns
  `product/PRODUCT_PLAN_<slug>.md`
- `TECHNICAL_PLAN_CAPTURE_PROMPT.md` owns
  `technical/TECHNICAL_PLAN_<slug>.md`
- `MILESTONE_DETAIL_PLAN_CAPTURE_PROMPT.md` owns optional flat milestone
  subplans in the slug root
- `CODING_START_READINESS_CHECK_PROMPT.md` is the only final start gate before
  coding

The slug plan set should usually include:

- `workflow/roadmaps/plans/<slug>/PLANSET_INDEX.md`
- `workflow/roadmaps/plans/<slug>/IMPLEMENTATION_PLAN_<slug>.md`
- `workflow/roadmaps/plans/<slug>/product/PRODUCT_PLAN_<slug>.md` if product
  changes are needed
- `workflow/roadmaps/plans/<slug>/technical/TECHNICAL_PLAN_<slug>.md` if
  technical changes are needed
- optional flat subplans in the slug root when one part of the work needs a
  separate execution plan
- `product/` and `technical/` hold planning artifacts for the slug, not the
  evergreen durable product or technical docs themselves

Example layout:

```text
workflow/roadmaps/plans/<slug>/
  PLANSET_INDEX.md
  IMPLEMENTATION_PLAN_<slug>.md
  MILESTONE_PLAN_<slug>_M1.md
  product/
    PRODUCT_PLAN_<slug>.md
    <user resources if any>
  technical/
    TECHNICAL_PLAN_<slug>.md
    <rules or user resources if any>
```

The milestone-specific plan should:

- refine one milestone only
- stay subordinate to the parent multi-milestone plan
- be optional rather than mandatory

See also:

- [`../../FLOW.md`](../../FLOW.md)
