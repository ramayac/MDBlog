#!/bin/sh
set -eu

wiki_dir=${1:-wiki}
diff_range=${2:-master...HEAD}
lines=${3:-10}
script_dir=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
candidate_file=$(mktemp)

trap 'rm -f "$candidate_file"' EXIT HUP INT TERM
sh "$script_dir/wiki-ingest-candidates.sh" "$diff_range" > "$candidate_file" || true

if [ ! -s "$candidate_file" ]; then
	echo "no wiki refresh needed: no ingest candidates for diff range $diff_range"
	exit 0
fi

echo "== wiki files =="
sh "$script_dir/wiki-list.sh" "$wiki_dir"
echo
echo "== recent log =="
sh "$script_dir/wiki-log-tail.sh" "$wiki_dir/log.md" "$lines"
echo
echo "== changed files =="
sh "$script_dir/wiki-changed.sh" "$diff_range" || true
echo
echo "== ingest candidates =="
cat "$candidate_file"
echo
echo "== lint =="
sh "$script_dir/wiki-lint.sh" "$wiki_dir"