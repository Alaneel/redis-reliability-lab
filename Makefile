.PHONY: build test integration check fmt clean
VERSION ?= dev
export VERSION

build:
	go build -trimpath -buildvcs=false -ldflags="-X main.version=$$VERSION" -o bin/redis-lab ./cmd/redis-lab
test:
	go test -race -count=1 ./...
integration:
	go test -race -tags=integration -count=1 -timeout=20m -v ./internal/experiment
check:
	test -z "$$(gofmt -l cmd internal)"
	go vet ./...
	go test -race -count=1 ./...
fmt:
	gofmt -w cmd internal
clean:
	rm -f bin/redis-lab
