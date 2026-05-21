# Tracker Minimalism Case Study

This note lives under `docs/` in the source repo. It is reference material
only; the install-facing tracker artifact is the optional source under
`cmd/lclp-track/`, and the tracker writes runtime state only into an installed
`workflow/` tree.

This note explains the shape of `lclp-track` and why it is an interesting
example of intentionally minimal audit logging plus unexpectedly non-trivial
locking.

The tool was not designed as a large subsystem. It exists to answer a small
operational need:

- record a lightweight append-only audit trail for one slug
- do it without a database
- do it without a schema migration story
- do it without cross-referenced tables, joins, or background services

That small scope is precisely what makes it a useful case study.

## What The Tracker Actually Is

`lclp-track` is a tiny CLI in [`cmd/lclp-track/`](../cmd/lclp-track/main.go).

It writes one JSONL log per slug under:

- `workflow/logs/audit/<slug>.audit.log.jsonl`

It also uses one lock file per slug under:

- `workflow/logs/audit/.locks/<slug>.lock`

The runtime model is deliberately simple:

- one slug, one append-only log
- one command invocation, one new record
- no in-place mutation of existing log entries
- no global database
- no event bus
- no secondary indexes

## Why The Audit Model Is Minimal

The tracker does not try to model the whole project.

It only records small operational facts such as:

- slug
- phase
- action
- status
- optional prompt path
- optional artifact paths
- optional free-form note
- optional `replaces` pointer to an earlier event

That means the tool avoids a large amount of complexity that more traditional
audit systems usually acquire:

- no normalized tables
- no foreign-key graph
- no state machine persisted in multiple places
- no transactional coordination between audit rows and process artifacts
- no history rewriting

Instead, the tracker treats the durable process artifacts under `workflow/`
as the real source material and the audit log as a thin operator-facing record
of what happened around them.

## Why `replace` Exists

The tracker still needed a way to correct mistakes without rewriting history.

That is what `replace` does.

It does not mutate an earlier line.
It appends a new event that supersedes an earlier event by ID.

That keeps the core model simple:

- the file stays append-only
- history remains visible
- corrections remain explicit
- the effective live state can be derived from the append-only history

This is a good example of minimalism done carefully.
The design avoids direct mutation, but still provides a practical way to fix
operator mistakes.

## Why Locking Turned Out To Matter More Than Expected

At first glance, this looks like a trivial local file append problem.

It is not.

The moment you need the tool to behave correctly across parallel development
work, you need coordination around a shared slug-scoped file.

Without locking, parallel commands can interfere with each other in several
ways:

- two writers can interleave appends
- one reader can observe a partially written or racing update
- two correction operations can resolve history from the same stale view
- multiple working paths to the same repo can accidentally treat the same log
  as different targets unless paths are canonicalized

For a tiny tool, that is already enough to make locking a real design concern.

## The Locking Solution

On Linux, the tracker uses `flock` via [`lock_linux.go`](../cmd/lclp-track/lock_linux.go).

Important properties:

- locking is per slug, not global
- the lock file lives inside the repo's runtime tree
- the lock path is repo-scoped and canonicalized
- readers take shared locks
- writers take exclusive locks
- the lock file is hidden runtime state under `.locks/`

That last point matters.
The visible audit log remains the durable append-only record.
The lock file is purely coordination state.

## Why Canonical Paths Matter

One subtle issue is that locking only works if all participants are really
locking the same file.

If one invocation points at the repo through one path and another invocation
points at the same repo through a symlink or alternate path, naive locking can
silently become two unrelated locks.

That is why the tracker resolves canonical paths and insists on a real install
root containing `workflow/`.

Without that rule, a typo or alias can produce the worst kind of failure:

- the command appears to succeed
- but it writes audit state into the wrong tree
- while the real install remains untouched

For a tiny audit tool, that is a serious integrity problem.

## Why The Tool Is Linux-Only

The tracker builds everywhere, but functional locking support is Linux-only in
the shipped design.

That is not an accident.
It is a deliberate choice to keep the implementation small and predictable
instead of pretending to support a cross-platform locking story that the
project does not actually want to maintain.

This is another form of minimalism:

- narrow the supported operating window
- make unsupported cases explicit
- do not hide a platform gap behind fake success

## What This Tool Deliberately Does Not Do

The tracker is intentionally not:

- a workflow engine
- a task database
- a general event store
- a concurrent multi-host distributed log
- a replacement for the real `workflow/roadmaps/` and checkpoint artifacts

It is just enough audit structure to help operators understand what was done,
in what phase, and what later corrective event superseded an earlier one.

## Why This Is A Useful Case Study

This small tool highlights two useful engineering lessons.

First: minimal data models can be powerful.

An append-only JSONL log with explicit correction events can go a long way
without introducing a database or a large operational footprint.

Second: even very small tools can hit hard boundary problems.

Here, the hard part was not the record format.
It was coordination:

- which install root is real
- whether different paths refer to the same runtime tree
- whether concurrent access is safe
- how to keep transient lock state separate from durable audit state

That combination makes `lclp-track` an accidental but useful example of
practical minimalism:

- keep the model tiny
- keep the audit trail immutable
- accept a narrow support window
- solve the locking problem honestly rather than pretending it does not exist
