// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/inovex/CalendarSync/internal/models"
	mock "github.com/stretchr/testify/mock"

	time "time"
)

// Sink is an autogenerated mock type for the Sink type
type Sink struct {
	mock.Mock
}

type Sink_Expecter struct {
	mock *mock.Mock
}

func (_m *Sink) EXPECT() *Sink_Expecter {
	return &Sink_Expecter{mock: &_m.Mock}
}

// CreateEvent provides a mock function with given fields: ctx, e
func (_m *Sink) CreateEvent(ctx context.Context, e models.Event) error {
	ret := _m.Called(ctx, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Event) error); ok {
		r0 = rf(ctx, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sink_CreateEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateEvent'
type Sink_CreateEvent_Call struct {
	*mock.Call
}

// CreateEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - e models.Event
func (_e *Sink_Expecter) CreateEvent(ctx interface{}, e interface{}) *Sink_CreateEvent_Call {
	return &Sink_CreateEvent_Call{Call: _e.mock.On("CreateEvent", ctx, e)}
}

func (_c *Sink_CreateEvent_Call) Run(run func(ctx context.Context, e models.Event)) *Sink_CreateEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Event))
	})
	return _c
}

func (_c *Sink_CreateEvent_Call) Return(_a0 error) *Sink_CreateEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Sink_CreateEvent_Call) RunAndReturn(run func(context.Context, models.Event) error) *Sink_CreateEvent_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteEvent provides a mock function with given fields: ctx, e
func (_m *Sink) DeleteEvent(ctx context.Context, e models.Event) error {
	ret := _m.Called(ctx, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Event) error); ok {
		r0 = rf(ctx, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sink_DeleteEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteEvent'
type Sink_DeleteEvent_Call struct {
	*mock.Call
}

// DeleteEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - e models.Event
func (_e *Sink_Expecter) DeleteEvent(ctx interface{}, e interface{}) *Sink_DeleteEvent_Call {
	return &Sink_DeleteEvent_Call{Call: _e.mock.On("DeleteEvent", ctx, e)}
}

func (_c *Sink_DeleteEvent_Call) Run(run func(ctx context.Context, e models.Event)) *Sink_DeleteEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Event))
	})
	return _c
}

func (_c *Sink_DeleteEvent_Call) Return(_a0 error) *Sink_DeleteEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Sink_DeleteEvent_Call) RunAndReturn(run func(context.Context, models.Event) error) *Sink_DeleteEvent_Call {
	_c.Call.Return(run)
	return _c
}

// EventsInTimeframe provides a mock function with given fields: ctx, start, end
func (_m *Sink) EventsInTimeframe(ctx context.Context, start time.Time, end time.Time) ([]models.Event, error) {
	ret := _m.Called(ctx, start, end)

	var r0 []models.Event
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) ([]models.Event, error)); ok {
		return rf(ctx, start, end)
	}
	if rf, ok := ret.Get(0).(func(context.Context, time.Time, time.Time) []models.Event); ok {
		r0 = rf(ctx, start, end)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Event)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, time.Time, time.Time) error); ok {
		r1 = rf(ctx, start, end)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCalendarID provides a mock function with given fields:
func (_m *Sink) GetCalendarID() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Sink_GetSourceID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSourceID'
type Sink_GetSourceID_Call struct {
	*mock.Call
}

// GetSourceID is a helper method to define mock.On call
func (_e *Sink_Expecter) GetSourceID() *Sink_GetSourceID_Call {
	return &Sink_GetSourceID_Call{Call: _e.mock.On("GetSourceID")}
}

func (_c *Sink_GetSourceID_Call) Run(run func()) *Sink_GetSourceID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Sink_GetSourceID_Call) Return(_a0 string) *Sink_GetSourceID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Sink_GetSourceID_Call) RunAndReturn(run func() string) *Sink_GetSourceID_Call {
	_c.Call.Return(run)
	return _c
}

// Name provides a mock function with given fields:
func (_m *Sink) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Sink_Name_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Name'
type Sink_Name_Call struct {
	*mock.Call
}

// Name is a helper method to define mock.On call
func (_e *Sink_Expecter) Name() *Sink_Name_Call {
	return &Sink_Name_Call{Call: _e.mock.On("Name")}
}

func (_c *Sink_Name_Call) Run(run func()) *Sink_Name_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *Sink_Name_Call) Return(_a0 string) *Sink_Name_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Sink_Name_Call) RunAndReturn(run func() string) *Sink_Name_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateEvent provides a mock function with given fields: ctx, e
func (_m *Sink) UpdateEvent(ctx context.Context, e models.Event) error {
	ret := _m.Called(ctx, e)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Event) error); ok {
		r0 = rf(ctx, e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sink_UpdateEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateEvent'
type Sink_UpdateEvent_Call struct {
	*mock.Call
}

// UpdateEvent is a helper method to define mock.On call
//   - ctx context.Context
//   - e models.Event
func (_e *Sink_Expecter) UpdateEvent(ctx interface{}, e interface{}) *Sink_UpdateEvent_Call {
	return &Sink_UpdateEvent_Call{Call: _e.mock.On("UpdateEvent", ctx, e)}
}

func (_c *Sink_UpdateEvent_Call) Run(run func(ctx context.Context, e models.Event)) *Sink_UpdateEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.Event))
	})
	return _c
}

func (_c *Sink_UpdateEvent_Call) Return(_a0 error) *Sink_UpdateEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Sink_UpdateEvent_Call) RunAndReturn(run func(context.Context, models.Event) error) *Sink_UpdateEvent_Call {
	_c.Call.Return(run)
	return _c
}

// NewSink creates a new instance of Sink. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewSink(t interface {
	mock.TestingT
	Cleanup(func())
}) *Sink {
	mock := &Sink{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
