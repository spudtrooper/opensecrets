#!/bin/sh
#
# Integration test
#
set -e

scripts=$(dirname $0)

$scripts/run.sh
$scripts/run.sh gl --id LA
$scripts/run.sh gl --id N00026823
$scripts/run.sh gmpfd --cid N00026823

echo "OK"