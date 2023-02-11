SHELL := /bin/bash
PROJECT_ROOT_DIR := $(shell git rev-parse --show-toplevel)

BIN_DIR := $(PROJECT_ROOT_DIR)/bin
BUILD_DIR := $(PROJECT_ROOT_DIR)/build

.DEFAULT_GOAL := help

include mk/dbmate.mk
include mk/ci.mk
include mk/generate.mk

.PHONY: gomod
gomod:          ## Run mod tidy and vendor
	go mod tidy -compat=1.19
	go mod vendor

.PHONY: setup
setup:          ## Setup Locally
	@ sh ./scripts/setup-local.sh

#
# Commands which are intended to run inside the development container
#
.PHONY: api-run
api-run:        ## Runs server with watch mode on in docker development container
	gin \
		--immediate \
		--port 8602 \
		--appPort 8080 \
		--build ./cmd/server \
		--bin ".gin-server" \
		run ./cmd/server/main.go

.PHONY: help
help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'
