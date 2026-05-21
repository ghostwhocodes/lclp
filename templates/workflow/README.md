# Layered Context Loop

This directory is the shipped runtime skeleton template for the Layered Context
Loop process.

When installed into a repository, this directory is the live `workflow/` tree.
In the source LCLP repo, the same content lives under `templates/workflow/` so
it can be copied into target repositories as `workflow/`.

The templates, prompts, and process description live under `process/`.

Path note:

- in an installed target repo, see `process/README.md` and `process/FLOW.md`
- in this source repo, those same files live at `../../process/README.md` and
  `../../process/FLOW.md`

The intended runtime stack is:

1. `bootstrap/`
2. `specs/product/`
3. `specs/technical/`
4. `roadmaps/proposals/`
5. `roadmaps/plans/`
6. `logs/checkpoints/`

Boundary:

- `process/` = process definition, templates, and prompts
- this directory becomes the live artifact tree after install
- if you are reading it inside the source LCLP repo, it is the starter tree
  that installers copy into target repos as `workflow/`

The working model is:

- higher layers hold durable context
- lower layers hold temporary shaping, planning, and execution artifacts
- completed temporary artifacts may later be deleted once their value has been
  absorbed into durable layers and git history
- `roadmaps/plans/<slug>/PLANSET_INDEX.md` is the slug landing page, not the
  final coding-start approval artifact
- `logs/checkpoints/` holds execution and review summaries even when a small
  milestone is executed directly against the parent plan without a separate
  checkpoint plan

For the canonical handoff between proposal shaping, plan splitting, coding
start, checkpoint loops, and rollup, see:

- `process/FLOW.md` in an installed repo
- `../../process/FLOW.md` in this source repo
