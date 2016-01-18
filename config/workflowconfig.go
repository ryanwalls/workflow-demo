package config

var Workflow *WorkflowConfig

const (
	rootKey = "supportOptimizationWorkflow"
)

func init() {
	Workflow = SetupConfiguration()
}

type DeciderConfig struct {
	TaskTimeout string
	Name        string
	Version     string
}

type StepConfig struct {
	Name                        string
	ScheduleToCloseTimeout      string
	ScheduleToStartTimeout      string
	StartToCloseTimeout         string
	DefaultTaskHeartbeatTimeout string
	ProductionVersion           string
	ExperimentalVersions        []string
}

type WorkflowConfig struct {
	Decider         DeciderConfig
	Version         string
	Name            string
	DefaultTaskList string
	StepOrder       []string
	Steps           map[string]StepConfig
}

func SetupConfiguration() *WorkflowConfig {
	workflowConfig := &WorkflowConfig{}
	if err := Viper.UnmarshalKey(rootKey, workflowConfig); err != nil {
		panic(err)
	}
	return workflowConfig
}
