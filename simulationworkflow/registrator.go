package simulationworkflow

import (
	"github.com/3dsim/workflow/config"
	"github.com/3dsim/workflow/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/aws/aws-sdk-go/service/swf/swfiface"
)

var log = logger.Log.New("pkg", "simulationworkflow")

// Registrator handles registration of domains and the simulation workflow itself
type Registrator interface {
	RegisterSimulationWorkflow() error
	RegisterDomain() error
	InitializeWorkflow() error
	RegisterActivity() error
}

type registrator struct {
	swfAPI swfiface.SWFAPI
}

// NewRegistrator instantiates registrator
func NewRegistrator() Registrator {
	return &registrator{swfAPI: swf.New(NewSession())}
}

func (r *registrator) InitializeWorkflow() error {
	if err := r.RegisterDomain(); err != nil {
		return err
	}
	if err := r.RegisterSimulationWorkflow(); err != nil {
		return err
	}
	return nil
}

func (r *registrator) RegisterSimulationWorkflow() error {
	params := &swf.RegisterWorkflowTypeInput{
		Domain:  aws.String(config.Viper.GetString("env")),
		Name:    aws.String(config.Workflow.Name),
		Version: aws.String(config.Workflow.Version),
		DefaultTaskStartToCloseTimeout: aws.String(config.Workflow.Decider.TaskTimeout),
		DefaultTaskList: &swf.TaskList{
			Name: aws.String(config.Workflow.DefaultTaskList),
		},
	}
	resp, err := r.swfAPI.RegisterWorkflowType(params)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && "TypeAlreadyExistsFault" == awsErr.Code() {
			log.Info("Workflow already registered", "name", *params.Name, "version", *params.Version)
			return nil
		}
		return err
	}
	log.Info("Successfully registered workflow", "response", resp, "name", *params.Name, "version", *params.Version)
	return nil
}

func (r *registrator) RegisterDomain() error {
	params := &swf.RegisterDomainInput{
		Name: aws.String(config.Viper.GetString("env")),
		WorkflowExecutionRetentionPeriodInDays: aws.String(config.Viper.GetString("DomainExecutionRetentionPeriod")),
	}
	resp, err := r.swfAPI.RegisterDomain(params)
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok && "DomainAlreadyExistsFault" == awsErr.Code() {
			log.Info("Domain already registered", "name", *params.Name)
			return nil
		}
		return err
	}
	log.Info("Successfully registered domain", "response", resp, "name", *params.Name)
	return nil
}

func (r *registrator) RegisterActivity() error {
	params := &swf.RegisterActivityTypeInput{
		Domain:                      aws.String(config.Viper.GetString("env")),
		Name:                        aws.String(config.Workflow.Steps["preproc"].Name),
		Version:                     aws.String(config.Workflow.Steps["preproc"].ProductionVersion),
		DefaultTaskHeartbeatTimeout: aws.String(config.Workflow.Steps["preproc"].DefaultTaskHeartbeatTimeout),
		DefaultTaskList:             &swf.TaskList{},
	}
	params.DefaultTaskList.Name = aws.String(*params.Name + "-" + *params.Version)

	_, err := r.swfAPI.RegisterActivityType(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		if awsErr, ok := err.(awserr.Error); ok {
			log.Info("Received an awsErr", "code", awsErr.Code(), "message", awsErr.Message(), "error", awsErr.Error())
			return nil
		}
		return err
	}

	return nil
}
