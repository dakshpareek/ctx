package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBundleWritesDefaultContextFile(t *testing.T) {
	dir := t.TempDir()
	writeTempFile(t, dir, "main.go", "package main\n")
	_, _ = executeCommand(t, dir, "init")
	_, _ = executeCommand(t, dir, "ask", "--quiet")

	idx := loadIndex(t, dir)
	entry := idx.Files["main.go"]
	writeTempFile(t, dir, entry.SkeletonPath, "** skeleton **\n")
	_, _ = executeCommand(t, dir, "update")

	stdout := execAndCaptureStdout(t, dir, "bundle")

	if !strings.Contains(stdout, "Exported") {
		t.Fatalf("expected export summary in output:\n%s", stdout)
	}
	if !strings.Contains(stdout, ".ctx/context.md") {
		t.Fatalf("expected default output path mentioned in stdout:\n%s", stdout)
	}

	contextPath := filepath.Join(dir, ".ctx", "context.md")
	data, err := os.ReadFile(contextPath)
	if err != nil {
		t.Fatalf("read bundle: %v", err)
	}
	if !strings.Contains(string(data), "** skeleton **") {
		t.Fatalf("expected bundle to contain skeleton content")
	}
}

func TestBundleWhenSkeletonsMissing(t *testing.T) {
	dir := t.TempDir()
	writeTempFile(t, dir, "main.go", "package main\n")
	_, _ = executeCommand(t, dir, "init")

	stdout, stderr, err := executeCommandAllowError(t, dir, "bundle")
	if err == nil {
		t.Fatalf("expected bundle to fail when skeletons missing")
	}
	if !strings.Contains(err.Error(), "no current skeletons") && !strings.Contains(stderr, "no current skeletons") && !strings.Contains(stdout, "no current skeletons") {
		t.Fatalf("expected message about missing skeletons, got:\nerr=%v\nstdout=%s\nstderr=%s", err, stdout, stderr)
	}
}
