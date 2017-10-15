// +build integration

package integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchArtists(t *testing.T) {
	result, resp, err := client.Search().Artists(context.Background(), "Oxxymiron", nil)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotZero(t, result)
	assert.NotEmpty(t, result.Result.Artists.Results)
}

func TestSearchAlbums(t *testing.T) {
	result, resp, err := client.Search().Albums(context.Background(), "Oxxymiron", nil)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotZero(t, result)
	assert.NotEmpty(t, result.Result.Albums.Results)
}

func TestSearchTracks(t *testing.T) {
	result, resp, err := client.Search().Tracks(context.Background(), "Oxxymiron", nil)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotZero(t, result)
	assert.NotEmpty(t, result.Result.Tracks.Results)
}

func TestSearchAll(t *testing.T) {
	result, resp, err := client.Search().All(context.Background(), "Oxxymiron", nil)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotZero(t, result)
	assert.NotEmpty(t, result.Result.Artists.Results)
	assert.NotEmpty(t, result.Result.Albums.Results)
	assert.NotEmpty(t, result.Result.Tracks.Results)
	assert.NotZero(t, result.Result.Best)
}
