# Bootstrap Prompt Pack

This folder contains prompts for bootstrapping the manual process into an
existing repository after a context clear or while installing it into a repo
that does not yet have durable process artifacts.

This directory assumes the prompt pack is already installed in the current
repository. It is not the installer itself.

This directory is prompt-only.
Temporary working notes, copied source proposals, and retrofit inventories
belong under `workflow/bootstrap/`.

These prompts are intentionally layer-by-layer.

Use them in this order:

1. `RETROFIT_SEQUENCE_PROMPT.md`
2. `PRODUCT_LAYER_RETROFIT_PROMPT.md`
3. `TECHNICAL_DESIGN_LAYER_RETROFIT_PROMPT.md`
4. `PENDING_PROPOSAL_RETROFIT_PROMPT.md` once per outstanding pending proposal
5. the normal planning prompts in `process/roadmaps/plans/` for the specific
   proposal you are actively taking forward

Important retrofit rules:

- build the evergreen layers first
- keep each layer minimal and high-signal
- document current stable truth first, not desired future changes
- do not process completed or historical proposals in this retrofit flow
- gather source proposal material from the existing repository into
  `workflow/bootstrap/` when that makes the retrofit easier to review
- treat user-supplied `*_PROPOSAL.md` style source proposals as proposal inputs
  and ignore implementation plans unless you are explicitly using them later for
  planning work
- do not over-document the repo in one pass; create enough structure to support
  the pending work cleanly

For this repository, the expected practical flow is:

1. bootstrap `workflow/specs/product/`
2. bootstrap `workflow/specs/technical/` as the durable technical-doc layer
3. retrofit each outstanding pending proposal into
   `workflow/roadmaps/proposals/`
4. plan only the active proposal you want to implement now
