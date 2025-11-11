package cmd

import "github.com/spf13/cobra"

const (
	coreGroupID     = "core"
	advancedGroupID = "advanced"
)

// NewRootCmd constructs the base CLI command for ctx.
func NewRootCmd(version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ctx",
		Short: "Extract and maintain your codebase's architecture",
		Long: `ctx helps you maintain a lightweight, up-to-date architectural 
snapshot of your project for AI-assisted development.`,
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.Version = version

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		_ = cmd.Help()
		return nil
	}

	cmd.AddGroup(
		&cobra.Group{
			ID:    coreGroupID,
			Title: "Core Workflow",
		},
		&cobra.Group{
			ID:    advancedGroupID,
			Title: "Advanced Commands",
		},
	)

	coreCommands := []*cobra.Command{
		newInitCmd(),
		newAskCmd(),
		newUpdateCmd(),
		newBundleCmd(),
		newStatusCmd(),
	}

	for _, coreCmd := range coreCommands {
		coreCmd.GroupID = coreGroupID
		cmd.AddCommand(coreCmd)
	}

	advancedCommands := []*cobra.Command{
		newSyncCmd(),
		newGenerateCmd(),
		newPipelineCmd(),
		newValidateCmd(),
		newExportCmd(),
		newCleanCmd(),
		newRebuildCmd(),
	}

	for _, advancedCmd := range advancedCommands {
		advancedCmd.GroupID = advancedGroupID
		cmd.AddCommand(advancedCmd)
	}

	return cmd
}
