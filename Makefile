GO = CGO_ENABLED=0 GO111MODULE=on GOPROXY=https://goproxy.cn,direct go
BUILD_DATE := $(shell date '+%Y-%m-%d %H:%M:%S')
GIT_BRANCH := $(shell git symbolic-ref --short -q HEAD)
GIT_COMMIT_HASH := $(shell git rev-parse HEAD|cut -c 1-8)
GO_FLAGS := -v -ldflags="-X 'github.com/voidint/tsdump/build.Built=$(BUILD_DATE)' -X 'github.com/voidint/tsdump/build.GitCommit=$(GIT_COMMIT_HASH)' -X 'github.com/voidint/tsdump/build.GitBranch=$(GIT_BRANCH)'"

all: install test

build:
	$(GO) build $(GO_FLAGS)

install: build
	$(GO) install $(GO_FLAGS)

test:
	$(GO) test -v ./...

clean:
	$(GO) clean -x

.PHONY: all build install test clean