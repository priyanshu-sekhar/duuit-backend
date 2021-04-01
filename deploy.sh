#!/usr/bin/env bash
set -xe

# build binary
GOARCH=amd64 GOOS=linux go build -o bin/application

eb deploy
