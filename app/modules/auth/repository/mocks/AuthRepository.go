// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/adamnasrudin03/go-template/app/modules/auth/dto"
	mock "github.com/stretchr/testify/mock"

	models "github.com/adamnasrudin03/go-template/app/models"
)

// AuthRepository is an autogenerated mock type for the AuthRepository type
type AuthRepository struct {
	mock.Mock
}

// Login provides a mock function with given fields: ctx, input
func (_m *AuthRepository) Login(ctx context.Context, input dto.LoginReq) (*models.User, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginReq) (*models.User, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dto.LoginReq) *models.User); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dto.LoginReq) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: ctx, input
func (_m *AuthRepository) Register(ctx context.Context, input models.User) (*models.User, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.User) (*models.User, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.User) *models.User); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.User) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAuthRepository creates a new instance of AuthRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAuthRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *AuthRepository {
	mock := &AuthRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
