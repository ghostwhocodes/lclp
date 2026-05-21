# System Overview v1: <product name>

## Document Status

- version: `v1`
- status: `<Draft | Working | Frozen | Superseded>`
- supersedes: `<none>`
- superseded by: `<none>`

## Summary

One short paragraph describing this system-overview version and the system shape
it represents.

## Current State

What exists in this system-overview version that matters structurally?

- `<system fact>`
- `<system fact>`

## System Invariants

Rules that should not be violated within this system-overview version unless
explicitly changed:

- `<invariant>`
- `<invariant>`

## Boundaries

What are the main bounded contexts, subsystems, or responsibility boundaries in
this system-overview version?

- `<boundary>: <responsibility>`
- `<boundary>: <responsibility>`

## Core Flows

What are the most important product or operational flows in this system-overview
version?

- `<flow>`
- `<flow>`

## Operating Context

Broad runtime and operational context, kept technology-neutral:

- deployment posture: `<standalone | managed service | enterprise-managed | hybrid | unknown>`
- operating environments: `<local | on-prem | private cloud | public cloud | mixed | unknown>`
- integration posture: `<isolated | integration-heavy | platform component | unknown>`
- data sensitivity: `<low | moderate | high | mixed | unknown>`

## Quality Attributes

The most important non-functional expectations for this system-overview version:

- scale expectations: `<brief summary>`
- reliability expectations: `<brief summary>`
- security expectations: `<brief summary>`
- operability expectations: `<brief summary>`
- latency/performance expectations: `<brief summary>`
- cost sensitivity: `<brief summary>`

## Compatibility And Migration Stance

How change should be handled relative to this system-overview version:

- compatibility stance: `<strict compatibility | compatibility only when explicitly requested | direct cutover preferred | mixed>`
- migration stance: `<in-place evolution | staged cutover | replacement buildout | unknown>`
- bridge/shim stance: `<allowed | discouraged | forbidden unless explicitly requested>`

## Stable Areas

What parts of this system-overview version should be treated as stable unless
explicitly revisited?

- `<stable area>`
- `<stable area>`

## Change-Sensitive Areas

What parts of this system-overview version are known to be more likely to change,
need close monitoring, or should be treated carefully during planning and
refactoring?

- `<change-sensitive area>`
- `<change-sensitive area>`

## Non-Goals

Technical-design approaches or outcomes that are explicitly not part of this
system-overview version:

- `<non-goal>`
- `<non-goal>`

## Decision Status

- `<decision> - <Frozen | Directional | Open>`
- `<decision> - <Frozen | Directional | Open>`

## Open Questions

- `<question>`
- `<question>`

## Related Docs

- `<product brief>`
- `<subsystem design docs>`
- `<implementation plans>`
- `<technical changes>`
- `<previous system-overview versions if any>`
