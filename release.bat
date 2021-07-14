@echo off
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -ldflags="-s -w" -o ./bin/win.exe
copy /y config.ini .\bin\config.ini
rd /s/q .\bin\templates
xcopy /i /e /h /y templates .\bin\templates
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -ldflags="-s -w" -o ./bin/mac