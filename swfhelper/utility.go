package swfhelper

import (
	"github.com/aws/aws-sdk-go/service/swf"
)

func getLastNonDecisionEvent(decisionTask *swf.PollForDecisionTaskOutput) *swf.HistoryEvent {
	return nil
}

func IsWorkflowExecutionJustStarted(decisionTask *swf.PollForDecisionTaskOutput) bool {
	return *decisionTask.PreviousStartedEventId == 0 && *decisionTask.StartedEventId != 0
}
