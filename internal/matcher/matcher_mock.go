// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/quay/claircore/libvuln/driver (interfaces: Matcher)

// Package matcher is a generated GoMock package.
package matcher

import (
	gomock "github.com/golang/mock/gomock"
	claircore "github.com/quay/claircore"
	driver "github.com/quay/claircore/libvuln/driver"
	reflect "reflect"
)

// MockMatcher is a mock of Matcher interface
type MockMatcher struct {
	ctrl     *gomock.Controller
	recorder *MockMatcherMockRecorder
}

// MockMatcherMockRecorder is the mock recorder for MockMatcher
type MockMatcherMockRecorder struct {
	mock *MockMatcher
}

// NewMockMatcher creates a new mock instance
func NewMockMatcher(ctrl *gomock.Controller) *MockMatcher {
	mock := &MockMatcher{ctrl: ctrl}
	mock.recorder = &MockMatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMatcher) EXPECT() *MockMatcherMockRecorder {
	return m.recorder
}

// Filter mocks base method
func (m *MockMatcher) Filter(arg0 *claircore.IndexRecord) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Filter", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Filter indicates an expected call of Filter
func (mr *MockMatcherMockRecorder) Filter(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockMatcher)(nil).Filter), arg0)
}

// Name mocks base method
func (m *MockMatcher) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockMatcherMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockMatcher)(nil).Name))
}

// Query mocks base method
func (m *MockMatcher) Query() []driver.MatchConstraint {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query")
	ret0, _ := ret[0].([]driver.MatchConstraint)
	return ret0
}

// Query indicates an expected call of Query
func (mr *MockMatcherMockRecorder) Query() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockMatcher)(nil).Query))
}

// Vulnerable mocks base method
func (m *MockMatcher) Vulnerable(arg0 *claircore.IndexRecord, arg1 *claircore.Vulnerability) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Vulnerable", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Vulnerable indicates an expected call of Vulnerable
func (mr *MockMatcherMockRecorder) Vulnerable(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Vulnerable", reflect.TypeOf((*MockMatcher)(nil).Vulnerable), arg0, arg1)
}
