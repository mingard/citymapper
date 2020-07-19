# CityMapper Coding Test
# Arthur Mingard

#
# Variables
#
OWNER ?= mingard
PROJECT ?= citymapper
VERSION ?= 1.0.0
BUILD ?= $(shell git rev-parse --short HEAD)
ENV ?= develop
ENVFILE ?= envfile.$(ENV)

# Current build settings
ARCH ?= amd64
NAME := $(PROJECT)-$(ARCH)
TARGET := $(shell pwd)/bin/$(NAME)
ENTRYPOINT := cmd/$(PROJECT).go

# Go command variables
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Linter variables
BIN_DIR := $(GOPATH)/bin
GOMETALINTER := $(BIN_DIR)/gometalinter

# Compiler variables
CFLAGS =
LDFLAGS = -ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"
SRC ?= $(shell find . -type f -name '*.go' -not -path "./vendor/*")
PKGS = $(shell go list ./... | grep -v /vendor)


# Determine compiler vars based on arch type
CVARS = CGO_ENABLED=0 GOARCH=$(ARCH)
ifeq ($(ARCH), arm)
	CVARS += GOARM=6
endif

# Environment variables
include $(ENVFILE)
export $(shell sed 's/=.*//' $(ENVFILE))

#
# Rules
#
.DEFAULT_GOAL: $(TARGET)
.PHONY: build clean run fmt lint

$(TARGET): $(SRC)
	$(CVARS) $(GOBUILD) $(CFLAGS) $(LDFLAGS) -o "$(TARGET)" $(ENTRYPOINT)

build: $(TARGET)
	@true

clean:
	rm -f $(TARGET)
	
run: build
	@$(TARGET) $(filter-out $@,$(MAKECMDGOALS))

%:
	@true

version:
	@echo $(VERSION)

test:
	$(GOTEST) $(PKGS)

lint: $(GOMETALINTER)
	$(GOMETALINTER) ./... --vendor --fast --disable=maligned

$(GOMETALINTER):
	$(GOGET) -u github.com/alecthomas/gometalinter
	$(GOMETALINTER) --install 1>/dev/null

fmt:
	gofmt -l -w $(SRC)