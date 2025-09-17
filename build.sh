#!/bin/bash

GOOS=$1
GOARCH=$2
EXT=

if [ "$GOOS" == "" ]; then
    GOOS=$(go env GOOS)
fi

if [ "$GOARCH" == "" ]; then
    GOARCH=$(go env GOARCH)
fi

if [ "$GOOS" == "windows" ]; then
    EXT=".exe"
fi

echo "Building from ./pkg/ to ./bin/portables/GoFileEncoder_portable_${GOOS}_${GOARCH}${EXT}"

mkdir -p "./bin/portables"

go-bindata -pkg "assets" -o "assets/bindata.go" "LICENSE"
go build -o "./bin/portables/GoFileEncoder_portable_${GOOS}_${GOARCH}${EXT}" "./pkg/GoFileEncoder.go"