// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	entities "simple-invitation/domain/entities"

	mock "github.com/stretchr/testify/mock"

	time "time"

	uuid "github.com/google/uuid"
)

// IGatheringUsecase is an autogenerated mock type for the IGatheringUsecase type
type IGatheringUsecase struct {
	mock.Mock
}

// CreateNewGathering provides a mock function with given fields: idAttendees, idCreator, gatheringType, name, location, scheduled_at
func (_m *IGatheringUsecase) CreateNewGathering(idAttendees []uuid.UUID, idCreator uuid.UUID, gatheringType string, name string, location string, scheduled_at time.Time) (*entities.Gathering, error) {
	ret := _m.Called(idAttendees, idCreator, gatheringType, name, location, scheduled_at)

	var r0 *entities.Gathering
	var r1 error
	if rf, ok := ret.Get(0).(func([]uuid.UUID, uuid.UUID, string, string, string, time.Time) (*entities.Gathering, error)); ok {
		return rf(idAttendees, idCreator, gatheringType, name, location, scheduled_at)
	}
	if rf, ok := ret.Get(0).(func([]uuid.UUID, uuid.UUID, string, string, string, time.Time) *entities.Gathering); ok {
		r0 = rf(idAttendees, idCreator, gatheringType, name, location, scheduled_at)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Gathering)
		}
	}

	if rf, ok := ret.Get(1).(func([]uuid.UUID, uuid.UUID, string, string, string, time.Time) error); ok {
		r1 = rf(idAttendees, idCreator, gatheringType, name, location, scheduled_at)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchGatheringTypeById provides a mock function with given fields: id
func (_m *IGatheringUsecase) FetchGatheringTypeById(id uuid.UUID) (*entities.GatheringType, error) {
	ret := _m.Called(id)

	var r0 *entities.GatheringType
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*entities.GatheringType, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *entities.GatheringType); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.GatheringType)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchGatheringTypes provides a mock function with given fields:
func (_m *IGatheringUsecase) FetchGatheringTypes() ([]*entities.GatheringType, error) {
	ret := _m.Called()

	var r0 []*entities.GatheringType
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*entities.GatheringType, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*entities.GatheringType); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.GatheringType)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchInvitationStatusById provides a mock function with given fields: id
func (_m *IGatheringUsecase) FetchInvitationStatusById(id uuid.UUID) (*entities.InvitationStatus, error) {
	ret := _m.Called(id)

	var r0 *entities.InvitationStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*entities.InvitationStatus, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *entities.InvitationStatus); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.InvitationStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchInvitationStatuses provides a mock function with given fields:
func (_m *IGatheringUsecase) FetchInvitationStatuses() ([]*entities.InvitationStatus, error) {
	ret := _m.Called()

	var r0 []*entities.InvitationStatus
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*entities.InvitationStatus, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*entities.InvitationStatus); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*entities.InvitationStatus)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreInvitationStatus provides a mock function with given fields: name
func (_m *IGatheringUsecase) StoreInvitationStatus(name string) (*entities.InvitationStatus, error) {
	ret := _m.Called(name)

	var r0 *entities.InvitationStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.InvitationStatus, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.InvitationStatus); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.InvitationStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StoreNewGatheringType provides a mock function with given fields: name
func (_m *IGatheringUsecase) StoreNewGatheringType(name string) (*entities.GatheringType, error) {
	ret := _m.Called(name)

	var r0 *entities.GatheringType
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*entities.GatheringType, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) *entities.GatheringType); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.GatheringType)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIGatheringUsecase creates a new instance of IGatheringUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIGatheringUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *IGatheringUsecase {
	mock := &IGatheringUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
