You are retroactively building the `workflow/specs/product/` layer for an existing
repository after a context clear.

Your job is to create or refine the minimum evergreen product-design-layer artifacts
needed for this repository under the manual process.

Target artifacts:

- `Product Brief`
- `Feature Index`
- `Product Decision Log` if needed

Rules:

- Keep this layer high level and stable.
- Do not turn this into a backlog, roadmap, or implementation plan.
- Use current repository truth and stable existing docs as input.
- Do not pull detailed technical-design content up into this
  layer.
- Add only the minimum product-level structure needed to support later
  proposal work.
- If the repository already implies product capabilities, capture them lightly
  in the feature index.
- If major reasoning/tradeoff decisions already exist and are worth preserving,
  capture them in the product decision log.
- Do not overfit this layer to one pending proposal.

Use these source signals:

- repository README/user docs
- current repo structure and naming
- existing pending proposals only as hints for product capabilities that are
  already clearly part of the product direction
- current implemented truth, not speculative target-state architecture

Process:

1. Inspect the current repository and existing docs.
2. Ask only the clarifying questions needed to create a minimal product-design layer.
3. Create or refine the product-design-layer artifacts.
4. Keep the output intentionally lightweight and durable.

Success criteria:

- the product brief is high-level and changes rarely
- the feature index exists as the evergreen capability ledger
- any durable product-level decisions worth preserving are recorded
- nothing in this layer reads like a plan or low-level technical spec
