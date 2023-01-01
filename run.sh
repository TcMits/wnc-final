#!/bin/bash

set -o errexit

go run ./cmd/migrate/main.go
go run ./cmd/createuser/main.go
go run ./main.go