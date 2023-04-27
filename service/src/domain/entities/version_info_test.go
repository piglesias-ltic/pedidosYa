package entities

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestVersionInfoMarshal(t *testing.T) {
	var v VersionInfo
	e := VersionInfo{Version: "1.0.0", ReleaseDate: time.Now().UTC()}

	j := fmt.Sprintf(`{"version": "%s", "release_date": "%s"}`, e.Version, e.ReleaseDate.Format(time.RFC3339Nano))
	err := json.Unmarshal([]byte(j), &v)

	if assert.NoError(t, err) {
		assert.Equal(t, e, v)
	}

}
