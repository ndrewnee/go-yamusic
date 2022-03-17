//go:build integration

package integration

import (
	"context"
	"github.com/stretchr/testify/require"
	"net/http"
	"strconv"
	"testing"
)

func TestTracks(t *testing.T) {
	var kind int = 1695506
	ctx := context.Background()
	t.Run("Get track information", func(t *testing.T) {
		result, resp, err := client.Tracks().Get(ctx, kind)
		require.NoError(t, err)
		require.NotZero(t, result)
		require.Equal(t, http.StatusOK, resp.StatusCode)
		require.NotZero(t, result.Result)
		require.Equal(t, result.Result[0].ID, strconv.Itoa(kind))
	})
	t.Run("Get track download URL", func(t *testing.T) {
		url, err := client.Tracks().GetDownloadURL(ctx, kind)
		require.NoError(t, err)
		require.NotZero(t, url)
	})
}
