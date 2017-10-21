// +build integration

package integration

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccountGetStatus(t *testing.T) {
	accountStatus, resp, err := client.Account().GetStatus(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.NotZero(t, accountStatus)
	assert.NotZero(t, accountStatus.Result)
}
