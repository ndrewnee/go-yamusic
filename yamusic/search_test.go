package yamusic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchService_Artists(t *testing.T) {
	setup()
	defer teardown()
	testSearch(t, "artist", client.Search().Artists)
}

func TestSearchService_Albums(t *testing.T) {
	setup()
	defer teardown()
	testSearch(t, "album", client.Search().Albums)
}

func TestSearchService_Tracks(t *testing.T) {
	setup()
	defer teardown()
	testSearch(t, "track", client.Search().Tracks)
}

func TestSearchService_All(t *testing.T) {
	setup()
	defer teardown()
	testSearch(t, "all", client.Search().All)
}

func testSearch(t *testing.T, searchType string, searchFunc func(
	ctx context.Context,
	query string,
	opts *SearchOptions,
) (*SearchResp, *http.Response, error)) {
	want := &SearchResp{}
	want.InvocationInfo.ReqID = searchType

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
		assert.Equal(t,
			"/search?nocorrect=true&page=2&text=blah&type="+searchType,
			r.URL.String())

		b, err := json.Marshal(want)
		assert.NoError(t, err)
		fmt.Fprint(w, string(b))
	})

	opts := &SearchOptions{Page: 2, NoCorrect: true}
	result, _, err := searchFunc(context.Background(), "blah", opts)
	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}
