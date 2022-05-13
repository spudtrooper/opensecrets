#!/bin/sh
#
# Run unit tests
#
set -e

scripts=$(dirname $0)

go test ./...