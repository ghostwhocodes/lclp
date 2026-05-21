# Layered Context Loop Templates

This folder contains lightweight templates and capture prompts for the
Layered Context Loop process used in this repository.

Use this README as the installed entrypoint for the prompt pack.
The other installed companion directories are:

- [../bootstrap/README.md](../bootstrap/README.md)

Live generated artifacts for this process belong under:

- the installed `workflow/` directory beside this `process/` directory

See the short end-to-end flow here:

- [FLOW.md](./FLOW.md)

Single-prompt AI entrypoint:

- [RUN_LCLP_FOR_CHANGE.md](./RUN_LCLP_FOR_CHANGE.md)

## Quick Run Order

Use this as the default operator path when the repository already has enough
durable context to take one change forward:

If you want the AI to route phases automatically instead of following this list
manually, start with `process/RUN_LCLP_FOR_CHANGE.md`.

1. allocate or confirm a slug with
   `process/roadmaps/proposals/ALLOCATE_PROPOSAL_SLUG_PROMPT.md` only if the
   user did not already supply one
2. draft `workflow/roadmaps/proposals/CHANGE_PROPOSAL_<slug>.md` with
   `process/roadmaps/proposals/CHANGE_PROPOSAL_CAPTURE_PROMPT.md`
3. run
   `process/roadmaps/proposals/PROPOSAL_TO_PLANSET_CAPTURE_PROMPT.md` to create
   or refine the slug directory, `PLANSET_INDEX.md`, and the required plan set
4. draft the root implementation plan with
   `process/roadmaps/plans/MULTI_MILESTONE_PLAN_CAPTURE_PROMPT.md`
5. draft `product/PRODUCT_PLAN_<slug>.md` with
   `process/roadmaps/plans/PRODUCT_PLAN_CAPTURE_PROMPT.md` if product impact is
   required
6. draft `technical/TECHNICAL_PLAN_<slug>.md` with
   `process/roadmaps/plans/TECHNICAL_PLAN_CAPTURE_PROMPT.md` if technical
   impact is required
7. run the final start gate with
   `process/roadmaps/plans/CODING_START_READINESS_CHECK_PROMPT.md`
8. optionally expand one milestone with
   `process/roadmaps/plans/MILESTONE_DETAIL_PLAN_CAPTURE_PROMPT.md` when the
   milestone is complex enough to justify it
9. create a checkpoint plan with
   `process/logs/checkpoints/CHECKPOINT_PLAN_CAPTURE_PROMPT.md`, or execute
   directly against the parent milestone only when planning explicitly says no
   separate checkpoint plan is warranted
10. execute with
    `process/logs/checkpoints/EXECUTE_CHECKPOINT_PLAN_PROMPT.md`
11. review and create any fix-forward artifact with
    `process/logs/checkpoints/REVIEW_AND_FIX_FORWARD_PROMPT.md`
12. close out with
    `process/logs/checkpoints/PROPOSAL_ROLLUP_PROMPT.md`

If the repository does not yet have durable product and technical context, use
[../bootstrap/README.md](../bootstrap/README.md) first instead of starting at
the proposal layer.

Prompt ownership at the planning handoff:

- `PROPOSAL_TO_PLANSET_CAPTURE_PROMPT.md` is the split-orchestration prompt; it
  confirms the slug, decides which plans are required, and ensures the slug
  directory and `PLANSET_INDEX.md` exist
- `MULTI_MILESTONE_PLAN_CAPTURE_PROMPT.md` owns the root
  `IMPLEMENTATION_PLAN_<slug>.md`
- `PRODUCT_PLAN_CAPTURE_PROMPT.md` owns
  `product/PRODUCT_PLAN_<slug>.md`
- `TECHNICAL_PLAN_CAPTURE_PROMPT.md` owns
  `technical/TECHNICAL_PLAN_<slug>.md`
- `PLANSET_INDEX.md` is a landing page only; the separate coding-start
  readiness prompt decides whether coding may start

The intended artifact stack is:

1. `Product Brief`
2. `Feature Index`
3. optional `Feature Spec`
4. level-specific `Decision Log`s
5. `System Overview`
6. `Subsystem Design`
7. optional `Technical Change`
8. `Technical Decision Log`
9. day-to-day `Change Proposal`
10. `Plan Set Index`
11. `Implementation Plan`
12. `Checkpoint Plan` and checkpoint execution/review artifacts
13. temporary attempt plans and review prompts

