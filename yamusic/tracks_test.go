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
