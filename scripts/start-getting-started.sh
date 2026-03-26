#!/bin/bash
#
# 1. ensure that docker_compose.dev.env file exists.
# 2. docker compose to ensure running the containers.
# 3. go build and run pttbbs-backend.

docker compose -f docker/docker-compose.dev.yaml up -d --no-recreate

ini_filename=docs/examples/etc/pttbbs-backend/production.ini
package=github.com/Ptt-official-app/pttbbs-backend/types
commit=`git rev-parse --short HEAD`
version=`git describe --tags`

go build -ldflags "-X ${package}.GIT_VERSION=${commit} -X ${package}.VERSION=${version}" && ./pttbbs-backend -ini ${ini_filename}
