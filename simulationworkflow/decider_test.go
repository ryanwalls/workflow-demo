package simulationworkflow

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDecider(t *testing.T) {
	// arrange
	taskListName := "someTaskList"
	identity := "Decider1.0"
	domain := "dev"

	// act
	d := NewDecider(taskListName, identity, domain)

	// assert
	deciderStruct := d.(*decider)
	assert.Equal(t, identity, *deciderStruct.identity)
	assert.Equal(t, taskListName, *deciderStruct.taskList.Name)
	assert.Equal(t, domain, *deciderStruct.domain)
}
