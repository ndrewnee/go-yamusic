package yamusic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestTracksSevice_Get(t *testing.T) {
	setup()
	defer teardown()

	want := &TrackResp{}
	want.InvocationInfo.ReqID = "Tracks.Get"

	kind := 42

	mux.HandleFunc(
		fmt.Sprintf("/tracks/%d", kind),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
			b, err := json.Marshal(want)
			assert.NoError(t, err)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Tracks().Get(
		context.Background(),
		kind,
	)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}

func TestTracksSevice_GetDownloadInfoResp(t *testing.T) {
	setup()
	defer teardown()

	want := &TrackResp{}
	want.InvocationInfo.ReqID = "Tracks.GetDownloadInfo"

	kind := 42

	mux.HandleFunc(
		fmt.Sprintf("/tracks/%d/download-info", kind),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
			b, err := json.Marshal(want)
			assert.NoError(t, err)
			fmt.Fprint(w, string(b))
		},
	)

	result, _, err := client.Tracks().GetDownloadInfoResp(
		context.Background(),
		kind,
	)

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}

func TestTracksSevice_GetDownloadInfo(t *testing.T) {
	setup()
	defer teardown()

	want1 := &DownloadInfoResp{}
	want1.InvocationInfo.ReqID = "Tracks.GetDownloadInfo"
	want1.Result = append(want1.Result, struct {
		Codec           string `json:"codec"`
		Gain            bool   `json:"gain"`
		Preview         bool   `json:"preview"`
		DownloadInfoURL string `json:"downloadInfoUrl"`
		Direct          bool   `json:"direct"`
		BitrateInKbps   int    `json:"bitrateInKbps"`
	}{
		DownloadInfoURL: "/dlinfourl",
	})

	kind := 42

	mux.HandleFunc(
		fmt.Sprintf("/tracks/%d/download-info", kind),
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
			b, err := json.Marshal(want1)
			assert.NoError(t, err)
			fmt.Fprint(w, string(b))
		},
	)

	want2 := &DownloadInfo{}

	mux.HandleFunc(
		want1.Result[0].DownloadInfoURL,
		func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, http.MethodGet, r.Method)
			assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
			b, err := json.Marshal(want2)
			assert.NoError(t, err)
			fmt.Fprint(w, string(b))
		},
	)

	_, _, err := client.Tracks().GetDownloadInfo(
		context.Background(),
		kind,
	)

	assert.NoError(t, err)
}
