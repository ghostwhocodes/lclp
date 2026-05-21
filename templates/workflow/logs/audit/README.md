# Audit Runtime Layer

Put optional slug-scoped audit logs here when using `lclp-track`.

Typical contents:

- `<slug>.audit.log.jsonl`
- `.locks/<slug>.lock`
- `.gitignore` to keep transient lockfiles out of normal git status

Rules:

- keep one append-only JSONL file per slug
- treat `.locks/` as transient runtime state owned by the tracker
- use canonical uppercase slugs such as `P-0001`
- use `replace` as a corrective event that supersedes an earlier entry instead
  of rewriting history, while preserving or overriding the logical event action
- `get` returns the live leaf state by default; use `get --all` for raw history
- keep the audit log lightweight; it is an operator aid, not a replacement for
  the process artifacts under `workflow/roadmaps/` or `workflow/logs/checkpoints/`
