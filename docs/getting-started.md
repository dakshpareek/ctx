# Getting Started

Welcome to `ctx`, the Code Context CLI. Follow this guide to bootstrap a new project and keep your AI companions fully informed.

## Prerequisites

- Go 1.21+ installed.
- A Git repository (recommended for fast change detection, but not required).

## Installation

```bash
go install github.com/yourusername/ctx@latest
```

This places the `ctx` binary in your `$GOBIN` (`$GOPATH/bin` by default).

## Initialize a Project

```bash
cd /path/to/project
ctx init
```

The command:

1. Creates `.ctx/`, `config.json`, `index.json`, and `skeletons/`.
2. Adds `.ctx/` to `.gitignore`.
3. Scans your project for trackable files and marks them as `missing`.

## Run the Guided Workflow

`ctx` breaks work into two simple phases.

### Phase 1 – First-Time Setup

```bash
ctx init
ctx ask
# ... share .ctx/prompt.md with your AI assistant ...
ctx update
```

What happens:

1. `ctx init` creates `.ctx/`, seeds the index, and adds the workspace to `.gitignore`.
2. `ctx ask` runs `sync` + `generate`, saving the prompt to `.ctx/prompt.md`.
3. Your AI assistant writes skeletons under `.ctx/skeletons/…`.
4. `ctx update` recomputes hashes and marks everything current.

### Phase 2 – Daily Refresh

```bash
ctx ask
# ... regenerate skeletons via AI ...
ctx update
```

- `ctx ask` only targets files flagged as `pending`, `stale`, or `missing`.
- `ctx update` surfaces before/after stats and warns if additional files need attention.

Optional extras:

- `ctx bundle` → package all current skeletons into `.ctx/context.md`.
- `ctx status --verbose` → inspect which files still need work.

## Advanced Controls (Optional)

If you prefer manual control or automation hooks:

- `ctx sync --full` – rescan the entire project.
- `ctx generate --files path/to/file.go` – build a targeted prompt.
- `ctx pipeline` – run `sync` and `generate` together in scripts.
- `ctx validate --fix --strict` – enforce integrity in CI.
- `ctx export --format json` – produce machine-readable bundles.

## Next Steps

- Explore the [Workflows](./workflows.md) guide for deeper scenarios.
- Browse [Examples](./examples.md) for scripted usage patterns.
- Keep your tests green with `go test ./... -cover`.
- Contribute improvements—see [Contributing](../CONTRIBUTING.md).
