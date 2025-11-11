## Relevant Files

- `cmd/root.go` - Registers CLI commands; needs to feature the new wrappers and reorder advanced commands.
- `cmd/ask.go` - New command orchestrating sync and prompt generation with friendly messaging.
- `cmd/update.go` - New command wrapping validation with `--fix` to mark skeletons current.
- `cmd/bundle.go` - New command wrapping export to produce full-context bundles.
- `cmd/workspace.go` - Shared helpers for checking `.ctx/` workspace and index availability.
- `cmd/generate.go` - Adjust default filters and output behavior to support the simplified flow.
- `cmd/pipeline.go` - Ensure combined sync/generate command inherits the new defaults.
- `cmd/validate.go` - Surface summary data for `update` and ensure statuses transition cleanly.
- `cmd/export.go` - Provide reusable export logic for the `bundle` command and default output path handling.
- `internal/index/index.go` - May expose helper(s) for stats retrieval used by wrapper commands.
- `main_test.go` - Integration coverage likely needs updates for the new command set.
- `cmd/ask_test.go` - New unit tests validating ask command flags and execution path.
- `cmd/update_test.go` - New unit tests for update command behavior.
- `cmd/bundle_test.go` - New unit tests for bundle command behavior.
- `cmd/advanced_cmds.go` - Shared advanced-command description string.
- `cmd/generate_errors_test.go` - Tests covering generate command defaults and quiet behavior.
- `README.md` - Update quick start and workflow documentation.
- `docs/workflows.md` - Rewrite daily flow around `init → ask → update`.
- `docs/commands.md` - Document the new command set and mark advanced commands.
- `docs/examples.md` - New file with practical usage scenarios.
- `docs/getting-started.md` - Align onboarding instructions with the guided workflow.
- `docs/troubleshooting.md` - Update resolutions and add new FAQ entries.
- `CHANGELOG.md` - Summarize behavior changes and migration notes (create if missing).

### Notes

- Unit tests should live alongside the code they exercise (e.g., `cmd/ask.go` with `cmd/ask_test.go`).
- Use `go test ./...` to exercise the full suite after changes.
- Implement incrementally and test each command independently before moving to the next.
- Wrapper commands can reuse existing helpers: `runSync`, `runGenerate`, `runValidate`, `runStatus`, and `runExport`, each already encapsulating core logic for their respective commands.

## Tasks

