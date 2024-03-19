#!/usr/bin/env bash

air="$HOME/lib/go/bin/air"

clear
goimports -w ./*.go
go vet .
if [[ -f $air ]]; then
    $air
else
    go run "$(pwd)/main.go"
fi
