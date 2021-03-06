package commands

import (
	"github.com/3dsim/workflow/config"
	"github.com/3dsim/workflow/logger"
	"github.com/3dsim/workflow/simulationworkflow"
	"github.com/spf13/cobra"
)

var workerCmd = &cobra.Command{
	// Eventually, we have more deciders and workflows.  May want to namespace commands
	// using https://github.com/gosuri/cmdns
	Use:   "worker",
	Short: "Start worker",
	Long:  `Start worker`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := simulationworkflow.NewRegistrator().InitializeWorkflow(); err != nil {
			logger.Log.Crit("Could not initialize the workflow", "error", err)
		}
		if err := simulationworkflow.NewRegistrator().RegisterActivity(); err != nil {
			logger.Log.Crit("Could not register the activity", "error", err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		preprocTaskList := config.Workflow.Steps["preproc"].Name + "-" +
			config.Workflow.Steps["preproc"].ProductionVersion
		taskListName := preprocTaskList
		identity := preprocTaskList
		domain := config.Viper.GetString("env")
		simulationworkflow.NewWorker(taskListName, identity, domain).Start()
		for {
		}
	},
}
