#!/bin/sh
set -eu

wiki_dir=${1:-wiki}
query=${2:-}

if [ -z "$query" ]; then
	echo "Usage: sh scripts/wiki-search.sh [wiki-dir] <query>" >&2
	exit 1
fi

grep -R -n -i -F -- "$query" "$wiki_dir"