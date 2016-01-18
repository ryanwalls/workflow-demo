package simulationworkflow

import (
	"github.com/aws/aws-sdk-go/service/swf"
)

func isWorkflowExecutionJustStarted(decisionTask *swf.PollForDecisionTaskOutput) bool {
	return *decisionTask.PreviousStartedEventId == 0 && *decisionTask.StartedEventId != 0
}
