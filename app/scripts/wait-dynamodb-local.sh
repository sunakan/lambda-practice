#!/bin/sh
set -eu

################################################################################
# How to run
# $ sh scripts/wait-dynamodb-local.sh http://localhost:8000
################################################################################
readonly MAX_TRY_COUNT=10
readonly SLEEP_SECOND=3
readonly URL="$1"

echo "Wait for ${URL}"
try=0
until curl -X POST --connect-timeout 10 --max-time 10 ${URL} > /dev/null 2>&1; do
  echo -n '.' >&2
  sleep ${SLEEP_SECOND}
  try=$(expr $try + 1)
  if [ ${try} -ge ${MAX_TRY_COUNT} ]; then
    echo ''
    echo "Failed to wait for ${URL}"
    exit 1
  fi
done
echo "You can connect to ${URL}"
