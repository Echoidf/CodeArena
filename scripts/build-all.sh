#!/usr/bin/env bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build  -o release/code-arena-mac main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build  -o release/code-arena-mac-arm64 main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o release/code-arena-linux main.go
CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build  -o release/code-arena-linux-arm64 main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  -o release/code-arena-windows.exe main.go
CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build  -o release/code-arena-windows-arm64.exe main.go