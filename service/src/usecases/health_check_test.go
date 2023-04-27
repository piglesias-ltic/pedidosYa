package usecases

import (
	"errors"
	"testing"
	"time"

	"github.com/peya/src/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHealthCheckRepo struct {
	mock.Mock
}

func (m *MockHealthCheckRepo) GetVersion() (*entities.VersionInfo, error) {
	args := m.Called()
	var v *entities.VersionInfo
	var ok bool
	if v, ok = args.Get(0).(*entities.VersionInfo); !ok {
		v = nil
	}
	return v, args.Error(1)
}

func TestHealthCheckUsecase(t *testing.T) {
	mockRepo := new(MockHealthCheckRepo)
	e := entities.VersionInfo{Version: "1.0.0", ReleaseDate: time.Now().UTC()}

	mockRepo.On("GetVersion").Return(&e, nil)

	u := HealthCheckUsecaseImpl{HealthCheckRepo: mockRepo}
	v, err := u.GetStatus()

	if assert.NoError(t, err) {
		assert.Equal(t, &e, v)
		mockRepo.AssertCalled(t, "GetVersion")
	}

}

func TestHealthCheckUsecaseError(t *testing.T) {
	mockRepo := new(MockHealthCheckRepo)
	e := errors.New("fake repo error")

	mockRepo.On("GetVersion").Return(nil, e)

	u := HealthCheckUsecaseImpl{HealthCheckRepo: mockRepo}
	v, err := u.GetStatus()

	if assert.Error(t, err) {
		assert.Equal(t, e, err)
		assert.Nil(t, v)
		mockRepo.AssertCalled(t, "GetVersion")
	}

}
