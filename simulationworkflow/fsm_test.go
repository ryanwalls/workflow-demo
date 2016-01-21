package simulationworkflow

import (
	"github.com/3dsim/workflow/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/sclasen/swfsm/fsm"
	"github.com/sclasen/swfsm/sugar"
	"github.com/stretchr/testify/assert"

	"time"

	"testing"
)

type TestData struct{}

func TestFSMWhenStartingNewWorkflow(t *testing.T) {
	// arrange
	simulationStateManager := &simulationStateManager{}
	simulationStateManager.FSM = simulationStateManager.setupFSM()
	events := []*swf.HistoryEvent{
		&swf.HistoryEvent{EventType: sugar.S("DecisionTaskStarted"), EventId: sugar.I(3)},
		&swf.HistoryEvent{EventType: sugar.S("DecisionTaskScheduled"), EventId: sugar.I(2)},
		sugar.EventFromPayload(1, &swf.WorkflowExecutionStartedEventAttributes{
			Input: fsm.StartFSMWorkflowInput(simulationStateManager.FSM, new(TestData)),
		}),
	}
	first := testDecisionTask(0, events)

	// act
	_, decisions, _, err := simulationStateManager.Tick(first)

	// assert
	assert.Nil(t, err, "Should be no errors")
	assert.NotNil(t, decisions, "Should have returned some decisions")
	logger.Log.Info("Made some decisions", "decisions", decisions)
	assert.True(t, Find(decisions, schedulePreprocPredicate), "Should have scheduled preproc")
}

func schedulePreprocPredicate(d *swf.Decision) bool {
	return *d.DecisionType == "ScheduleActivityTask" && *d.ScheduleActivityTaskDecisionAttributes.ActivityType.Name == "preproc"
}

func Find(decisions []*swf.Decision, predicate func(*swf.Decision) bool) bool {
	return FindDecision(decisions, predicate) != nil
}

func FindDecision(decisions []*swf.Decision, predicate func(*swf.Decision) bool) *swf.Decision {
	for _, d := range decisions {
		if predicate(d) {
			return d
		}
	}
	return nil
}

func testDecisionTask(prevStarted int, events []*swf.HistoryEvent) *swf.PollForDecisionTaskOutput {

	d := &swf.PollForDecisionTaskOutput{
		Events:                 events,
		PreviousStartedEventId: sugar.I(prevStarted),
		StartedEventId:         sugar.I(prevStarted + len(events)),
		WorkflowExecution:      testWorkflowExecution,
		WorkflowType:           testWorkflowType,
	}
	for i, e := range d.Events {
		if e.EventId == nil {
			e.EventId = sugar.L(*d.StartedEventId - int64(i))
		}
		e.EventTimestamp = aws.Time(time.Unix(0, 0))
		d.Events[i] = e
	}
	return d
}

var testWorkflowExecution = &swf.WorkflowExecution{WorkflowId: sugar.S("workflow-id"), RunId: sugar.S("run-id")}
var testWorkflowType = &swf.WorkflowType{Name: sugar.S("workflow-name"), Version: sugar.S("workflow-version")}
