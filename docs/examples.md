# Examples

Practical scenarios showing how to combine `ctx` commands in real projects.

## 1. First-Time Setup Script

```bash
#!/usr/bin/env bash
set -e

ctx init
ctx ask --quiet          # writes .ctx/prompt.md without printing it
open .ctx/prompt.md      # or use your editor to copy the prompt
```

**Why it helps:** Automates the initial bootstrap so you can copy the prompt immediately and hand it to your AI assistant.

## 2. Daily Refresh Alias

Add this to your shell profile:

```bash
alias ctx-refresh='ctx ask --quiet && echo "Prompt: .ctx/prompt.md"'
```

Follow up with:

```bash
# After AI updates skeletons
ctx update
```

**Why it helps:** Reduces the daily loop to two short commands (`ctx-refresh` + `ctx update`).

## 3. Targeting Specific Files

```bash
ctx ask --files internal/api/handler.go,internal/api/router.go
# ... regenerate only those skeletons ...
ctx update
```

**Why it helps:** Keeps prompts manageable when you are refactoring a focused area.

## 4. Producing a JSON Bundle for CI

```bash
ctx ask --quiet
# trigger AI to update skeletons (human or automated step)
ctx update
ctx bundle --format json --output artifacts/context.json
```

**Why it helps:** Stores the current architecture snapshot as a CI artifact that teammates or bots can reuse.

## 5. Recovery After Major Refactor

```bash
ctx sync --full
ctx ask --quiet
# regenerate skeletons via AI
ctx update
```

**Why it helps:** Forces a full rescan after large file moves or renames, while keeping the guided workflow intact.

---

Looking for more background? Dive into the [Workflows](./workflows.md) guide or the [Troubleshooting](./troubleshooting.md) section when something feels off.
