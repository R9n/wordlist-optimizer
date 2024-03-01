@echo off

set OUTPUT_FILE=wlopt.exe

echo Compiling the main program...
go build -o %OUTPUT_FILE% main.go

set "SCRIPT_DIR=%~dp0"

echo Adicionando %SCRIPT_DIR% ao PATH...
setx PATH "%PATH%%SCRIPT_DIR%%OUTPUT_FILE%" /m

echo Done!
echo Por favor, reabra o terminal para que as alterações tenham efeito.
