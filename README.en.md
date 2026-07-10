# pttbbs-backend

[![Go](https://github.com/Ptt-official-app/pttbbs-backend/actions/workflows/go.yml/badge.svg)](https://github.com/Ptt-official-app/pttbbs-backend/actions/workflows/go.yml)
[![GoDoc](https://pkg.go.dev/badge/github.com/Ptt-official-app/pttbbs-backend?status.svg)](https://pkg.go.dev/github.com/Ptt-official-app/pttbbs-backend?tab=doc)
[![codecov](https://codecov.io/gh/Ptt-official-app/pttbbs-backend/branch/main/graph/badge.svg)](https://codecov.io/gh/Ptt-official-app/pttbbs-backend)

## README Translation

* [English](https://github.com/Ptt-official-app/pttbbs-backend/blob/main/README.en.md)
* [正體中文](https://github.com/Ptt-official-app/pttbbs-backend/blob/main/README.zh-TW.md)

## Overview

go implementation of [the middleware of PTT](https://hackmd.io/@twbbs/Root#%E6%9E%B6%E6%A7%8B%E5%9C%96).

With [Ptt-official-app pttbbs](https://github.com/ptt-official-app/go-pttbbs), pttbbs-backend intends to be the modern BBS framework.

## Demo Site

* [dev](https://www.devptt.dev)
* [term (PttChrome)](https://term.devptt.dev)

## Getting Started

You can start this project with the following steps:

1. install [docker](https://www.docker.com/)
2. `./scripts/start-getting-started.sh`
3. try the api at [http://localhost:8080](http://localhost:8080)

### [Legacy PTTBBS](https://github.com/ptt/pttbbs)

You can go to the legacy PTTBBS with the following steps:

4. `telnet localhost 8888`
5. login with SYSOP/123123.

## API

You can try the api at [https://doc.devptt.dev](https://doc.devptt.dev).

You can copy the `curl` command from the api website if you encounter `CORS` issue.

## Coding Convention

We use the following libraries for coding convention:

* [gotests](https://github.com/cweill/gotests) for test-generation
* [gofumpt](https://github.com/mvdan/gofumpt) for formatting

## Discussing / Reviewing / Questioning the code

Besides creating issues, you can do the following
to discuss / review / question the code:

* `git clone` this repo
* create a review-\[topic\] branch
* commenting at the specific code-region.
* pull-request
* start discussion.

## Develop

You can start developing by forking this repository.

## Unit-Test

You can do unit-test with:

* `./scripts/test.sh`

You can check coverage with:

* `./scripts/coverage.sh`

## Swagger

You can run swagger with:

* setup python virtualenv.
* `cd apidoc; pip install .; cd ..`
* `./scripts/swagger.sh [host]`
* go to [http://localhost:8080](http://localhost:8080)

## Schema definition

* [https://github.com/Ptt-official-app/pttbbs-backend/tree/main/schema](https://github.com/Ptt-official-app/pttbbs-backend/tree/main/schema)

## Repository Naming

* 2024-06-26: This repo is mainly based on pttbbs. We rename the repo as pttbbs-backend.
* 2022-12-16: The reason why this repo is called go-openbbsmiddleware is because previously the [.NET ASP](https://github.com/Ptt-official-app/AspCoreOpenBBSMiddleware) developers envisioned that the scope of this [middleware](https://hackmd.io/@twbbs/Root#%E6%9E%B6%E6%A7%8B%E5%9C%96) can include other versions of bbs (Maple/中山之島). The naming of this repo followed the naming convention at that time.
