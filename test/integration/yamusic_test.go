//go:build integration

package integration

import (
	"os"
	"strconv"

	"github.com/ndrewnee/go-yamusic/yamusic"
)

var client *yamusic.Client

func init() {
	var userID int
	var err error

	userIDString := os.Getenv("YANDEX_USER_ID")
	if userIDString != "" {
		userID, err = strconv.Atoi(userIDString)
		if err != nil {
			panic(err)
		}
	}

	accessToken := os.Getenv("YANDEX_ACCESS_TOKEN")
	client = yamusic.NewClient(yamusic.AccessToken(userID, accessToken))
}
