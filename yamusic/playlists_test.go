package yamusic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaylistsService_List(t *testing.T) {
	setup()
	defer teardown()

	want := &PlaylistsList{}
	want.InvocationInfo.ReqID = "Playlists.List"

	mux.HandleFunc(
		"/users/1000/playlists/list",
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			b, _ := json.Marshal(want)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Playlists().List(context.Background(), 1000)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}

func TestPlaylistsService_Get(t *testing.T) {
	setup()
	defer teardown()

	want := &PlaylistsGet{}
	want.InvocationInfo.ReqID = "Playlists.Get"

	mux.HandleFunc(
		"/users/1000/playlists/2000",
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			b, _ := json.Marshal(want)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Playlists().Get(
		context.Background(),
		1000,
		2000,
	)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}

func TestPlaylistsService_GetByKinds(t *testing.T) {
	setup()
	defer teardown()

	want := &PlaylistsGetByKinds{}
	want.InvocationInfo.ReqID = "Playlists.GetByKinds"

	mux.HandleFunc(
		"/users/1000/playlists",
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(
				t,
				"/users/1000/playlists?kinds=101%2C102&mixed=true&rich-tracks=true",
				r.URL.String(),
			)
			b, _ := json.Marshal(want)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Playlists().GetByKinds(
		context.Background(),
		1000,
		&PlaylistsGetByKindOptions{
			Kinds:      []int{101, 102},
			Mixed:      true,
			RichTracks: true,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}
