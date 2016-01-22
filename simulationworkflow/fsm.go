package simulationworkflow

import (
	// "bytes"
	// "encoding/json"
	"github.com/3dsim/workflow/config"
	"github.com/3dsim/workflow/logger"
	// "github.com/3dsim/workflow/swfhelper"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/aws/aws-sdk-go/service/swf/swfiface"
	"github.com/inconshreveable/log15"
	"github.com/sclasen/swfsm/fsm"
	"github.com/sclasen/swfsm/sugar"
	"github.com/twinj/uuid"
	"sync"
)

type SimulationStateManager interface {
	Start()
}

type simulationStateManager struct {
	swfAPI   swfiface.SWFAPI
	domain   *string
	taskList *swf.TaskList
	identity *string
	*fsm.FSM
	log15.Logger
}

// {"CustomerId":1,"SimulationResultId":172,"S3SimulationFolder":"Customer_1/simulation172_2016-01-22",
// "FileLocation":"Customer_1/simulation172_2016-01-22/b308147a-a73c-493f-bf10-478219419057_scanpattern.zip",
// "ZoxFileLocation":"Customer_1/simulation172_2016-01-22/035c0f01-d53a-4b36-bbaa-21650a0e2a64_voxel.zip",
// "CoarseZoxFileLocation":"Customer_1/simulation172_2016-01-22/035c0f01-d53a-4b36-bbaa-21650a0e2a64_voxel.zip",
// "MediumZoxFileLocation":"Customer_1/simulation172_2016-01-22/035c0f01-d53a-4b36-bbaa-21650a0e2a64_voxel.zip",
// "FineZoxFileLocation":"Customer_1/simulation172_2016-01-22/035c0f01-d53a-4b36-bbaa-21650a0e2a64_voxel.zip",
// "sizeX":0.001,"sizeY":0.001,"sizeZ":0.001}
type StateData struct {
	CustomerID         int `json:"CustomerId"`
	SimulationResultID int `json:"SimulationResultId"`
	S3SimulationFolder string
}

func NewSimulationStateManager(taskListName, identity, domain string) SimulationStateManager {
	simulationStateManager := &simulationStateManager{
		swfAPI:   swf.New(NewSession()),
		domain:   aws.String(domain),
		taskList: &swf.TaskList{Name: aws.String(taskListName)},
		identity: aws.String(identity),
	}
	simulationStateManager.FSM = simulationStateManager.setupFSM()
	simulationStateManager.Logger = logger.Log.New("component", "decider", "domain", *simulationStateManager.domain,
		"taskListName", taskListName, "identity", identity)
	return simulationStateManager
}

func (s *simulationStateManager) setupFSM() *fsm.FSM {
	simulationFSM := &fsm.FSM{
		Name:       "example-fsm",
		SWF:        s.swfAPI,
		DataType:   StateData{},
		Identity:   "example-fsm",
		Domain:     "dev",
		TaskList:   config.Workflow.DefaultTaskList,
		Serializer: &fsm.JSONStateSerializer{},
	}
	typed := fsm.Typed(new(StateData))

	startState := &fsm.FSMState{Name: "start", Decider: typed.Decider(waitForStart)}
	preprocState := &fsm.FSMState{Name: "preproc", Decider: typed.Decider(waitForPreproc)}

	simulationFSM.AddInitialState(startState)
	simulationFSM.AddState(preprocState)
	simulationFSM.Init()
	return simulationFSM
}

func (s *simulationStateManager) Start() {
	log.Info("Starting decider")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.FSM.Start()
		// s.startWorkflowExecution()
		for {
		}
	}()
	wg.Wait()
}

var waitForStart = func(f *fsm.FSMContext, h *swf.HistoryEvent, stateData *StateData) fsm.Outcome {
	decisions := f.EmptyDecisions()
	switch *h.EventType {
	case swf.EventTypeWorkflowExecutionStarted:
		logger.Log.Info("Workflow execution started", "stateData", *stateData)
		decision := &swf.Decision{}
		decision.DecisionType = aws.String(swf.DecisionTypeScheduleActivityTask)
		decision.ScheduleActivityTaskDecisionAttributes = &swf.ScheduleActivityTaskDecisionAttributes{
			ActivityId: aws.String(uuid.NewV4().String()), // Required
			ActivityType: &swf.ActivityType{ // Required
				Name:    aws.String(config.Workflow.Steps["preproc"].Name),              // Required
				Version: aws.String(config.Workflow.Steps["preproc"].ProductionVersion), // Required
			},
		}
		// Input: {"CustomerId": 1, "SimulationResultId": 172, "S3SimulationFolder": "Customer_1/simulation172_2016-01-22"}
		decisions = append(decisions, decision)
		return f.Goto("preproc", stateData, decisions)
	}
	//if the event was unexpected just stay here
	return f.Stay(stateData, decisions)
}

var waitForPreproc = func(f *fsm.FSMContext, lastEvent *swf.HistoryEvent, stateData *StateData) fsm.Outcome {
	decisions := f.EmptyDecisions()
	switch *lastEvent.EventType {
	case swf.EventTypeActivityTaskCompleted:
		logger.Log.Info("Previous scheduled activity completed successfully")
		decision := &swf.Decision{}
		decision.DecisionType = aws.String(swf.DecisionTypeCompleteWorkflowExecution)
		decision.CompleteWorkflowExecutionDecisionAttributes = &swf.CompleteWorkflowExecutionDecisionAttributes{
			Result: lastEvent.ActivityTaskCompletedEventAttributes.Result,
		}
		decisions = append(decisions, decision)
	case swf.EventTypeActivityTaskFailed:
		logger.Log.Error("Activity task failed", "lastEvent", lastEvent)
		decision := &swf.Decision{}
		decision.DecisionType = aws.String(swf.DecisionTypeFailWorkflowExecution)
		decision.FailWorkflowExecutionDecisionAttributes = &swf.FailWorkflowExecutionDecisionAttributes{
			Reason: sugar.S("Preproc failed"),
		}
		decisions = append(decisions, decision)
	}
	//if the event was unexpected just stay here
	return f.Stay(stateData, decisions)
}

func (s *simulationStateManager) startWorkflowExecution() error {
	_, err := s.swfAPI.StartWorkflowExecution(&swf.StartWorkflowExecutionInput{
		Domain:                       sugar.S(config.Viper.GetString("env")),
		WorkflowId:                   sugar.S(uuid.NewV4().String()),
		ExecutionStartToCloseTimeout: sugar.S("120"),
		TaskStartToCloseTimeout:      sugar.S("120"),
		ChildPolicy:                  sugar.S(swf.ChildPolicyTerminate),
		//you will have previously regiestered a WorkflowType that this FSM will work.
		WorkflowType: &swf.WorkflowType{Name: sugar.S(config.Workflow.Name), Version: sugar.S(config.Workflow.Version)},
		Input:        fsm.StartFSMWorkflowInput(s.FSM, &StateData{SimulationResultID: 10, CustomerID: 1, S3SimulationFolder: "folder/test"}),
	})
	if err != nil {
		panic(err)
	}
	return nil
}
