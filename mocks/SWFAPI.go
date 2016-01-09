package mocks

import "github.com/stretchr/testify/mock"

import "github.com/aws/aws-sdk-go/aws/request"
import "github.com/aws/aws-sdk-go/service/swf"

type SWFAPI struct {
	mock.Mock
}

func (_m *SWFAPI) CountClosedWorkflowExecutionsRequest(_a0 *swf.CountClosedWorkflowExecutionsInput) (*request.Request, *swf.WorkflowExecutionCount) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.CountClosedWorkflowExecutionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.WorkflowExecutionCount
	if rf, ok := ret.Get(1).(func(*swf.CountClosedWorkflowExecutionsInput) *swf.WorkflowExecutionCount); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.WorkflowExecutionCount)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) CountClosedWorkflowExecutions(_a0 *swf.CountClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error) {
	ret := _m.Called(_a0)

	var r0 *swf.WorkflowExecutionCount
	if rf, ok := ret.Get(0).(func(*swf.CountClosedWorkflowExecutionsInput) *swf.WorkflowExecutionCount); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.WorkflowExecutionCount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.CountClosedWorkflowExecutionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) CountOpenWorkflowExecutionsRequest(_a0 *swf.CountOpenWorkflowExecutionsInput) (*request.Request, *swf.WorkflowExecutionCount) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.CountOpenWorkflowExecutionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.WorkflowExecutionCount
	if rf, ok := ret.Get(1).(func(*swf.CountOpenWorkflowExecutionsInput) *swf.WorkflowExecutionCount); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.WorkflowExecutionCount)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) CountOpenWorkflowExecutions(_a0 *swf.CountOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error) {
	ret := _m.Called(_a0)

	var r0 *swf.WorkflowExecutionCount
	if rf, ok := ret.Get(0).(func(*swf.CountOpenWorkflowExecutionsInput) *swf.WorkflowExecutionCount); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.WorkflowExecutionCount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.CountOpenWorkflowExecutionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) CountPendingActivityTasksRequest(_a0 *swf.CountPendingActivityTasksInput) (*request.Request, *swf.PendingTaskCount) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.CountPendingActivityTasksInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.PendingTaskCount
	if rf, ok := ret.Get(1).(func(*swf.CountPendingActivityTasksInput) *swf.PendingTaskCount); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.PendingTaskCount)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) CountPendingActivityTasks(_a0 *swf.CountPendingActivityTasksInput) (*swf.PendingTaskCount, error) {
	ret := _m.Called(_a0)

	var r0 *swf.PendingTaskCount
	if rf, ok := ret.Get(0).(func(*swf.CountPendingActivityTasksInput) *swf.PendingTaskCount); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.PendingTaskCount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.CountPendingActivityTasksInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) CountPendingDecisionTasksRequest(_a0 *swf.CountPendingDecisionTasksInput) (*request.Request, *swf.PendingTaskCount) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.CountPendingDecisionTasksInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.PendingTaskCount
	if rf, ok := ret.Get(1).(func(*swf.CountPendingDecisionTasksInput) *swf.PendingTaskCount); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.PendingTaskCount)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) CountPendingDecisionTasks(_a0 *swf.CountPendingDecisionTasksInput) (*swf.PendingTaskCount, error) {
	ret := _m.Called(_a0)

	var r0 *swf.PendingTaskCount
	if rf, ok := ret.Get(0).(func(*swf.CountPendingDecisionTasksInput) *swf.PendingTaskCount); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.PendingTaskCount)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.CountPendingDecisionTasksInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) DeprecateActivityTypeRequest(_a0 *swf.DeprecateActivityTypeInput) (*request.Request, *swf.DeprecateActivityTypeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.DeprecateActivityTypeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.DeprecateActivityTypeOutput
	if rf, ok := ret.Get(1).(func(*swf.DeprecateActivityTypeInput) *swf.DeprecateActivityTypeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.DeprecateActivityTypeOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) DeprecateActivityType(_a0 *swf.DeprecateActivityTypeInput) (*swf.DeprecateActivityTypeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.DeprecateActivityTypeOutput
	if rf, ok := ret.Get(0).(func(*swf.DeprecateActivityTypeInput) *swf.DeprecateActivityTypeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.DeprecateActivityTypeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.DeprecateActivityTypeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) DeprecateDomainRequest(_a0 *swf.DeprecateDomainInput) (*request.Request, *swf.DeprecateDomainOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.DeprecateDomainInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.DeprecateDomainOutput
	if rf, ok := ret.Get(1).(func(*swf.DeprecateDomainInput) *swf.DeprecateDomainOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.DeprecateDomainOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) DeprecateDomain(_a0 *swf.DeprecateDomainInput) (*swf.DeprecateDomainOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.DeprecateDomainOutput
	if rf, ok := ret.Get(0).(func(*swf.DeprecateDomainInput) *swf.DeprecateDomainOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.DeprecateDomainOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.DeprecateDomainInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) DeprecateWorkflowTypeRequest(_a0 *swf.DeprecateWorkflowTypeInput) (*request.Request, *swf.DeprecateWorkflowTypeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.DeprecateWorkflowTypeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.DeprecateWorkflowTypeOutput
	if rf, ok := ret.Get(1).(func(*swf.DeprecateWorkflowTypeInput) *swf.DeprecateWorkflowTypeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.DeprecateWorkflowTypeOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) DeprecateWorkflowType(_a0 *swf.DeprecateWorkflowTypeInput) (*swf.DeprecateWorkflowTypeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.DeprecateWorkflowTypeOutput
	if rf, ok := ret.Get(0).(func(*swf.DeprecateWorkflowTypeInput) *swf.DeprecateWorkflowTypeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.DeprecateWorkflowTypeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.DeprecateWorkflowTypeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) DescribeActivityTypeRequest(_a0 *swf.DescribeActivityTypeInput) (*request.Request, *swf.DescribeActivityTypeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.DescribeActivityTypeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.DescribeActivityTypeOutput
	if rf, ok := ret.Get(1).(func(*swf.DescribeActivityTypeInput) *swf.DescribeActivityTypeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.DescribeActivityTypeOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) DescribeActivityType(_a0 *swf.DescribeActivityTypeInput) (*swf.DescribeActivityTypeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.DescribeActivityTypeOutput
	if rf, ok := ret.Get(0).(func(*swf.DescribeActivityTypeInput) *swf.DescribeActivityTypeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.DescribeActivityTypeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.DescribeActivityTypeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) DescribeDomainRequest(_a0 *swf.DescribeDomainInput) (*request.Request, *swf.DescribeDomainOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.DescribeDomainInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.DescribeDomainOutput
	if rf, ok := ret.Get(1).(func(*swf.DescribeDomainInput) *swf.DescribeDomainOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.DescribeDomainOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) DescribeDomain(_a0 *swf.DescribeDomainInput) (*swf.DescribeDomainOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.DescribeDomainOutput
	if rf, ok := ret.Get(0).(func(*swf.DescribeDomainInput) *swf.DescribeDomainOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.DescribeDomainOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.DescribeDomainInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) DescribeWorkflowExecutionRequest(_a0 *swf.DescribeWorkflowExecutionInput) (*request.Request, *swf.DescribeWorkflowExecutionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.DescribeWorkflowExecutionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.DescribeWorkflowExecutionOutput
	if rf, ok := ret.Get(1).(func(*swf.DescribeWorkflowExecutionInput) *swf.DescribeWorkflowExecutionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.DescribeWorkflowExecutionOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) DescribeWorkflowExecution(_a0 *swf.DescribeWorkflowExecutionInput) (*swf.DescribeWorkflowExecutionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.DescribeWorkflowExecutionOutput
	if rf, ok := ret.Get(0).(func(*swf.DescribeWorkflowExecutionInput) *swf.DescribeWorkflowExecutionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.DescribeWorkflowExecutionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.DescribeWorkflowExecutionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) DescribeWorkflowTypeRequest(_a0 *swf.DescribeWorkflowTypeInput) (*request.Request, *swf.DescribeWorkflowTypeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.DescribeWorkflowTypeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.DescribeWorkflowTypeOutput
	if rf, ok := ret.Get(1).(func(*swf.DescribeWorkflowTypeInput) *swf.DescribeWorkflowTypeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.DescribeWorkflowTypeOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) DescribeWorkflowType(_a0 *swf.DescribeWorkflowTypeInput) (*swf.DescribeWorkflowTypeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.DescribeWorkflowTypeOutput
	if rf, ok := ret.Get(0).(func(*swf.DescribeWorkflowTypeInput) *swf.DescribeWorkflowTypeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.DescribeWorkflowTypeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.DescribeWorkflowTypeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) GetWorkflowExecutionHistoryRequest(_a0 *swf.GetWorkflowExecutionHistoryInput) (*request.Request, *swf.GetWorkflowExecutionHistoryOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.GetWorkflowExecutionHistoryInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.GetWorkflowExecutionHistoryOutput
	if rf, ok := ret.Get(1).(func(*swf.GetWorkflowExecutionHistoryInput) *swf.GetWorkflowExecutionHistoryOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.GetWorkflowExecutionHistoryOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) GetWorkflowExecutionHistory(_a0 *swf.GetWorkflowExecutionHistoryInput) (*swf.GetWorkflowExecutionHistoryOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.GetWorkflowExecutionHistoryOutput
	if rf, ok := ret.Get(0).(func(*swf.GetWorkflowExecutionHistoryInput) *swf.GetWorkflowExecutionHistoryOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.GetWorkflowExecutionHistoryOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.GetWorkflowExecutionHistoryInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) GetWorkflowExecutionHistoryPages(_a0 *swf.GetWorkflowExecutionHistoryInput, _a1 func(*swf.GetWorkflowExecutionHistoryOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*swf.GetWorkflowExecutionHistoryInput, func(*swf.GetWorkflowExecutionHistoryOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *SWFAPI) ListActivityTypesRequest(_a0 *swf.ListActivityTypesInput) (*request.Request, *swf.ListActivityTypesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.ListActivityTypesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.ListActivityTypesOutput
	if rf, ok := ret.Get(1).(func(*swf.ListActivityTypesInput) *swf.ListActivityTypesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.ListActivityTypesOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) ListActivityTypes(_a0 *swf.ListActivityTypesInput) (*swf.ListActivityTypesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.ListActivityTypesOutput
	if rf, ok := ret.Get(0).(func(*swf.ListActivityTypesInput) *swf.ListActivityTypesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.ListActivityTypesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.ListActivityTypesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) ListActivityTypesPages(_a0 *swf.ListActivityTypesInput, _a1 func(*swf.ListActivityTypesOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*swf.ListActivityTypesInput, func(*swf.ListActivityTypesOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *SWFAPI) ListClosedWorkflowExecutionsRequest(_a0 *swf.ListClosedWorkflowExecutionsInput) (*request.Request, *swf.WorkflowExecutionInfos) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.ListClosedWorkflowExecutionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.WorkflowExecutionInfos
	if rf, ok := ret.Get(1).(func(*swf.ListClosedWorkflowExecutionsInput) *swf.WorkflowExecutionInfos); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.WorkflowExecutionInfos)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) ListClosedWorkflowExecutions(_a0 *swf.ListClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error) {
	ret := _m.Called(_a0)

	var r0 *swf.WorkflowExecutionInfos
	if rf, ok := ret.Get(0).(func(*swf.ListClosedWorkflowExecutionsInput) *swf.WorkflowExecutionInfos); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.WorkflowExecutionInfos)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.ListClosedWorkflowExecutionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) ListClosedWorkflowExecutionsPages(_a0 *swf.ListClosedWorkflowExecutionsInput, _a1 func(*swf.WorkflowExecutionInfos, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*swf.ListClosedWorkflowExecutionsInput, func(*swf.WorkflowExecutionInfos, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *SWFAPI) ListDomainsRequest(_a0 *swf.ListDomainsInput) (*request.Request, *swf.ListDomainsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.ListDomainsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.ListDomainsOutput
	if rf, ok := ret.Get(1).(func(*swf.ListDomainsInput) *swf.ListDomainsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.ListDomainsOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) ListDomains(_a0 *swf.ListDomainsInput) (*swf.ListDomainsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.ListDomainsOutput
	if rf, ok := ret.Get(0).(func(*swf.ListDomainsInput) *swf.ListDomainsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.ListDomainsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.ListDomainsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) ListDomainsPages(_a0 *swf.ListDomainsInput, _a1 func(*swf.ListDomainsOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*swf.ListDomainsInput, func(*swf.ListDomainsOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *SWFAPI) ListOpenWorkflowExecutionsRequest(_a0 *swf.ListOpenWorkflowExecutionsInput) (*request.Request, *swf.WorkflowExecutionInfos) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.ListOpenWorkflowExecutionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.WorkflowExecutionInfos
	if rf, ok := ret.Get(1).(func(*swf.ListOpenWorkflowExecutionsInput) *swf.WorkflowExecutionInfos); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.WorkflowExecutionInfos)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) ListOpenWorkflowExecutions(_a0 *swf.ListOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error) {
	ret := _m.Called(_a0)

	var r0 *swf.WorkflowExecutionInfos
	if rf, ok := ret.Get(0).(func(*swf.ListOpenWorkflowExecutionsInput) *swf.WorkflowExecutionInfos); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.WorkflowExecutionInfos)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.ListOpenWorkflowExecutionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) ListOpenWorkflowExecutionsPages(_a0 *swf.ListOpenWorkflowExecutionsInput, _a1 func(*swf.WorkflowExecutionInfos, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*swf.ListOpenWorkflowExecutionsInput, func(*swf.WorkflowExecutionInfos, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *SWFAPI) ListWorkflowTypesRequest(_a0 *swf.ListWorkflowTypesInput) (*request.Request, *swf.ListWorkflowTypesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.ListWorkflowTypesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.ListWorkflowTypesOutput
	if rf, ok := ret.Get(1).(func(*swf.ListWorkflowTypesInput) *swf.ListWorkflowTypesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.ListWorkflowTypesOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) ListWorkflowTypes(_a0 *swf.ListWorkflowTypesInput) (*swf.ListWorkflowTypesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.ListWorkflowTypesOutput
	if rf, ok := ret.Get(0).(func(*swf.ListWorkflowTypesInput) *swf.ListWorkflowTypesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.ListWorkflowTypesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.ListWorkflowTypesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) ListWorkflowTypesPages(_a0 *swf.ListWorkflowTypesInput, _a1 func(*swf.ListWorkflowTypesOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*swf.ListWorkflowTypesInput, func(*swf.ListWorkflowTypesOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *SWFAPI) PollForActivityTaskRequest(_a0 *swf.PollForActivityTaskInput) (*request.Request, *swf.PollForActivityTaskOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.PollForActivityTaskInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.PollForActivityTaskOutput
	if rf, ok := ret.Get(1).(func(*swf.PollForActivityTaskInput) *swf.PollForActivityTaskOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.PollForActivityTaskOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) PollForActivityTask(_a0 *swf.PollForActivityTaskInput) (*swf.PollForActivityTaskOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.PollForActivityTaskOutput
	if rf, ok := ret.Get(0).(func(*swf.PollForActivityTaskInput) *swf.PollForActivityTaskOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.PollForActivityTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.PollForActivityTaskInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) PollForDecisionTaskRequest(_a0 *swf.PollForDecisionTaskInput) (*request.Request, *swf.PollForDecisionTaskOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.PollForDecisionTaskInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.PollForDecisionTaskOutput
	if rf, ok := ret.Get(1).(func(*swf.PollForDecisionTaskInput) *swf.PollForDecisionTaskOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.PollForDecisionTaskOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) PollForDecisionTask(_a0 *swf.PollForDecisionTaskInput) (*swf.PollForDecisionTaskOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.PollForDecisionTaskOutput
	if rf, ok := ret.Get(0).(func(*swf.PollForDecisionTaskInput) *swf.PollForDecisionTaskOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.PollForDecisionTaskOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.PollForDecisionTaskInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) PollForDecisionTaskPages(_a0 *swf.PollForDecisionTaskInput, _a1 func(*swf.PollForDecisionTaskOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*swf.PollForDecisionTaskInput, func(*swf.PollForDecisionTaskOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
func (_m *SWFAPI) RecordActivityTaskHeartbeatRequest(_a0 *swf.RecordActivityTaskHeartbeatInput) (*request.Request, *swf.RecordActivityTaskHeartbeatOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.RecordActivityTaskHeartbeatInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.RecordActivityTaskHeartbeatOutput
	if rf, ok := ret.Get(1).(func(*swf.RecordActivityTaskHeartbeatInput) *swf.RecordActivityTaskHeartbeatOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.RecordActivityTaskHeartbeatOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) RecordActivityTaskHeartbeat(_a0 *swf.RecordActivityTaskHeartbeatInput) (*swf.RecordActivityTaskHeartbeatOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.RecordActivityTaskHeartbeatOutput
	if rf, ok := ret.Get(0).(func(*swf.RecordActivityTaskHeartbeatInput) *swf.RecordActivityTaskHeartbeatOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.RecordActivityTaskHeartbeatOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.RecordActivityTaskHeartbeatInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) RegisterActivityTypeRequest(_a0 *swf.RegisterActivityTypeInput) (*request.Request, *swf.RegisterActivityTypeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.RegisterActivityTypeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.RegisterActivityTypeOutput
	if rf, ok := ret.Get(1).(func(*swf.RegisterActivityTypeInput) *swf.RegisterActivityTypeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.RegisterActivityTypeOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) RegisterActivityType(_a0 *swf.RegisterActivityTypeInput) (*swf.RegisterActivityTypeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.RegisterActivityTypeOutput
	if rf, ok := ret.Get(0).(func(*swf.RegisterActivityTypeInput) *swf.RegisterActivityTypeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.RegisterActivityTypeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.RegisterActivityTypeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) RegisterDomainRequest(_a0 *swf.RegisterDomainInput) (*request.Request, *swf.RegisterDomainOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.RegisterDomainInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.RegisterDomainOutput
	if rf, ok := ret.Get(1).(func(*swf.RegisterDomainInput) *swf.RegisterDomainOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.RegisterDomainOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) RegisterDomain(_a0 *swf.RegisterDomainInput) (*swf.RegisterDomainOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.RegisterDomainOutput
	if rf, ok := ret.Get(0).(func(*swf.RegisterDomainInput) *swf.RegisterDomainOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.RegisterDomainOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.RegisterDomainInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) RegisterWorkflowTypeRequest(_a0 *swf.RegisterWorkflowTypeInput) (*request.Request, *swf.RegisterWorkflowTypeOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.RegisterWorkflowTypeInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.RegisterWorkflowTypeOutput
	if rf, ok := ret.Get(1).(func(*swf.RegisterWorkflowTypeInput) *swf.RegisterWorkflowTypeOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.RegisterWorkflowTypeOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) RegisterWorkflowType(_a0 *swf.RegisterWorkflowTypeInput) (*swf.RegisterWorkflowTypeOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.RegisterWorkflowTypeOutput
	if rf, ok := ret.Get(0).(func(*swf.RegisterWorkflowTypeInput) *swf.RegisterWorkflowTypeOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.RegisterWorkflowTypeOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.RegisterWorkflowTypeInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) RequestCancelWorkflowExecutionRequest(_a0 *swf.RequestCancelWorkflowExecutionInput) (*request.Request, *swf.RequestCancelWorkflowExecutionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.RequestCancelWorkflowExecutionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.RequestCancelWorkflowExecutionOutput
	if rf, ok := ret.Get(1).(func(*swf.RequestCancelWorkflowExecutionInput) *swf.RequestCancelWorkflowExecutionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.RequestCancelWorkflowExecutionOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) RequestCancelWorkflowExecution(_a0 *swf.RequestCancelWorkflowExecutionInput) (*swf.RequestCancelWorkflowExecutionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.RequestCancelWorkflowExecutionOutput
	if rf, ok := ret.Get(0).(func(*swf.RequestCancelWorkflowExecutionInput) *swf.RequestCancelWorkflowExecutionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.RequestCancelWorkflowExecutionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.RequestCancelWorkflowExecutionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) RespondActivityTaskCanceledRequest(_a0 *swf.RespondActivityTaskCanceledInput) (*request.Request, *swf.RespondActivityTaskCanceledOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.RespondActivityTaskCanceledInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.RespondActivityTaskCanceledOutput
	if rf, ok := ret.Get(1).(func(*swf.RespondActivityTaskCanceledInput) *swf.RespondActivityTaskCanceledOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.RespondActivityTaskCanceledOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) RespondActivityTaskCanceled(_a0 *swf.RespondActivityTaskCanceledInput) (*swf.RespondActivityTaskCanceledOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.RespondActivityTaskCanceledOutput
	if rf, ok := ret.Get(0).(func(*swf.RespondActivityTaskCanceledInput) *swf.RespondActivityTaskCanceledOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.RespondActivityTaskCanceledOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.RespondActivityTaskCanceledInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) RespondActivityTaskCompletedRequest(_a0 *swf.RespondActivityTaskCompletedInput) (*request.Request, *swf.RespondActivityTaskCompletedOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.RespondActivityTaskCompletedInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.RespondActivityTaskCompletedOutput
	if rf, ok := ret.Get(1).(func(*swf.RespondActivityTaskCompletedInput) *swf.RespondActivityTaskCompletedOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.RespondActivityTaskCompletedOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) RespondActivityTaskCompleted(_a0 *swf.RespondActivityTaskCompletedInput) (*swf.RespondActivityTaskCompletedOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.RespondActivityTaskCompletedOutput
	if rf, ok := ret.Get(0).(func(*swf.RespondActivityTaskCompletedInput) *swf.RespondActivityTaskCompletedOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.RespondActivityTaskCompletedOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.RespondActivityTaskCompletedInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) RespondActivityTaskFailedRequest(_a0 *swf.RespondActivityTaskFailedInput) (*request.Request, *swf.RespondActivityTaskFailedOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.RespondActivityTaskFailedInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.RespondActivityTaskFailedOutput
	if rf, ok := ret.Get(1).(func(*swf.RespondActivityTaskFailedInput) *swf.RespondActivityTaskFailedOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.RespondActivityTaskFailedOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) RespondActivityTaskFailed(_a0 *swf.RespondActivityTaskFailedInput) (*swf.RespondActivityTaskFailedOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.RespondActivityTaskFailedOutput
	if rf, ok := ret.Get(0).(func(*swf.RespondActivityTaskFailedInput) *swf.RespondActivityTaskFailedOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.RespondActivityTaskFailedOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.RespondActivityTaskFailedInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) RespondDecisionTaskCompletedRequest(_a0 *swf.RespondDecisionTaskCompletedInput) (*request.Request, *swf.RespondDecisionTaskCompletedOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.RespondDecisionTaskCompletedInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.RespondDecisionTaskCompletedOutput
	if rf, ok := ret.Get(1).(func(*swf.RespondDecisionTaskCompletedInput) *swf.RespondDecisionTaskCompletedOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.RespondDecisionTaskCompletedOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) RespondDecisionTaskCompleted(_a0 *swf.RespondDecisionTaskCompletedInput) (*swf.RespondDecisionTaskCompletedOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.RespondDecisionTaskCompletedOutput
	if rf, ok := ret.Get(0).(func(*swf.RespondDecisionTaskCompletedInput) *swf.RespondDecisionTaskCompletedOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.RespondDecisionTaskCompletedOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.RespondDecisionTaskCompletedInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) SignalWorkflowExecutionRequest(_a0 *swf.SignalWorkflowExecutionInput) (*request.Request, *swf.SignalWorkflowExecutionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.SignalWorkflowExecutionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.SignalWorkflowExecutionOutput
	if rf, ok := ret.Get(1).(func(*swf.SignalWorkflowExecutionInput) *swf.SignalWorkflowExecutionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.SignalWorkflowExecutionOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) SignalWorkflowExecution(_a0 *swf.SignalWorkflowExecutionInput) (*swf.SignalWorkflowExecutionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.SignalWorkflowExecutionOutput
	if rf, ok := ret.Get(0).(func(*swf.SignalWorkflowExecutionInput) *swf.SignalWorkflowExecutionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.SignalWorkflowExecutionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.SignalWorkflowExecutionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) StartWorkflowExecutionRequest(_a0 *swf.StartWorkflowExecutionInput) (*request.Request, *swf.StartWorkflowExecutionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.StartWorkflowExecutionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.StartWorkflowExecutionOutput
	if rf, ok := ret.Get(1).(func(*swf.StartWorkflowExecutionInput) *swf.StartWorkflowExecutionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.StartWorkflowExecutionOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) StartWorkflowExecution(_a0 *swf.StartWorkflowExecutionInput) (*swf.StartWorkflowExecutionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.StartWorkflowExecutionOutput
	if rf, ok := ret.Get(0).(func(*swf.StartWorkflowExecutionInput) *swf.StartWorkflowExecutionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.StartWorkflowExecutionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.StartWorkflowExecutionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
func (_m *SWFAPI) TerminateWorkflowExecutionRequest(_a0 *swf.TerminateWorkflowExecutionInput) (*request.Request, *swf.TerminateWorkflowExecutionOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*swf.TerminateWorkflowExecutionInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *swf.TerminateWorkflowExecutionOutput
	if rf, ok := ret.Get(1).(func(*swf.TerminateWorkflowExecutionInput) *swf.TerminateWorkflowExecutionOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*swf.TerminateWorkflowExecutionOutput)
		}
	}

	return r0, r1
}
func (_m *SWFAPI) TerminateWorkflowExecution(_a0 *swf.TerminateWorkflowExecutionInput) (*swf.TerminateWorkflowExecutionOutput, error) {
	ret := _m.Called(_a0)

	var r0 *swf.TerminateWorkflowExecutionOutput
	if rf, ok := ret.Get(0).(func(*swf.TerminateWorkflowExecutionInput) *swf.TerminateWorkflowExecutionOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*swf.TerminateWorkflowExecutionOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*swf.TerminateWorkflowExecutionInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
