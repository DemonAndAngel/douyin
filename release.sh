#!/bin/sh
CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/mac
cp config.ini ./bin/config.ini
rm -rf ./bin/templates
cp -r templates ./bin/templates
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o ./bin/win.exe