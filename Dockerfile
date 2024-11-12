FROM docker:dind

RUN apk add --no-cache \
    ca-certificates \
    bash \
    alpine-sdk \
    musl-dev \
    openssl \
    go \
    git \
    mingw-w64-gcc \
    python3 \
    py3-pip

ENV GOLANG_VERSION 1.21.3
ENV PYTHONUNBUFFERED=1
ENV GO111MODULE=on 

WORKDIR /app

USER root

COPY main .

RUN go mod download 

ENTRYPOINT go run ./benchmark/main.go
