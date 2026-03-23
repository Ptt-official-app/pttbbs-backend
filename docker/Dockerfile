FROM golang:1.25-bookworm

ENV GOROOT=/usr/local/go
ENV PATH=${PATH}:/usr/local/go/bin

ARG GO_PTTBBS_VERSION
ARG TAG=docker

# go-pttbbs
COPY . /srv/pttbbs-backend

RUN echo "GO_PTTBBS_VERSION: ${GO_PTTBBS_VERSION} TAG: ${TAG}"

WORKDIR /srv/pttbbs-backend
RUN mkdir -p /etc/pttbbs-backend && cp docs/config/01-config.docker.ini /etc/pttbbs-backend/production.ini

WORKDIR /srv/pttbbs-backend
RUN go mod download
RUN cp /srv/pttbbs-backend/docs/config/00-config-production.go.template /go/pkg/mod/github.com/\!ptt-official-app/go-pttbbs@${GO_PTTBBS_VERSION}/ptttype/00-config-production.go
RUN go build -ldflags "-X github.com/Ptt-official-app/pttbbs-backend/types.GIT_VERSION=`git rev-parse --short HEAD` -X github.com/Ptt-official-app/pttbbs-backend/types.VERSION=`git describe --tags`" -tags ${TAG}

RUN mkdir -p /static

# cmd
WORKDIR /srv/pttbbs-backend
CMD ["/srv/pttbbs-backend/pttbbs-backend", "-ini", "production.ini"]

EXPOSE 3457
