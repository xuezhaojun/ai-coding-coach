#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
PROGRESS_DIR="$REPO_ROOT/my-progress"

echo "=== AIgoCoach Reset ==="
echo "This will reset ALL progress, solutions, and tracking files."
echo ""

if [ ! -d "$PROGRESS_DIR" ]; then
    echo "No my-progress/ directory found. Run ./scripts/init.sh first."
    exit 1
fi

read -p "Are you sure? (y/N) " -n 1 -r
echo ""

if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Aborted."
    exit 0
fi

echo "Resetting workspace..."
rm -rf "$PROGRESS_DIR"

# Re-initialize from templates
"$REPO_ROOT/scripts/init.sh"
