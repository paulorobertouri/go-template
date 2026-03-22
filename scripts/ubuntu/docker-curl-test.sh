#!/bin/bash

cd "$(dirname "$0")/../../"

bash tests/docker/test_with_curl.sh
