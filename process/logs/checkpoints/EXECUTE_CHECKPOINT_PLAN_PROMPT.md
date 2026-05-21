You are executing or resuming one saved checkpoint plan, or executing directly
against one parent-plan milestone when planning explicitly said no separate
checkpoint plan is warranted.

The caller will invoke you in this style:

`please start or continue with: <checkpoint plan name>`

or:

`please start or continue with: <parent plan name> / Milestone <n>`

Assume:

- you may be restarting after a context clear
- there may be uncommitted files in the working tree
- do not assume existing partial work is yours, intentional for this run, or in
  scope until the checkpoint plan or parent plan plus repository state make that
  clear

Your job is to:

1. load the named checkpoint plan from `workflow/logs/checkpoints/`, or load
   the referenced parent plan milestone when planning explicitly said no
   separate checkpoint plan is warranted
2. inspect the current repo state and working tree
3. classify existing changes as one of:
   - clearly in-scope and safe to reuse
   - unrelated and must be left untouched
   - ownership-unclear or scope-unclear and therefore blocking
4. determine what remains under the current authorized scope
5. execute or continue the work
6. if the checkpoint or direct lightweight milestone run is complete or
   materially progressed, write a concise
   execution summary log to `workflow/logs/checkpoints/`

Execution rules:

- Stay inside the checkpoint's authorized scope.
- For the lightweight direct path, stay inside the referenced parent plan and
  milestone scope exactly as if a separate checkpoint plan had existed.
- Existing unrelated changes must be left untouched.
- Existing in-scope changes may be reused only when repository state makes them
  clearly compatible with the current authorized scope.
- If an in-scope file has ownership-unclear or scope-unclear existing changes,
  stop and report the blocker instead of building on top of them.
- Do not silently continue into the next milestone or adjacent work once the
  current checkpoint scope is complete.
- If the current checkpoint is done and the plan does not explicitly allow
  auto-advance, stop and report completion.
- Ask for approval only when a blocker affects a core requirement, would
  violate a stated constraint or invariant, would require an irreversible
  external choice, or would force a decision between materially different
  tradeoff paths.
- Otherwise, make a reasonable decision, document it in the execution summary,
  and continue.

Use this execution summary structure:

# Execution Summary: <checkpoint plan name | parent plan name / Milestone <n>>

## Outcome

- status: `<completed | partial | blocked>`
- scope result: `<completed within authorized scope | partial within authorized scope | blocked>`

## Working Tree Assessment

- reused in-scope existing changes: `<none | files or short note>`
- unrelated existing changes left untouched: `<none | files or short note>`
- ownership-unclear or blocking existing changes: `<none | files or short note>`

## What Landed

- `<landed item>`
- `<landed item>`

## Decisions Made

- `<decision>`
- `<decision>`

## What Differed From Plan

- `<difference>`
- `<difference>`

## Docs Likely Needing Update

- `<doc>`
- `<doc>`

## Unresolved Deviations

- `<deviation>`
- `<deviation>`

## Remaining Work

- `<remaining item>`
- `<remaining item>`

## Notes

- `<note>`
- `<note about whether auto-advance was allowed or intentionally not used>`
