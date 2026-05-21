# Planning Runtime Layer

Put live slug-based planning artifacts here.

Each proposal should map to:

```text
workflow/roadmaps/plans/<slug>/
  PLANSET_INDEX.md
  IMPLEMENTATION_PLAN_<slug>.md
  MILESTONE_PLAN_<slug>_M<n>.md
  product/
    PRODUCT_PLAN_<slug>.md
    <user resources if any>
  technical/
    TECHNICAL_PLAN_<slug>.md
    <rules or user resources if any>
```

Rules:

- `PLANSET_INDEX.md` is the landing page for the slug plan set
- the root slug directory holds the overall implementation plan
- `product/` holds separated product plans and product-side resources
- `technical/` holds separated technical plans and technical-side resources
- flat subplans may sit in the root slug directory when they coordinate the
  whole change
