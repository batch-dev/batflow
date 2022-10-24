package cmd

import "github.com/spf13/cobra"

func GetRootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "batflow",
		Short: "Manages Batch Flow",
	}
	cmd.AddCommand(getWorkerCommand())
	cmd.AddCommand(getExecCommand())
	return cmd
}
