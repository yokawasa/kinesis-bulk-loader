.PHONY: clean all kinesis-bulk-loader

.DEFAULT_GOAL := all

TARGETS=kinesis-bulk-loader

CUR := $(shell pwd)
OS := $(shell uname)
VERSION := $(shell cat ${CUR}/VERSION)

kinesis-bulk-loader:
	golint ${CUR}
	GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=${VERSION}" -o ${CUR}/dist/kinesis-bulk-loader_linux_amd64 ${CUR}/src
	GOOS=darwin GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=${VERSION}" -o ${CUR}/dist/kinesis-bulk-loader_darwin_amd64 ${CUR}/src
	GOOS=windows GOARCH=amd64 GO111MODULE=on go build -ldflags "-X main.buildVersion=${VERSION}" -o ${CUR}/dist/kinesis-bulk-loader_windows_amd64 ${CUR}/src

all: $(TARGETS)

clean:
	rm -rf dist
