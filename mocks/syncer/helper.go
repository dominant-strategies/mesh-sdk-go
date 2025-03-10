// Code generated by mockery v2.13.1. DO NOT EDIT.

package syncer

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/dominant-strategies/mesh-sdk-go/types"
)

// Helper is an autogenerated mock type for the Helper type
type Helper struct {
	mock.Mock
}

// Block provides a mock function with given fields: _a0, _a1, _a2
func (_m *Helper) Block(_a0 context.Context, _a1 *types.NetworkIdentifier, _a2 *types.PartialBlockIdentifier) (*types.Block, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *types.Block
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier, *types.PartialBlockIdentifier) *types.Block); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Block)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier, *types.PartialBlockIdentifier) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NetworkStatus provides a mock function with given fields: _a0, _a1
func (_m *Helper) NetworkStatus(_a0 context.Context, _a1 *types.NetworkIdentifier) (*types.NetworkStatusResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.NetworkStatusResponse
	if rf, ok := ret.Get(0).(func(context.Context, *types.NetworkIdentifier) *types.NetworkStatusResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NetworkStatusResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.NetworkIdentifier) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHelper interface {
	mock.TestingT
	Cleanup(func())
}

// NewHelper creates a new instance of Helper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHelper(t mockConstructorTestingTNewHelper) *Helper {
	mock := &Helper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
