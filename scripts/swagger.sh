#!/bin/bash

if [ "$1" == "" ]; then
    echo "usage: swagger.sh [host]"
    exit 255
fi

host=$1
currentDir=`pwd`

mkdir -p swagger/v3

python -m apidoc.pttbbs_backend.gen_openapi --out swagger/v3/openapi.json --host "${host}"

docker container stop swagger-pttbbs-backend-v3
docker container rm swagger-pttbbs-backend-v3
docker run -itd --restart always --name swagger-pttbbs-backend-v3 -p 127.0.0.1:8080:8080 -e SWAGGER_JSON=/foo/v3/openapi.json -v ${PWD}/swagger:/foo swaggerapi/swagger-ui
