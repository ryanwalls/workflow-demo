package commands

import (
	"github.com/3dsim/workflow/config"
	"github.com/3dsim/workflow/logger"
	"github.com/3dsim/workflow/simulationworkflow"
	"github.com/spf13/cobra"
)

var deciderCmd = &cobra.Command{
	// Eventually, we have more deciders and workflows.  May want to namespace commands
	// using https://github.com/gosuri/cmdns
	Use:   "decider",
	Short: "Start EXASIM decider",
	Long:  `decider starts a decider for the support optimization workflow`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := simulationworkflow.NewRegistrator().InitializeWorkflow(); err != nil {
			logger.Log.Crit("Could not initialize the workflow", "error", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		taskListName := config.Viper.GetString("supportOptimizationWorkflow.defaultTaskList")
		identity := config.Viper.GetString("supportOptimizationWorkflow.deciderIdentity")
		domain := config.Viper.GetString("env")
		simulationworkflow.NewDecider(taskListName, identity, domain).Start()
	},
}
