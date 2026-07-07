#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
TEMPLATES_DIR="$REPO_ROOT/templates"
CONFIG_FILE="$REPO_ROOT/.aicodingcoach.yaml"

# Determine round number: argument > config file > default 1
if [ $# -ge 1 ]; then
    ROUND="$1"
else
    ROUND=$(grep 'current_round:' "$CONFIG_FILE" 2>/dev/null | awk '{print $2}' || echo "1")
fi

ROUND_DIR="$REPO_ROOT/my-progress/round-${ROUND}"

if [ ! -d "$ROUND_DIR" ]; then
    echo "my-progress/round-${ROUND}/ does not exist. Run ./scripts/init.sh ${ROUND} first."
    exit 1
fi

echo "Syncing templates → my-progress/round-${ROUND}/ (incremental)..."

added=0
updated=0
skipped=0

cd "$TEMPLATES_DIR/problems"

# Ensure all directories exist
find . -type d | while read -r dir; do
    mkdir -p "$ROUND_DIR/problems/$dir"
done

# Sync assets (always overwrite — these are reference images, not user work)
find . -path "*/assets/*" -type f | while read -r file; do
    dest="$ROUND_DIR/problems/$file"
    if [ ! -f "$dest" ] || ! cmp -s "$file" "$dest"; then
        cp "$file" "$dest"
        echo "  + $file"
    fi
done

# Sync READMEs (always overwrite — these are problem descriptions, not user work)
find . \( -name "README.md" -o -name "README_zh.md" \) | while read -r file; do
    dest="$ROUND_DIR/problems/$file"
    if [ ! -f "$dest" ] || ! cmp -s "$file" "$dest"; then
        cp "$file" "$dest"
        echo "  ~ $file"
    fi
done

# Sync test files for all languages (always overwrite — tests are source of truth)
# Go tests
find . -path "*/go/*_test.go" -type f 2>/dev/null | while read -r file; do
    dest="$ROUND_DIR/problems/$file"
    if [ ! -f "$dest" ] || ! cmp -s "$file" "$dest"; then
        cp "$file" "$dest"
        echo "  ~ $file"
    fi
done

# Python tests
find . -path "*/python/test_*.py" -type f 2>/dev/null | while read -r file; do
    dest="$ROUND_DIR/problems/$file"
    if [ ! -f "$dest" ] || ! cmp -s "$file" "$dest"; then
        cp "$file" "$dest"
        echo "  ~ $file"
    fi
done

# TypeScript tests
find . -path "*/typescript/*.test.ts" -type f 2>/dev/null | while read -r file; do
    dest="$ROUND_DIR/problems/$file"
    if [ ! -f "$dest" ] || ! cmp -s "$file" "$dest"; then
        cp "$file" "$dest"
        echo "  ~ $file"
    fi
done

# Sync solution stubs and solutions — only if missing (never overwrite user work)
# Go
find . -path "*/go/*.go" ! -name "*_test.go" -type f 2>/dev/null | while read -r file; do
    dest="$ROUND_DIR/problems/$file"
    if [ ! -f "$dest" ]; then
        cp "$file" "$dest"
        echo "  + $file"
    fi
done

# Python
find . -path "*/python/*.py" ! -name "test_*.py" -type f 2>/dev/null | while read -r file; do
    dest="$ROUND_DIR/problems/$file"
    if [ ! -f "$dest" ]; then
        cp "$file" "$dest"
        echo "  + $file"
    fi
done

# TypeScript
find . -path "*/typescript/*.ts" ! -name "*.test.ts" -type f 2>/dev/null | while read -r file; do
    dest="$ROUND_DIR/problems/$file"
    if [ ! -f "$dest" ]; then
        cp "$file" "$dest"
        echo "  + $file"
    fi
done

# Sync tracking files only if missing
for f in checklist.md progress.md; do
    if [ ! -f "$ROUND_DIR/$f" ]; then
        cp "$TEMPLATES_DIR/$f" "$ROUND_DIR/$f"
        echo "  + $f"
    fi
done

echo ""
echo "Sync complete. Your solutions in my-progress/round-${ROUND}/ were not touched."
