# yamusic

[![Go Reference](https://pkg.go.dev/badge/github.com/ndrewnee/go-yamusic/yamusic/yamusic.svg)](https://pkg.go.dev/github.com/ndrewnee/go-yamusic/yamusic)
[![Go](https://github.com/ndrewnee/go-yamusic/actions/workflows/go.yml/badge.svg?branch=master)](https://github.com/ndrewnee/go-yamusic/actions/workflows/go.yml)

## Description

Unofficial Go client library for [Yandex.Music API](https://music.yandex.ru).

Golang fork of [Node.js library](https://github.com/itsmepetrov/yandex-music-api).
Client style based on [google/go-github](https://github.com/google/go-github).

## Usage

```go
import "github.com/ndrewnee/go-yamusic/yamusic"
```

Construct a new Yandex.Music client, then use the various services on the client to access different parts of the Yandex.Music API.

Using [functional options for friendly APIs](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis).

```go
package main

import (
    "github.com/ndrewnee/go-yamusic/yamusic"
    "github.com/rubyist/circuitbreaker"
    "context"
    "log"
    "net/http"
    "time"
)

func main() {
    // constructing http client with circuit breaker
    // it implements yamusic.Doer interface
    circuitClient := circuit.NewHTTPClient(time.Second * 5, 10, nil)
    client := yamusic.NewClient(
        // if you want http client with circuit breaker
        yamusic.HTTPClient(circuitClient),
        // provide user_id and access_token (needed by some methods)
        yamusic.AccessToken(100, "some_access_token"),
    )
    // list all genres
    genres, resp, err := client.Genres().List(context.Background())
    if err != nil {
        log.Fatal(err)
    }
    // resp is general type *http.Response
    if resp.StatusCode != http.StatusOK {
        log.Fatal("http status is not 200")
    }
    log.Println("Genres: ", genres)
    // create new public playlist. Need access token
    createdPlaylist, _, err := client.Playlists().Create(context.Background(), "New Playlist", true)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Created playlist: ", createdPlaylist)
}
```

## 👷 Build

Build package

```bash
make build
```

## 🧪 Testing

Run unit tests:

```sh
make test
```

Run integration tests:

Note that you should set `YANDEX_USER_ID` and `YANDEX_ACCESS_TOKEN` environment variables.

```sh
make test_integration
```

## 🖍 Lint

Run linters

```bash
make lint
```
