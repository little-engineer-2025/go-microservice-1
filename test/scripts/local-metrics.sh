#!/bin/bash
set -eo pipefail

source "$(dirname "${BASH_SOURCE[0]}")/local.inc"

# TODO Update this BASE_URL for your API
BASE_URL="http://localhost:9000"

exec "${REPOBASEDIR}/test/scripts/curl.sh" -i "${BASE_URL}/metrics"
