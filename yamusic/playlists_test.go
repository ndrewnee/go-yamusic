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
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
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
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
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
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
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

func TestPlaylistsService_Rename(t *testing.T) {
	setup()
	defer teardown()

	want := &PlaylistsRename{}
	want.InvocationInfo.ReqID = "Playlists.Rename"

	kind := 1004
	newValue := "newValue"

	mux.HandleFunc(
		fmt.Sprintf("/users/%v/playlists/%v/name", userID, kind),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			err := r.ParseForm()
			assert.NoError(t, err)
			assert.Equal(t, newValue, r.FormValue("value"))
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))

			b, _ := json.Marshal(want)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Playlists().Rename(
		context.Background(),
		kind,
		newValue,
	)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}

func TestPlaylistsService_Create(t *testing.T) {
	setup()
	defer teardown()

	want := &PlaylistsCreate{}
	want.InvocationInfo.ReqID = "Playlists.Create"

	title := "title"

	mux.HandleFunc(
		fmt.Sprintf("/users/%v/playlists/create", userID),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			err := r.ParseForm()
			assert.NoError(t, err)
			assert.Equal(t, title, r.FormValue("title"))
			assert.Equal(t, "public", r.FormValue("visibility"))
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))

			b, _ := json.Marshal(want)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Playlists().Create(
		context.Background(),
		title,
		true,
	)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}

func TestPlaylistsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	want := &PlaylistsDelete{}
	want.InvocationInfo.ReqID = "Playlists.Delete"

	kind := 1004

	mux.HandleFunc(
		fmt.Sprintf("/users/%v/playlists/%v/delete", userID, kind),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))

			b, _ := json.Marshal(want)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Playlists().Delete(
		context.Background(),
		kind,
	)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}
