package yamusic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeedService_Get(t *testing.T) {
	setup()
	defer teardown()

	want := &FeedResp{}
	want.InvocationInfo.ReqID = "Feed.Get"

	mux.HandleFunc("/feed", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Feed().Get(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}
