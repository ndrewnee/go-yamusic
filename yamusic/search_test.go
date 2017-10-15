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

	want := &Search{}
	want.InvocationInfo.ReqID = "Search.Artists"

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t,
			"/search?nocorrect=true&page=2&text=blah&type=artist",
			r.RequestURI)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	opts := &SearchOptions{Page: 2, NoCorrect: true}
	result, _, err := client.Search().Artists(context.Background(), "blah", opts)
	assert.NoError(t, err)
	assert.Equal(t, want, result)
}

func TestSearchService_Albums(t *testing.T) {
	setup()
	defer teardown()

	want := &Search{}
	want.InvocationInfo.ReqID = "Search.Albums"

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t,
			"/search?nocorrect=true&page=2&text=blah&type=album",
			r.RequestURI)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	opts := &SearchOptions{Page: 2, NoCorrect: true}
	result, _, err := client.Search().Albums(context.Background(), "blah", opts)
	assert.NoError(t, err)
	assert.Equal(t, want, result)
}

func TestSearchService_Tracks(t *testing.T) {
	setup()
	defer teardown()

	want := &Search{}
	want.InvocationInfo.ReqID = "Search.Tracks"

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t,
			"/search?nocorrect=true&page=2&text=blah&type=track",
			r.RequestURI)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	opts := &SearchOptions{Page: 2, NoCorrect: true}
	result, _, err := client.Search().Tracks(context.Background(), "blah", opts)
	assert.NoError(t, err)
	assert.Equal(t, want, result)
}

func TestSearchService_All(t *testing.T) {
	setup()
	defer teardown()

	want := &Search{}
	want.InvocationInfo.ReqID = "Search.All"

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t,
			"/search?nocorrect=true&page=2&text=blah&type=all",
			r.RequestURI)

		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	opts := &SearchOptions{Page: 2, NoCorrect: true}
	result, _, err := client.Search().All(context.Background(), "blah", opts)
	assert.NoError(t, err)
	assert.Equal(t, want, result)
}
