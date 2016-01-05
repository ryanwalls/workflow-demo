package commands

import (
	"github.com/spf13/cobra"
	"os"
)

// SolverSvcCmd is the root command to run solvers
var SolverSvcCmd = &cobra.Command{
	Use: "canary",
}

//Execute adds all child commands to the root command SolverSvcCmd and sets flags appropriately.
func Execute() {
	addCommands()
	if err := SolverSvcCmd.Execute(); err != nil {
		// the err is already logged by Cobra
		os.Exit(-1)
	}
}

func addCommands() {
	SolverSvcCmd.AddCommand(end2EndCmd)
}
