package routes

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/peya/src/di"
	"github.com/peya/src/domain/entities"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockHealthCheckUseCase struct {
	mock.Mock
}

func (m *MockHealthCheckUseCase) GetStatus() (*entities.VersionInfo, error) {
	args := m.Called()
	return &entities.VersionInfo{}, args.Error(1)
}

func TestRouter_Health(t *testing.T) {
	mock := MockHealthCheckUseCase{}
	i := di.Injector{HealthCheckUsecase: &mock}
	injector = &i
	r := SetupRouter()

	mock.On("GetStatus").Return(nil, nil)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mock.AssertCalled(t, "GetStatus")
}

func TestRouter_HealthError(t *testing.T) {
	mock := MockHealthCheckUseCase{}
	i := di.Injector{HealthCheckUsecase: &mock}
	injector = &i
	r := SetupRouter()

	mock.On("GetStatus").Return(nil, errors.New("some error"))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	mock.AssertCalled(t, "GetStatus")
}
