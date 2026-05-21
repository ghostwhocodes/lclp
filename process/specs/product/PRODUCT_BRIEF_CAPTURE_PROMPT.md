You are helping create a lightweight `Product Brief` for a software product.

Your job is to interview the user and then produce a concise, high-signal
document using the supplied template structure.

Save the resulting document to:

- `workflow/specs/product/PRODUCT_BRIEF.md`

Rules:

- Keep this technology-neutral.
- Do not choose languages, frameworks, clouds, databases, or implementation
  stacks.
- Do not turn this into a feature backlog or implementation plan.
- Keep this brief high level and durable.
- Do not stuff detailed or fast-changing feature information into this
  document; that belongs in a feature index or feature spec.
- Do not ask dozens of low-value product-marketing questions.
- Ask only the highest-leverage clarifying questions needed to remove major
  ambiguity.
- If something is genuinely unknown, mark it as `Open` instead of guessing.
- Prefer broad product shape and constraints over detailed feature lists.

What you are trying to capture:

- what the product is
- who it is for
- who operates it
- what broad need or opportunity it addresses
- whether it is SaaS, self-hosted, local-first, enterprise-managed, or hybrid
- broad domain, regulatory, privacy, residency, or procurement constraints
- what success looks like
- what it is explicitly not trying to do
- what broad capability themes it is likely to accumulate over time
- what is frozen versus still open

Process:

1. Start by asking a small first round of clarifying questions.
2. Ask no more than 7 questions in one round.
3. Prioritize questions that change the shape of the product brief materially.
4. If the user already provided enough detail for a section, do not ask again.
5. After the user answers, ask a short second round only if needed.
6. Once there is enough information, draft the full `Product Brief`.
7. Keep the brief short, readable, and suitable as a front-door context doc
   for both humans and coding agents.

Use this output structure exactly:

# Product Brief: <product name>

## Summary

## Users And Operators

## Product Intent

## Product Shape

## Domain Constraints

## Success Shape

## Non-Goals

## Stage And Scope

## Capability Themes

## Decision Status

## Open Questions

## Related Docs

When drafting:

- separate `Frozen`, `Directional`, and `Open` decisions clearly
- use short paragraphs and flat bullets
- prefer clarity over completeness
- optimize for a brief that should rarely need major rewriting
- leave unknowns explicit instead of filling gaps with assumptions
