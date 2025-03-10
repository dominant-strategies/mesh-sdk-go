// Code generated by mockery v2.13.1. DO NOT EDIT.

package modules

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	parser "github.com/dominant-strategies/mesh-sdk-go/parser"
	database "github.com/dominant-strategies/mesh-sdk-go/storage/database"
	types "github.com/dominant-strategies/mesh-sdk-go/types"
)

// BalanceStorageHandler is an autogenerated mock type for the BalanceStorageHandler type
type BalanceStorageHandler struct {
	mock.Mock
}

// AccountsReconciled provides a mock function with given fields: ctx, dbTx, count
func (_m *BalanceStorageHandler) AccountsReconciled(ctx context.Context, dbTx database.Transaction, count int) error {
	ret := _m.Called(ctx, dbTx, count)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, int) error); ok {
		r0 = rf(ctx, dbTx, count)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AccountsSeen provides a mock function with given fields: ctx, dbTx, count
func (_m *BalanceStorageHandler) AccountsSeen(ctx context.Context, dbTx database.Transaction, count int) error {
	ret := _m.Called(ctx, dbTx, count)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.Transaction, int) error); ok {
		r0 = rf(ctx, dbTx, count)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlockAdded provides a mock function with given fields: ctx, block, changes
func (_m *BalanceStorageHandler) BlockAdded(ctx context.Context, block *types.Block, changes []*parser.BalanceChange) error {
	ret := _m.Called(ctx, block, changes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Block, []*parser.BalanceChange) error); ok {
		r0 = rf(ctx, block, changes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BlockRemoved provides a mock function with given fields: ctx, block, changes
func (_m *BalanceStorageHandler) BlockRemoved(ctx context.Context, block *types.Block, changes []*parser.BalanceChange) error {
	ret := _m.Called(ctx, block, changes)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.Block, []*parser.BalanceChange) error); ok {
		r0 = rf(ctx, block, changes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBalanceStorageHandler interface {
	mock.TestingT
	Cleanup(func())
}

// NewBalanceStorageHandler creates a new instance of BalanceStorageHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBalanceStorageHandler(t mockConstructorTestingTNewBalanceStorageHandler) *BalanceStorageHandler {
	mock := &BalanceStorageHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
