# Product Plan: Plan-Set Index And Coding-Start Gate Cleanup

> Artifact type: `temporary planning artifact`
>
> Curated source-repo example only. In a live installed repo, the equivalent
> artifact would live under `workflow/...` rather than `examples/...`.

## Slug

- slug: `P-0001`
- directory: `examples/fix-forward/roadmaps/plans/P-0001/product/`

## Summary

Update the operator-facing process language so the split between proposal
shaping, plan creation, and coding-start gating is easy to understand.

## Scope

- operator-facing wording
- gate names and teaching language

## Timing And Dependency

- expected landing point: `early`
- timing rationale: `the operator-facing wording should land in the early planning milestone, but it is not itself a separate start gate outside the root plan`

## Requested Product Changes

- teach the proposal gate as a split-planning gate rather than a coding gate
- teach the coding-start gate as a plan-set completeness gate

## Inputs And Resources

- `process/FLOW.md`
- existing proposal and plan readmes

## Deliverables

- updated process copy in the core flow docs
- example artifacts that demonstrate the intended operator-facing language

## Validation

- the flow docs can be read in order without collapsing proposal and
  coding-start gating into one hard gate
- the worked example uses the same operator-facing terms

## Links

- plan-set index: `examples/fix-forward/roadmaps/plans/P-0001/PLANSET_INDEX.md`
- root implementation plan: `examples/fix-forward/roadmaps/plans/P-0001/IMPLEMENTATION_PLAN_P-0001.md`
- proposal: `examples/fix-forward/roadmaps/proposals/CHANGE_PROPOSAL_P-0001.md`
- related product docs: `process/FLOW.md`
