#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "$0")/../../"

bash tests/docker/test_with_curl.sh
