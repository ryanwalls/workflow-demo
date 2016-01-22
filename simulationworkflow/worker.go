package simulationworkflow

import (
	"github.com/3dsim/workflow/logger"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/inconshreveable/log15"
	"github.com/sclasen/swfsm/activity"
)

var workerLog log15.Logger

// Worker handles all the coordination logic for the workflow.  See
// http://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-dev-workers.html
type Worker interface {
	Start()
}

// NewWorker creates a new worker
func NewWorker(taskListName, identity, domain string) Worker {
	worker := &activity.ActivityWorker{
		SWF:      swf.New(NewSession()),
		Domain:   domain,
		TaskList: taskListName,
		Identity: identity,
	}
	taskHandler := &activity.ActivityHandler{Activity: "preproc", HandlerFunc: handleTask}
	worker.AddHandler(taskHandler)
	workerLog = logger.Log.New("component", "worker", "domain", domain,
		"taskListName", taskListName, "identity", identity)
	return worker
}

func handleTask(activityTask *swf.PollForActivityTaskOutput, input interface{}) (interface{}, error) {
	workerLog.Info("Received task", "task", activityTask, "input", input)
	return input, nil
}
