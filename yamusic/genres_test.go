package yamusic

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenresService_List(t *testing.T) {
	setup()
	defer teardown()

	want := &Genres{}
	want.InvocationInfo.ReqID = "Genres.List"

	mux.HandleFunc("/genres", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		b, _ := json.Marshal(want)
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Genres.List(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, want, result)
}
