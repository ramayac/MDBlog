#!/bin/sh
set -eu

wiki_dir=${1:-wiki}
index_file="$wiki_dir/index.md"
log_file="$wiki_dir/log.md"
status=0

for required in \
	"$wiki_dir/README.md" \
	"$wiki_dir/index.md" \
	"$wiki_dir/log.md" \
	"$wiki_dir/schema.md" \
	"$wiki_dir/phases.md" \
	"$wiki_dir/repo-map.md" \
	"$wiki_dir/operations/ingest.md" \
	"$wiki_dir/operations/query.md" \
	"$wiki_dir/operations/lint.md"
do
	if [ ! -f "$required" ]; then
		echo "missing required wiki file: $required" >&2
		status=1
	fi
done

if [ -f "$index_file" ]; then
	target_file=$(mktemp)
	trap 'rm -f "$target_file"' EXIT HUP INT TERM
	sed -n 's/.*](\([^)]*\)).*/\1/p' "$index_file" > "$target_file"
	while IFS= read -r target
	do
		[ -n "$target" ] || continue
		case "$target" in
			http://*|https://*|\#*)
				continue
				;;
		esac
		if [ ! -f "$wiki_dir/$target" ]; then
			echo "broken index link: $target" >&2
			status=1
		fi
	done < "$target_file"
	rm -f "$target_file"
	trap - EXIT HUP INT TERM
fi

if [ -f "$log_file" ]; then
	invalid_log_headings=$(grep '^## \[' "$log_file" | grep -Ev '^## \[[0-9]{4}-[0-9]{2}-[0-9]{2}\] [^|]+ \| .+$' || true)
	if [ -n "$invalid_log_headings" ]; then
		echo "invalid log headings:" >&2
		echo "$invalid_log_headings" >&2
		status=1
	fi
fi

markers=$(grep -R -n -E 'TODO:|TBD:|UNKNOWN:' "$wiki_dir" | grep -v 'grep -R "TODO:' || true)
if [ -n "$markers" ]; then
	echo "marker findings:" >&2
	echo "$markers" >&2
	status=1
fi

if [ "$status" -eq 0 ]; then
	echo "wiki lint OK"
fi

exit "$status"