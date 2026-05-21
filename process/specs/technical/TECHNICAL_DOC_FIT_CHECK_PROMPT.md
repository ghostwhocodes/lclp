You are reviewing whether a document is suitable to live in the
`workflow/specs/technical/` layer as an evergreen technical-design document.

Your job is to inspect one document, or a set of documents in the technical-
design folder, and determine whether they are written appropriately for this
layer.

This is a review task, not a rewrite task unless the user explicitly asks for a
rewrite afterward.

Rules:

- Evaluate the document as a long-lived technical context artifact.
- A good technical-design doc is more concrete than product docs and more
  evergreen than a plan.
- It may contain concrete technical detail, including schemas, messages,
  contracts, naming conventions, and subsystem specifics.
- It should not contain milestone ladders, attempt history, temporary progress
  notes, or broad product positioning.
- It should not duplicate proposal sequencing or execution-log content.
- It should not simply narrate the current code path line-by-line.
- It should explain intended technical structure and current relevant
  technical reality in a way that remains useful after individual
  implementation runs.

Review against these checks:

1. Is the document evergreen enough to belong in the technical-design layer?
2. Does it avoid milestone, attempt, and execution-log content?
3. Does it avoid duplicating product-level intent that belongs in
   `workflow/specs/product/`?
4. Does it fit one of the main technical-design shapes such as `System
   Overview`, `Subsystem Design`, `Technical Change`, or `Technical Decision
   Log`?
5. Does it contain concrete technical detail that is useful to
   future implementation work?
6. Is it written as technical guidance and durable reference rather than as a
   plan or a code walkthrough?
7. If it contains transitional `v1`/`v2` detail, is that still justified by
   the current codebase state?

Output format:

## Verdict

- overall fit: `<good fit | partial fit | wrong layer>`
- recommendation: `<keep | revise | move | split>`

## Findings

- `<finding>`
- `<finding>`

## Suggested Actions

- `<action>`
- `<action>`

## Layer Decision

- if it belongs in `workflow/specs/technical/`, explain why
- if it should move to `workflow/specs/product/` or a plan artifact,
  explain why

When writing findings:

- be specific
- focus on layer fit, drift, and evergreen usefulness
- prefer actionable guidance over generic criticism
