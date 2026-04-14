#!/bin/sh
set -eu

diff_range=${1:-master...HEAD}

git rev-parse --is-inside-work-tree >/dev/null 2>&1
git diff --name-only "$diff_range" | awk '
	NF == 0 { next }
	$0 ~ /^wiki\// { next }
	$0 ~ /^posts\// { next }
	$0 ~ /^bin\// { next }
	$0 ~ /^render\// { next }
	$0 == "feed.xml" { next }
	$0 == "robots.txt" { next }
	$0 == "sitemap.xml" { next }
	$0 ~ /\.log$/ { next }
	$0 ~ /\.tmp$/ { next }
	{ print }
'