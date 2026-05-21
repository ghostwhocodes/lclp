You are retroactively building the `workflow/specs/technical/` layer for an existing
repository after a context clear.

Your job is to create or refine the minimum current-state technical-design-layer
artifacts needed for this repository under the manual process.

Target artifacts:

- current `System Overview v1` if no durable system-overview document exists yet
- one or more `Subsystem Design` docs only where they materially help
- `Technical Decision Log` if needed

Rules:

- Capture current technical truth first.
- Do not treat pending proposals as if they have already changed the
  system design.
- Keep this layer current-state and durable.
- Do not mix implementation-plan sequencing into this layer.
- Do not replace durable technical docs with proposal text.
- Create only the minimum useful technical context needed for later pending
  proposal work.
- If major technical reasoning or tradeoffs already exist and are worth
  preserving, capture them in the technical decision log.

Use these source signals:

- current repository structure
- current code boundaries and major flows
- schemas, contracts, and interfaces that are already durable
- stable docs
- current implemented task/runtime model
- existing pending proposals only as indicators of pressure areas, not as
  already-landed technical changes

Process:

1. Inspect the current repository and existing docs/code.
2. Ask only the clarifying questions needed to create a minimal current-state
   technical-design layer.
3. Create or refine the technical-design-layer artifacts.
4. Keep the result minimal, current-state, and practical for later coding work.

Success criteria:

- there is a usable current system-overview baseline
- major boundaries, invariants, and subsystem contracts are visible where needed
- durable technical reasoning is captured where needed
- target-state change remains for later proposal work, not baked into the
  current technical baseline
