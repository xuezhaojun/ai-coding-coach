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

# Determine language: argument --lang > config file > default "go"
LANG="go"
for arg in "$@"; do
    case $arg in
        --lang=go) LANG="go" ;;
        --lang=python) LANG="python" ;;
        --lang=typescript) LANG="typescript" ;;
        --lang=all) LANG="all" ;;
    esac
done

ROUND_DIR="$REPO_ROOT/my-progress/round-${ROUND}"

echo "=== AI Coding Coach — Round $ROUND (lang: $LANG) ==="

if [ -d "$ROUND_DIR" ]; then
    echo "my-progress/round-${ROUND}/ already exists."
    read -p "Overwrite and start fresh? (y/N) " -n 1 -r
    echo ""
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "Aborted."
        exit 0
    fi
    rm -rf "$ROUND_DIR"
fi

echo "Initializing my-progress/round-${ROUND}/ from templates..."
mkdir -p "$ROUND_DIR"

# Copy checklist (per-round)
cp "$TEMPLATES_DIR/checklist.md" "$ROUND_DIR/checklist.md"

# Copy progress.md to my-progress/ only if it doesn't exist yet (shared across rounds)
PROGRESS_FILE="$REPO_ROOT/my-progress/progress.md"
if [ ! -f "$PROGRESS_FILE" ]; then
    cp "$TEMPLATES_DIR/progress.md" "$PROGRESS_FILE"
    echo "Created my-progress/progress.md (shared across all rounds)"
fi

# Copy problem files for the selected language(s)
echo "Copying problem files (language: $LANG)..."
cd "$TEMPLATES_DIR/problems"

# Create all problem directories
find . -type d | while read -r dir; do
    mkdir -p "$ROUND_DIR/problems/$dir"
done

# Always copy READMEs and assets
find . \( -name "README.md" -o -name "README_zh.md" \) | while read -r file; do
    cp "$file" "$ROUND_DIR/problems/$file"
done

find . -path "*/assets/*" -type f 2>/dev/null | while read -r file; do
    cp "$file" "$ROUND_DIR/problems/$file"
done

# Copy language-specific files
copy_lang_files() {
    local lang="$1"
    local ext="$2"
    local test_pattern="$3"
    local solution_pattern="$4"

    find . -path "*/$lang/*" -type f 2>/dev/null | while read -r file; do
        cp "$file" "$ROUND_DIR/problems/$file"
    done
}

case "$LANG" in
    go)
        copy_lang_files "go" "go" "*_test.go" "solution.go"
        ;;
    python)
        copy_lang_files "python" "py" "test_*.py" "solution.py"
        ;;
    typescript)
        copy_lang_files "typescript" "ts" "*.test.ts" "solution.ts"
        ;;
    all)
        copy_lang_files "go" "go" "*_test.go" "solution.go"
        copy_lang_files "python" "py" "test_*.py" "solution.py"
        copy_lang_files "typescript" "ts" "*.test.ts" "solution.ts"
        ;;
esac

# Update current_round in config file
if [ -f "$CONFIG_FILE" ]; then
    sed -i '' "s/^current_round:.*/current_round: ${ROUND}/" "$CONFIG_FILE"
fi

echo ""
echo "Done! Your workspace is ready at my-progress/round-${ROUND}/"
echo ""
echo "  my-progress/round-${ROUND}/"
echo "  ├── checklist.md   — your problem checklist (per round)"
echo "  └── problems/      — write your solutions here"
echo ""
echo "  my-progress/progress.md — your progress tracker (shared across all rounds)"
echo ""
echo "Start a new round:   ./scripts/init.sh <round-number> [--lang=go|python|typescript|all]"
echo "Reference solutions: templates/problems/<category>/<problem>/<lang>/solution.*"
echo "Problem info:        templates/problems/<category>/<problem>/README.md"
echo ""
echo "This directory is gitignored — it's your personal workspace."
