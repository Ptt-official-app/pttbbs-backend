#!/bin/bash
#
# 1. prepare docs/examples/var/mail
# 2. docker compose to ensure running the containers.
# 3. swagger to generate the doc and the docker containers for localhost:8080
# 4. go build and run pttbbs-backend.

echo -e "\x1b[1;32m[Info] 1. prepare docs/examples/var/mail\x1b[m"
mkdir -p docs/examples/var/mail
chmod 777 docs/examples/var/mail

echo -e "\x1b[1;32m[Info] 2. docker compose\x1b[m"
docker compose -f docker/docker-compose.dev.yaml up -d --no-recreate

echo -e "\x1b[1;32m[Info] 3. swagger for localhost:8080\x1b[m"
./scripts/swagger.sh localhost:3457

echo -e "\x1b[1;32m[Info] 4. go build and run ./pttbbs-backend\x1b[m"
ini_filename=docs/examples/etc/pttbbs-backend/production.ini
package=github.com/Ptt-official-app/pttbbs-backend/types
commit=`git rev-parse --short HEAD`
version=`git describe --tags`

go build -ldflags "-X ${package}.GIT_VERSION=${commit} -X ${package}.VERSION=${version}" && ./pttbbs-backend -ini ${ini_filename}
