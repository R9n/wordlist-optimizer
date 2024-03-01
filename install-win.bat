@echo off

set OUTPUT_FILE=wlopt.exe

echo Compiling the main program...
go build -o %OUTPUT_FILE% main.go

set "SCRIPT_DIR=%~dp0"

echo Adicionando %SCRIPT_DIR% ao PATH...
setx PATH "%PATH%;%SCRIPT_DIR%" /m

echo Done!
echo Please now reopen your terminal to make changes take effect.
