package configuration_test

import (
	"testing"

	mocks "simple-invitation/tests/mocks/configuration"

	"github.com/stretchr/testify/assert"
)

func TestGetUserIDatabaseReader(t *testing.T) {
	mock := new(mocks.IDatabaseReader)
	mock.On("GetUser").Return("rani")
	assert.Equal(t, "rani", mock.GetUser())
	mock.AssertExpectations(t)
}
func TestGetPasswordIDatabaseReader(t *testing.T) {
	mock := new(mocks.IDatabaseReader)
	mock.On("GetPassword").Return("rani")
	assert.Equal(t, "rani", mock.GetPassword())

	mock.AssertExpectations(t)
}

func TestGetProtocolIDatabaseReader(t *testing.T) {
	mock := new(mocks.IDatabaseReader)
	mock.On("GetProtocol").Return("tcp")
	assert.Equal(t, "tcp", mock.GetProtocol())

	mock.AssertExpectations(t)
}

func TestGetPortIDatabaseReader(t *testing.T) {
	mock := new(mocks.IDatabaseReader)
	mock.On("GetPort").Return("3306")
	assert.Equal(t, "3306", mock.GetPort())

	mock.AssertExpectations(t)
}

func TestGetHostIDatabaseReader(t *testing.T) {
	mock := new(mocks.IDatabaseReader)
	mock.On("GetHost").Return("127.0.0.1")
	assert.Equal(t, "127.0.0.1", mock.GetHost())

	mock.AssertExpectations(t)
}

func TestGetNameIDatabaseReader(t *testing.T) {
	mock := new(mocks.IDatabaseReader)
	mock.On("GetName").Return("simpleinvitation")
	assert.Equal(t, "simpleinvitation", mock.GetName())

	mock.AssertExpectations(t)
}
