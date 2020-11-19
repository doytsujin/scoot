// Code generated by MockGen. DO NOT EDIT.
// Source: scheduler.go

// Package server is a generated GoMock package.
package server

import (
	gomock "github.com/golang/mock/gomock"
	stats "github.com/twitter/scoot/common/stats"
	saga "github.com/twitter/scoot/saga"
	domain "github.com/twitter/scoot/scheduler/domain"
	reflect "reflect"
)

// MockScheduler is a mock of Scheduler interface
type MockScheduler struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulerMockRecorder
}

// MockSchedulerMockRecorder is the mock recorder for MockScheduler
type MockSchedulerMockRecorder struct {
	mock *MockScheduler
}

// NewMockScheduler creates a new mock instance
func NewMockScheduler(ctrl *gomock.Controller) *MockScheduler {
	mock := &MockScheduler{ctrl: ctrl}
	mock.recorder = &MockSchedulerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockScheduler) EXPECT() *MockSchedulerMockRecorder {
	return m.recorder
}

// ScheduleJob mocks base method
func (m *MockScheduler) ScheduleJob(jobDef domain.JobDefinition) (string, error) {
	ret := m.ctrl.Call(m, "ScheduleJob", jobDef)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ScheduleJob indicates an expected call of ScheduleJob
func (mr *MockSchedulerMockRecorder) ScheduleJob(jobDef interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScheduleJob", reflect.TypeOf((*MockScheduler)(nil).ScheduleJob), jobDef)
}

// KillJob mocks base method
func (m *MockScheduler) KillJob(jobId string) error {
	ret := m.ctrl.Call(m, "KillJob", jobId)
	ret0, _ := ret[0].(error)
	return ret0
}

// KillJob indicates an expected call of KillJob
func (mr *MockSchedulerMockRecorder) KillJob(jobId interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KillJob", reflect.TypeOf((*MockScheduler)(nil).KillJob), jobId)
}

// GetSagaCoord mocks base method
func (m *MockScheduler) GetSagaCoord() saga.SagaCoordinator {
	ret := m.ctrl.Call(m, "GetSagaCoord")
	ret0, _ := ret[0].(saga.SagaCoordinator)
	return ret0
}

// GetSagaCoord indicates an expected call of GetSagaCoord
func (mr *MockSchedulerMockRecorder) GetSagaCoord() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSagaCoord", reflect.TypeOf((*MockScheduler)(nil).GetSagaCoord))
}

