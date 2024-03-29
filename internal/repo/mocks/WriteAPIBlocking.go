// Code generated by mockery v2.12.2. DO NOT EDIT.

package mocks

import (
	context "context"
	testing "testing"

	mock "github.com/stretchr/testify/mock"

	write "github.com/influxdata/influxdb-client-go/v2/api/write"
)

// WriteAPIBlocking is an autogenerated mock type for the WriteAPIBlocking type
type WriteAPIBlocking struct {
	mock.Mock
}

// WritePoint provides a mock function with given fields: ctx, point
func (_m *WriteAPIBlocking) WritePoint(ctx context.Context, point ...*write.Point) error {
	_va := make([]interface{}, len(point))
	for _i := range point {
		_va[_i] = point[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*write.Point) error); ok {
		r0 = rf(ctx, point...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteRecord provides a mock function with given fields: ctx, line
func (_m *WriteAPIBlocking) WriteRecord(ctx context.Context, line ...string) error {
	_va := make([]interface{}, len(line))
	for _i := range line {
		_va[_i] = line[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, ...string) error); ok {
		r0 = rf(ctx, line...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewWriteAPIBlocking creates a new instance of WriteAPIBlocking. It also registers the testing.TB interface on the mock and a cleanup function to assert the mocks expectations.
func NewWriteAPIBlocking(t testing.TB) *WriteAPIBlocking {
	mock := &WriteAPIBlocking{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
