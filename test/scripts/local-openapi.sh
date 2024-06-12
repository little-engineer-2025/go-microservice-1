#!/bin/bash
set -eo pipefail

source "$(dirname "${BASH_SOURCE[0]}")/local.inc"

# TODO Update this base URL for your API
BASE_URL="http://localhost:8000/api/todo/v1"

exec "${REPOBASEDIR}/test/scripts/curl.sh" -i "${BASE_URL}/openapi.json"
