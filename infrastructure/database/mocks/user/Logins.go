// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package user

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	user "godem/domain/models/user"
)

// Logins is an autogenerated mock type for the Logins type
type Logins struct {
	mock.Mock
}

// Authenticate provides a mock function with given fields: ctx, requestData
func (_m *Logins) Authenticate(ctx context.Context, requestData *user.LoginRequest) (*user.Users, error) {
	ret := _m.Called(ctx, requestData)

	var r0 *user.Users
	if rf, ok := ret.Get(0).(func(context.Context, *user.LoginRequest) *user.Users); ok {
		r0 = rf(ctx, requestData)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.Users)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *user.LoginRequest) error); ok {
		r1 = rf(ctx, requestData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
