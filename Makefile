GO111MODULE=on
GOPROXY=https://goproxy.cn,direct

BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
GIT_COMMIT_HASH := $(shell git rev-parse --verify HEAD)
GO_BUILD_FLAGS := -v -trimpath -ldflags '-s -w -X "github.com/voidint/tsdump/build.Date=$(BUILD_DATE)" -X "github.com/voidint/tsdump/build.Commit=$(GIT_COMMIT_HASH)"'

all: install test

build:
	@echo "GO111MODULE=$(GO111MODULE)"
	@echo "GOPROXY=$(GOPROXY)"
	go build $(GO_BUILD_FLAGS)

install: build
	go install $(GO_BUILD_FLAGS)

test:
	go test -v ./...

clean:
	go clean -x

.PHONY: all build test