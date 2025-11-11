package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/dakshpareek/ctx/internal/display"
)

func newPipelineCmd() *cobra.Command {
	syncOpts := syncOptions{}
	genOpts := generateOptions{
		filter: "pending,stale,missing",
	}

	cmd := &cobra.Command{
		Use:   "pipeline",
		Short: "Run sync and generate in one step",
		Long:  advancedDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			printAdvancedNotice("ctx ask")
			if err := runSync(syncOpts); err != nil {
				return err
			}
			if err := runGenerate(genOpts); err != nil {
				return err
			}

			if genOpts.quiet {
				fmt.Println(display.Info("Prompt generation finished silently (--quiet). Run 'ctx update' after saving skeletons."))
			} else {
				fmt.Println()
				fmt.Println(display.Info("Next steps: share the prompt with your AI assistant, update skeletons, then run 'ctx update'."))
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&syncOpts.full, "full", false, "force full scan (ignore git diff)")
	cmd.Flags().BoolVarP(&syncOpts.verbose, "verbose", "v", false, "show detailed file changes during sync")
	cmd.Flags().StringVar(&genOpts.filter, "filter", genOpts.filter, "comma-separated statuses to include")
	cmd.Flags().StringVar(&genOpts.files, "files", "", "comma-separated list of specific files to include in the prompt")
	cmd.Flags().StringVarP(&genOpts.output, "output", "o", "", "write prompt to a specific file")
	cmd.Flags().BoolVarP(&genOpts.quiet, "quiet", "q", false, "Suppress prompt body (still writes to file)")

	return cmd
}
