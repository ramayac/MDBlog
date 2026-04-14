#!/bin/sh
set -eu

wiki_dir=${1:-wiki}

find "$wiki_dir" -maxdepth 2 -type f | LC_ALL=C sort