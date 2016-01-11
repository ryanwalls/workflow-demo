package simulationworkflow

import (
	"errors"
	"github.com/3dsim/workflow/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/aws/aws-sdk-go/service/swf/swfiface"
	"github.com/inconshreveable/log15"
	"sync"
	"time"
)

// Decider handles all the coordination logic for the workflow.  See
// http://docs.aws.amazon.com/amazonswf/latest/developerguide/swf-dg-dev-deciders.html
type Decider interface {
	Start()
}

type decider struct {
	swfAPI   swfiface.SWFAPI
	domain   *string
	taskList *swf.TaskList
	identity *string
	log15.Logger
}

// NewDecider creates a new decider
func NewDecider(taskListName, identity, domain string) Decider {
	decider := &decider{
		swfAPI:   swf.New(NewSession()),
		domain:   aws.String(domain),
		taskList: &swf.TaskList{Name: aws.String(taskListName)},
		identity: aws.String(identity),
	}
	decider.Logger = logger.Log.New("component", "decider", "domain", *decider.domain,
		"taskListName", taskListName, "identity", identity)
	return decider
}

func (d *decider) Start() {
	log.Info("Starting decider")
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		d.listenForDecisionTasks(d.pollForDecisionTasks())
	}()
	wg.Wait()
}

func (d *decider) decide(decisionTask *swf.PollForDecisionTaskOutput) error {
	return errors.New("Not implemented yet")
}

func (d *decider) pollForDecisionTasks() chan *swf.PollForDecisionTaskOutput {
	decisionTaskChannel := make(chan *swf.PollForDecisionTaskOutput)
	go func() {
		for {
			d.Debug("Polling for decision task")
			if decisionTask, err := d.pollForDecisionTask(); err != nil {
				d.Warn("Error polling for decision task.  Waiting 1 minute before trying again", "error", err)
				time.Sleep(1 * time.Minute)
			} else {
				decisionTaskChannel <- decisionTask
			}
		}
	}()
	return decisionTaskChannel
}

func (d *decider) pollForDecisionTask() (*swf.PollForDecisionTaskOutput, error) {
	params := &swf.PollForDecisionTaskInput{
		Domain:   d.domain,
		TaskList: d.taskList,
		Identity: d.identity,
	}
	return d.swfAPI.PollForDecisionTask(params)
}

func (d *decider) listenForDecisionTasks(decisionTaskChannel <-chan *swf.PollForDecisionTaskOutput) {
	for decisionTask := range decisionTaskChannel {
		d.Info("Handling decision", "decision", decisionTask)
		if err := d.decide(decisionTask); err != nil {
			// Don't do anything because decision will timeout after the time specified in
			// SupportOptimizationWorkflowDeciderTaskTimeout in config-*.yml files.
			d.Error("Error making decision", "error", err, "decisionTask", decisionTask)
		}
	}
}
