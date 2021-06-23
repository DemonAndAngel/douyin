#!/bin/sh
go build -o ./bin/run
cp config.ini ./bin/config.ini
rm -rf ./bin/templates
cp -r templates ./bin/templates