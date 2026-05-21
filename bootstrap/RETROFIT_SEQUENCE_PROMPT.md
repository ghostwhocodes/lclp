You are resuming work after a context clear and must retroactively apply the
manual process to an existing repository.

Your job is not to implement a feature immediately.
Your job is to rebuild the process layers in a disciplined order so later work
has clean context.

Follow this retrofit sequence:

1. Build the product-design layer in minimal evergreen form.
2. Build the technical-design layer in minimal current-state form.
3. Retrofit each outstanding pending proposal into the new
   `workflow/roadmaps/proposals/` model.
4. Only then use the planning prompts for the specific proposal currently being
   taken forward.

Retrofit constraints:

- Do not process completed or historical proposals in this flow.
- Use the existing repository, current docs, and current code as the source of
  truth for current state.
- Keep every layer minimal and high-signal.
- Do not pull future desired changes into the evergreen layers unless they are
  already durable truths in the repository.
- Pending proposals can later drive updates upward into the evergreen layers,
  but the first pass should avoid speculative over-modeling.
- At proposal level, handle one pending proposal at a time.
- Use source proposal material from the existing repository as the input set.
- Copy or summarize that source material into `workflow/bootstrap/` when a
  temporary retrofit workspace is useful.
- Ignore implementation-plan files during the retrofit pass unless explicitly
  asked to use them for later planning.

For each retrofit phase:

- ask only the clarifying questions needed for that layer
- create or refine the minimum durable artifacts needed
- avoid over-documenting the whole system at once

When the retrofit sequence is complete, the planning prompts in
`process/roadmaps/plans/` should be used only for the currently active
proposal.
