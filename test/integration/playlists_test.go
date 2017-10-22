// +build integration

package integration

import (
	"context"
	"github.com/ndrewnee/go-yamusic/yamusic"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaylistsList(t *testing.T) {
	playlists, resp, err := client.Playlists().List(
		context.Background(),
		495301201,
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotZero(t, playlists)
	assert.NotEmpty(t, playlists.Result)
}

func TestPlaylistsGet(t *testing.T) {
	playlist, resp, err := client.Playlists().Get(
		context.Background(),
		495301201,
		101,
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotZero(t, playlist)
	assert.NotZero(t, playlist.Result)
}

func TestPlaylistsGetByKinds(t *testing.T) {
	playlists, resp, err := client.Playlists().GetByKinds(
		context.Background(),
		495301201,
		&yamusic.PlaylistsGetByKindOptions{
			Kinds:      []int{101, 1004},
			Mixed:      false,
			RichTracks: false,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotZero(t, playlists)
	assert.NotEmpty(t, playlists.Result)
}

func TestPlaylistsRename(t *testing.T) {
	playlist, resp, err := client.Playlists().Rename(
		context.Background(),
		1004,
		"New name",
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotZero(t, playlist)
	assert.NotZero(t, playlist.Result)
}
