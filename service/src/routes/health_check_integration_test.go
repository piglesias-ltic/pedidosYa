package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/peya/src/domain/entities"
	"github.com/stretchr/testify/assert"
)

func Test_Integration_Health(t *testing.T) {
	r := SetupRouter()
	v := entities.VersionInfo{}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	res := w.Body.Bytes()
	if err := json.Unmarshal(res, &v); assert.NoError(t, err) {
		assert.NotNil(t, v)
		assert.NotNil(t, v.Version)
		assert.NotNil(t, v.ReleaseDate)
	}
}
