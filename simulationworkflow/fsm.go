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
		Domain:     "dev",
		TaskList:   config.Workflow.DefaultTaskList,
		Serializer: &fsm.JSONStateSerializer{},
	}
	typed := fsm.Typed(new(StateData))

	startState := &fsm.FSMState{Name: "start", Decider: typed.Decider(waitForStart)}
	preprocState := &fsm.FSMState{Name: "preproc", Decider: fsm.CompleteWorkflow()}

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
			// TODO this should be a unique ID see http://stackoverflow.com/questions/28992136/what-is-the-activity-id-in-aws-swf
			ActivityId: aws.String("3"), // Required
			// TODO don't hard code these values
			ActivityType: &swf.ActivityType{ // Required
				Name:    aws.String("preproc"), // Required
				Version: aws.String("1.0"),     // Required
			},
			ScheduleToCloseTimeout: aws.String("NONE"),
			ScheduleToStartTimeout: aws.String("NONE"),
			StartToCloseTimeout:    aws.String("NONE"),
		}
		decisions = append(decisions, decision)
		return f.Goto("preproc", stateData, decisions)
	}
	//if the event was unexpected just stay here
	return f.Stay(stateData, decisions)
}

func (s *simulationStateManager) startWorkflowExecution() error {
	_, err := s.swfAPI.StartWorkflowExecution(&swf.StartWorkflowExecutionInput{
		Domain:                       sugar.S(config.Viper.GetString("env")),
		WorkflowId:                   sugar.S(uuid.NewV4().String()),
		ExecutionStartToCloseTimeout: sugar.S("60"),
		TaskStartToCloseTimeout:      sugar.S("60"),
		ChildPolicy:                  sugar.S(swf.ChildPolicyTerminate),
		//you will have previously regiestered a WorkflowType that this FSM will work.
		WorkflowType: &swf.WorkflowType{Name: sugar.S(config.Workflow.Name), Version: sugar.S(config.Workflow.Version)},
		Input:        fsm.StartFSMWorkflowInput(s.FSM, &StateData{Count: 10, Message: "start"}),
	})
	if err != nil {
		panic(err)
	}
	return nil
}
