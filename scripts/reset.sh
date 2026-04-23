#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"

# Determine round number: argument > .aigocoach.yaml > default 1
if [ $# -ge 1 ]; then
    ROUND="$1"
else
    ROUND=$(grep 'current_round:' "$REPO_ROOT/.aigocoach.yaml" 2>/dev/null | awk '{print $2}' || echo "1")
fi

ROUND_DIR="$REPO_ROOT/my-progress/round-${ROUND}"

echo "=== AIgoCoach Reset — Round $ROUND ==="
echo "This will reset ALL progress and solutions for round ${ROUND}."
echo ""

if [ ! -d "$ROUND_DIR" ]; then
    echo "No my-progress/round-${ROUND}/ directory found. Nothing to reset."
    exit 1
fi

read -p "Are you sure? (y/N) " -n 1 -r
echo ""

if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "Aborted."
    exit 0
fi

echo "Resetting round ${ROUND}..."
rm -rf "$ROUND_DIR"

# Re-initialize from templates
"$REPO_ROOT/scripts/init.sh" "$ROUND"
