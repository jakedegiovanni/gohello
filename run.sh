#!/bin/bash

set -e

PORT=$1

./build.sh

if [[ -z $PORT ]]; then
    echo "No port passed in; defaulting to 8080"
    PORT=8080
fi

./bin/hello.exe -port $PORT
