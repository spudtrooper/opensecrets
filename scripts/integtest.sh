#!/bin/sh
#
# Integration test
#
set -e

scripts=$(dirname $0)

# TODO: Add these as unit tests, also.
$scripts/run.sh
$scripts/run.sh GetLegislators --id LA
$scripts/run.sh GetLegislators --id N00026823
$scripts/run.sh GetMemPFDprofile --cid N00026823
$scripts/run.sh GetMemPFDprofile --cid N00026823 --year 2016
$scripts/run.sh GetCandSummary --cid N00026823
$scripts/run.sh GetCandSummary --cid N00026823 --cycle 2016
$scripts/run.sh GetCandContrib --cid N00026823
$scripts/run.sh GetCandContrib --cid N00026823 --cycle 2016
$scripts/run.sh GetCandIndustry --cid N00026823
$scripts/run.sh GetCandIndustry --cid N00026823 --cycle 2016
$scripts/run.sh GetCandByInd --cid N00026823 --ind M02
$scripts/run.sh GetCandByInd --cid N00026823 --ind M02 --cycle 2016
$scripts/run.sh GetCandSector --cid N00026823
$scripts/run.sh GetCandSector --cid N00026823 --cycle 2016
$scripts/run.sh GetCongCmteIndus -cmte HARM -ind F10
$scripts/run.sh GetCongCmteIndus -cmte HARM -ind F10  --congno 11
$scripts/run.sh GetOrgs --org Goldman
$scripts/run.sh GetOrgSummary --org_id D000000125

echo "OK"