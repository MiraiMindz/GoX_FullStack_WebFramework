#!/usr/bin/env bash

goimports -w ../*/*.go
go vet ../...
go run $1
