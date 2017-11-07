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

	want := &GenresListResp{}
	want.InvocationInfo.ReqID = "Genres.List"

	mux.HandleFunc("/genres", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "OAuth "+accessToken, r.Header.Get("Authorization"))
		b, err := json.Marshal(want)
		assert.NoError(t, err)
		fmt.Fprint(w, string(b))
	})

	result, _, err := client.Genres().List(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, want.InvocationInfo.ReqID, result.InvocationInfo.ReqID)
}
