#!/bin/bash
set -eo pipefail

source "$(dirname "${BASH_SOURCE[0]}")/local.inc"

UUID="$1"
[ "${UUID}" != "" ] || error "UUID is empty"

exec "${REPOBASEDIR}/test/scripts/curl.sh" -i -X DELETE "${BASE_URL}/todos/${UUID}"
