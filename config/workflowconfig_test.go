package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkflowConfigInitialization(t *testing.T) {

	// act

	// assert
	assert.NotNil(t, Workflow, "Expected workflow config to not be nil")
	assert.Equal(t, "1.0", Workflow.Version)
	assert.Equal(t, "SupportOptimizationWorkflowDecider", Workflow.Decider.Name)
	assert.Equal(t, "preproc", Workflow.StepOrder[0])
	assert.Equal(t, "thermalsolver", Workflow.StepOrder[1])
	assert.Equal(t, "mechanicssolver", Workflow.StepOrder[2])
	assert.Equal(t, "Preprocessor", Workflow.Steps[Workflow.StepOrder[0]].Name)
	assert.Equal(t, "NONE", Workflow.Steps[Workflow.StepOrder[0]].ScheduleToCloseTimeout)
}

func TestStepOrdersMatchSteps(t *testing.T) {
	stepOrder := Viper.GetStringSlice("supportOptimizationWorkflow.stepOrder")
	preprocValue := Viper.GetStringMap("supportOptimizationWorkflow.steps." + stepOrder[0])
	assert.Equal(t, "Preprocessor", preprocValue["name"])
}
