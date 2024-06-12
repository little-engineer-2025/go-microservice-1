#!/bin/bash
set -eo pipefail

source "$(dirname "${BASH_SOURCE[0]}")/local.inc"

exec "${REPOBASEDIR}/test/scripts/curl.sh" -i "${BASE_URL}/todos"
