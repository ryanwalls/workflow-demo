package commands

import (
	"github.com/spf13/cobra"
	"os"
)

// WorkflowCmd is the root command
var WorkflowCmd = &cobra.Command{
	Use: "workflow",
}

//Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	addCommands()
	if err := WorkflowCmd.Execute(); err != nil {
		// the err is already logged by Cobra
		os.Exit(-1)
	}
}

func addCommands() {
	WorkflowCmd.AddCommand(deciderCmd)
	WorkflowCmd.AddCommand(workerCmd)
	WorkflowCmd.AddCommand(fsmCmd)
}
