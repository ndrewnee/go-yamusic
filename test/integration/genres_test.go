//go:build integration

package integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenresList(t *testing.T) {
	genres, resp, err := client.Genres().List(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotZero(t, genres)
	assert.NotEmpty(t, genres.Result)
}
