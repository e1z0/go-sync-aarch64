# This fork

This fork enables you to run Brave Sync Server v2 on ARM64 (aarch64) docker containers. Docker swarm supported by default. Please note that nfs file systems are currently unsupported.

```
version: '3.4'

services:
  web:
    image: nulldevil/brave-go-sync:aarch64
    depends_on:
      - dynamo-local
      - redis
    networks:
      - sync
      - traefik_public
    command: "./main"
    deploy:
      resources:
        limits:
          cpus: '2.75' # one cpu
          memory: 480M
      placement:
        constraints: [node.labels.name == node0]
      labels:
        # traefik common
        - traefik.enable=true
        - traefik.docker.network=traefik_public
        # traefikv1
        - traefik.frontend.rule=Host:brave.domain.com
        - traefik.port=8295
        # traefikv2
        - "traefik.http.routers.brave.rule=Host(`brave.domain.com`)"
        - "traefik.http.services.brave.loadbalancer.server.port=8295"
    environment:
      - PPROF_ENABLED=true
      - SENTRY_DSN
      - ENV=local
      - DEBUG=1
      - AWS_ACCESS_KEY_ID=#
      - AWS_SECRET_ACCESS_KEY=#
      - AWS_REGION=us-west-2
      - AWS_ENDPOINT=http://dynamo-local:8000
      - TABLE_NAME=client-entity-dev
      - REDIS_URL=redis:6379
  dynamo-local:
    image: nulldevil/go-sync_dynamo:aarch64
    volumes:
     - ./data/db:/db
    user: root
    deploy:
      resources:
        limits:
          cpus: '2.75' # one cpu
          memory: 480M
      placement:
        constraints: [node.labels.name == node0]
    networks:
      - sync
  redis:
    image: public.ecr.aws/ubuntu/redis:latest
    volumes:
     - ./data/redis:/var/lib/redis
    deploy:
      resources:
        limits:
          cpus: '2.75' # one cpu
          memory: 480M
      placement:
        constraints: [node.labels.name == node0]
    environment:
      - ALLOW_EMPTY_PASSWORD=yes
    networks:
      - sync

networks:
  traefik_public:
    external: true
  sync:
    driver: overlay
    ipam:
      config:
        - subnet: 172.16.21.0/24
```
# Brave browser configuration MacOS

You can simple run browser with parameter:

```
/Applications/Brave\ Browser.app/Contents/MacOS/Brave\ Browser --sync-url=https://brave.domain.com/v2
```
But if you need more simplicity just follow these instructions:

Add file named "launcher" to **/Applications/Brave Browser.app/Contents/MacOS/launcher** with these contents
```
#!/bin/sh
RealBin="Brave Browser"
AppDir="$(dirname "$0")"
exec "$AppDir/$RealBin" --sync-url=https://brave.domain.com/v2 "$@"
```
Now edit **Info.plist**:
```
/usr/libexec/PlistBuddy -c "Set CFBundleExecutable launcher" /Applications/Brave\ Browser.app/Contents/Info.plist
```
Check if it's changed:
```
/usr/libexec/PlistBuddy -c "Print CFBundleExecutable" /Applications/Brave\ Browser.app/Contents/Info.plist
```
If everything is ok. Now we need to move "Brave Browser" from **/Applications** to your **Desktop** (use **Shift+CMD** while dragging). This will trick MacOS cache to re-read **Info.plist**.

# Brave browser configuration MacOS (another way)

Make an apple script with contents:
```
do shell script "\"/Applications/Brave Browser.app/Contents/MacOS/Brave Browser\" --sync-url=https://brave.domain.com/v2"
quit
```
Save it as **apple script** and then export as **app bundle** on the desktop. Now open **~/Desktop/Your saved bundle.app/Contents/Info.plist**
And after the first <dict> add these two lines:
```
<key>NSUIElement</key>
<true/>
```
Save it and whola!


# Brave browser configuration Windows

On the windows system it's never been so simple, just right click on the desktop shortcut of "Brave Browser" and select properties, or just press alt+enter. Now under the **target** field add parameter after brave.exe"
```
--sync-url=https://brave.domain.com/v2
```
Full line should look like this:
```
"C:\Program Files\BraveSoftware\Brave-Browser\Application\brave.exe" --sync-url=https://brave.domain.com/v2
```

# Brave Sync Server v2

A sync server implemented in go to communicate with Brave sync clients using
[components/sync/protocol/sync.proto](https://cs.chromium.org/chromium/src/components/sync/protocol/sync.proto).
Current Chromium version for sync protocol buffer files used in this repo is Chromium 116.0.5845.183.

This server supports endpoints as bellow.
- The `POST /v2/command/` endpoint handles Commit and GetUpdates requests from sync clients and return corresponding responses both in protobuf format. Detailed of requests and their corresponding responses are defined in `schema/protobuf/sync_pb/sync.proto`. Sync clients are responsible for generating valid access tokens and present them to the server in the Authorization header of requests.

Currently we use dynamoDB as the datastore, the schema could be found in `schema/dynamodb/table.json`.

## Developer Setup
1. [Install Go 1.18](https://golang.org/doc/install)
2. [Install GolangCI-Lint](https://github.com/golangci/golangci-lint#install)
3. [Install gowrap](https://github.com/hexdigest/gowrap#installation)
4. Clone this repo
5. [Install protobuf protocol compiler](https://github.com/protocolbuffers/protobuf#protocol-compiler-installation) if you need to compile protobuf files, which could be built using `make protobuf`.
6. Build via `make`

## Local development using Docker and DynamoDB Local
1. Clone this repo
2. Run `make docker`
3. Run `make docker-up`
4. For running unit tests, run `make docker-test`

# Updating protocol definitions
1. Copy the `.proto` files from `components/sync/protocol` in `chromium` to `schema/protobuf/sync_pb` in `go-sync`.
2. Copy the `.proto` files from `components/sync/protocol` in `brave-core` to `schema/protobuf/sync_pb` in `go-sync`.
3. Run `make repath-proto` to set correct import paths in `.proto` files.
4. Run `make proto-go-module` to add the `go_module` option to `.proto` files.
5. Run `make protobuf` to generate the Go code from `.proto` definitions.

## Prometheus Instrumentation
The instrumented datastore and redis interfaces are generated, providing integration with Prometheus.  The following will re-generate the instrumented code, required when updating protocl definitions:

```
make instrumented
```

Changes to `datastore/datastore.go` or `cache/cache.go` should be followed with the above command.
