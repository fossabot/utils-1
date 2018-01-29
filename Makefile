MMAKE := $(shell command -v mmake 2> /dev/null)
GORETURNS := $(shell command -v goreturns 2> /dev/null)
GOMETALINTER := $(shell command -v gometalinter 2> /dev/null)
MARKDOWNLINT := $(shell command -v markdownlint 2> /dev/null)

# dispaly make task usage
help:
ifndef MMAKE
	@echo "Please run 'make boostrap'"
else
	mmake help -v
	@echo "You may want to 'alias make=mmake'"
endif
.PHONY: help
.DEFAULT_GOAL := help

# run all the checks; vet, lint, test
all: vet lint test
.PHONY: all

# check golang syntax and format
#
# 	- runs go vet, fmt with -s, goimports, and goreturns
vet:
	go vet ./...
	gofmt -s -l .
	goreturns -b -i -l .
.PHONY: vet

# check style and lint
#
# 	- runs gometalinter with all linters enable except gas
lint:
	gometalinter \
		--enable-all \
		--sort=path \
		--aggregate \
		--tests \
		./...
.PHONY: lint

# run golang tests
#
# 	- runs go test verbosely
test:
	go test -v ./...
.PHONY: test

# build the static linked binary
build:
	CGO_ENABLED=0 go build ./...
.PHONY: build

precommit: vet lint test
	markdownlint $(shell find . -type f -name '*.md')
.PHONY: name

# installs development tools
bootstrap:
ifndef MMAKE
	go get -u github.com/tj/mmake/cmd/mmake
	@echo "Alias make=mmake"
else
	@echo "Already Installed: mmake. Skipping."
endif
ifndef GORETURNS
	go get -u sourcegraph.com/sqs/goreturns
else
	@echo "Already Installed: goreturns. Skipping."
endif
ifndef GOMETALINTER
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install
else
	@echo "Already Installed: gometalinter. Skipping."
endif
ifndef MARKDOWNLINT
	npm install --global markdownlint
else
	@echo "Already Installed: markdownlint. Skipping."
endif
.PHONY: bootstrap

MAKEFLAGS += --no-builtin-rules
.SUFFIXES:
