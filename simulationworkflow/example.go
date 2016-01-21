package simulationworkflow

import (
	// "bytes"
	// "encoding/json"

	// "github.com/3dsim/workflow/swfhelper"
	"github.com/3dsim/workflow/config"
	"github.com/3dsim/workflow/logger"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/sclasen/swfsm/fsm"
	"github.com/sclasen/swfsm/sugar"

	"strconv"
	"time"
)

// create with swf.NewClient
var client *swf.SWF = swf.New(NewSession())

// data type that will be managed by the FSM
type StateData struct {
	Message string `json:"message,omitempty"`
	Count   int    `json:"count,omitempty"`
}

//event type that will be signalled to the FSM with signal name "hello"
type Hello struct {
	Message string `json:"message,omitempty"`
}

func Example() {
	//This is an example of building Deciders without using decider composition.
	//the FSM we will create will oscillate between 2 states,
	//waitForSignal -> will wait till the workflow is started or signalled, and update the StateData based on the Hello message received, set a timer, and transition to waitForTimer
	//waitForTimer -> will wait till the timer set by waitForSignal fires, and will signal the workflow with a Hello message, and transition to waitFotSignal
	waitForSignal := func(f *fsm.FSMContext, h *swf.HistoryEvent, d *StateData) fsm.Outcome {
		logger.Log.Info("In waitForSignal")
		decisions := f.EmptyDecisions()
		switch *h.EventType {
		case swf.EventTypeWorkflowExecutionStarted, swf.EventTypeWorkflowExecutionSignaled:
			if *h.EventType == swf.EventTypeWorkflowExecutionSignaled && *h.WorkflowExecutionSignaledEventAttributes.SignalName == "hello" {
				hello := &Hello{}
				f.EventData(h, &Hello{})
				d.Count++
				d.Message = hello.Message
			}
			timeoutSeconds := "5" //swf uses stringy numbers in many places
			timerDecision := &swf.Decision{
				DecisionType: sugar.S(swf.DecisionTypeStartTimer),
				StartTimerDecisionAttributes: &swf.StartTimerDecisionAttributes{
					StartToFireTimeout: sugar.S(timeoutSeconds),
					TimerId:            sugar.S("timeToSignal"),
				},
			}
			decisions = append(decisions, timerDecision)
			return f.Goto("waitForTimer", d, decisions)
		}
		//if the event was unexpected just stay here
		return f.Stay(d, decisions)

	}

	waitForTimer := func(f *fsm.FSMContext, h *swf.HistoryEvent, d *StateData) fsm.Outcome {
		logger.Log.Info("In waitForTimer")
		decisions := f.EmptyDecisions()
		switch *h.EventType {
		case swf.EventTypeTimerFired:
			//every time the timer fires, signal the workflow with a Hello
			message := strconv.FormatInt(time.Now().Unix(), 10)
			signalInput := &Hello{message}
			signalDecision := &swf.Decision{
				DecisionType: sugar.S(swf.DecisionTypeSignalExternalWorkflowExecution),
				SignalExternalWorkflowExecutionDecisionAttributes: &swf.SignalExternalWorkflowExecutionDecisionAttributes{
					SignalName: sugar.S("hello"),
					Input:      sugar.S(f.Serialize(signalInput)),
					RunId:      f.RunId,
					WorkflowId: f.WorkflowId,
				},
			}
			decisions = append(decisions, signalDecision)

			return f.Goto("waitForSignal", d, decisions)
		}
		//if the event was unexpected just stay here
		return f.Stay(d, decisions)
	}

	//These 2 deciders are the same as the ones above, but use composable decider bits.
	typed := fsm.Typed(new(StateData))

	// updateState := typed.StateFunc(func(f *fsm.FSMContext, h *swf.HistoryEvent, d *StateData) {
	// 	hello := &Hello{}
	// 	f.EventData(h, &Hello{})
	// 	d.Count++
	// 	d.Message = hello.Message
	// })
	//
	// setTimer := typed.DecisionFunc(func(f *fsm.FSMContext, h *swf.HistoryEvent, d *StateData) *swf.Decision {
	// 	logger.Log.Info("In setTimer")
	// 	timeoutSeconds := "5" //swf uses stringy numbers in many places
	// 	return &swf.Decision{
	// 		DecisionType: sugar.S(swf.DecisionTypeStartTimer),
	// 		StartTimerDecisionAttributes: &swf.StartTimerDecisionAttributes{
	// 			StartToFireTimeout: sugar.S(timeoutSeconds),
	// 			TimerId:            sugar.S("timeToSignal"),
	// 		},
	// 	}
	// })
	//
	// sendSignal := typed.DecisionFunc(func(f *fsm.FSMContext, h *swf.HistoryEvent, d *StateData) *swf.Decision {
	// 	logger.Log.Info("In sendSignal")
	// 	message := strconv.FormatInt(time.Now().Unix(), 10)
	// 	signalInput := &Hello{message}
	// 	return &swf.Decision{
	// 		DecisionType: sugar.S(swf.DecisionTypeSignalExternalWorkflowExecution),
	// 		SignalExternalWorkflowExecutionDecisionAttributes: &swf.SignalExternalWorkflowExecutionDecisionAttributes{
	// 			SignalName: sugar.S("hello"),
	// 			Input:      sugar.S(f.Serialize(signalInput)),
	// 			RunId:      f.RunId,
	// 			WorkflowId: f.WorkflowId,
	// 		},
	// 	}
	// })

	// waitForSignalComposedDecider := fsm.NewComposedDecider(
	// 	fsm.OnStarted(fsm.UpdateState(updateState), fsm.AddDecision(setTimer), fsm.Transition("waitForTimer")),
	// 	fsm.OnSignalReceived("hello", fsm.UpdateState(updateState), fsm.AddDecision(setTimer), fsm.Transition("waitForTimer")),
	// 	fsm.DefaultDecider(),
	// )
	//
	// waitForTimerComposedDecider := fsm.NewComposedDecider(
	// 	fsm.OnTimerFired("timeToSignal", fsm.AddDecision(sendSignal), fsm.Transition("waitForSignal")),
	// 	fsm.DefaultDecider(),
	// )

	//create the fsm.FSMState by passing the decider function through TypedDecider(),
	//which lets you use d *StateData rather than d interface{} in your decider.
	waitForSignalState := &fsm.FSMState{Name: "waitForSignal", Decider: typed.Decider(waitForSignal)}
	waitForTimerState := &fsm.FSMState{Name: "waitForTimer", Decider: typed.Decider(waitForTimer)}
	//or with the composed deciders
	// waitForSignalState = &fsm.FSMState{Name: "waitForSignal", Decider: waitForSignalComposedDecider}
	// waitForTimerState = &fsm.FSMState{Name: "waitForTimer", Decider: waitForTimerComposedDecider}
	//wire it up in an fsm
	myFSM := &fsm.FSM{
		Name:       "example-fsm-decider",
		SWF:        client,
		DataType:   StateData{},
		Domain:     config.Viper.GetString("env"),
		TaskList:   "defaultTaskList",
		Serializer: &fsm.JSONStateSerializer{},
	}
	//add states to FSM
	myFSM.AddInitialState(waitForSignalState)
	myFSM.AddState(waitForTimerState)

	// To start workflows using this fsm
	logger.Log.Info("Starting workflow execution")
	_, err := client.StartWorkflowExecution(&swf.StartWorkflowExecutionInput{
		Domain:                       sugar.S("dev"),
		WorkflowId:                   sugar.S("130"),
		ExecutionStartToCloseTimeout: sugar.S("60"),
		TaskStartToCloseTimeout:      sugar.S("60"),
		ChildPolicy:                  sugar.S(swf.ChildPolicyTerminate),
		//you will have previously regiestered a WorkflowType that this FSM will work.
		WorkflowType: &swf.WorkflowType{Name: sugar.S(config.Workflow.Name), Version: sugar.S(config.Workflow.Version)},
		Input:        fsm.StartFSMWorkflowInput(myFSM, &StateData{Count: 10, Message: "start"}),
	})
	if err != nil {
		panic(err)
	}

	//start it up!
	myFSM.Start()

	for {
	}
}
