#!/bin/sh

OS="linux"
ARCH="amd64"

if [ "$1" == "mac" -o "$1" == "windows" ]; then
    OS=$1
    if [ "$OS" == "mac" ]; then
        OS="darwin"
    fi
fi
echo "build for ${OS}, arch: ${ARCH}..."

export GOOS=$OS
export GOARCH=$ARCH
export CGO_ENABLED=0
go build  -ldflags '-extldflags "-static"'  -o bin/gin-api-frame_server gin-api-frame/cmd/api/
