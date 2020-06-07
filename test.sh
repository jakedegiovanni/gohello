#!/bin/bash

VERBOSE_FLAG=$1

go test ./cmd/... ./internal/app/... -count=1 -cover $VERBOSE_FLAG
