// Code generated by MockGen. DO NOT EDIT.
// Source: ./handlers/interface.go

// Package mock_handlers is a generated GoMock package.
package mock_handlers

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entities "github.com/goltsev/task-management/entities"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// ChangeTaskPriority mocks base method.
func (m *MockService) ChangeTaskPriority(tid, prio int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeTaskPriority", tid, prio)
	ret0, _ := ret[0].(error)
	return ret0
}

// ChangeTaskPriority indicates an expected call of ChangeTaskPriority.
func (mr *MockServiceMockRecorder) ChangeTaskPriority(tid, prio interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeTaskPriority", reflect.TypeOf((*MockService)(nil).ChangeTaskPriority), tid, prio)
}

// CreateColumn mocks base method.
func (m *MockService) CreateColumn(pid int, name string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateColumn", pid, name)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateColumn indicates an expected call of CreateColumn.
func (mr *MockServiceMockRecorder) CreateColumn(pid, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateColumn", reflect.TypeOf((*MockService)(nil).CreateColumn), pid, name)
}

// CreateComment mocks base method.
func (m *MockService) CreateComment(tid int, text string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateComment", tid, text)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockServiceMockRecorder) CreateComment(tid, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockService)(nil).CreateComment), tid, text)
}

// CreateProject mocks base method.
func (m *MockService) CreateProject(name, desc string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", name, desc)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockServiceMockRecorder) CreateProject(name, desc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockService)(nil).CreateProject), name, desc)
}

// CreateTask mocks base method.
func (m *MockService) CreateTask(cid int, name, desc string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", cid, name, desc)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockServiceMockRecorder) CreateTask(cid, name, desc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockService)(nil).CreateTask), cid, name, desc)
}

// DeleteColumn mocks base method.
func (m *MockService) DeleteColumn(cid int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteColumn", cid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteColumn indicates an expected call of DeleteColumn.
func (mr *MockServiceMockRecorder) DeleteColumn(cid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteColumn", reflect.TypeOf((*MockService)(nil).DeleteColumn), cid)
}

// DeleteComment mocks base method.
func (m *MockService) DeleteComment(cid int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteComment", cid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteComment indicates an expected call of DeleteComment.
func (mr *MockServiceMockRecorder) DeleteComment(cid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteComment", reflect.TypeOf((*MockService)(nil).DeleteComment), cid)
}

// DeleteProject mocks base method.
func (m *MockService) DeleteProject(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockServiceMockRecorder) DeleteProject(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockService)(nil).DeleteProject), id)
}

// DeleteTask mocks base method.
func (m *MockService) DeleteTask(tid int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", tid)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockServiceMockRecorder) DeleteTask(tid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockService)(nil).DeleteTask), tid)
}

