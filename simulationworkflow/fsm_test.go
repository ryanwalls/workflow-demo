package simulationworkflow

import (
	// "github.com/3dsim/workflow/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/sclasen/swfsm/fsm"
	"github.com/sclasen/swfsm/sugar"
	"github.com/stretchr/testify/assert"

	"time"

	"testing"
)

type TestData struct{}

var testActivityInfo = fsm.ActivityInfo{ActivityId: "activityId", ActivityType: &swf.ActivityType{Name: sugar.S("activity"), Version: sugar.S("activityVersion")}}
var testWorkflowExecution = &swf.WorkflowExecution{WorkflowId: sugar.S("workflow-id"), RunId: sugar.S("run-id")}
var testWorkflowType = &swf.WorkflowType{Name: sugar.S("workflow-name"), Version: sugar.S("workflow-version")}

func TestFSMWhenStartingNewWorkflowExpectsPreprocScheduled(t *testing.T) {
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
	firstDecisionTask := testDecisionTask(0, events)

	// act
	_, decisions, _, err := simulationStateManager.Tick(firstDecisionTask)

	// assert
	assert.Nil(t, err, "Should be no errors")
	assert.NotNil(t, decisions, "Should have returned some decisions")
	assert.True(t, Find(decisions, scheduleActivityPredicateFunc("preproc")), "Should have scheduled preproc")
}

func TestFSMWhenInPreprocStateAndPreprocSuccessfullyCompletesExpectsWorkflowComplete(t *testing.T) {
	// arrange
	simulationStateManager := &simulationStateManager{}
	simulationStateManager.FSM = simulationStateManager.setupFSM()

	serializedState := &fsm.SerializedState{}
	serializedState.StateName = "preproc"
	serializedState.StateData = simulationStateManager.FSM.Serialize(&TestData{})
	serializedState.StateVersion = 1

	markerRecordedEvent := sugar.EventFromPayload(5, &swf.MarkerRecordedEventAttributes{
		MarkerName: sugar.S(fsm.StateMarker),
		Details:    sugar.S(simulationStateManager.FSM.Serialize(serializedState)),
	})

	events := []*swf.HistoryEvent{
		&swf.HistoryEvent{EventType: sugar.S("DecisionTaskStarted"), EventId: sugar.I(9)},
		&swf.HistoryEvent{EventType: sugar.S("DecisionTaskScheduled"), EventId: sugar.I(8)},
		sugar.EventFromPayload(7, &swf.ActivityTaskCompletedEventAttributes{
			ScheduledEventId: sugar.I(6),
		}),
		sugar.EventFromPayload(6, &swf.ActivityTaskScheduledEventAttributes{
			ActivityId:   sugar.S(testActivityInfo.ActivityId),
			ActivityType: testActivityInfo.ActivityType,
		}),
		markerRecordedEvent,
	}
	first := testDecisionTask(5, events)

	// act
	_, decisions, _, err := simulationStateManager.Tick(first)

	// assert
	assert.Nil(t, err, "Should be no errors")
	assert.NotNil(t, decisions, "Should have returned some decisions")
	assert.True(t, Find(decisions, workflowCompletedPredicate), "Should have completed the workflow")
}

func TestFSMWhenInPreprocStateAndPreprocFailsExpectsWorkflowFailed(t *testing.T) {
	// arrange
	simulationStateManager := &simulationStateManager{}
	simulationStateManager.FSM = simulationStateManager.setupFSM()

	serializedState := &fsm.SerializedState{}
	serializedState.StateName = "preproc"
	serializedState.StateData = simulationStateManager.FSM.Serialize(&TestData{})
	serializedState.StateVersion = 1

	markerRecordedEvent := sugar.EventFromPayload(5, &swf.MarkerRecordedEventAttributes{
		MarkerName: sugar.S(fsm.StateMarker),
		Details:    sugar.S(simulationStateManager.FSM.Serialize(serializedState)),
	})

	events := []*swf.HistoryEvent{
		&swf.HistoryEvent{EventType: sugar.S("DecisionTaskStarted"), EventId: sugar.I(9)},
		&swf.HistoryEvent{EventType: sugar.S("DecisionTaskScheduled"), EventId: sugar.I(8)},
		sugar.EventFromPayload(7, &swf.ActivityTaskFailedEventAttributes{
			ScheduledEventId: sugar.I(6),
		}),
		sugar.EventFromPayload(6, &swf.ActivityTaskScheduledEventAttributes{
			ActivityId:   sugar.S(testActivityInfo.ActivityId),
			ActivityType: testActivityInfo.ActivityType,
		}),
		markerRecordedEvent,
	}
	first := testDecisionTask(5, events)

	// act
	_, decisions, _, err := simulationStateManager.Tick(first)

	// assert
	assert.Nil(t, err, "Should be no errors")
	assert.NotNil(t, decisions, "Should have returned some decisions")
	assert.True(t, Find(decisions, workflowCancelledPredicate), "Should have failed the workflow because preproc failed")
}

func workflowCompletedPredicate(d *swf.Decision) bool {
	return *d.DecisionType == swf.DecisionTypeCompleteWorkflowExecution
}

func workflowCancelledPredicate(d *swf.Decision) bool {
	return *d.DecisionType == swf.DecisionTypeFailWorkflowExecution
}

func scheduleActivityPredicateFunc(activityName string) func(*swf.Decision) bool {
	return func(d *swf.Decision) bool {
		return *d.DecisionType == swf.DecisionTypeScheduleActivityTask && *d.ScheduleActivityTaskDecisionAttributes.ActivityType.Name == activityName
	}
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

func DecisionsToEvents(decisions []*swf.Decision) []*swf.HistoryEvent {
	var events []*swf.HistoryEvent
	for _, d := range decisions {
		if scheduleActivityPredicate(d) {
			event := &swf.HistoryEvent{
				EventType: sugar.S("ActivityTaskCompleted"),
				EventId:   sugar.I(7),
				ActivityTaskCompletedEventAttributes: &swf.ActivityTaskCompletedEventAttributes{
					ScheduledEventId: sugar.I(6),
				},
			}
			events = append(events, event)
			event = &swf.HistoryEvent{
				EventType: sugar.S("ActivityTaskScheduled"),
				EventId:   sugar.I(6),
				ActivityTaskScheduledEventAttributes: &swf.ActivityTaskScheduledEventAttributes{
					ActivityId:   sugar.S(testActivityInfo.ActivityId),
					ActivityType: testActivityInfo.ActivityType,
				},
			}
			events = append(events, event)
		}
		if stateMarkerPredicate(d) {
			event := &swf.HistoryEvent{
				EventType: sugar.S("MarkerRecorded"),
				EventId:   sugar.I(5),
				MarkerRecordedEventAttributes: &swf.MarkerRecordedEventAttributes{
					MarkerName: sugar.S(fsm.StateMarker),
					Details:    d.RecordMarkerDecisionAttributes.Details,
				},
			}
			events = append(events, event)

		}
	}
	return events
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

func scheduleActivityPredicate(d *swf.Decision) bool {
	return *d.DecisionType == "ScheduleActivityTask"
}

func stateMarkerPredicate(d *swf.Decision) bool {
	return *d.DecisionType == "RecordMarker" && *d.RecordMarkerDecisionAttributes.MarkerName == fsm.StateMarker
}