// OfflineWorker mocks base method
func (m *MockScheduler) OfflineWorker(req domain.OfflineWorkerReq) error {
	ret := m.ctrl.Call(m, "OfflineWorker", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// OfflineWorker indicates an expected call of OfflineWorker
func (mr *MockSchedulerMockRecorder) OfflineWorker(req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfflineWorker", reflect.TypeOf((*MockScheduler)(nil).OfflineWorker), req)
}

// ReinstateWorker mocks base method
func (m *MockScheduler) ReinstateWorker(req domain.ReinstateWorkerReq) error {
	ret := m.ctrl.Call(m, "ReinstateWorker", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReinstateWorker indicates an expected call of ReinstateWorker
func (mr *MockSchedulerMockRecorder) ReinstateWorker(req interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReinstateWorker", reflect.TypeOf((*MockScheduler)(nil).ReinstateWorker), req)
}

// SetSchedulerStatus mocks base method
func (m *MockScheduler) SetSchedulerStatus(maxTasks int) error {
	ret := m.ctrl.Call(m, "SetSchedulerStatus", maxTasks)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSchedulerStatus indicates an expected call of SetSchedulerStatus
func (mr *MockSchedulerMockRecorder) SetSchedulerStatus(maxTasks interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSchedulerStatus", reflect.TypeOf((*MockScheduler)(nil).SetSchedulerStatus), maxTasks)
}

// GetSchedulerStatus mocks base method
func (m *MockScheduler) GetSchedulerStatus() (int, int) {
	ret := m.ctrl.Call(m, "GetSchedulerStatus")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	return ret0, ret1
}

// GetSchedulerStatus indicates an expected call of GetSchedulerStatus
func (mr *MockSchedulerMockRecorder) GetSchedulerStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSchedulerStatus", reflect.TypeOf((*MockScheduler)(nil).GetSchedulerStatus))
}

// GetClassLoadPcts mocks base method
func (m *MockScheduler) GetClassLoadPcts() (map[string]int32, error) {
	ret := m.ctrl.Call(m, "GetClassLoadPcts")
	ret0, _ := ret[0].(map[string]int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClassLoadPcts indicates an expected call of GetClassLoadPcts
func (mr *MockSchedulerMockRecorder) GetClassLoadPcts() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClassLoadPcts", reflect.TypeOf((*MockScheduler)(nil).GetClassLoadPcts))
}

// SetClassLoadPcts mocks base method
func (m *MockScheduler) SetClassLoadPcts(classLoads map[string]int32) error {
	ret := m.ctrl.Call(m, "SetClassLoadPcts", classLoads)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetClassLoadPcts indicates an expected call of SetClassLoadPcts
func (mr *MockSchedulerMockRecorder) SetClassLoadPcts(classLoads interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClassLoadPcts", reflect.TypeOf((*MockScheduler)(nil).SetClassLoadPcts), classLoads)
}

// GetRequestorToClassMap mocks base method
func (m *MockScheduler) GetRequestorToClassMap() (map[string]string, error) {
	ret := m.ctrl.Call(m, "GetRequestorToClassMap")
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRequestorToClassMap indicates an expected call of GetRequestorToClassMap
func (mr *MockSchedulerMockRecorder) GetRequestorToClassMap() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRequestorToClassMap", reflect.TypeOf((*MockScheduler)(nil).GetRequestorToClassMap))
}

// SetRequestorToClassMap mocks base method
func (m *MockScheduler) SetRequestorToClassMap(requestorToClassMap map[string]string) error {
	ret := m.ctrl.Call(m, "SetRequestorToClassMap", requestorToClassMap)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRequestorToClassMap indicates an expected call of SetRequestorToClassMap
func (mr *MockSchedulerMockRecorder) SetRequestorToClassMap(requestorToClassMap interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRequestorToClassMap", reflect.TypeOf((*MockScheduler)(nil).SetRequestorToClassMap), requestorToClassMap)
}

// GetRebalanceMinDuration mocks base method
func (m *MockScheduler) GetRebalanceMinDuration() (int32, error) {
	ret := m.ctrl.Call(m, "GetRebalanceMinDuration")
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRebalanceMinDuration indicates an expected call of GetRebalanceMinDuration
func (mr *MockSchedulerMockRecorder) GetRebalanceMinDuration() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRebalanceMinDuration", reflect.TypeOf((*MockScheduler)(nil).GetRebalanceMinDuration))
}

// SetRebalanceMinDuration mocks base method
func (m *MockScheduler) SetRebalanceMinDuration(durationMin int32) error {
	ret := m.ctrl.Call(m, "SetRebalanceMinDuration", durationMin)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRebalanceMinDuration indicates an expected call of SetRebalanceMinDuration
func (mr *MockSchedulerMockRecorder) SetRebalanceMinDuration(durationMin interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRebalanceMinDuration", reflect.TypeOf((*MockScheduler)(nil).SetRebalanceMinDuration), durationMin)
}

// GetRebalanceThreshold mocks base method
func (m *MockScheduler) GetRebalanceThreshold() (int32, error) {
	ret := m.ctrl.Call(m, "GetRebalanceThreshold")
	ret0, _ := ret[0].(int32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRebalanceThreshold indicates an expected call of GetRebalanceThreshold
func (mr *MockSchedulerMockRecorder) GetRebalanceThreshold() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRebalanceThreshold", reflect.TypeOf((*MockScheduler)(nil).GetRebalanceThreshold))
}

// SetRebalanceThreshold mocks base method
func (m *MockScheduler) SetRebalanceThreshold(durationMin int32) error {
	ret := m.ctrl.Call(m, "SetRebalanceThreshold", durationMin)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetRebalanceThreshold indicates an expected call of SetRebalanceThreshold
func (mr *MockSchedulerMockRecorder) SetRebalanceThreshold(durationMin interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetRebalanceThreshold", reflect.TypeOf((*MockScheduler)(nil).SetRebalanceThreshold), durationMin)
}

// MockSchedulingAlgorithm is a mock of SchedulingAlgorithm interface
type MockSchedulingAlgorithm struct {
	ctrl     *gomock.Controller
	recorder *MockSchedulingAlgorithmMockRecorder
}

// MockSchedulingAlgorithmMockRecorder is the mock recorder for MockSchedulingAlgorithm
type MockSchedulingAlgorithmMockRecorder struct {
	mock *MockSchedulingAlgorithm
}

// NewMockSchedulingAlgorithm creates a new mock instance
func NewMockSchedulingAlgorithm(ctrl *gomock.Controller) *MockSchedulingAlgorithm {
	mock := &MockSchedulingAlgorithm{ctrl: ctrl}
	mock.recorder = &MockSchedulingAlgorithmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSchedulingAlgorithm) EXPECT() *MockSchedulingAlgorithmMockRecorder {
	return m.recorder
}

// GetTasksToBeAssigned mocks base method
func (m *MockSchedulingAlgorithm) GetTasksToBeAssigned(jobs []*jobState, stat stats.StatsReceiver, cs *clusterState, requestors map[string][]*jobState) ([]*taskState, []*taskState) {
	ret := m.ctrl.Call(m, "GetTasksToBeAssigned", jobs, stat, cs, requestors)
	ret0, _ := ret[0].([]*taskState)
	ret1, _ := ret[1].([]*taskState)
	return ret0, ret1
}

// GetTasksToBeAssigned indicates an expected call of GetTasksToBeAssigned
func (mr *MockSchedulingAlgorithmMockRecorder) GetTasksToBeAssigned(jobs, stat, cs, requestors interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksToBeAssigned", reflect.TypeOf((*MockSchedulingAlgorithm)(nil).GetTasksToBeAssigned), jobs, stat, cs, requestors)
}
