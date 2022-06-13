// Code generated by mockery v1.0.0. DO NOT EDIT.

package run

import (
	context "context"
	"github.com/shiningrush/fastflow/pkg/utils"

	mock "github.com/stretchr/testify/mock"
)

// MockExecuteContext is an autogenerated mock type for the ExecuteContext type
type MockExecuteContext struct {
	mock.Mock
}

// Context provides a mock function with given fields:
func (_m *MockExecuteContext) Context() context.Context {
	ret := _m.Called()

	var r0 context.Context
	if rf, ok := ret.Get(0).(func() context.Context); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(context.Context)
		}
	}

	return r0
}

// GetVar provides a mock function with given fields: varName
func (_m *MockExecuteContext) GetVar(varName string) (string, bool) {
	ret := _m.Called(varName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(varName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(varName)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// IterateVars provides a mock function with given fields: iterateFunc
func (_m *MockExecuteContext) IterateVars(iterateFunc utils.KeyValueIterateFunc) {
	_m.Called(iterateFunc)
}

// ShareData provides a mock function with given fields:
func (_m *MockExecuteContext) ShareData() ShareDataOperator {
	ret := _m.Called()

	var r0 ShareDataOperator
	if rf, ok := ret.Get(0).(func() ShareDataOperator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ShareDataOperator)
		}
	}

	return r0
}

// Trace provides a mock function with given fields: msg, opt
func (_m *MockExecuteContext) Trace(msg string, opt ...TraceOp) {
	_va := make([]interface{}, len(opt))
	for _i := range opt {
		_va[_i] = opt[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// Tracef provides a mock function with given fields: msg, a
func (_m *MockExecuteContext) Tracef(msg string, a ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, msg)
	_ca = append(_ca, a...)
	_m.Called(_ca...)
}

// WithValue provides a mock function with given fields: key, value
func (_m *MockExecuteContext) WithValue(key interface{}, value interface{}) {
	_m.Called(key, value)
}
