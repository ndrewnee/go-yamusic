// +build integration

package integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeedGet(t *testing.T) {
	feed, resp, err := client.Feed().Get(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotZero(t, feed)
	assert.NotZero(t, feed.Result)
}
