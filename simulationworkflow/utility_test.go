package simulationworkflow

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsWorkflowExecutionJustStartedWhenStartedExpectsTrue(t *testing.T) {
	// arrange
	decisionTask := &swf.PollForDecisionTaskOutput{PreviousStartedEventId: aws.Int64(0), StartedEventId: aws.Int64(3)}

	// act
	isJustStarted := isWorkflowExecutionJustStarted(decisionTask)

	// assert
	assert.True(t, isJustStarted, "Should return true because the previous started id is 0 and started id is nonzero")
}

func TestIsWorkflowExecutionJustStartedWhenNotStartedExpectsFalse(t *testing.T) {
	// arrange
	decisionTask := &swf.PollForDecisionTaskOutput{PreviousStartedEventId: aws.Int64(0), StartedEventId: aws.Int64(0)}

	// act
	isJustStarted := isWorkflowExecutionJustStarted(decisionTask)

	// assert
	assert.False(t, isJustStarted, "Should return false because the previous started id is 0 and started id is 0")
}
