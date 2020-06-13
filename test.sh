#!/bin/bash

set -e

VERBOSE_FLAG=$1

go test ./internal/app/... -count=1 -cover $VERBOSE_FLAG
go test ./cmd/... -p 1 -parallel 1 -count=1 -cover $VERBOSE_FLAG
