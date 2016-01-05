package commands

import (
	"github.com/3dsim/canary/config"
	"github.com/3dsim/canary/logger"
	"github.com/spf13/cobra"
)

var end2EndCmd = &cobra.Command{
	Use:   "end2end",
	Short: "Start simulation, check for errors.",
	Long: `end2end starts a default 1x1x1 cube simulation by calling the studio api
  start simulation endpoint.  It then polls the simulation results and checks
  progress.  If an error or timeout occurs, the command exits with an error code.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Log.Info("Started", "env", config.Viper.GetString("env"))
	},
}
