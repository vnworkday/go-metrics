#!/usr/bin/env bash

set -eufo pipefail

echo "🏁 Checking that code complies with static analysis requirements..."

# Feel free to add more directories to skip here, separated by pipes. For example: "vendor|mocks"
skip_dirs="scripts"

packages=$(go list ./... | grep -v -E "$skip_dirs")

echo "🔍 Checking the following packages:"
for package in $packages; do
  echo "  - $package"
  # Note that we globally disable some checks. The list is controlled by the
  # top-level staticcheck.conf file in this repo.
  go run honnef.co/go/tools/cmd/staticcheck "${package}"
  exit_code=$?

  if [[ $exit_code -ne 0 ]]; then
    echo "🚫 Static analysis failed."
    exit 1
  fi
done

echo "✅ All packages pass static analysis."