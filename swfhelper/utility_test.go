package swfhelper

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLastNonDecisionEvent(t *testing.T) {
	// arrange
	// historyEvents := []*swf.HistoryEvent{
	// 	&swf.HistoryEvent{EventId: aws.Int64(1), EventType: aws.String(swf.EventTypeWorkflowExecutionStarted)},
	// 	&swf.HistoryEvent{EventId: aws.Int64(2), EventType: aws.String(swf.EventTypeDecisionTaskScheduled)},
	// 	&swf.HistoryEvent{EventId: aws.Int64(3), EventType: aws.String(swf.EventTypeDecisionTaskStarted)},
	// 	&swf.HistoryEvent{EventId: aws.Int64(4), EventType: aws.String(swf.EventTypeDecisionTaskCompleted)},
	// 	&swf.HistoryEvent{EventId: aws.Int64(5), EventType: aws.String(swf.EventTypeActivityTaskScheduled)},
	// 	&swf.HistoryEvent{EventId: aws.Int64(6), EventType: aws.String(swf.EventTypeActivityTaskStarted)},
	// 	&swf.HistoryEvent{EventId: aws.Int64(7), EventType: aws.String(swf.EventTypeActivityTaskCompleted)},
	// 	&swf.HistoryEvent{EventId: aws.Int64(8), EventType: aws.String(swf.EventTypeDecisionTaskScheduled)},
	// 	&swf.HistoryEvent{EventId: aws.Int64(9), EventType: aws.String(swf.EventTypeDecisionTaskStarted)},
	// }

	// act

	// assert

}

func TestIsWorkflowExecutionJustStartedWhenStartedExpectsTrue(t *testing.T) {
	// arrange
	decisionTask := &swf.PollForDecisionTaskOutput{PreviousStartedEventId: aws.Int64(0), StartedEventId: aws.Int64(3)}

	// act
	isJustStarted := IsWorkflowExecutionJustStarted(decisionTask)

	// assert
	assert.True(t, isJustStarted, "Should return true because the previous started id is 0 and started id is nonzero")
}

func TestIsWorkflowExecutionJustStartedWhenNotStartedExpectsFalse(t *testing.T) {
	// arrange
	decisionTask := &swf.PollForDecisionTaskOutput{PreviousStartedEventId: aws.Int64(0), StartedEventId: aws.Int64(0)}

	// act
	isJustStarted := IsWorkflowExecutionJustStarted(decisionTask)

	// assert
	assert.False(t, isJustStarted, "Should return false because the previous started id is 0 and started id is 0")
}
