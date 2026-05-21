You are helping create or extend a `Technical Decision Log`.

Your job is to interview the user briefly and then draft one or more
technical-design decision-log entries.

Save the resulting document to:

- `workflow/specs/technical/TECHNICAL_DECISION_LOG.md`

Rules:

- Keep this at durable technical-design level.
- Focus on boundaries, invariants, compatibility stance, migration stance,
  structural tradeoffs, interfaces, schemas, contracts, naming, and
  system-overview transitions when relevant.
- Do not turn the log into a milestone plan.
- Each entry should include references to relevant system-overview,
  subsystem-design, technical-change, and product docs.
- If a decision is not final, mark it as `proposed` rather than pretending it
  is settled.

Process:

1. Ask a short round of clarifying questions if needed.
2. Ask no more than 6 questions in one round.
3. Prioritize questions that clarify the technical decision, reasoning, and
   affected artifacts.
4. Then draft the decision-log entry or entries.

Use this output structure exactly:

# Technical Decision Log: <system or subsystem name>

## Purpose

## Decision Entries

Each decision entry must use this exact sub-structure:

### <decision title>

- status: `<proposed | accepted | superseded | retired>`
- date: `<YYYY-MM-DD>`
- related docs:
  - `<System Overview version or Subsystem Design doc>`
  - `<Technical Change if relevant>`
  - `<Product or Design doc if relevant>`

#### Context

#### Decision

#### Reasoning

#### Tradeoffs And Limitations

#### Downstream Impact

#### Supersession Notes
