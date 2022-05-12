#!/bin/sh
#
# All tests
#
set -e

scripts=$(dirname $0)

go test ./...
$scripts/integtest.sh