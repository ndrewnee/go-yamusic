.PHONY: build test test_integration lint

build:
	go build -race ./...

test:
	go test -race ./yamusic/

# Should set `YANDEX_USER_ID` and `YANDEX_ACCESS_TOKEN` before testing.
test_integration:
	go test -race -tags=integration ./test/integration/

lint:
	golangci-lint run
