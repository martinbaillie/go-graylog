SHELL 	:= /bin/bash
DIR 	:=$(shell pwd)
TIME 	:=$(shell date '+%Y-%m-%dT%T%z')
HASH 	:=$(shell git rev-parse --verify --short HEAD || echo "")
BRANCH 	:=$(shell git rev-parse --abbrev-ref HEAD)
API_VER :=$(shell grep "\"version\": \"" api/swagger.json | awk '{print $$2}' \
	| sed 's/\"//g')

SRC 		?=$(shell go list ./...)
SRCFILES	?=$(shell find . -type f -name '*.go' -not -path './vendor/*')
SRCDIRS		?=$(shell find . -mindepth 1 -maxdepth 1 -type d | \
			  grep -Ev 'vendor|pkg')

LDFLAGS ?=-s -w -extld ld -extldflags -static \
		  -X 'main.BuildTime=$(TIME)' \
		  -X 'main.Project=$(PROJECT)' \
		  -X 'main.Version=$(VERSION)' \
		  -X 'main.APIVersion=$(API_VER)'
FLAGS	?=-a -installsuffix cgo -ldflags "$(LDFLAGS)"
GOOS 	?=darwin windows linux
GOARCH 	?=amd64

VERSION ?=$(BRANCH)-$(HASH)
PROJECT	?=graylog

.PHONY: clean fix fmt metalint test build help

clean:: ## Removes binary and generated files
	@echo ">> cleaning"
	@rm -f cmd/graylog/graylog-* 

fix:: ## Runs the Golang fix tool
	@echo ">> fixing"
	@go fix $(SRC)

fmt:: ## Formats source code according to the Go standard
	@echo ">> formatting"
	@gofmt -w -s -l $(SRCFILES)

metalint: ## Heavyweight linting
	@echo ">> metalinting"
	@declare -i res=0; inc(){ res+=1; }; trap inc ERR; for d in $(SRCDIRS); do \
		gometalinter --config=.metalinter.json $$d/...; done; exit $$res

test: clean fix fmt metalint

build:: test ## Builds for all arch ($GOARCH) and OS ($GOOS)
	@echo ">> building"
	@cd cmd/graylog && for arch in ${GOARCH}; do \
		for os in ${GOOS}; do \
			echo ">>>> $${os}/$${arch}"; \
			env GOOS=$${os} GOARCH=$${arch} \
			CGO_ENABLED=0 go build $(FLAGS) \
			-o $(PROJECT)-$${os}-$${arch}; \
		done; \
	done

# A help target including self-documenting targets (see the awk statement)
help: ## This help target
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / \
		{printf "\033[36m%-30s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)
