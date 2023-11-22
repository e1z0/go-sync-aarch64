GIT_VERSION := $(shell git describe --abbrev=8 --dirty --always --tags)
GIT_COMMIT := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date +%s)

.PHONY: all build test lint clean

all: lint test build

repath-proto:
	find ./schema/protobuf/sync_pb | grep "\.proto$\" | xargs sed -i 's/import \"components\/sync\/protocol\//import \"/g'
	find ./schema/protobuf/sync_pb | grep "\.proto$\" | xargs sed -i 's/import \"brave\/components\/sync\/protocol\//import \"/g'

proto-go-module:
	find ./schema/protobuf/sync_pb | grep "\.proto$\" | xargs sed -i 's/syntax = \"proto2\";/syntax = \"proto2\";\n\noption go_package = \"\.\/sync_pb\";/'

protobuf:
	protoc -I schema/protobuf/sync_pb/ schema/protobuf/sync_pb/*.proto --go_out=schema/protobuf/

build:
	go run main.go

test:
	go test -v ./...

lint:
	docker run -t --rm -v "$$(pwd):/app" -w /app golangci/golangci-lint golangci-lint run -v

clean:
	rm -f sync-server

docker:
	COMMIT=$(GIT_COMMIT) VERSION=$(GIT_VERSION) BUILD_TIME=$(BUILD_TIME) docker-compose build
docker_aarch64:
	COMMIT=$(GIT_COMMIT) VERSION=$(GIT_VERSION) BUILD_TIME=$(BUILD_TIME) docker-compose -f docker-compose-aarch64.yml build

docker-up:
	COMMIT=$(GIT_COMMIT) VERSION=$(GIT_VERSION) BUILD_TIME=$(BUILD_TIME) docker-compose up
docker-up_aarch64:
	COMMIT=$(GIT_COMMIT) VERSION=$(GIT_VERSION) BUILD_TIME=$(BUILD_TIME) docker-compose -f docker-compose-aarch64.yml up -d
docker-down_aarch64:
	COMMIT=$(GIT_COMMIT) VERSION=$(GIT_VERSION) BUILD_TIME=$(BUILD_TIME) docker-compose -f docker-compose-aarch64.yml down
docker-upload_aarch64:
	docker tag go-sync-aarch64_dynamo-local:latest nulldevil/go-sync_dynamo:aarch64
	docker push nulldevil/go-sync_dynamo:aarch64
	docker tag go-sync-aarch64_dev:latest nulldevil/brave-go-sync:aarch64
	docker push nulldevil/brave-go-sync:aarch64

docker-test:
	COMMIT=$(GIT_COMMIT) VERSION=$(GIT_VERSION) BUILD_TIME=$(BUILD_TIME) docker-compose -f docker-compose.yml run --rm dev make test

instrumented:
	gowrap gen -p github.com/brave/go-sync/datastore -i Datastore -t ./.prom-gowrap.tmpl -o ./datastore/instrumented_datastore.go
	gowrap gen -p github.com/brave/go-sync/cache -i RedisClient -t ./.prom-gowrap.tmpl -o ./cache/instrumented_redis.go
clean_docker:
	docker image prune -a -f
	docker volume prune -f
	docker network prune -f
	docker system prune --volumes -f
