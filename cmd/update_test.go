package cmd

import (
	"strings"
	"testing"

	"github.com/dakshpareek/ctx/internal/types"
)

func TestUpdateMarksFilesCurrentAndShowsSummary(t *testing.T) {
	dir := t.TempDir()
	writeTempFile(t, dir, "main.go", "package main\n")
	_, _ = executeCommand(t, dir, "init")

	_, _ = executeCommand(t, dir, "ask", "--quiet")

	idx := loadIndex(t, dir)
	entry, ok := idx.Files["main.go"]
	if !ok {
		t.Fatalf("expected main.go tracked in index")
	}

	writeTempFile(t, dir, entry.SkeletonPath, "** skeleton **\n")

	stdout := execAndCaptureStdout(t, dir, "update")
	if !strings.Contains(stdout, "Update summary") {
		t.Fatalf("expected update summary in output:\n%s", stdout)
	}
	if !strings.Contains(stdout, "Current: 0 → 1") {
		t.Fatalf("expected current count transition, got:\n%s", stdout)
	}

	idx = loadIndex(t, dir)
	if idx.Files["main.go"].Status != types.StatusCurrent {
		t.Fatalf("expected status current after update, got %s", idx.Files["main.go"].Status)
	}
}

func TestUpdateHandlesMissingSkeletonsGracefully(t *testing.T) {
	dir := t.TempDir()
	writeTempFile(t, dir, "main.go", "package main\n")
	_, _ = executeCommand(t, dir, "init")
	_, _ = executeCommand(t, dir, "ask", "--quiet")

	stdout := execAndCaptureStdout(t, dir, "update")
	if !strings.Contains(stdout, "Remaining work detected") {
		t.Fatalf("expected reminder about remaining work when skeletons missing:\n%s", stdout)
	}

	idx := loadIndex(t, dir)
	entry := idx.Files["main.go"]
	if entry.Status != types.StatusMissing && entry.Status != types.StatusPendingGeneration {
		t.Fatalf("expected file to remain missing or pending, got %s", entry.Status)
	}

	stdout = execAndCaptureStdout(t, dir, "update")
	if !strings.Contains(stdout, "Remaining work detected") {
		t.Fatalf("expected reminder about remaining work on subsequent update:\n%s", stdout)
	}
}
