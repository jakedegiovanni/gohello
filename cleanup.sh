#!/bin/bash

set -e

echo "Starting clean operation..."

if [[ -d bin/ ]]; then
    echo "Removing bin/ directory."
    rm -rf bin/
fi

echo "Cleaned up successfully."
