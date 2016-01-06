package decider

import (
	"github.com/3dsim/workflow/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/swf"
)

func New() {

}

var log = logger.Log.New("pkg", "decider")

func RegisterWorkflow() {
	session := session.New(&aws.Config{Region: aws.String("us-west-2")})

	svc := swf.New(session)
	params := &swf.RegisterWorkflowTypeInput{
		Domain:  aws.String("simulation"),                  // Required
		Name:    aws.String("SupportOptimizationWorkflow"), // Required
		Version: aws.String("1.0"),                         // Required
		DefaultTaskList: &swf.TaskList{
			Name: aws.String("mainTaskList"), // Required
		},
	}

	resp, err := svc.RegisterWorkflowType(params)
	if err != nil {
		log.Error("Error registering workflow", "error", err, "params", params)
	}
	log.Info("Successfully registered workflow", "response", resp, "params", params)
}

func RegisterDomain() {
	session := session.New(&aws.Config{Region: aws.String("us-west-2")})

	svc := swf.New(session)

	params := &swf.RegisterDomainInput{
		Name: aws.String("simulation"), // Required
		WorkflowExecutionRetentionPeriodInDays: aws.String("90"), // Required
		Description:                            aws.String("Domain for handling simulation workflows"),
	}
	_, err := svc.RegisterDomain(params)
	if err != nil {
		log.Error("Problem registering domain", "error", err)
	}
}
