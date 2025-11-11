package cmd

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunGenerateErrors(t *testing.T) {
	dir := t.TempDir()
	cleanup := changeDir(t, dir)
	err := runGenerate(generateOptions{})
	if err == nil {
		t.Fatalf("expected error when not initialized")
	}
	cleanup()

	_, _ = executeCommand(t, dir, "init")

	cleanup = changeDir(t, dir)
	defer cleanup()

	if err := runGenerate(generateOptions{files: "unknown.go"}); err == nil {
		t.Fatalf("expected error for untracked file")
	}
}

func TestRunGenerateStdout(t *testing.T) {
	dir := t.TempDir()
	writeTempFile(t, dir, "main.go", "package main\n")
	_, _ = executeCommand(t, dir, "init")

	cleanup := changeDir(t, dir)
	defer cleanup()

	output := captureOutput(t, func() {
		if err := runGenerate(generateOptions{}); err != nil {
			t.Fatalf("runGenerate: %v", err)
		}
	})
	if !strings.Contains(output, "Code Context Skeleton Generation") {
		t.Fatalf("expected prompt output")
	}
}

func TestGenerateCommandDefaultFilterIncludesPending(t *testing.T) {
	dir := t.TempDir()
	writeTempFile(t, dir, "main.go", "package main\n")
	_, _ = executeCommand(t, dir, "init")

	_, _ = executeCommand(t, dir, "generate")
	_, _ = executeCommand(t, dir, "generate")
}

func TestRunGenerateWritesDefaultPromptFile(t *testing.T) {
	dir := t.TempDir()
	writeTempFile(t, dir, "main.go", "package main\n")
	_, _ = executeCommand(t, dir, "init")

	cleanup := changeDir(t, dir)
	defer cleanup()

	if err := runGenerate(generateOptions{quiet: true}); err != nil {
		t.Fatalf("runGenerate: %v", err)
	}

	defaultPath := filepath.Join(".ctx", "prompt.md")
	data, err := os.ReadFile(defaultPath)
	if err != nil {
		t.Fatalf("expected default prompt file: %v", err)
	}
	if !strings.Contains(string(data), "Code Context Skeleton Generation") {
		t.Fatalf("expected prompt content in %s", defaultPath)
	}
}

func TestRunGenerateQuietSuppressesPromptPrint(t *testing.T) {
	dir := t.TempDir()
	writeTempFile(t, dir, "main.go", "package main\n")
	_, _ = executeCommand(t, dir, "init")

	cleanup := changeDir(t, dir)
	defer cleanup()

	output := captureOutput(t, func() {
		if err := runGenerate(generateOptions{quiet: true}); err != nil {
			t.Fatalf("runGenerate: %v", err)
		}
	})

	if strings.Contains(output, "Code Context Skeleton Generation") {
		t.Fatalf("expected quiet mode to suppress prompt body, got: %s", output)
	}
	if !strings.Contains(output, "Generated prompts for 1 file(s)") {
		t.Fatalf("expected summary output even in quiet mode")
	}
}
