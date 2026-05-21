# Worked Example: Happy Path

This example stays under `examples/` in the source repo. It is reference
material only and is not copied into installed target repositories.

This is the primary worked example.
Read this one first.

This example shows one complete process-only artifact set for a change that:

- has one proposal slug
- splits into root, product, and technical plans
- uses a slug landing page
- makes durable update obligations visible with timing tags
- uses checkpoint summaries as rollup inputs
- reaches accepted review and a clean rollup closeout

This is not live repository work.
It is a curated source-repo example to demonstrate the intended shape of the
process.
In a live installation, the equivalent runtime artifacts would live under
`workflow/`.
The artifact bodies below use example-local paths so the example can be read
mechanically in place.
This example shows the checkpointed execution path with a clean completion in
one checkpoint.
In live use, a very small milestone may execute directly against the parent
plan, but it should still produce execution and review summaries under
`workflow/logs/checkpoints/`.

After this one, read [../fix-forward/README.md](../fix-forward/README.md) if
you want to see the recovery path for a partial milestone.

Read it in this order:

1. [`roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md`](./roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md)
2. [`roadmaps/plans/P-0001/PLANSET_INDEX.md`](./roadmaps/plans/P-0001/PLANSET_INDEX.md)
3. [`roadmaps/plans/P-0001/IMPLEMENTATION_PLAN_P-0001.md`](./roadmaps/plans/P-0001/IMPLEMENTATION_PLAN_P-0001.md)
4. [`roadmaps/plans/P-0001/product/PRODUCT_PLAN_P-0001.md`](./roadmaps/plans/P-0001/product/PRODUCT_PLAN_P-0001.md)
5. [`roadmaps/plans/P-0001/technical/TECHNICAL_PLAN_P-0001.md`](./roadmaps/plans/P-0001/technical/TECHNICAL_PLAN_P-0001.md)
6. [`logs/checkpoints/P-0001_M1_CHECKPOINT_PLAN.md`](./logs/checkpoints/P-0001_M1_CHECKPOINT_PLAN.md)
7. [`logs/checkpoints/P-0001_M1_EXECUTION_SUMMARY.md`](./logs/checkpoints/P-0001_M1_EXECUTION_SUMMARY.md)
8. [`logs/checkpoints/P-0001_M1_REVIEW_SUMMARY.md`](./logs/checkpoints/P-0001_M1_REVIEW_SUMMARY.md)
9. [`logs/checkpoints/P-0001_PROPOSAL_ROLLUP.md`](./logs/checkpoints/P-0001_PROPOSAL_ROLLUP.md)