Within `process/`, the core definition is organized under:

- `process/specs/product/`
- `process/specs/technical/`
- `process/roadmaps/proposals/`
- `process/roadmaps/plans/`
- `process/logs/checkpoints/`

The boundary between the main artifact families is:

- `Product Brief`: high-level product intent, users, and operating context
- `Feature Index`: the evergreen record of capabilities being added over time
- `Feature Spec`: optional deeper product-level detail for one feature
- `Decision Log`: durable reasoning, tradeoffs, and downstream-impact records
  at the relevant layer
- `System Overview`: the main durable technical-design surface; keep it current
  and use versioning only when it materially helps
- `Subsystem Design`: a detailed durable technical reference for one subsystem,
  contract surface, or implementation area
- `Technical Change`: an optional bounded durable technical-change contract
- `Technical Decision Log`: the single durable decision log for the
  technical-design layer
- `Change Proposal`: a temporary day-to-day change record used before planning
- `Bootstrap Workspace`: a temporary runtime area used while installing or
  retrofitting the process into an existing repository
- `Plan Set Index`: the landing page for one slug-specific plan set
- `Implementation Plan`: the root plan inside one slug-specific plan directory
- `Milestone Plan`: an optional deeper plan for one specific milestone
- `Checkpoint Plan`: a saved internal Codex execution plan derived from a
  milestone

The preferred technical-doc boundary is:

- use `workflow/specs/technical/` as the durable technical-doc layer
- do not split one topic across parallel technical-doc types that restate the
  same thing
- keep one durable technical doc path per topic whenever possible
- use `workflow/bootstrap/` only as a temporary retrofit/install workspace
- each proposal must have one unique slug
- proposals belong under `workflow/roadmaps/proposals/`
- planning artifacts for a slug live under `workflow/roadmaps/plans/<slug>/`
- if product work is needed, it should be planned under
  `workflow/roadmaps/plans/<slug>/product/`
- if technical work is needed, it should be planned under
  `workflow/roadmaps/plans/<slug>/technical/`
- the root slug directory should hold the overall implementation plan and any
  flat subplans that coordinate the whole change
- `workflow/roadmaps/plans/<slug>/PLANSET_INDEX.md` should be the
  human-readable landing page for the slug plan set

Important rule:

- the product brief should remain high level and stable
- new features should normally be captured in the feature index rather than by
  rewriting the product brief
- durable reasoning and tradeoffs should be captured in the relevant decision
  logs
- `workflow/roadmaps/proposals/` may cross layers, but must not replace durable
  product or technical-doc updates
- a proposal is ready to split into planning only when it has a clear slug and
  a clear plan split
- a proposal should make core requirements and durable update obligations
  visible and timing-tagged before coding starts
- `workflow/roadmaps/plans/<slug>/PLANSET_INDEX.md` is a landing page and
  planning snapshot, not the final coding-start verdict
- multi-milestone plans belong in slug plan sets under
  `workflow/roadmaps/plans/<slug>/`
- milestone-specific plans are optional and subordinate to the parent
  multi-milestone plan
- checkpoint plans belong in `workflow/logs/checkpoints/` and preserve
  execution shaping across runs and context clears
- a small milestone may execute directly against its parent plan only when
  planning explicitly says no separate checkpoint plan is warranted, but it must
  still write execution and review summaries under `workflow/logs/checkpoints/`
- proposals do not own detailed implementation sequencing
- milestone ladders belong in implementation plans, not proposals
- technical changes should capture durable technical changes, constraints, and
  guardrails without forcing a current -> gap -> target model
- coding should start only after the planning-side readiness prompt confirms
  the slug plan set is complete, linked coherently, has visible durable update
  obligations, and has no unlanded or unverifiable `precondition` updates in
  the referenced durable docs or logs
- durable update timing should use `precondition`, `early`, `rollup`, or
  `not applicable`
- milestone completion does not authorize silent auto-advance into the next
  milestone unless the plan explicitly says so
- add a new artifact only when it changes a gate, reduces ambiguity, or
  materially improves handoff

See also:

- [`PROCESS_DESCRIPTION.md`](./PROCESS_DESCRIPTION.md)
- [`FLOW.md`](./FLOW.md)

Worked examples remain source-repo reference material and are intentionally not
part of the installed prompt pack.

Each capture prompt is designed for use with an LLM. The prompts instruct the
model to ask clarifying questions when important details are vague instead of
guessing.
