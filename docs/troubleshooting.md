# Troubleshooting

Common problems and how to fix them.

## `ctx init` Fails with “Already initialized”

`ctx init` aborts if `.ctx/` exists.

- Run `ctx rebuild --confirm` to reset the workspace.
- Or delete `.ctx/` manually if you are sure it’s stale.

## `ctx sync` Doesn’t Detect Changes

- Ensure you are running from the project root.
- For git repositories, confirm the project is a valid Git worktree (`git status`).
- Use `ctx sync --full` to force a complete scan.

## `ctx ask` Shows “All skeletons are current” But I Changed Files

- Ensure the files you edited are tracked by `.ctx/config.json` include/exclude rules.
- If you recently added new globs to `.ctxignore`, run `ctx sync --full` once to refresh the index.
- Check `ctx status --verbose`—if files appear as `current`, rerun `ctx ask --files path/to/file.go` for a targeted prompt.

## `ctx generate` Returns “No files match the requested filters”

- Run `ctx ask` or `ctx sync` first so the index reflects recent edits.
- Check status with `ctx status --verbose`.
- Use `--filter pending,stale,missing` (the default) unless you need other statuses.

## Skeleton Files Missing After AI Update

- Run `ctx validate --fix` to mark missing skeletons and restore index health.
- Rerun `ctx ask --files` for the affected paths and regenerate skeletons.

## Export Fails with “no current skeletons to export”

- Ensure you have at least one file marked `current`.
- Use the daily loop (`ctx ask` → regenerate skeletons → `ctx update`) to bring files current.

## Hash Mismatch Warnings

`ctx validate` or `ctx sync` may flag mismatched hashes when skeleton content changes without updating `index.json`.

1. Rerun `ctx ask --files` for the affected file.
2. Update the skeleton file.
3. Run `ctx update` (or `ctx validate --fix`) to refresh hashes.

## Still Stuck?

Capture the command output and index snippet, then open an issue or start a discussion in your repository. Include:

- Command run and flags.
- Relevant status output.
- Excerpts from `.ctx/index.json`.
