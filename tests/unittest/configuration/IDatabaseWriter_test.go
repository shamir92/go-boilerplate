package configuration_test

import (
	"testing"

	mocks "simple-invitation/tests/mocks/configuration"

	"github.com/stretchr/testify/assert"
)

func TestGetUserIDatabaseWriter(t *testing.T) {
	mock := new(mocks.IDatabaseWriter)
	mock.On("GetUser").Return("rani")
	assert.Equal(t, "rani", mock.GetUser())
	mock.AssertExpectations(t)
}
func TestGetPasswordIDatabaseWriter(t *testing.T) {
	mock := new(mocks.IDatabaseWriter)
	mock.On("GetPassword").Return("rani")
	assert.Equal(t, "rani", mock.GetPassword())

	mock.AssertExpectations(t)
}

func TestGetProtocolIDatabaseWriter(t *testing.T) {
	mock := new(mocks.IDatabaseWriter)
	mock.On("GetProtocol").Return("tcp")
	assert.Equal(t, "tcp", mock.GetProtocol())

	mock.AssertExpectations(t)
}

func TestGetPortIDatabaseWriter(t *testing.T) {
	mock := new(mocks.IDatabaseWriter)
	mock.On("GetPort").Return("3306")
	assert.Equal(t, "3306", mock.GetPort())

	mock.AssertExpectations(t)
}

func TestGetHostIDatabaseWriter(t *testing.T) {
	mock := new(mocks.IDatabaseWriter)
	mock.On("GetHost").Return("127.0.0.1")
	assert.Equal(t, "127.0.0.1", mock.GetHost())

	mock.AssertExpectations(t)
}

func TestGetNameIDatabaseWriter(t *testing.T) {
	mock := new(mocks.IDatabaseWriter)
	mock.On("GetName").Return("simpleinvitation")
	assert.Equal(t, "simpleinvitation", mock.GetName())

	mock.AssertExpectations(t)
}
