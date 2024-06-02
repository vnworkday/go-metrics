#!/usr/bin/env bash

set -eufo pipefail

echo "🏁 Checking that code complies with go fmt requirements..."

files=$(go fmt ./...)

if [[ -n "$files" ]]; then
  echo "🚫 The following files are not formatted correctly:"
  for file in $files; do
    echo "  - $file"
  done
  echo "🚫 Please run \"go fmt\" to fix their formatting."
  exit 1
fi

echo "✅ All files are formatted correctly."
exit 0