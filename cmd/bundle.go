package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/dakshpareek/ctx/internal/display"
)

type bundleOptions struct {
	output string
	format string
}

func newBundleCmd() *cobra.Command {
	opts := bundleOptions{}

	cmd := &cobra.Command{
		Use:   "bundle",
		Short: "Export all current skeletons into a single context file",
		Long: `Bundle current skeletons before a pairing or AI session.

ctx bundle will export all current skeletons along with index stats into .ctx/context.md by default.
Use --output to override the destination or --format to export JSON.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runBundle(opts)
		},
	}

	cmd.Flags().StringVarP(&opts.output, "output", "o", "", "write export to file instead of default .ctx/context.<ext>")
	cmd.Flags().StringVar(&opts.format, "format", "markdown", "output format: markdown or json")

	return cmd
}

func runBundle(opts bundleOptions) error {
	ctxDir, _, err := ensureWorkspace(true)
	if err != nil {
		return err
	}

	exportOpts := exportOptions{
		format: opts.format,
		output: opts.output,
	}

	if exportOpts.output == "" {
		extension := "md"
		if opts.format == "json" {
			extension = "json"
		}
		exportOpts.output = filepath.Join(ctxDir, fmt.Sprintf("context.%s", extension))
	}

	if err := runExport(exportOpts); err != nil {
		return err
	}

	fmt.Println(display.Info("Next: share this bundle with your AI assistant for full-project context."))
	return nil
}
