#!/usr/bin/env bash

set -e

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hello main.go
docker build -t golang-hello .
rm -d hello