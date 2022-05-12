#!/bin/sh
#
# Integration test
#
set -e

scripts=$(dirname $0)

$scripts/run.sh
$scripts/run.sh GetLegislators --id LA
$scripts/run.sh GetLegislators --id N00026823
$scripts/run.sh GetMemPFDprofile --cid N00026823
$scripts/run.sh GetMemPFDprofile --cid N00026823 --year 2016
$scripts/run.sh GetCandSummary --cid N00026823
$scripts/run.sh GetCandSummary --cid N00026823 --cycle 2016
$scripts/run.sh GetCandContrib --cid N00026823
$scripts/run.sh GetCandContrib --cid N00026823 --cycle 2016

echo "OK"