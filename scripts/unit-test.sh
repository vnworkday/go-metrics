#!/usr/bin/env bash

set -euo pipefail

echo "🏁 Running unit tests..."

go test ./... -covermode=count

echo "✅ All unit tests passed."
exit 0