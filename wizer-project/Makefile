BROKER_BINARY=brokerApp
AUTH_BINARY= authApp
LISTENER_BINARY = listenerApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting micro-test images..."
	docker compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_auth build_listener
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker compose down
	@echo "Done!"

## build_broker: builds the broker binary as a linux executable
build_broker:
	@echo "Building broker binary..."
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o bin/${BROKER_BINARY} ./cmd/api
	@echo "Build Broker done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building auth binary..."
	cd ../auth-service && env GOOS=linux CGO_ENABLED=0 go build -o bin/${AUTH_BINARY} ./cmd/api
	@echo "Build Auth done!"

## build_listener: builds the listener binary as a linux executable
build_listener:
	@echo "Building listener binary..."
	cd ../listener-service && env GOOS=linux CGO_ENABLED=0 go build -o bin/${LISTENER_BINARY} .
	@echo "Build listener done!"



###################################################
############### PROTOBUF COMMANDS #################
###################################################
BIN_DIR = bin
PROTO_DIR = proto

SHELL := bash
	SHELL_VERSION = $(shell echo $$BASH_VERSION)
	UNAME := $(shell uname -s)
	VERSION_AND_ARCH = $(shell uname -rm)
	ifeq ($(UNAME),Darwin)
		OS = macos ${VERSION_AND_ARCH}
	else ifeq ($(UNAME),Linux)
		OS = linux ${VERSION_AND_ARCH}
	else
		$(error OS not supported by this Makefile)
	endif
	PACKAGE = $(shell head -1 ./../$@/go.mod | awk '{print $$2}')
	CHECK_DIR_CMD = test -d ./../$@ || (echo "\033[31mDirectory $@ doesn't exist\033[0m" && false)
	HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	RM_F_CMD = rm -f
	RM_RF_CMD = ${RM_F_CMD} -r

project := broker-service listener-servive auth-service

proto-all: $(project) ## generate Pbs

proto broker-service: $@ ## Generate Pbs
proto auth-service: $@ ## Generate Pbs
proto listener-service: $@ ## Generate Pbs

$(project):
	@${CHECK_DIR_CMD}
	@echo "Building proto ${@}..."
	protoc -I./../$@/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=./../$@ --go-grpc_opt=module=${PACKAGE} --go-grpc_out=./../$@ ./../$@/${PROTO_DIR}/*.proto
	@echo "Build proto ${@}!"
