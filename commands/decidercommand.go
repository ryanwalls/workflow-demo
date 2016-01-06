package commands

import (
	"github.com/3dsim/workflow/config"
	"github.com/3dsim/workflow/decider"
	"github.com/3dsim/workflow/logger"
	"github.com/spf13/cobra"
)

var deciderCmd = &cobra.Command{
	Use:   "decider",
	Short: "Start decider",
	Long:  `decider starts a decider for AWS SWF`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Log.Info("Started", "env", config.Viper.GetString("env"))
		decider.RegisterDomain()
	},
}
