#!/bin/sh
go build -o ./bin/run
cp config.ini ./bin/config.ini
cp -r templates ./bin/templates