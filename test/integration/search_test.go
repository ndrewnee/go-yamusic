//go:build integration

package integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/ndrewnee/go-yamusic/yamusic"

	"github.com/stretchr/testify/assert"
)

func TestSearchArtists(t *testing.T) {
	result := testSearch(t, client.Search().Artists)
	assert.NotEmpty(t, result.Result.Artists.Results)
}

func TestSearchAlbums(t *testing.T) {
	result := testSearch(t, client.Search().Albums)
	assert.NotEmpty(t, result.Result.Albums.Results)
}

func TestSearchTracks(t *testing.T) {
	result := testSearch(t, client.Search().Tracks)
	assert.NotEmpty(t, result.Result.Tracks.Results)
}

func TestSearchAll(t *testing.T) {
	result := testSearch(t, client.Search().All)
	assert.NotEmpty(t, result.Result.Artists.Results)
	assert.NotEmpty(t, result.Result.Albums.Results)
	assert.NotEmpty(t, result.Result.Tracks.Results)
	assert.NotZero(t, result.Result.Best)
}

func testSearch(t *testing.T, searchFunc func(
	ctx context.Context,
	query string,
	opts *yamusic.SearchOptions,
) (*yamusic.SearchResp, *http.Response, error)) *yamusic.SearchResp {
	ctx := context.Background()
	result, resp, err := searchFunc(ctx, "Oxxymiron", nil)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotZero(t, result)

	return result
}
