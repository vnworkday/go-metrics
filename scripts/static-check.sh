#!/usr/bin/env bash

set -eufo pipefail

echo "🏁 Checking that code complies with static analysis requirements..."

skip_dirs="vendor|node_modules|public|storage|bootstrap"

packages=$(go list ./... | grep -v -E "$skip_dirs")

echo "🔍 Checking the following packages:"
for package in $packages; do
  echo "  - $package"
done

# Note that we globally disable some checks. The list is controlled by the
# top-level staticcheck.conf file in this repo.
go run honnef.co/go/tools/cmd/staticcheck "${packages}"
exit_code=$?

if [[ $exit_code -ne 0 ]]; then
  echo "🚫 Static analysis failed."
  exit 1
fi

echo "✅ All packages pass static analysis."