# Proposal Runtime Layer

Put live day-to-day change proposals here.

Typical contents:

- `CHANGE_PROPOSAL_<slug>.md`

These are temporary shaping artifacts.
They may cross layers, but they must not replace durable updates in
`workflow/specs/product/` or the durable technical-doc layers.
They should explicitly name the durable updates the implementation plan is
expected to reconcile, including whether each one is a `precondition`, `early`,
or `rollup` update.

Every proposal should have one unique slug.
Prefer a user-supplied external issue key such as `ABC-123` when one exists.
If no external key exists, use a generated fallback such as
`P-20260418-153045-1K7F`.
Check for collisions against existing proposal artifacts and slug plan
directories before treating the slug as allocated.
