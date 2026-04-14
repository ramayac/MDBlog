#!/bin/sh
set -eu

wiki_dir=${1:-wiki}

find "$wiki_dir" -type f -name '*.md' | LC_ALL=C sort | while IFS= read -r file
do
	grep -n -E '^#{1,6} ' "$file" || true
done