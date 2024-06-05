#!/usr/bin/env bash

set -euo pipefail

echo "🏁 Running unit tests..."

exclude_packages=(
    "internal/mocks"
)

cmd="go list ./..."

for exclude_package in "${exclude_packages[@]}"; do
    cmd+=" | grep -v ${exclude_package}"
done

cmd+=" | xargs go test -covermode=atomic -vet=all"

eval "$cmd"

echo "✅ All unit tests passed."
exit 0