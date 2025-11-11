# Workflows

`ctx` centers on two repeatable phases that keep your AI-ready skeletons fresh.

```
┌─────────────┐     ┌─────────────┐
│  Phase 1    │     │   Phase 2   │
│  Setup      │──►──│  Daily Run  │
└─────────────┘     └─────────────┘
```

## Phase 1 – First-Time Setup

Goal: bootstrap `.ctx/`, capture an initial prompt, and mark skeletons current.

```bash
ctx init
ctx ask
# ... share .ctx/prompt.md with your AI and save skeletons ...
ctx update
ctx status
```

- `ctx init` scaffolds `.ctx/`, seeds the index, and adds `.ctx/` to `.gitignore`.
- `ctx ask` runs `sync` and writes the prompt to `.ctx/prompt.md`.
- `ctx update` validates skeletons, marking everything current once saved.
- `ctx status` confirms the mirror is clean.

## Phase 2 – Daily Refresh

Goal: keep skeletons aligned with code changes as you develop.

```bash
ctx ask
# ... AI regenerates skeletons ...
ctx update
ctx bundle   # optional, share .ctx/context.md with collaborators
```

- `ctx ask` only targets files that changed since the last update.
- `ctx update` recomputes hashes and highlights anything still pending.
- `ctx bundle` packages all current skeletons before pairing sessions.

## When to Use Advanced Commands

The guided workflow covers most cases. Reach for advanced commands when you need finer control:

- `ctx sync --full` – rescan the entire repo after massive refactors.
- `ctx generate --files path/to/file.go` – build a prompt for a specific subset.
- `ctx pipeline` – scriptable combo of sync + generate.
- `ctx validate --fix --strict` – tighten CI gates.
- `ctx export --format json` – automate context publishing.

## Recovery Playbook

If `.ctx/` drifts out of sync:

```bash
ctx rebuild --confirm
ctx ask
# regenerate skeletons via AI
ctx update
```

Need to reset only a few files? Use `ctx generate --files` followed by your usual `ask → update` loop.

## Automation Ideas

Until dedicated CI recipes land, consider:

- `ctx ask --quiet` in pre-commit hooks to surface pending skeleton work.
- `ctx validate --fix --strict` in CI to prevent stale mirrors.
- Publishing the output of `ctx bundle` for asynchronous collaborators.

See [`docs/examples.md`](./examples.md) for concrete scenarios.
