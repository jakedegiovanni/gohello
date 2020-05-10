#!/bin/bash

if [[ ! -d bin/ ]]; then
    echo "Creating bin/ directory."
    mkdir bin
fi

go build -o bin/ cmd/**/*.go
