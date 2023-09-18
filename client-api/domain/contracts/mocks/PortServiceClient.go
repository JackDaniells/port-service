// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	port "github.com/JackDaniells/port-service/proto"
)

// PortServiceClient is an autogenerated mock type for the PortServiceClient type
type PortServiceClient struct {
	mock.Mock
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *PortServiceClient) FindByID(ctx context.Context, id string) (*port.Port, error) {
	ret := _m.Called(ctx, id)

	var r0 *port.Port
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*port.Port, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *port.Port); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*port.Port)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StreamClose provides a mock function with given fields: ctx, stream
func (_m *PortServiceClient) StreamClose(ctx context.Context, stream port.PortService_CreateClient) (*port.CreateResponse, error) {
	ret := _m.Called(ctx, stream)

	var r0 *port.CreateResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, port.PortService_CreateClient) (*port.CreateResponse, error)); ok {
		return rf(ctx, stream)
	}
	if rf, ok := ret.Get(0).(func(context.Context, port.PortService_CreateClient) *port.CreateResponse); ok {
		r0 = rf(ctx, stream)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*port.CreateResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, port.PortService_CreateClient) error); ok {
		r1 = rf(ctx, stream)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StreamCreate provides a mock function with given fields: ctx
func (_m *PortServiceClient) StreamCreate(ctx context.Context) (port.PortService_CreateClient, error) {
	ret := _m.Called(ctx)

	var r0 port.PortService_CreateClient
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (port.PortService_CreateClient, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) port.PortService_CreateClient); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(port.PortService_CreateClient)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StreamSendPortFile provides a mock function with given fields: stream, _a1
func (_m *PortServiceClient) StreamSendPortFile(stream port.PortService_CreateClient, _a1 *port.Port) error {
	ret := _m.Called(stream, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(port.PortService_CreateClient, *port.Port) error); ok {
		r0 = rf(stream, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPortServiceClient creates a new instance of PortServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPortServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *PortServiceClient {
	mock := &PortServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
