// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import apperrors "github.com/kyma-project/kyma/components/compass-runtime-agent/internal/apperrors"

import mock "github.com/stretchr/testify/mock"
import types "k8s.io/apimachinery/pkg/types"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// CreateHandler provides a mock function with given fields: application, appUID, serviceId, name
func (_m *Repository) CreateHandler(application string, appUID types.UID, serviceId string, name string) apperrors.AppError {
	ret := _m.Called(application, appUID, serviceId, name)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, types.UID, string, string) apperrors.AppError); ok {
		r0 = rf(application, appUID, serviceId, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// CreateInstance provides a mock function with given fields: application, appUID, serviceId, name
func (_m *Repository) CreateInstance(application string, appUID types.UID, serviceId string, name string) apperrors.AppError {
	ret := _m.Called(application, appUID, serviceId, name)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, types.UID, string, string) apperrors.AppError); ok {
		r0 = rf(application, appUID, serviceId, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// CreateRule provides a mock function with given fields: application, appUID, serviceId, name
func (_m *Repository) CreateRule(application string, appUID types.UID, serviceId string, name string) apperrors.AppError {
	ret := _m.Called(application, appUID, serviceId, name)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, types.UID, string, string) apperrors.AppError); ok {
		r0 = rf(application, appUID, serviceId, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// DeleteHandler provides a mock function with given fields: name
func (_m *Repository) DeleteHandler(name string) apperrors.AppError {
	ret := _m.Called(name)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string) apperrors.AppError); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// DeleteInstance provides a mock function with given fields: name
func (_m *Repository) DeleteInstance(name string) apperrors.AppError {
	ret := _m.Called(name)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string) apperrors.AppError); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// DeleteRule provides a mock function with given fields: name
func (_m *Repository) DeleteRule(name string) apperrors.AppError {
	ret := _m.Called(name)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string) apperrors.AppError); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// UpsertHandler provides a mock function with given fields: application, appUID, serviceId, name
func (_m *Repository) UpsertHandler(application string, appUID types.UID, serviceId string, name string) apperrors.AppError {
	ret := _m.Called(application, appUID, serviceId, name)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, types.UID, string, string) apperrors.AppError); ok {
		r0 = rf(application, appUID, serviceId, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// UpsertInstance provides a mock function with given fields: application, appUID, serviceId, name
func (_m *Repository) UpsertInstance(application string, appUID types.UID, serviceId string, name string) apperrors.AppError {
	ret := _m.Called(application, appUID, serviceId, name)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, types.UID, string, string) apperrors.AppError); ok {
		r0 = rf(application, appUID, serviceId, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}

// UpsertRule provides a mock function with given fields: application, appUID, serviceId, name
func (_m *Repository) UpsertRule(application string, appUID types.UID, serviceId string, name string) apperrors.AppError {
	ret := _m.Called(application, appUID, serviceId, name)

	var r0 apperrors.AppError
	if rf, ok := ret.Get(0).(func(string, types.UID, string, string) apperrors.AppError); ok {
		r0 = rf(application, appUID, serviceId, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(apperrors.AppError)
		}
	}

	return r0
}
