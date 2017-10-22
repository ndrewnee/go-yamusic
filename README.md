# yamusic

[![GoDoc](https://godoc.org/github.com/ndrewnee/go-yamusic/yamusic?status.svg)](https://godoc.org/github.com/ndrewnee/go-yamusic/yamusic)
[![Go Report Card](https://goreportcard.com/badge/github.com/ndrewnee/go-yamusic)](https://goreportcard.com/report/github.com/ndrewnee/go-yamusic)
[![Build Status](https://travis-ci.org/ndrewnee/go-yamusic.svg?branch=master)](https://travis-ci.org/ndrewnee/go-yamusic)
[![Coverage Status](https://coveralls.io/repos/github/ndrewnee/go-yamusic/badge.svg)](https://coveralls.io/github/ndrewnee/go-yamusic)
[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/ndrewnee/go-yamusic/issues)

## Description

Unofficial Go client library for [Yandex.Music API](https://music.yandex.ru).

Golang fork of [Node.js library](https://github.com/itsmepetrov/yandex-music-api).

Client style based on [google/go-github](https://github.com/google/go-github).

## Usage

```go
import "github.com/ndrewnee/go-yamusic/yamusic"
```

Construct a new Yandex.Music client, then use the various services on the client to access different parts of the Yandex.Music API. For example:

```go
client := yamusic.NewClient()

// list all genres
genres, _, err := client.Genres.List(context.Background())
if err != nil {
    log.Fatal(err)
}

// search for artists
artists, _, err := client.Search.Artists(context.Background(), "Oxxymiron", nil)
if err != nil {
    log.Fatal(err)
}
```

## Tests

Running unit tests:

```sh
go test ./yamusic/
```

Running integration tests:

```sh
go test -tags=integration ./test/integration/
```
