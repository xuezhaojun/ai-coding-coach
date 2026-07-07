# AI Coding Coach 🤖

English | [中文](README_zh.md)

> Your AI algorithm coach — 150 classic problems in Go, Python, and TypeScript, guided step by step.

## What Is This?

An interview prep repo designed to work with AI coding assistants (Claude Code, Cursor, Windsurf, etc.). Clone it, open it, and start practicing algorithm problems with an AI coach that:

- Guides you through 150 problems (based on NeetCode 150) without giving away answers
- Supports Go, Python, and TypeScript — practice in the language of your choice
- Tracks your progress in `my-progress/checklist.md`
- Records insights and mistakes in `my-progress/progress.md` and identifies weak areas
- Generates variant problems based on your error patterns
- Provides optimal reference solutions for every problem
- Supports both English and Chinese interaction

## Quick Start

```bash
git clone https://github.com/xuezhaojun/ai-coding-coach.git
cd ai-coding-coach

# Initialize your personal workspace
# Choose your practice language: go, python, typescript, or all
./scripts/init.sh --lang=go

# (Optional) Set your preferred interaction language
# Edit .aicodingcoach.yaml → interaction_language: "zh" or "en"

# Open in VS Code with Claude Code / Cursor / Windsurf
code .

# Start talking to the agent:
# "Let's start with an easy problem"
# "I want to practice binary search"
# "Show me my weak areas"
```

## How It Works

### 1. Pick a Problem

Open `my-progress/checklist.md` to see all 150 problems organized by category and difficulty (numbered `01.01`, `01.02`, etc.). Each links to your working file in your chosen language.

### 2. Write Your Solution

Edit the file in `my-progress/problems/<category>/<problem>/<lang>/` — the function signature is already there, just fill in the body.

### 3. Run Tests

```bash
# Go
go test ./my-progress/problems/01_arrays_hashing/03_two_sum/go/... -v

# Python
pytest my-progress/problems/01_arrays_hashing/03_two_sum/python/ -v

# TypeScript
npx vitest run my-progress/problems/01_arrays_hashing/03_two_sum/typescript/
```

### 4. Track Progress

The agent automatically updates `my-progress/checklist.md` (checkboxes) and `my-progress/progress.md` (session log with insights, mistakes, and knowledge points) as you work. The progress tracker is shared across all languages.

### 5. Check Reference Solutions

Each problem has a reference solution in `templates/problems/<category>/<problem>/<lang>/solution.*`. Ask the agent or read it directly when you're stuck.

### 6. Reset & Start Fresh

```bash
# Reset all progress and solutions to blank state
./scripts/reset.sh
```

## Project Structure

```
ai-coding-coach/
├── CLAUDE.md                  # Agent behavior instructions
├── AGENTS.md                  # Symlink → CLAUDE.md (for other AI tools)
├── .aicodingcoach.yaml        # User configuration (gitignored, created by init.sh)
├── templates/                 # READ-ONLY (do not modify)
│   ├── checklist.md           # Problem checklist template
│   ├── progress.md            # Progress tracker template
│   └── problems/              # Problem stubs, tests, solutions, metadata
│       ├── 01_arrays_hashing/
│       │   ├── 01_contains_duplicate/
│       │   │   ├── README.md                   # Difficulty, topics, approach
│       │   │   ├── go/
│       │   │   │   ├── contains_duplicate.go   # Go function stub
│       │   │   │   ├── contains_duplicate_test.go  # Go tests
│       │   │   │   └── solution.go             # Optimal reference solution
│       │   │   ├── python/
│       │   │   │   ├── contains_duplicate.py   # Python function stub
│       │   │   │   ├── test_contains_duplicate.py  # Python tests (pytest)
│       │   │   │   └── solution.py             # Optimal reference solution
│       │   │   └── typescript/
│       │   │       ├── contains_duplicate.ts   # TypeScript function stub
│       │   │       ├── contains_duplicate.test.ts  # TypeScript tests (vitest)
│       │   │       └── solution.ts             # Optimal reference solution
│       │   └── ...
│       ├── 02_two_pointers/       # 5 problems
│       ├── 03_sliding_window/     # 6 problems
│       ├── 04_stack/              # 7 problems
│       ├── 05_binary_search/      # 7 problems
│       ├── 06_linked_list/        # 11 problems
│       ├── 07_trees/              # 15 problems
│       ├── 08_tries/              # 3 problems
│       ├── 09_heap/               # 7 problems
│       ├── 10_backtracking/       # 9 problems
│       ├── 11_graphs/             # 13 problems
│       ├── 12_advanced_graphs/    # 6 problems
│       ├── 13_dp_1d/              # 12 problems
│       ├── 14_dp_2d/              # 11 problems
│       ├── 15_greedy/             # 8 problems
│       ├── 16_intervals/          # 6 problems
│       ├── 17_math_geometry/      # 8 problems
│       └── 18_bit_manipulation/   # 7 problems
├── my-progress/               # Your personal workspace (gitignored)
│   ├── round-1/               # Round 1
│   │   ├── checklist.md       # Your problem checklist with progress
│   │   ├── progress.md        # Your progress tracker (insights, mistakes, stats)
│   │   └── problems/          # Your working copies (write solutions here)
│   ├── round-2/               # Round 2 (fresh copy)
│   └── ...
├── tmp/                       # Variant problems (gitignored)
└── scripts/
    ├── init.sh                # Initialize my-progress/ from templates (--lang flag)
    └── reset.sh               # Reset to fresh state
```

## Configuration

`.aicodingcoach.yaml` is your personal configuration file (gitignored). It is automatically created when you run `./scripts/init.sh`. You can also create it manually:

```yaml
# Interaction language: "en" or "zh"
interaction_language: en

# Practice language: "go", "python", or "typescript"
practice_language: go

# Current practice round (1, 2, 3, ...)
# Each round creates a fresh copy of all problems under my-progress/round-N/
current_round: 1
```

| Field | Description | Default |
|-------|-------------|---------|
| `interaction_language` | Agent interaction language. `"en"` for English, `"zh"` for Chinese. | `en` |
| `practice_language` | Programming language for practice. `"go"`, `"python"`, or `"typescript"`. | `go` |
| `current_round` | Active practice round number. Determines which `my-progress/round-N/` directory the agent uses. Updated automatically by `./scripts/init.sh <N>`. | `1` |

### Multi-Round Practice

You can practice all 150 problems multiple times. Each round gets a fresh workspace:

```bash
# Start round 1 (default, Go)
./scripts/init.sh --lang=go

# Start round 2 (fresh copy, Python)
./scripts/init.sh 2 --lang=python

# Start round 3 (all languages)
./scripts/init.sh 3 --lang=all

# Start round 3 (TypeScript)
./scripts/init.sh 3 --lang=typescript
```

Each round is stored independently under `my-progress/round-N/` with its own checklist and progress tracker. Previous rounds are preserved so you can compare improvement over time.

## License

MIT
