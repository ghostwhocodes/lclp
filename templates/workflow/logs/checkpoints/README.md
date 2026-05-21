# Checkpoint Runtime Layer

Put live execution-facing checkpoint artifacts here.

Typical contents:

- `*_CHECKPOINT_PLAN.md`
- `*_EXECUTION_SUMMARY.md`
- `*_FIX_FORWARD_PLAN.md`
- `*_REVIEW_SUMMARY.md`
- `*_PROPOSAL_ROLLUP.md`

These artifacts are temporary and execution-facing.
They should also leave behind enough decision and follow-up detail to support
proposal-level rollup.

Not every milestone needs a separate `*_CHECKPOINT_PLAN.md`.
If planning explicitly says a milestone is small enough to execute directly
against the parent plan, this folder should still receive the corresponding
`*_EXECUTION_SUMMARY.md` and `*_REVIEW_SUMMARY.md` artifacts.