// GetColumn mocks base method.
func (m *MockService) GetColumn(cid int) (*entities.Column, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetColumn", cid)
	ret0, _ := ret[0].(*entities.Column)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetColumn indicates an expected call of GetColumn.
func (mr *MockServiceMockRecorder) GetColumn(cid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetColumn", reflect.TypeOf((*MockService)(nil).GetColumn), cid)
}

// GetColumnList mocks base method.
func (m *MockService) GetColumnList(pid int) ([]*entities.Column, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetColumnList", pid)
	ret0, _ := ret[0].([]*entities.Column)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetColumnList indicates an expected call of GetColumnList.
func (mr *MockServiceMockRecorder) GetColumnList(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetColumnList", reflect.TypeOf((*MockService)(nil).GetColumnList), pid)
}

// GetComment mocks base method.
func (m *MockService) GetComment(cid int) (*entities.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComment", cid)
	ret0, _ := ret[0].(*entities.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment.
func (mr *MockServiceMockRecorder) GetComment(cid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockService)(nil).GetComment), cid)
}

// GetCommentList mocks base method.
func (m *MockService) GetCommentList(tid int) ([]*entities.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentList", tid)
	ret0, _ := ret[0].([]*entities.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentList indicates an expected call of GetCommentList.
func (mr *MockServiceMockRecorder) GetCommentList(tid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentList", reflect.TypeOf((*MockService)(nil).GetCommentList), tid)
}

// GetProject mocks base method.
func (m *MockService) GetProject(id int) (*entities.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProject", id)
	ret0, _ := ret[0].(*entities.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockServiceMockRecorder) GetProject(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockService)(nil).GetProject), id)
}

// GetProjectList mocks base method.
func (m *MockService) GetProjectList() ([]*entities.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectList")
	ret0, _ := ret[0].([]*entities.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectList indicates an expected call of GetProjectList.
func (mr *MockServiceMockRecorder) GetProjectList() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectList", reflect.TypeOf((*MockService)(nil).GetProjectList))
}

// GetProjectTaskList mocks base method.
func (m *MockService) GetProjectTaskList(pid int) ([]*entities.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectTaskList", pid)
	ret0, _ := ret[0].([]*entities.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectTaskList indicates an expected call of GetProjectTaskList.
func (mr *MockServiceMockRecorder) GetProjectTaskList(pid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectTaskList", reflect.TypeOf((*MockService)(nil).GetProjectTaskList), pid)
}

// GetTask mocks base method.
func (m *MockService) GetTask(tid int) (*entities.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTask", tid)
	ret0, _ := ret[0].(*entities.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTask indicates an expected call of GetTask.
func (mr *MockServiceMockRecorder) GetTask(tid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTask", reflect.TypeOf((*MockService)(nil).GetTask), tid)
}

// GetTaskList mocks base method.
func (m *MockService) GetTaskList(cid int) ([]*entities.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskList", cid)
	ret0, _ := ret[0].([]*entities.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskList indicates an expected call of GetTaskList.
func (mr *MockServiceMockRecorder) GetTaskList(cid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskList", reflect.TypeOf((*MockService)(nil).GetTaskList), cid)
}

// MoveColumn mocks base method.
func (m *MockService) MoveColumn(cid, pos int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveColumn", cid, pos)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveColumn indicates an expected call of MoveColumn.
func (mr *MockServiceMockRecorder) MoveColumn(cid, pos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveColumn", reflect.TypeOf((*MockService)(nil).MoveColumn), cid, pos)
}

// MoveTaskTo mocks base method.
func (m *MockService) MoveTaskTo(tid, colpos int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveTaskTo", tid, colpos)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveTaskTo indicates an expected call of MoveTaskTo.
func (mr *MockServiceMockRecorder) MoveTaskTo(tid, colpos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveTaskTo", reflect.TypeOf((*MockService)(nil).MoveTaskTo), tid, colpos)
}

// UpdateColumn mocks base method.
func (m *MockService) UpdateColumn(c *entities.Column) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumn", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateColumn indicates an expected call of UpdateColumn.
func (mr *MockServiceMockRecorder) UpdateColumn(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumn", reflect.TypeOf((*MockService)(nil).UpdateColumn), c)
}

// UpdateColumnInfo mocks base method.
func (m *MockService) UpdateColumnInfo(id int, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateColumnInfo", id, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateColumnInfo indicates an expected call of UpdateColumnInfo.
func (mr *MockServiceMockRecorder) UpdateColumnInfo(id, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateColumnInfo", reflect.TypeOf((*MockService)(nil).UpdateColumnInfo), id, name)
}

// UpdateComment mocks base method.
func (m *MockService) UpdateComment(c *entities.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateComment", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateComment indicates an expected call of UpdateComment.
func (mr *MockServiceMockRecorder) UpdateComment(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateComment", reflect.TypeOf((*MockService)(nil).UpdateComment), c)
}

// UpdateCommentInfo mocks base method.
func (m *MockService) UpdateCommentInfo(cid int, text string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCommentInfo", cid, text)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCommentInfo indicates an expected call of UpdateCommentInfo.
func (mr *MockServiceMockRecorder) UpdateCommentInfo(cid, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCommentInfo", reflect.TypeOf((*MockService)(nil).UpdateCommentInfo), cid, text)
}

// UpdateProject mocks base method.
func (m *MockService) UpdateProject(p *entities.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", p)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockServiceMockRecorder) UpdateProject(p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockService)(nil).UpdateProject), p)
}

// UpdateProjectInfo mocks base method.
func (m *MockService) UpdateProjectInfo(id int, name, desc string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProjectInfo", id, name, desc)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProjectInfo indicates an expected call of UpdateProjectInfo.
func (mr *MockServiceMockRecorder) UpdateProjectInfo(id, name, desc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProjectInfo", reflect.TypeOf((*MockService)(nil).UpdateProjectInfo), id, name, desc)
}

// UpdateTask mocks base method.
func (m *MockService) UpdateTask(t *entities.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", t)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockServiceMockRecorder) UpdateTask(t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockService)(nil).UpdateTask), t)
}

// UpdateTaskInfo mocks base method.
func (m *MockService) UpdateTaskInfo(tid int, name, desc string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTaskInfo", tid, name, desc)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTaskInfo indicates an expected call of UpdateTaskInfo.
func (mr *MockServiceMockRecorder) UpdateTaskInfo(tid, name, desc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTaskInfo", reflect.TypeOf((*MockService)(nil).UpdateTaskInfo), tid, name, desc)
}
