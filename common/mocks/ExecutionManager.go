// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package mocks

import "github.com/uber/cadence/common/persistence"
import "github.com/stretchr/testify/mock"

type ExecutionManager struct {
	mock.Mock
}

// CreateWorkflowExecution provides a mock function with given fields: request
func (_m *ExecutionManager) CreateWorkflowExecution(request *persistence.CreateWorkflowExecutionRequest) (*persistence.CreateWorkflowExecutionResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.CreateWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(*persistence.CreateWorkflowExecutionRequest) *persistence.CreateWorkflowExecutionResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.CreateWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.CreateWorkflowExecutionRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWorkflowExecution provides a mock function with given fields: request
func (_m *ExecutionManager) GetWorkflowExecution(request *persistence.GetWorkflowExecutionRequest) (*persistence.GetWorkflowExecutionResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.GetWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(*persistence.GetWorkflowExecutionRequest) *persistence.GetWorkflowExecutionResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.GetWorkflowExecutionRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWorkflowExecution provides a mock function with given fields: request
func (_m *ExecutionManager) UpdateWorkflowExecution(request *persistence.UpdateWorkflowExecutionRequest) (*persistence.UpdateWorkflowExecutionResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.UpdateWorkflowExecutionResponse
	if rf, ok := ret.Get(0).(func(*persistence.UpdateWorkflowExecutionRequest) *persistence.UpdateWorkflowExecutionResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.UpdateWorkflowExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.UpdateWorkflowExecutionRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetMutableState provides a mock function with given fields: request
func (_m *ExecutionManager) ResetMutableState(request *persistence.ResetMutableStateRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.ResetMutableStateRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteWorkflowExecution provides a mock function with given fields: request
func (_m *ExecutionManager) DeleteWorkflowExecution(request *persistence.DeleteWorkflowExecutionRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.DeleteWorkflowExecutionRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCurrentExecution provides a mock function with given fields: request
func (_m *ExecutionManager) GetCurrentExecution(request *persistence.GetCurrentExecutionRequest) (*persistence.GetCurrentExecutionResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.GetCurrentExecutionResponse
	if rf, ok := ret.Get(0).(func(*persistence.GetCurrentExecutionRequest) *persistence.GetCurrentExecutionResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetCurrentExecutionResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.GetCurrentExecutionRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransferTasks provides a mock function with given fields: request
func (_m *ExecutionManager) GetTransferTasks(request *persistence.GetTransferTasksRequest) (*persistence.GetTransferTasksResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.GetTransferTasksResponse
	if rf, ok := ret.Get(0).(func(*persistence.GetTransferTasksRequest) *persistence.GetTransferTasksResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetTransferTasksResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.GetTransferTasksRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompleteTransferTask provides a mock function with given fields: request
func (_m *ExecutionManager) CompleteTransferTask(request *persistence.CompleteTransferTaskRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.CompleteTransferTaskRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RangeCompleteTransferTask provides a mock function with given fields: request
func (_m *ExecutionManager) RangeCompleteTransferTask(request *persistence.RangeCompleteTransferTaskRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.RangeCompleteTransferTaskRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetReplicationTasks provides a mock function with given fields: request
func (_m *ExecutionManager) GetReplicationTasks(request *persistence.GetReplicationTasksRequest) (*persistence.GetReplicationTasksResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.GetReplicationTasksResponse
	if rf, ok := ret.Get(0).(func(*persistence.GetReplicationTasksRequest) *persistence.GetReplicationTasksResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetReplicationTasksResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.GetReplicationTasksRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompleteReplicationTask provides a mock function with given fields: request
func (_m *ExecutionManager) CompleteReplicationTask(request *persistence.CompleteReplicationTaskRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.CompleteReplicationTaskRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTimerIndexTasks provides a mock function with given fields: request
func (_m *ExecutionManager) GetTimerIndexTasks(request *persistence.GetTimerIndexTasksRequest) (*persistence.GetTimerIndexTasksResponse, error) {
	ret := _m.Called(request)

	var r0 *persistence.GetTimerIndexTasksResponse
	if rf, ok := ret.Get(0).(func(*persistence.GetTimerIndexTasksRequest) *persistence.GetTimerIndexTasksResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*persistence.GetTimerIndexTasksResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*persistence.GetTimerIndexTasksRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CompleteTimerTask provides a mock function with given fields: request
func (_m *ExecutionManager) CompleteTimerTask(request *persistence.CompleteTimerTaskRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.CompleteTimerTaskRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RangeCompleteTimerTask provides a mock function with given fields: request
func (_m *ExecutionManager) RangeCompleteTimerTask(request *persistence.RangeCompleteTimerTaskRequest) error {
	ret := _m.Called(request)

	var r0 error
	if rf, ok := ret.Get(0).(func(*persistence.RangeCompleteTimerTaskRequest) error); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *ExecutionManager) Close() {
	_m.Called()
}

var _ persistence.ExecutionManager = (*ExecutionManager)(nil)
