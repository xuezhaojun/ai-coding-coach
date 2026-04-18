#!/usr/bin/env bash
set -euo pipefail

REPO_ROOT="$(cd "$(dirname "$0")/.." && pwd)"
PROGRESS_DIR="$REPO_ROOT/my-progress"
TEMPLATES_DIR="$REPO_ROOT/templates"

if [ -d "$PROGRESS_DIR" ]; then
    echo "my-progress/ already exists."
    read -p "Overwrite and start fresh? (y/N) " -n 1 -r
    echo ""
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        echo "Aborted."
        exit 0
    fi
    rm -rf "$PROGRESS_DIR"
fi

echo "Initializing my-progress/ from templates..."
mkdir -p "$PROGRESS_DIR"

# Copy tracking files
cp "$TEMPLATES_DIR/checklist.md" "$PROGRESS_DIR/checklist.md"
cp "$TEMPLATES_DIR/progress.md" "$PROGRESS_DIR/progress.md"

# Copy problem stubs, tests, solutions, and READMEs
echo "Copying problem files..."
cd "$TEMPLATES_DIR/problems"
find . -type d | while read -r dir; do
    mkdir -p "$PROGRESS_DIR/problems/$dir"
done

find . -name "*.go" -o -name "README.md" -o -name "README_zh.md" \
    | while read -r file; do
    cp "$file" "$PROGRESS_DIR/problems/$file"
done

echo ""
echo "Done! Your workspace is ready at my-progress/"
echo ""
echo "  my-progress/"
echo "  ├── checklist.md   — your problem checklist"
echo "  ├── progress.md   — your progress tracker (insights, mistakes, stats)"
echo "  └── problems/      — write your solutions here"
echo ""
echo "Reference solutions: templates/problems/<category>/<problem>/solution.go"
echo "Problem info:        templates/problems/<category>/<problem>/README.md"
echo ""
echo "This directory is gitignored — it's your personal workspace."
