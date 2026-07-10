#!/bin/bash

user=$1
uid=$2

echo -e "\x1b[1;32m[INFO] to create user: ${user} uid: ${uid} in pttbbs-backend-dev-postfix-1\x1b[m"
docker exec pttbbs-backend-dev-postfix-1 useradd -s /bin/bash --uid ${uid} ${user}
docker exec pttbbs-backend-dev-postfix-1 id ${user}
