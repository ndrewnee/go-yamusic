package yamusic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaylistsService_List(t *testing.T) {
	setup()
	defer teardown()

	want := &PlaylistsListResp{}
	want.InvocationInfo.ReqID = "Playlists.List"

	mux.HandleFunc(
		fmt.Sprintf("/users/%v/playlists/list", client.UserID()),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
			b, err := json.Marshal(want)
			assert.NoError(t, err)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Playlists().List(context.Background(), 0)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}

func TestPlaylistsService_Get(t *testing.T) {
	setup()
	defer teardown()

	want := &PlaylistsGetResp{}
	want.InvocationInfo.ReqID = "Playlists.Get"

	kind := 2000

	mux.HandleFunc(
		fmt.Sprintf("/users/%v/playlists/%v", client.UserID(), kind),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
			b, err := json.Marshal(want)
			assert.NoError(t, err)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Playlists().Get(
		context.Background(),
		0,
		kind,
	)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}

func TestPlaylistsService_GetByKinds(t *testing.T) {
	setup()
	defer teardown()

	want := &PlaylistsGetByKindsResp{}
	want.InvocationInfo.ReqID = "Playlists.GetByKinds"

	kind1 := 101
	kind2 := 102
	mixed := true
	richTracks := true

	mux.HandleFunc(
		fmt.Sprintf("/users/%v/playlists", client.UserID()),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
			assert.Equal(
				t,
				fmt.Sprintf(
					"/users/%v/playlists?kinds=%v%%2C%v&mixed=%v&rich-tracks=%v",
					client.UserID(),
					kind1,
					kind2,
					mixed,
					richTracks,
				),
				r.URL.String(),
			)
			b, err := json.Marshal(want)
			assert.NoError(t, err)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Playlists().GetByKinds(
		context.Background(),
		0,
		&PlaylistsGetByKindOptions{
			Kinds:      []int{kind1, kind2},
			Mixed:      mixed,
			RichTracks: richTracks,
		},
	)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}

func TestPlaylistsService_Rename(t *testing.T) {
	setup()
	defer teardown()

	want := &PlaylistsRenameResp{}
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

			b, err := json.Marshal(want)
			assert.NoError(t, err)
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

	want := &PlaylistsCreateResp{}
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

			b, err := json.Marshal(want)
			assert.NoError(t, err)
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

	want := &PlaylistsDeleteResp{}
	want.InvocationInfo.ReqID = "Playlists.Delete"

	kind := 1004

	mux.HandleFunc(
		fmt.Sprintf("/users/%v/playlists/%v/delete", userID, kind),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))

			b, err := json.Marshal(want)
			assert.NoError(t, err)
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

func TestPlaylistsService_AddTracks(t *testing.T) {
	setup()
	defer teardown()

	want := &PlaylistsAddTracksResp{}
	want.InvocationInfo.ReqID = "Playlists.AddTracks"

	revision := 1
	kind := 1004
	trackID := 1
	albumID := 1

	mux.HandleFunc(
		fmt.Sprintf("/users/%v/playlists/%v/change-relative", userID, kind),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			err := r.ParseForm()
			assert.NoError(t, err)

			diff := fmt.Sprintf(
				`[{"op":"insert","at":0,"tracks":[{"id":%v,"albumId":%v}]}]`,
				trackID,
				albumID,
			)

			assert.Equal(t, diff, r.FormValue("diff"))
			assert.Equal(t, strconv.Itoa(revision), r.FormValue("revision"))
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))

			b, err := json.Marshal(want)
			assert.NoError(t, err)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Playlists().AddTracks(
		context.Background(),
		kind,
		revision,
		[]PlaylistsTrack{
			{
				ID:      trackID,
				AlbumID: albumID,
			},
		},
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}

func TestPlaylistsService_RemoveTracks(t *testing.T) {
	setup()
	defer teardown()

	want := &PlaylistsRemoveTracksResp{}
	want.InvocationInfo.ReqID = "Playlists.RemoveTracks"

	revision := 1
	kind := 1004
	trackID := 1
	albumID := 1

	mux.HandleFunc(
		fmt.Sprintf("/users/%v/playlists/%v/change-relative", userID, kind),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodPost, r.Method)

			err := r.ParseForm()
			assert.NoError(t, err)

			diff := fmt.Sprintf(
				`[{"op":"delete","from":0,"to":1,"tracks":[{"id":%v,"albumId":%v}]}]`,
				trackID,
				albumID,
			)

			assert.Equal(t, diff, r.FormValue("diff"))
			assert.Equal(t, strconv.Itoa(revision), r.FormValue("revision"))
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))

			b, err := json.Marshal(want)
			assert.NoError(t, err)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Playlists().RemoveTracks(
		context.Background(),
		kind,
		revision,
		[]PlaylistsTrack{
			{
				ID:      trackID,
				AlbumID: albumID,
			},
		},
		nil,
	)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}
