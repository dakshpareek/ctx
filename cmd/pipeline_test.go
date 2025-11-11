package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPipelineGeneratesPrompt(t *testing.T) {
	tempDir := t.TempDir()
	writeTempFile(t, tempDir, "main.go", "package main\n")

	_, _ = executeCommand(t, tempDir, "init")

	if err := os.WriteFile(filepath.Join(tempDir, "main.go"), []byte("package main\n// change\n"), 0o644); err != nil {
		t.Fatalf("write change: %v", err)
	}

	_, _ = executeCommand(t, tempDir, "pipeline", "--output", "pipeline.md")

	data, err := os.ReadFile(filepath.Join(tempDir, "pipeline.md"))
	if err != nil {
		t.Fatalf("read pipeline prompt: %v", err)
	}
	if len(data) == 0 {
		t.Fatalf("expected pipeline prompt to have content")
	}
}

func TestPipelineDefaultOutputAndQuietMode(t *testing.T) {
	dir := t.TempDir()
	writeTempFile(t, dir, "main.go", "package main\n")
	_, _ = executeCommand(t, dir, "init")

	if err := os.WriteFile(filepath.Join(dir, "main.go"), []byte("package main\n// change\n"), 0o644); err != nil {
		t.Fatalf("write change: %v", err)
	}

	stdout := execAndCaptureStdout(t, dir, "pipeline")
	defaultPrompt := filepath.Join(dir, ".ctx", "prompt.md")
	if _, err := os.Stat(defaultPrompt); err != nil {
		t.Fatalf("expected default prompt at %s: %v", defaultPrompt, err)
	}
	if !strings.Contains(stdout, ".ctx/prompt.md") {
		t.Fatalf("expected stdout to mention default prompt path, got:\n%s", stdout)
	}

	stdout = execAndCaptureStdout(t, dir, "pipeline", "--quiet")
	if strings.Contains(stdout, "Code Context Skeleton Generation") {
		t.Fatalf("quiet pipeline should suppress prompt body, got:\n%s", stdout)
	}
	if !strings.Contains(stdout, "Prompt generation finished silently") {
		t.Fatalf("expected quiet message in stdout, got:\n%s", stdout)
	}
}