- [x] 1.0 Implement friendly wrapper commands for core workflow
  - [x] 1.1 Update `cmd/root.go` to register new commands prominently and categorize legacy commands as advanced options.
  - [x] 1.2 Review existing command wiring and helper functions to identify reusable pieces for wrappers.
  - [x] 1.3 Create `cmd/ask.go` implementing the sync + generate flow with default filter (`pending,stale,missing`) and default output path `.ctx/prompt.md`.
  - [x] 1.4 Add `--quiet` support to `ask` to suppress prompt printing and ensure user-facing messaging guides next steps.
  - [x] 1.5 Create `cmd/update.go` that wraps validation with `--fix`, prints a summary, and handles partial success messaging.
  - [x] 1.6 Create `cmd/bundle.go` that wraps export, writes to `.ctx/context.md` when no `--output` is provided, and reports included skeleton count.
  - [x] 1.7 Add intelligent error messages when `.ctx/` doesn't exist (guide user to run `ctx init` first).
  - [x] 1.8 Handle edge case where `ctx ask` finds no changes (print success message, don't create empty prompt file).
  - [x] 1.9 Add confirmation/summary after `ctx update` showing before/after stats.

- [x] 2.0 Align existing commands with the simplified UX defaults
  - [x] 2.1 Change the default filter in `cmd/generate.go`/`cmd/pipeline.go` to include `pending,stale,missing`.
  - [x] 2.2a Add default output file behavior in `runGenerate` (write to `.ctx/prompt.md` when `--output` is omitted).
  - [x] 2.2b Retain stdout printing by default; only suppress with `--quiet` flag.
  - [x] 2.2c Add tests confirming both file and stdout outputs work correctly.
  - [x] 2.3 Ensure `pipeline` inherits the new output behavior and prints follow-up guidance consistent with `ask`.
  - [x] 2.4 Update help/usage strings for `sync`, `generate`, `validate`, and `pipeline` to label them as advanced workflows.
  - [x] 2.5 Keep `sync`, `generate`, `validate`, `export` fully functional but mark as "Advanced" in help.
  - [x] 2.6 Add deprecation warnings (optional, for future major version) to steer users toward new commands.
  - [x] 2.7 Ensure `--help` output shows new commands first, advanced commands in separate section.
  - [x] 2.8 Manual testing checkpoint: Run through complete workflow (`init → ask → update`) on sample project.
  - [x] 2.9 Verify error messages are helpful and guide user to next action.

- [x] 3.0 Refresh documentation to reflect the new command structure
  - [x] 3.1 Update `README.md` quick start and daily workflow sections to showcase `init → ask → update`.
  - [x] 3.2 Revise `docs/workflows.md` to center the two-phase story (setup vs. daily refresh) using new command names.
  - [x] 3.3 Update `docs/commands.md` to add entries for `ask`, `update`, `bundle`, and reposition legacy commands under an advanced heading.
  - [x] 3.4 Verify ancillary docs (e.g., troubleshooting, contributing) don't contradict the new UX.
  - [x] 3.5 Create `docs/examples.md` with common scenarios (first-time setup, daily refresh, specific file updates).
  - [x] 3.6 Add troubleshooting entry for "Why doesn't `ctx ask` show my changes?" (need to run `ctx sync` after ignoring files, etc.).
  - [x] 3.7 Add visual diagram showing the workflow cycle in `docs/workflows.md`.

### 4.0 Expand automated test coverage for new behaviors

- [x] 4.1 Add unit tests for `ask`, `update`, and `bundle` covering flag parsing, default paths, and stdout messaging.
- [x] 4.2 Update integration tests (e.g., `main_test.go`) to exercise the new commands and confirm end-to-end flows succeed.
- [x] 4.3 Extend existing tests for `runGenerate` and `pipeline` to assert the revised defaults.
- [x] 4.4 Test `ctx ask` when all files are current (should print success, not error).
- [x] 4.5 Test `ctx update` when no skeletons exist yet (should handle gracefully).
- [x] 4.6 Test `ctx bundle` when `.ctx/skeletons/` is empty.
- [x] 4.7 Add smoke test that exercises full cycle: `init → ask → update → status`.

### 5.0 Communicate changes and migration guidance for existing users

- [ ] 5.1 Draft release notes or update `CHANGELOG.md` detailing command renames, new defaults, and backward-compatibility nuances.
- [ ] 5.2 Note any necessary migration steps (e.g., scripts that call `ctx generate`) and recommend `ctx ask` in documentation.
- [ ] 5.3 Coordinate version bumping and ensure `--version` output reflects the release containing these UX changes.
- [ ] 5.4 Update any analytics/telemetry to track new command usage vs old commands (if applicable).
- [ ] 5.5 Create `ctx migrate` command that shows current index status and validates workspace (optional, nice-to-have).

### 6.0 User experience polish

- [ ] 6.1 Add color/emoji to terminal output for better scannability (✅ ❌ 🔄 📝 📦 etc.).
- [ ] 6.2 Implement progress indicators for long operations (sync on large repos).
- [ ] 6.3 Add `--dry-run` flag to `ctx update` so users can preview what will change.
- [ ] 6.4 Consider adding `ctx reset` command to clear `.ctx/` and start fresh (useful for testing/recovery).
- [ ] 6.5 Ensure proper exit codes (0 for success, 1 for errors) for script usage across all commands.
- [ ] 6.6 Verify pipe-ability: `ctx ask --quiet` should work cleanly in scripts (e.g., `ctx ask --quiet | pbcopy`).

### 7.0 Quality assurance and edge cases

- [ ] 7.1 Add validation for corrupted `.ctx/index.json` with helpful recovery instructions.
- [ ] 7.2 Test performance on large repos (1000+ files) to ensure `sync` completes in reasonable time.
- [ ] 7.3 Verify help text consistency across `ctx --help`, `ctx ask --help`, etc. (same tone and formatting).
- [ ] 7.4 Test behavior when `.ctxignore` contains invalid patterns or syntax errors.
- [ ] 7.5 Ensure commands fail gracefully when run outside a Git repository or project root.

## Implementation Order (Recommended)

1. **Core plumbing** (Tasks 1.1-1.9) - Get commands wired up with basic functionality
2. **Quick validation** (Tasks 2.8-2.9) - Manually test on a small project before going deeper
3. **Defaults and alignment** (Tasks 2.1-2.7) - Align behavior across old and new commands
4. **Automated tests** (Section 4.0) - Lock in the behavior with comprehensive coverage
5. **Documentation** (Section 3.0) - Document what actually works
6. **UX polish** (Section 6.0) - Add colors, progress indicators, better error messages
7. **Quality assurance** (Section 7.0) - Edge cases and performance testing
8. **Release preparation** (Section 5.0) - Communicate changes and migration path

## Critical Success Factors

- **Incremental testing**: Test each new command independently before moving to the next
- **Backward compatibility**: Existing scripts using old commands must continue to work
- **Clear migration path**: Users should understand why and how to adopt new commands
- **Helpful errors**: Every error should guide the user toward a solution
- **Performance**: Commands should feel instant on typical projects (<500 files)

## Future Enhancements (Post-Release)

- [ ] Configuration file support (`.ctxrc`) for user preferences
- [ ] Interactive mode: `ctx ask --interactive` to select specific files
- [ ] Git hooks integration to auto-run `ctx sync` on commit
- [ ] VS Code extension for one-click skeleton generation
- [ ] `ctx doctor` command to diagnose common issues
