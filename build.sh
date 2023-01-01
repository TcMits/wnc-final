#!/bin/bash

set -o errexit

go mod tidy
go mod download
go build -tags netgo -ldflags '-s -w' -o createuser ./cmd/createuser
go build -tags netgo -ldflags '-s -w' -o migrate ./cmd/migrate
go build -tags netgo -ldflags '-s -w' -o app
./migrate
./createuser