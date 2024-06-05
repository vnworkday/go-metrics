#!/usr/bin/env bash

set -euo pipefail

echo "🏁 Running unit tests..."

go test ./... -covermode=atomic -vet=all -coverprofile=profile.cov

echo "✅ All unit tests passed."
exit 0