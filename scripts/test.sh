#!/bin/sh
#
# All tests
#
set -e

scripts=$(dirname $0)

$scripts/unittest.sh
$scripts/integtest.sh