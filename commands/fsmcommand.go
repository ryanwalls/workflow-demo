package commands

import (
	"github.com/3dsim/workflow/config"
	"github.com/3dsim/workflow/logger"
	"github.com/3dsim/workflow/simulationworkflow"
	"github.com/spf13/cobra"
)

var fsmCmd = &cobra.Command{
	Use:   "fsm",
	Short: "Start finite state manager",
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := simulationworkflow.NewRegistrator().InitializeWorkflow(); err != nil {
			logger.Log.Crit("Could not initialize the workflow", "error", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		taskListName := config.Viper.GetString("supportOptimizationWorkflow.defaultTaskList")
		identity := config.Viper.GetString("supportOptimizationWorkflow.deciderIdentity")
		domain := config.Viper.GetString("env")
		simulationworkflow.NewSimulationStateManager(taskListName, identity, domain).Start()
	},
}
