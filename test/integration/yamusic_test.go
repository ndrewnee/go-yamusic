// +build integration

package integration

import (
	"github.com/ndrewnee/go-yamusic/yamusic"
)

var (
	client *yamusic.Client
)

func init() {
	client = yamusic.NewClient()
}
