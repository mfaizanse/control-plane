// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import context "context"

import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/director/internal/model"

// EventAPIRepository is an autogenerated mock type for the EventAPIRepository type
type EventAPIRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, item
func (_m *EventAPIRepository) Create(ctx context.Context, item *model.EventAPIDefinition) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.EventAPIDefinition) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateMany provides a mock function with given fields: ctx, items
func (_m *EventAPIRepository) CreateMany(ctx context.Context, items []*model.EventAPIDefinition) error {
	ret := _m.Called(ctx, items)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []*model.EventAPIDefinition) error); ok {
		r0 = rf(ctx, items)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: ctx, tenantID, id
func (_m *EventAPIRepository) Delete(ctx context.Context, tenantID string, id string) error {
	ret := _m.Called(ctx, tenantID, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenantID, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteAllByApplicationID provides a mock function with given fields: ctx, tenantID, appID
func (_m *EventAPIRepository) DeleteAllByApplicationID(ctx context.Context, tenantID string, appID string) error {
	ret := _m.Called(ctx, tenantID, appID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, tenantID, appID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: ctx, tenantID, id
func (_m *EventAPIRepository) Exists(ctx context.Context, tenantID string, id string) (bool, error) {
	ret := _m.Called(ctx, tenantID, id)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, tenantID, id)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenantID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: ctx, tenantID, id
func (_m *EventAPIRepository) GetByID(ctx context.Context, tenantID string, id string) (*model.EventAPIDefinition, error) {
	ret := _m.Called(ctx, tenantID, id)

	var r0 *model.EventAPIDefinition
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.EventAPIDefinition); ok {
		r0 = rf(ctx, tenantID, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.EventAPIDefinition)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, tenantID, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListByApplicationID provides a mock function with given fields: ctx, tenantID, applicationID, pageSize, cursor
func (_m *EventAPIRepository) ListByApplicationID(ctx context.Context, tenantID string, applicationID string, pageSize int, cursor string) (*model.EventAPIDefinitionPage, error) {
	ret := _m.Called(ctx, tenantID, applicationID, pageSize, cursor)

	var r0 *model.EventAPIDefinitionPage
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, string) *model.EventAPIDefinitionPage); ok {
		r0 = rf(ctx, tenantID, applicationID, pageSize, cursor)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.EventAPIDefinitionPage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int, string) error); ok {
		r1 = rf(ctx, tenantID, applicationID, pageSize, cursor)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, item
func (_m *EventAPIRepository) Update(ctx context.Context, item *model.EventAPIDefinition) error {
	ret := _m.Called(ctx, item)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.EventAPIDefinition) error); ok {
		r0 = rf(ctx, item)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
