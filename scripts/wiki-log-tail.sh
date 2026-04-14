#!/bin/sh
set -eu

log_file=${1:-wiki/log.md}
lines=${2:-10}

grep '^## \[' "$log_file" | tail -n "$lines"