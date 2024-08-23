BROKER_BINARY=brokerApp
USER_BINARY = userApp
AUTH_BINARY= authApp
LISTENER_BINARY = listenerApp
PROFILE_BINARY = profileApp
MIGRATION_BINARY = migrationApp
ROUTINE_BINARY = routineApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting micro-test images..."
	docker compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_broker build_user build_auth build_profile build_routine build_migration build_front
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
	cd ../broker-service && env GOOS=linux CGO_ENABLED=0 go build -o bin/${BROKER_BINARY} ./cmd/app
	@echo "Build Broker done!"

## build_auth: builds the auth binary as a linux executable
build_auth:
	@echo "Building Auth service binary..."
	cd ../authentication-service && env GOOS=linux CGO_ENABLED=0 go build -o bin/${AUTH_BINARY} ./cmd/app
	@echo "Build Auth done!"

build_user:
	@echo "Building User Service binary..."
	cd ../user-service && env GOOS=linux CGO_ENABLED=0 go build -o bin/${USER_BINARY} ./cmd/app
	@echo "Build User done!"

build_profile:
	@echo "Building Profile Service binary..."
	cd ../profile-service && env GOOS=linux CGO_ENABLED=0 go build -o bin/${PROFILE_BINARY} ./cmd/app
	@echo "Build Profile done!"

build_routine:
		@echo "Building Routine Service binary..."
		cd ../routine-service && env GOOS=linux CGO_ENABLED=0 go build -o bin/${ROUTINE_BINARY} ./cmd/app
		@echo "Build Routine done!"

build_migration:
	@echo "Buildingn Migration Service Binary"
	cd ../migration-service && env GOOS=linux CGO_ENABLED=0 go build -o bin/${MIGRATION_BINARY} ./cmd/app
	@echo "Build Migration done!"

build_front:
	@echo "Building frontend in TS react..."
	cd ../wizer-frontend && npm run build
	@echo "Build frontend done!"



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

project := broker-service user-service auth-service profile-service routine-service

proto-all: $(project) ## generate Pbs

proto broker-service: $@ ## Generate Pbs
proto user-service: $@ ## Generate Pbs 
proto auth-service: $@ ## Generate Pbs
proto profile-service: $@ ## Generate Pbs
proto routine-service: $@ ## Generate Pbs

$(project):
	@${CHECK_DIR_CMD}
	@echo "Building proto ${@}..."
	protoc -I./../$@/${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=./../$@ --go-grpc_opt=module=${PACKAGE} --go-grpc_out=./../$@ ./../$@/${PROTO_DIR}/*.proto
	@echo "Build proto ${@}!"
