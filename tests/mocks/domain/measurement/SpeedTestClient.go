// Code generated by mockery 2.9.4. DO NOT EDIT.

package mocks

import (
	measurement "github.com/mariuscristian/speedtest/domain/measurement"
	mock "github.com/stretchr/testify/mock"
)

// SpeedTestClient is an autogenerated mock type for the SpeedTestClient type
type SpeedTestClient struct {
	mock.Mock
}

// GetMethod provides a mock function with given fields:
func (_m *SpeedTestClient) GetMethod() measurement.SpeedTestServerType {
	ret := _m.Called()

	var r0 measurement.SpeedTestServerType
	if rf, ok := ret.Get(0).(func() measurement.SpeedTestServerType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(measurement.SpeedTestServerType)
	}

	return r0
}

// Measure provides a mock function with given fields:
func (_m *SpeedTestClient) Measure() measurement.SpeedTestResult {
	ret := _m.Called()

	var r0 measurement.SpeedTestResult
	if rf, ok := ret.Get(0).(func() measurement.SpeedTestResult); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(measurement.SpeedTestResult)
	}

	return r0
}