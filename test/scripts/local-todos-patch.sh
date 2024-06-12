#!/bin/bash
set -eo pipefail

source "$(dirname "${BASH_SOURCE[0]}")/local.inc"

UUID="$1"
[ "${UUID}" != "" ] || error "UUID is empty"

exec "${REPOBASEDIR}/test/scripts/curl.sh" -i -X PATCH -d @<(sed -e "s/{{createDomain.response.body.domain_id}}/${UUID}/g" < "${REPOBASEDIR}/test/data/http/patch-rhel-idm-domain.json") "${BASE_URL}/todos/${UUID}"
