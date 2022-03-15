//go:build integration
// +build integration

package integration

import (
	"context"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestTracks(t *testing.T) {
	var kind int = 1695506
	ctx := context.Background()
	t.Run("Get track information", func(t *testing.T) {
		result, resp, err := client.Tracks().Get(ctx, kind)
		assert.NoError(t, err)
		assert.NotZero(t, result)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotZero(t, result.Result)
		assert.Equal(t, result.Result[0].ID, kind)
	})
	t.Run("Get track download URL", func(t *testing.T) {
		url, err := client.Tracks.Get(ctx, kind)
		assert.NoError(t, err)
		assert.NotZero(t, url)
	})
}
