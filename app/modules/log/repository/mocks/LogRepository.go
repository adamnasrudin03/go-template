// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/adamnasrudin03/go-template/app/modules/log/dto"
	mock "github.com/stretchr/testify/mock"

	models "github.com/adamnasrudin03/go-template/app/models"
)

// LogRepository is an autogenerated mock type for the LogRepository type
type LogRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, input
func (_m *LogRepository) Create(ctx context.Context, input models.Log) error {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Log) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateCache provides a mock function with given fields: ctx, key, data
func (_m *LogRepository) CreateCache(ctx context.Context, key string, data interface{}) {
	_m.Called(ctx, key, data)
}

// CreateLogActivity provides a mock function with given fields: ctx, input
func (_m *LogRepository) CreateLogActivity(ctx context.Context, input models.Log) error {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for CreateLogActivity")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Log) error); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DelCache provides a mock function with given fields: ctx, key
func (_m *LogRepository) DelCache(ctx context.Context, key string) error {
	ret := _m.Called(ctx, key)

	if len(ret) == 0 {
		panic("no return value specified for DelCache")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCache provides a mock function with given fields: ctx, key, res
func (_m *LogRepository) GetCache(ctx context.Context, key string, res interface{}) {
	_m.Called(ctx, key, res)
}

// GetList provides a mock function with given fields: ctx, params
func (_m *LogRepository) GetList(ctx context.Context, params dto.ListLogReq) ([]models.Log, error) {
	ret := _m.Called(ctx, params)

	if len(ret) == 0 {
		panic("no return value specified for GetList")
	}

	var r0 []models.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.ListLogReq) ([]models.Log, error)); ok {
		return rf(ctx, params)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.ListLogReq) []models.Log); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.ListLogReq) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewLogRepository creates a new instance of LogRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLogRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *LogRepository {
	mock := &LogRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
