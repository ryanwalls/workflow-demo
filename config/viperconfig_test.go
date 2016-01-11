package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadingMultipleConfigFilesExpectsSuccess(t *testing.T) {
	assert.Equal(t, "1.0", Viper.GetString("SupportOptimizationWorkflow.version"), "Could not read from remote config file")
	assert.Equal(t, "us-west-2", Viper.GetString("AwsRegion"), "Could not read from environment common config file")
	assert.Equal(t, "debug", Viper.GetString("LogLevel"), "Could not read from environment specific config file")
}
