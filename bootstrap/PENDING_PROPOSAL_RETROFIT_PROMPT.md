You are retroactively processing one outstanding pending proposal into the new
Layered Context Loop structure after a context clear.

Your job is to take one source proposal from the existing repository and map it
into the new system cleanly.

This prompt is for one pending proposal at a time.

Input selection rule:

- use a pending source proposal document, usually a `*_PROPOSAL.md` file
- do not process completed proposals here
- do not treat `*_IMPLEMENTATION_PLAN.md` files as source proposals for this
  retrofit step
- copy or summarize the selected source material into `workflow/bootstrap/`
  when that makes the retrofit easier to inspect

Target outcome:

- a clean `workflow/roadmaps/proposals/` change proposal for the current change
- clearly identified required durable updates in the evergreen higher layers
  (`workflow/specs/product/`, `workflow/specs/technical/`, and decision logs)
- a clear statement of what the first planning milestones should land

Rules:

- Do not jump straight into planning.
- First determine which higher layers this proposal affects.
- If the proposal implies product change, make or request the corresponding
  product-design-layer update.
- If it implies technical-design change, make or request the corresponding
  technical-design-layer update.
- If durable reasoning should be preserved, update the appropriate decision
  log.
- Keep the new `workflow/roadmaps/proposals/` artifact volatile and
  current-change oriented.
- Do not let the day-to-day proposal become the only place where durable
  product or technical-design changes are recorded.
- Planning may begin before those durable docs are edited, but only if the
  proposal names the required durable updates clearly enough to sequence them
  early.

Process:

1. Read the selected pending source proposal.
2. Map its contents into:
   - what belongs in higher evergreen layers
   - what belongs in the temporary `workflow/roadmaps/proposals/` layer
   - what belongs only in a later plan
3. Ask only the clarifying questions needed to cleanly perform that mapping.
4. Create or refine the corresponding proposal-layer and higher-layer updates.
5. State whether the proposal is now ready for the planning prompts in
   `process/roadmaps/plans/`.

Success criteria:

- the source proposal has been decomposed along the new cognitive seams
- durable changes are captured in durable layers
- temporary current-change context remains in `workflow/roadmaps/proposals/`
- planning can proceed once the required durable updates are explicit enough to
  become early milestones
