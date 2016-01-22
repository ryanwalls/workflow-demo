package simulationworkflow

import (
	"github.com/3dsim/workflow/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// NewSession creates a new session for the workflow in the appropriate region
func NewSession() *session.Session {
	return session.New(&aws.Config{LogLevel: aws.LogLevel(aws.LogDebugWithRequestErrors), Region: aws.String(config.Viper.GetString("AwsRegion"))})
}
