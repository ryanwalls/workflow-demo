package simulationworkflow

import (
	"github.com/3dsim/workflow/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDecider(t *testing.T) {
	// arrange

	// act
	d := NewDecider()

	// assert
	deciderStruct := d.(*decider)
	assert.Equal(t, config.Viper.GetString("SupportOptimizationWorkflowDeciderIdentity"), *deciderStruct.identity)
	assert.Equal(t, config.Viper.GetString("SupportOptimizationWorkflowDefaultTaskList"), *deciderStruct.taskList.Name)
	assert.Equal(t, config.Viper.GetString("env"), *deciderStruct.domain)
}
