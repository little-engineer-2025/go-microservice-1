#!/bin/bash
set -eo pipefail

source "$(dirname "${BASH_SOURCE[0]}")/local.inc"

exec "${REPOBASEDIR}/test/scripts/curl.sh" -i -X POST -d @"${REPOBASEDIR}/test/data/http/todo-create.json" "${BASE_URL}/todos"
