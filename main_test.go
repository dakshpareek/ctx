package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	rootcmd "github.com/dakshpareek/ctx/cmd"
	"github.com/dakshpareek/ctx/internal/index"
	"github.com/dakshpareek/ctx/internal/types"
)

func TestGuidedWorkflowEndToEnd(t *testing.T) {
	tempDir := t.TempDir()

	srcDir := filepath.Join("testdata", "sample-project")
	if err := copyTree(srcDir, tempDir); err != nil {
		t.Fatalf("copyTree: %v", err)
	}

	withWorkingDir(t, tempDir)

	runCommand(t, "init")
	runCommand(t, "ask", "--quiet")

	if _, err := os.Stat(".ctx/index.json"); err != nil {
		t.Fatalf("expected index.json to exist: %v", err)
	}

	promptPath := filepath.Join(".ctx", "prompt.md")
	if _, err := os.Stat(promptPath); err != nil {
		t.Fatalf("expected prompt.md to be generated at default location: %v", err)
	}
	verifyPromptContains(t, promptPath, "## Files to Process")

	idx := loadIndex(t, ".")
	entry, ok := idx.Files["src/app.go"]
	if !ok {
		t.Fatalf("expected src/app.go tracked in index")
	}
	if entry.Status != types.StatusPendingGeneration {
		t.Fatalf("expected status pendingGeneration after ask, got %s", entry.Status)
	}

	skeletonContent := "** skeleton **\n"
	if err := os.MkdirAll(filepath.Dir(entry.SkeletonPath), 0o755); err != nil {
		t.Fatalf("mkdir skeleton dir: %v", err)
	}
	if err := os.WriteFile(entry.SkeletonPath, []byte(skeletonContent), 0o644); err != nil {
		t.Fatalf("write skeleton: %v", err)
	}

	runCommand(t, "update")

	idx = loadIndex(t, ".")
	entry = idx.Files["src/app.go"]
	if entry.Status != types.StatusCurrent {
		t.Fatalf("expected status current after update, got %s", entry.Status)
	}

	runCommand(t, "bundle")

	contextPath := filepath.Join(".ctx", "context.md")
	if _, err := os.Stat(contextPath); err != nil {
		t.Fatalf("expected context bundle at %s: %v", contextPath, err)
	}
	data, err := os.ReadFile(contextPath)
	if err != nil {
		t.Fatalf("read context bundle: %v", err)
	}
	if !bytes.Contains(data, []byte(skeletonContent)) {
		t.Fatalf("expected bundle to include skeleton content")
	}
}

func runCommand(t *testing.T, args ...string) {
	t.Helper()

	cmd := rootcmd.NewRootCmd("test")
	var stdout, stderr bytes.Buffer
	cmd.SetOut(&stdout)
	cmd.SetErr(&stderr)
	cmd.SetArgs(args)

	if err := cmd.Execute(); err != nil {
		t.Fatalf("command %v failed: %v\nstdout:\n%s\nstderr:\n%s", args, err, stdout.String(), stderr.String())
	}
}

func withWorkingDir(t *testing.T, dir string) {
	t.Helper()
	original, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(original)
	})
}

func copyTree(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)

		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}

		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()

		out, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0o644)
		if err != nil {
			return err
		}
		defer out.Close()

		if _, err := io.Copy(out, in); err != nil {
			return err
		}
		return nil
	})
}

func verifyPromptContains(t *testing.T, path, substring string) {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read prompt: %v", err)
	}
	if !bytes.Contains(data, []byte(substring)) {
		t.Fatalf("expected prompt to contain %q", substring)
	}
}

func loadIndex(t *testing.T, dir string) *types.Index {
	t.Helper()
	idx, err := index.LoadIndex(filepath.Join(dir, ".ctx", "index.json"))
	if err != nil {
		t.Fatalf("LoadIndex: %v", err)
	}
	return idx
}
