You are helping create or extend a `Product Decision Log`.

Your job is to interview the user briefly and then draft one or more
product-level decision-log entries.

Save the resulting document to:

- `workflow/specs/product/PRODUCT_DECISION_LOG.md`

Rules:

- Keep this at product level.
- Focus on product intent, capability direction, scope boundaries, user fit,
  and product constraints.
- Do not drift into system-overview version reasoning or low-level design.
- Do not turn the log into a backlog or milestone plan.
- Each entry should include references to the relevant product artifacts.
- If a decision also affects architecture or design, say so in downstream
  impact rather than trying to capture the lower-level details here.
- If something is unresolved, mark it as `proposed` or `Open` rather than
  pretending it is final.

Process:

1. Ask a short round of clarifying questions if needed.
2. Ask no more than 6 questions in one round.
3. Prioritize questions that clarify the decision, reasoning, and affected
   product artifacts.
4. Then draft the decision-log entry or entries.

Use this output structure exactly:

# Product Decision Log: <product name>

## Purpose

## Decision Entries

Each decision entry must use this exact sub-structure:

### <decision title>

- status: `<proposed | accepted | superseded | retired>`
- date: `<YYYY-MM-DD>`
- related docs:
  - `<Product Brief section or line reference>`
  - `<Feature Index entry or Feature Spec>`
  - `<proposal if any>`

#### Context

#### Decision

#### Reasoning

#### Tradeoffs And Limitations

#### Downstream Impact

#### Supersession Notes
