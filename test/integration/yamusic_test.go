// +build integration

package integration

import (
	"github.com/ndrewnee/go-yamusic/yamusic"
	"os"
)

var (
	client *yamusic.Client
)

func init() {
	accessToken := os.Getenv("YANDEX_ACCESS_TOKEN")
	client = yamusic.NewClient(yamusic.AccessToken(accessToken))
}
