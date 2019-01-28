@echo off

set GOOS=js
set GOARCH=wasm

rmdir /Q /S dist
mkdir dist

call :Compile basic_triangle
call :Compile rotating_cube
call :Compile splashy
exit /B %ERRORLEVEL%

:Compile
echo Compiling "%~1%"...
mkdir ".\dist\%~1%"
copy %GOROOT%\misc\wasm\wasm_exec.html ".\dist\%~1%" >NUL
copy %GOROOT%\misc\wasm\wasm_exec.js ".\dist\%~1%" >NUL
go build -o .\dist\%~1\test.wasm .\%~1\main.go