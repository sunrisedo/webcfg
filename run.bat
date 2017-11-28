@echo off
SET GOPATH=%GOPATH%;%CD%\..\..

REM @echo on

Title webcfg
go build -o webcfg.exe sunrise/webcfg && webcfg.exe
