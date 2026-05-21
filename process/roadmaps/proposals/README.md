# Proposal Layer

This folder holds volatile day-to-day change proposals.

These proposals are not evergreen context artifacts.
They are temporary change records used to shape, discuss, and prepare a change
before planning and implementation.

They are closest to:

- RFC-style change discussion
- change-board style working proposals
- bounded change records that may affect multiple higher layers

These proposals may reference:

- current code structure
- specific implementation realities
- local constraints in the current repository state
- concrete scope boundaries for the change being proposed

These proposals should not contain:

- milestone ladders
- attempt-by-attempt execution history
- permanent product intent
- permanent technical-context reference docs that belong in
  `workflow/specs/technical/`

Important rule:

- a proposal in this folder is allowed to cross product and technical-design
  boundaries
- each proposal must have one unique slug
- the proposal must say whether it needs product plans, technical plans, or
  both
- the proposal must split into a slug-based plan set under
  `workflow/roadmaps/plans/<slug>/`
- the planning outputs for that slug should usually include:
  - one plan-set index in `workflow/roadmaps/plans/<slug>/`
  - one root implementation plan in `workflow/roadmaps/plans/<slug>/`
  - product plans or resources in `workflow/roadmaps/plans/<slug>/product/`
    when product changes are needed
  - technical plans or resources in `workflow/roadmaps/plans/<slug>/technical/`
    when technical changes are needed
- proposals do not give the final “ready to start coding” signal
- that signal belongs to a planning-side readiness check after the split plan
  set exists
- a proposal is complete enough to split only when it can drive the required
  slug plan set without guesswork
- add a new planning artifact only when it changes a gate, reduces ambiguity,
  or materially improves handoff

These proposals are temporary.

Once completed, they can be deleted from this folder and left in git history.
The durable records should remain in the higher evergreen layers and in any
completed planning artifacts.

See also:

- [`../../FLOW.md`](../../FLOW.md)
