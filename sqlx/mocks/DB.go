// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	sqlx "github.com/jmoiron/sqlx"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// GetMaster provides a mock function with given fields:
func (_m *DB) GetMaster() *sqlx.DB {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetMaster")
	}

	var r0 *sqlx.DB
	if rf, ok := ret.Get(0).(func() *sqlx.DB); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.DB)
		}
	}

	return r0
}

// StartTransaction provides a mock function with given fields: ctx
func (_m *DB) StartTransaction(ctx context.Context) (*sqlx.Tx, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for StartTransaction")
	}

	var r0 *sqlx.Tx
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*sqlx.Tx, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *sqlx.Tx); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sqlx.Tx)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewDB creates a new instance of DB. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDB(t interface {
	mock.TestingT
	Cleanup(func())
}) *DB {
	mock := &DB{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
