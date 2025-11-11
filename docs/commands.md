# Command Reference

`ctx` exposes two groups of commands: the guided “core workflow” that most developers use day-to-day, and “advanced” commands for fine-grained control or scripting. Every command accepts `--help` for detailed usage.

---

## Core Workflow

### `ctx init`

Bootstrap `.ctx/` in the current directory.

- Creates `config.json`, `index.json`, and `skeletons/`.
- Adds `.ctx/` to `.gitignore`.
- Performs the initial scan, marking files as `missing`.
- Safe to run once per project; use `ctx rebuild --confirm` to start over.

### `ctx ask`

Syncs the index and generates a prompt for files needing skeleton work.

- Writes the prompt to `.ctx/prompt.md` (override with `--output`).
- Prints the prompt unless `--quiet` is supplied.
- Automatically marks targeted files as `pendingGeneration`.

Flags:

- `--files` – focus on a specific comma-separated list.
- `--filter` – override the default statuses (`pending,stale,missing`).
- `--output`, `-o` – explicit prompt path.
- `--quiet`, `-q` – suppress prompt body on stdout.

### `ctx update`

Marks skeletons current after you save AI output.

- Wraps `ctx validate --fix`.
- Shows before/after stats (`current`, `stale`, `missing`, `pending`).
- Warns if work remains so you can rerun `ctx ask`.

### `ctx bundle`

Exports all current skeletons into a single artifact.

- Default path: `.ctx/context.md` (or `.ctx/context.json` with `--format json`).
- Helpful before pairing sessions or when handing context to a teammate.

Flags:

- `--output`, `-o` – custom destination.
- `--format` – `markdown` (default) or `json`.

### `ctx status`

Displays index summary and optional file-level details.

Flags:

- `--verbose`, `-v` – list files grouped by status.
- `--json` – emit machine-readable JSON.

---

## Advanced Commands

These commands remain available for scripts or specialized flows. The guided workflow (`init → ask → update → bundle`) should cover daily use.

### `ctx sync`

Scan the project for changes and update the index only.

Flags:

- `--full` – ignore Git hints and rescan the entire repo.
- `--verbose`, `-v` – print file-by-file changes.

### `ctx generate`

Build a prompt without the convenience wrapper provided by `ctx ask`.

Flags:

- `--filter` – statuses to include (`pending`, `stale`, `missing`, `current`).
- `--files` – comma-separated paths.
- `--output`, `-o` – prompt destination (default `.ctx/prompt.md`).
- `--quiet`, `-q` – suppress prompt body.

### `ctx pipeline`

Run `sync` then `generate` in one command—useful in automation.

Flags mirror `ctx sync` and `ctx generate`.

### `ctx validate`

Check consistency between source files, skeletons, and index metadata.

Flags:

- `--fix` – repair common issues (hash mismatches, missing skeletons).
- `--strict` – exit non-zero when issues remain.

### `ctx clean`

Remove orphaned skeleton files from `.ctx/skeletons/` and prune empty folders.

### `ctx rebuild`

Reset the entire `.ctx/` workspace (destructive).

Steps:

1. Delete all skeletons.
2. Reset the index.
3. Perform a full scan.

Requires `--confirm` to proceed.

### `ctx export`

Collect all `current` skeletons into one bundle.

Flags:

- `--format` – `markdown` (default) or `json`.
- `--output`, `-o` – write export to a file (stdout when omitted).

---

For automation-friendly recipes using these advanced commands, check out [`docs/examples.md`](./examples.md).
