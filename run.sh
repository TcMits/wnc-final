#!/bin/bash

set -o errexit

go run ./cmd/migrate
go run ./cmd/createuser
go run /main.go