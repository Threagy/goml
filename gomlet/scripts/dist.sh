#!/bin/bash

if [ -d "/dist/gomlet" ]; then
  rm -rf /dist/gomlet
fi

mkdir -p /dist/gomlet

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /dist/gomlet/gomlet .
cp -r ./proto /dist/gomlet/proto
