package simulationworkflow

import (
	"errors"
	"github.com/3dsim/workflow/config"
	"github.com/3dsim/workflow/mocks"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/swf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestRegisterSimulationWorkflowWhenFirstRegistrationExpectsSuccess(t *testing.T) {
	// arrange
	mockSWFAPI := &mocks.SWFAPI{}
	registerWorkflowTypeInput := &swf.RegisterWorkflowTypeInput{
		Domain:  aws.String(config.Viper.GetString("env")),
		Name:    aws.String(config.Workflow.Name),
		Version: aws.String(config.Workflow.Version),
		DefaultTaskStartToCloseTimeout: aws.String(config.Workflow.Decider.TaskTimeout),
		DefaultTaskList: &swf.TaskList{
			Name: aws.String(config.Workflow.DefaultTaskList),
		},
	}
	mockSWFAPI.On("RegisterWorkflowType", registerWorkflowTypeInput).Return(nil, nil)
	registrator := &registrator{swfAPI: mockSWFAPI}

	// act
	err := registrator.RegisterSimulationWorkflow()
	if err != nil {
		t.Fatal(err)
	}

	// assert
	mockSWFAPI.AssertCalled(t, "RegisterWorkflowType", mock.AnythingOfType("*swf.RegisterWorkflowTypeInput"))
}

func TestRegisterSimulationWorkflowWhenRegistrationErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	mockSWFAPI := &mocks.SWFAPI{}
	err := errors.New("Some registration error")
	mockSWFAPI.On("RegisterWorkflowType", mock.Anything).Return(nil, err)
	registrator := &registrator{swfAPI: mockSWFAPI}

	// act
	returnedErr := registrator.RegisterSimulationWorkflow()

	// assert
	mockSWFAPI.AssertCalled(t, "RegisterWorkflowType", mock.Anything)
	assert.Equal(t, err, returnedErr, "Should have returned the exact error that was returned when attempting to register")
}

func TestRegisterSimulationWorkflowWhenRegistrationErrorsBecauseAlreadyExistsExpectsSuccess(t *testing.T) {
	// arrange
	mockSWFAPI := &mocks.SWFAPI{}
	err := awserr.New("TypeAlreadyExistsFault", "Type already exists", nil)
	mockSWFAPI.On("RegisterWorkflowType", mock.Anything).Return(nil, err)
	registrator := &registrator{swfAPI: mockSWFAPI}

	// act
	returnedErr := registrator.RegisterSimulationWorkflow()

	// assert
	mockSWFAPI.AssertCalled(t, "RegisterWorkflowType", mock.Anything)
	assert.Nil(t, returnedErr, "Should not have returned an error because the workflow already exists, so we should just continue.")
}

func TestRegisterDomainWhenNotRegisteredExpectsSuccess(t *testing.T) {
	// arrange
	mockSWFAPI := &mocks.SWFAPI{}
	registerDomainInput := &swf.RegisterDomainInput{
		Name: aws.String(config.Viper.GetString("env")),
		WorkflowExecutionRetentionPeriodInDays: aws.String(config.Viper.GetString("DomainExecutionRetentionPeriod")),
	}
	mockSWFAPI.On("RegisterDomain", registerDomainInput).Return(nil, nil)
	registrator := &registrator{swfAPI: mockSWFAPI}

	// act
	err := registrator.RegisterDomain()
	if err != nil {
		t.Fatal(err)
	}

	// assert
	mockSWFAPI.AssertCalled(t, "RegisterDomain", mock.AnythingOfType("*swf.RegisterDomainInput"))
}

func TestRegisterDomainWhenAlreadyRegisteredExpectsSuccess(t *testing.T) {
	// arrange
	mockSWFAPI := &mocks.SWFAPI{}
	err := awserr.New("DomainAlreadyExistsFault", "Type already exists", nil)
	mockSWFAPI.On("RegisterDomain", mock.Anything).Return(nil, err)
	registrator := &registrator{swfAPI: mockSWFAPI}

	// act
	returnedErr := registrator.RegisterDomain()
	if returnedErr != nil {
		t.Fatal(returnedErr)
	}

	// assert
	mockSWFAPI.AssertCalled(t, "RegisterDomain", mock.Anything)
	assert.Nil(t, returnedErr, "Should not have returned an error because the domain already exists, so we should just continue.")
}

func TestRegisterDomainWhenRegistrationErrorsExpectsErrorReturned(t *testing.T) {
	// arrange
	mockSWFAPI := &mocks.SWFAPI{}
	err := awserr.New("SomeOtherFault", "Some other fault", nil)
	mockSWFAPI.On("RegisterDomain", mock.Anything).Return(nil, err)
	registrator := &registrator{swfAPI: mockSWFAPI}

	// act
	returnedErr := registrator.RegisterDomain()

	// assert
	mockSWFAPI.AssertCalled(t, "RegisterDomain", mock.Anything)
	assert.Equal(t, err, returnedErr, "Should have returned the exact error that was returned when attempting to register")
}

func TestInitializeWorkflowWhenNoErrorsExpectsDomainAndWorkflowRegistered(t *testing.T) {
	// arrange
	mockSWFAPI := &mocks.SWFAPI{}
	mockSWFAPI.On("RegisterDomain", mock.Anything).Return(nil, nil)
	mockSWFAPI.On("RegisterWorkflowType", mock.Anything).Return(nil, nil)
	registrator := &registrator{swfAPI: mockSWFAPI}

	// act
	err := registrator.InitializeWorkflow()
	if err != nil {
		t.Fatal(err)
	}

	// assert
	mockSWFAPI.AssertExpectations(t)
}

func TestInitializeWorkflowWhenDomainRegistrationErrorsExpectsDomainRegistrationErrorReturned(t *testing.T) {
	// arrange
	mockSWFAPI := &mocks.SWFAPI{}
	domainErr := errors.New("Some domain registration error")
	mockSWFAPI.On("RegisterDomain", mock.Anything).Return(nil, domainErr)
	registrator := &registrator{swfAPI: mockSWFAPI}

	// act
	err := registrator.InitializeWorkflow()

	// assert
	mockSWFAPI.AssertExpectations(t)
	assert.Equal(t, domainErr, err, "Should have returned the exact error that was returned when attempting to register")
}

func TestInitializeWorkflowWhenWorkflowRegistrationErrorsExpectsWorkflowRegistrationErrorReturned(t *testing.T) {
	// arrange
	mockSWFAPI := &mocks.SWFAPI{}
	workflowErr := errors.New("Some workflow registration error")
	mockSWFAPI.On("RegisterDomain", mock.Anything).Return(nil, nil)
	mockSWFAPI.On("RegisterWorkflowType", mock.Anything).Return(nil, workflowErr)
	registrator := &registrator{swfAPI: mockSWFAPI}

	// act
	err := registrator.InitializeWorkflow()

	// assert
	mockSWFAPI.AssertExpectations(t)
	assert.Equal(t, workflowErr, err, "Should have returned the exact error that was returned when attempting to register")
}
