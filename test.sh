#!/bin/bash

VERBOSE_FLAG=$1

go test ./cmd/... ./pkg/... -count=1 -cover $VERBOSE_FLAG
