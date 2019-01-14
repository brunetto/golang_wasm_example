#!/usr/bin/env bash

echo "building wasm lib" 

GOARCH=wasm GOOS=js go build -o main.wasm main.go
