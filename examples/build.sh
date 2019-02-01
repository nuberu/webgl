#!/usr/bin/env bash

MISC_PATH=${GOROOT}/misc/wasm
MISC_PATH=/usr/share/go-1.11/misc/wasm/

function compile {
    echo Compiling "$1"...
    mkdir ./dist/$1
    cp ${MISC_PATH}/wasm_exec.html "./dist/$1/" > /dev/null
    cp ${MISC_PATH}/wasm_exec.js "./dist/$1/" > /dev/null
    go build -o ./dist/$1/test.wasm ./$1/main.go
}

export GOOS=js
export GOARCH=wasm

rm -Rf ./dist
mkdir ./dist

compile basic_triangle
compile rotating_cube