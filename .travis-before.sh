#!/usr/bin/env bash

set -e

rm -rf ${GCR_REPO}
git clone https://scott-wong:${GITHUB_TOKEN}@github.com/scott-wong/gcr.git ${GCR_REPO}
