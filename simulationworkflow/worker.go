package simulationworkflow

import (
	"bytes"
	"encoding/json"
	"github.com/3dsim/workflow/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/aws/aws-sdk-go/service/swf/swfiface"
	"github.com/inconshreveable/log15"
	"sync"
	"time"
)

// Worker handles all the coordination logic for the workflow.  See
// http://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-dev-workers.html
type Worker interface {
	Start()
}

type worker struct {
	swfAPI   swfiface.SWFAPI
	domain   *string
	taskList *swf.TaskList
	identity *string
	log15.Logger
}

// NewWorker creates a new worker
func NewWorker(taskListName, identity, domain string) Worker {
	worker := &worker{
		swfAPI:   swf.New(NewSession()),
		domain:   aws.String(domain),
		taskList: &swf.TaskList{Name: aws.String(taskListName)},
		identity: aws.String(identity),
	}
	worker.Logger = logger.Log.New("component", "worker", "domain", *worker.domain,
		"taskListName", taskListName, "identity", identity)
	return worker
}

func (d *worker) Start() {
	log.Info("Starting worker")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		d.listenForActivityTasks(d.pollForActivityTasks())
	}()
	wg.Wait()
}

func (d *worker) doWork(activityTask *swf.PollForActivityTaskOutput) error {
	if *activityTask.StartedEventId == 0 {
		return nil
	}
	taskBytes, err := json.Marshal(activityTask)
	if err != nil {
		d.Error("Error marshalling", "err", err)
	}
	d.Info("Activity task received", "task", bytes.NewBuffer(taskBytes))
	params := &swf.RespondActivityTaskCompletedInput{
		TaskToken: activityTask.TaskToken,
		Result:    aws.String("Success"),
	}
	if _, err := d.swfAPI.RespondActivityTaskCompleted(params); err != nil {
		return err
	}
	return nil
}

func (d *worker) pollForActivityTasks() chan *swf.PollForActivityTaskOutput {
	activityTaskChannel := make(chan *swf.PollForActivityTaskOutput)
	go func() {
		for {
			d.Debug("Polling for activity task")
			if activityTask, err := d.pollForActivityTask(); err != nil {
				d.Warn("Error polling for activity task.  Waiting 1 minute before trying again", "error", err)
				time.Sleep(1 * time.Minute)
			} else {
				activityTaskChannel <- activityTask
			}
		}
	}()
	return activityTaskChannel
}

func (d *worker) pollForActivityTask() (*swf.PollForActivityTaskOutput, error) {
	params := &swf.PollForActivityTaskInput{
		Domain:   d.domain,
		TaskList: d.taskList,
		Identity: d.identity,
	}
	return d.swfAPI.PollForActivityTask(params)
}

func (d *worker) listenForActivityTasks(activityTaskChannel <-chan *swf.PollForActivityTaskOutput) {
	for activityTask := range activityTaskChannel {
		if err := d.doWork(activityTask); err != nil {
			// Don't do anything because activity will timeout after the time specified in
			// SupportOptimizationWorkflowWorkerTaskTimeout in config-*.yml files.
			d.Error("Error making activity", "error", err, "activityTask", activityTask)
		}
	}
}
