// Code generated by mockery v1.0.0. DO NOT EDIT.
package storage

import mock "github.com/stretchr/testify/mock"

// MockBookingStorage is an autogenerated mock type for the BookingStorage type
type MockBookingStorage struct {
	mock.Mock
}

// ByID provides a mock function with given fields:
func (_m *MockBookingStorage) ByID() {
	_m.Called()
}

// Create provides a mock function with given fields: _a0
func (_m *MockBookingStorage) Create(_a0 interface{}) {
	_m.Called(_a0)
}

// Delete provides a mock function with given fields:
func (_m *MockBookingStorage) Delete() {
	_m.Called()
}

// List provides a mock function with given fields:
func (_m *MockBookingStorage) List() {
	_m.Called()
}

// Update provides a mock function with given fields: _a0
func (_m *MockBookingStorage) Update(_a0 interface{}) {
	_m.Called(_a0)
}
