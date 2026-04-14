#!/bin/sh
set -eu

diff_range=${1:-master...HEAD}

git rev-parse --is-inside-work-tree >/dev/null 2>&1
git diff --name-only "$diff_range" | awk 'NF && $0 !~ /^wiki\// { print }'