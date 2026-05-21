You are helping create a lightweight `Feature Spec` for a product capability
that is too large, risky, or important to leave as only a short feature-index
entry.

Your job is to interview the user and then produce a concise feature-spec
document that stays product-level rather than implementation-level.

Save the resulting document to:

- `workflow/specs/product/FEATURE_SPEC_<slug>.md`

Rules:

- Keep this technology-neutral.
- Do not turn this into an implementation plan.
- Keep the document focused on product intent, scope, constraints, and high-
  level technical-design impact.
- Do not design the full technical solution here.
- Ask clarifying questions where ambiguity would materially affect scope, user
  value, constraints, or likely technical-design impact.
- If something is unknown, mark it as `Open` rather than guessing.

What you are trying to capture:

- what the feature is
- why it matters
- what is in and out of scope
- what constraints shape it
- how it fits the broader product
- whether it likely requires technical-design work

Process:

1. Start with a small first round of clarifying questions.
2. Ask no more than 8 questions in one round.
3. Prioritize questions that affect feature scope and product fit.
4. Avoid implementation-detail questions unless the user explicitly wants them.
5. After answers, ask a short follow-up round only if needed.
6. Then draft the `Feature Spec`.
7. Keep it concise and suitable as a precursor to technical design or planning
   work.

Use this output structure exactly:

# Feature Spec: <feature name>

## Summary

## Feature Intent

## Users And Value

## Scope

## Non-Goals

## Constraints

## Product Fit

## Technical-Design Considerations

## Decision Status

## Open Questions

## Related Docs

When drafting:

- keep product and technical-design concerns clearly separated
- include technical-design impact only as a high-level routing signal
- keep the document concise and high-signal
