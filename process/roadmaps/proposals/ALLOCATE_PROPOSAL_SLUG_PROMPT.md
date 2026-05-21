You are allocating a unique slug for one proposal.

Your job is to create or confirm one slug that is good enough to act as the
planning anchor for the proposal.

Rules:

- The slug does not need cryptographic uniqueness.
- It does need to be unique enough within this repository's workflow.
- Prefer a user-supplied external issue key such as `ABC-123` when one exists.
- If the user already supplied a slug, keep it unless it clearly collides.
- If there is no external issue key, generate a fallback slug.
- Keep the slug short enough to use as a planning anchor and readable enough
  for humans.
- The default generated fallback format is `P-YYYYMMDD-HHMMSS-XXXX` in UTC.
- `XXXX` should be a short uppercase alphanumeric suffix.
- Do not allocate slugs from a shared mutable counter file.
- Check for collisions in both:
  - `workflow/roadmaps/proposals/`
  - `workflow/roadmaps/plans/<slug>/`

Process:

1. Check whether the user already supplied a slug or an external issue key.
2. If an external issue key exists, use it as the preferred slug candidate.
3. If there is no external issue key, generate a fallback slug in the default
   format.
4. Check whether the candidate collides with an existing proposal artifact or
   slug plan directory.
5. If a user-supplied or external-key slug collides, stop and ask whether this
   is meant to continue the existing proposal or use a generated fallback slug.
6. If a generated fallback collides, generate a new suffix and re-check.
7. Return the selected slug and the basis for it.

Output format:

## Slug Allocation

- slug: `<slug>`
- source: `<external issue key | user supplied | generated fallback>`
- collision check:
  - `workflow/roadmaps/proposals/`
  - `workflow/roadmaps/plans/<slug>/`

## Notes

- `<note>`
- `<note>`
