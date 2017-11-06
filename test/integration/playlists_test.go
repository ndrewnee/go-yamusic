// +build integration

package integration

import (
	"context"
	"github.com/ndrewnee/go-yamusic/yamusic"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaylists(t *testing.T) {
	var kind int
	t.Run("Create playlist", func(t *testing.T) {
		playlist, resp, err := client.Playlists().Create(
			context.Background(),
			"New playlist",
			true,
		)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotZero(t, playlist)
		assert.NotZero(t, playlist.Result)

		kind = playlist.Result.Kind
	})

	t.Run("Rename playlist", func(t *testing.T) {
		playlist, resp, err := client.Playlists().Rename(
			context.Background(),
			kind,
			"New name",
		)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotZero(t, playlist)
		assert.NotZero(t, playlist.Result)
	})

	t.Run("Get list of playlists", func(t *testing.T) {
		playlists, resp, err := client.Playlists().List(
			context.Background(),
			client.UserID(),
		)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotZero(t, playlists)
		assert.NotEmpty(t, playlists.Result)
	})

	var revision int
	t.Run("Get one playlist by kind", func(t *testing.T) {
		playlist, resp, err := client.Playlists().Get(
			context.Background(),
			client.UserID(),
			kind,
		)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotZero(t, playlist)
		assert.NotZero(t, playlist.Result)

		revision = playlist.Result.Revision
	})

	t.Run("Add tracks to playlist", func(t *testing.T) {
		playlist, resp, err := client.Playlists().AddTracks(
			context.Background(),
			kind,
			revision,
			[]yamusic.PlaylistsTrack{
				{
					ID:      232419,
					AlbumID: 42206,
				},
			},
			nil,
		)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotZero(t, playlist)
		assert.NotZero(t, playlist.Result)

		revision = playlist.Result.Revision
	})

	t.Run("Remove tracks from playlist", func(t *testing.T) {
		playlist, resp, err := client.Playlists().RemoveTracks(
			context.Background(),
			kind,
			revision,
			[]yamusic.PlaylistsTrack{
				{
					ID:      232419,
					AlbumID: 42206,
				},
			},
			nil,
		)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotZero(t, playlist)
		assert.NotZero(t, playlist.Result)
	})

	t.Run("Get playlists by kinds", func(t *testing.T) {
		playlists, resp, err := client.Playlists().GetByKinds(
			context.Background(),
			client.UserID(),
			&yamusic.PlaylistsGetByKindOptions{
				Kinds:      []int{kind},
				Mixed:      false,
				RichTracks: false,
			},
		)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotZero(t, playlists)
		assert.NotEmpty(t, playlists.Result)
	})

	t.Run("Delete playlist", func(t *testing.T) {
		playlist, resp, err := client.Playlists().Delete(
			context.Background(),
			kind,
		)

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.NotZero(t, playlist)
		assert.NotZero(t, playlist.Result)
	})
}
