// Code generated by mockery v2.39.1. DO NOT EDIT.

package mocks

import (
	context "context"

	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"

	pb "github.com/keshu12345/blog/pb"
	"testing"

)

// BlogServiceClient is an autogenerated mock type for the BlogServiceClient type
type BlogServiceClient struct {
	mock.Mock
}

func NewMockBlogServiceClient(t *testing.T)BlogServiceClient{
   return BlogServiceClient{}
}

// CreatePost provides a mock function with given fields: ctx, in, opts
func (_m *BlogServiceClient) CreatePost(ctx context.Context, in *pb.CreatePostRequest, opts ...grpc.CallOption) (*pb.Post, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreatePost")
	}

	var r0 *pb.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreatePostRequest, ...grpc.CallOption) (*pb.Post, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.CreatePostRequest, ...grpc.CallOption) *pb.Post); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.CreatePostRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePost provides a mock function with given fields: ctx, in, opts
func (_m *BlogServiceClient) DeletePost(ctx context.Context, in *pb.DeletePostRequest, opts ...grpc.CallOption) (*pb.DeletePostResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for DeletePost")
	}

	var r0 *pb.DeletePostResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeletePostRequest, ...grpc.CallOption) (*pb.DeletePostResponse, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.DeletePostRequest, ...grpc.CallOption) *pb.DeletePostResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.DeletePostResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.DeletePostRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadPost provides a mock function with given fields: ctx, in, opts
func (_m *BlogServiceClient) ReadPost(ctx context.Context, in *pb.ReadPostRequest, opts ...grpc.CallOption) (*pb.Post, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for ReadPost")
	}

	var r0 *pb.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ReadPostRequest, ...grpc.CallOption) (*pb.Post, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.ReadPostRequest, ...grpc.CallOption) *pb.Post); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.ReadPostRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePost provides a mock function with given fields: ctx, in, opts
func (_m *BlogServiceClient) UpdatePost(ctx context.Context, in *pb.UpdatePostRequest, opts ...grpc.CallOption) (*pb.Post, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePost")
	}

	var r0 *pb.Post
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdatePostRequest, ...grpc.CallOption) (*pb.Post, error)); ok {
		return rf(ctx, in, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pb.UpdatePostRequest, ...grpc.CallOption) *pb.Post); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pb.Post)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pb.UpdatePostRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewBlogServiceClient creates a new instance of BlogServiceClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBlogServiceClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *BlogServiceClient {
	mock := &BlogServiceClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
