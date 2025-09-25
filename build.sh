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

TARGET="./bin/portables/GoFileEncoder_portable_${GOOS}_${GOARCH}${EXT}"

rm $TARGET

echo "Building from ./pkg/ to $TARGET"

echo Creating assets
go-bindata -pkg "assets" -o "assets/bindata.go" "LICENSE"

echo Building
go build -o "./bin/portables/GoFileEncoder_portable_${GOOS}_${GOARCH}${EXT}" "./pkg/GoFileEncoder.go"