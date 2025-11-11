package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/dakshpareek/ctx/internal/display"
	"github.com/dakshpareek/ctx/internal/index"
	"github.com/dakshpareek/ctx/internal/types"
)

type askOptions struct {
	quiet bool
}

func newAskCmd() *cobra.Command {
	opts := askOptions{}

	cmd := &cobra.Command{
		Use:   "ask",
		Short: "Sync project and generate a prompt for files needing skeleton updates",
		Long: `Run this after writing code to refresh your AI prompt.

ctx ask will:
  1. Sync the index with your latest code changes
  2. Identify files that need new or updated skeletons
  3. Generate a prompt saved to .ctx/prompt.md (and print it unless --quiet)
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runAsk(opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.quiet, "quiet", "q", false, "Do not print the generated prompt to stdout")

	return cmd
}

func runAsk(opts askOptions) error {
	ctxDir, _, err := ensureWorkspace(false)
	if err != nil {
		return err
	}

	if err := runSync(syncOptions{}); err != nil {
		return err
	}

	idx, err := index.LoadIndex(filepath.Join(ctxDir, indexFileName))
	if err != nil {
		return &types.Error{Code: types.ExitCodeData, Err: err}
	}

	stats := idx.Stats
	pending := stats.PendingGeneration
	stale := stats.Stale
	missing := stats.Missing
	total := pending + stale + missing

	if total == 0 {
		fmt.Println(display.Success("All skeletons are current! Nothing to generate right now."))
		fmt.Println(display.Info("Next: run 'ctx update' after adding new skeletons, or 'ctx status' to inspect the index."))
		return nil
	}

	fmt.Println(display.Bold("Files needing skeleton updates:"))
	if pending > 0 {
		fmt.Printf("  • %d pending generation\n", pending)
	}
	if stale > 0 {
		fmt.Printf("  • %d stale\n", stale)
	}
	if missing > 0 {
		fmt.Printf("  • %d missing\n", missing)
	}

	promptPath := filepath.Join(ctxDir, "prompt.md")
	genOpts := generateOptions{
		filter: "pending,stale,missing",
		output: promptPath,
		quiet:  opts.quiet,
	}

	if err := runGenerate(genOpts); err != nil {
		return err
	}

	fmt.Println()
	fmt.Println(display.Success("Prompt saved to %s", promptPath))
	fmt.Println(display.Info("Next steps: 1) share the prompt with your AI assistant, 2) save skeletons under .ctx/skeletons/, 3) run 'ctx update'."))

	if opts.quiet {
		return nil
	}

	data, err := os.ReadFile(promptPath)
	if err != nil {
		return &types.Error{Code: types.ExitCodeFileSystem, Err: fmt.Errorf("read prompt: %w", err)}
	}

	fmt.Println()
	fmt.Println(display.Bold("Prompt Preview:"))
	fmt.Println(string(data))

	return nil
}
