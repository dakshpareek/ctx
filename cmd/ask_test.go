package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestAskCreatesDefaultPromptAndPrintsSummary(t *testing.T) {
	dir := t.TempDir()
	writeTempFile(t, dir, "main.go", "package main\n")
	_, _ = executeCommand(t, dir, "init")

	stdout := execAndCaptureStdout(t, dir, "ask")

	promptPath := filepath.Join(dir, ".ctx", "prompt.md")
	data, err := os.ReadFile(promptPath)
	if err != nil {
		t.Fatalf("read prompt: %v", err)
	}
	if !strings.Contains(string(data), "Code Context Skeleton Generation") {
		t.Fatalf("expected prompt content in %s", promptPath)
	}

	if !strings.Contains(stdout, "Prompt saved to") {
		t.Fatalf("expected summary mentioning prompt path, got:\n%s", stdout)
	}
	if !strings.Contains(stdout, "Files needing skeleton updates") {
		t.Fatalf("expected list of files needing updates")
	}
}

func TestAskQuietSuppressesPromptBody(t *testing.T) {
	dir := t.TempDir()
	writeTempFile(t, dir, "main.go", "package main\n")
	_, _ = executeCommand(t, dir, "init")

	stdout := execAndCaptureStdout(t, dir, "ask", "--quiet")

	if strings.Contains(stdout, "Code Context Skeleton Generation") {
		t.Fatalf("expected quiet mode to suppress prompt body, got:\n%s", stdout)
	}
	if !strings.Contains(stdout, "Prompt saved to") {
		t.Fatalf("expected summary mentioning prompt path even in quiet mode")
	}

	promptPath := filepath.Join(dir, ".ctx", "prompt.md")
	if _, err := os.Stat(promptPath); err != nil {
		t.Fatalf("expected prompt file at %s", promptPath)
	}
}

func TestAskWhenAllCurrent(t *testing.T) {
	dir := t.TempDir()
	writeTempFile(t, dir, "main.go", "package main\n")
	_, _ = executeCommand(t, dir, "init")

	_, _ = executeCommand(t, dir, "ask", "--quiet")
	idx := loadIndex(t, dir)
	entry := idx.Files["main.go"]
	writeTempFile(t, dir, entry.SkeletonPath, "** skeleton **\n")
	_, _ = executeCommand(t, dir, "update")

	stdout := execAndCaptureStdout(t, dir, "ask", "--quiet")
	if !strings.Contains(stdout, "All skeletons are current") {
		t.Fatalf("expected success message when nothing needs generation:\n%s", stdout)
	}
	if strings.Contains(stdout, "Prompt saved to") {
		t.Fatalf("should not create prompt when everything current:\n%s", stdout)
	}
}
