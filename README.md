# GoAlgo Espresso ☕

English | [中文](README_zh.md)

> A shot of espresso for your algorithm brain — 150 classic problems in Go, guided by an AI coach.

## What Is This?

An interview prep repo designed to work with AI coding assistants (Claude Code, Cursor, Windsurf, etc.). Clone it, open it, and start practicing algorithm problems with an AI coach that:

- Guides you through 150 problems (based on NeetCode 150) without giving away answers
- Tracks your progress in `my-progress/plan.md`
- Records your mistakes in `my-progress/mistakes.md` and identifies weak areas
- Generates variant problems based on your error patterns
- Provides optimal reference solutions for every problem
- Supports both English and Chinese interaction

## Quick Start

```bash
git clone https://github.com/goalgo-espresso/go-algo-espresso.git
cd go-algo-espresso

# Initialize your personal workspace
./scripts/init.sh

# (Optional) Set your preferred language
# Edit .algo-espresso.yaml → language: "zh" or "en"

# Open in VS Code with Claude Code / Cursor / Windsurf
code .

# Start talking to the agent:
# "Let's start with an easy problem"
# "I want to practice binary search"
# "Show me my weak areas"
```

## How It Works

### 1. Pick a Problem

Open `my-progress/plan.md` to see all 150 problems organized by category and difficulty (numbered `01.01`, `01.02`, etc.). Each links to your working Go file.

### 2. Write Your Solution

Edit the `.go` file in `my-progress/problems/` — the function signature is already there, just fill in the body.

### 3. Run Tests

```bash
# Test a specific problem
go test ./my-progress/problems/01_arrays_hashing/03_two_sum/... -v

# Test all problems in a category
go test ./my-progress/problems/07_trees/... -v
```

### 4. Track Progress

The agent automatically updates `my-progress/plan.md` (checkboxes), `my-progress/mistakes.md` (error log), and `my-progress/progress.md` (knowledge points) as you work.

### 5. Check Reference Solutions

Each problem has a reference solution in `templates/problems/<category>/<problem>/solution.go`. Ask the agent or read it directly when you're stuck.

### 6. Reset & Start Fresh

```bash
# Reset all progress and solutions to blank state
./scripts/reset.sh
```

## Project Structure

```
go-algo-espresso/
├── CLAUDE.md                  # Agent behavior instructions
├── AGENTS.md                  # Symlink → CLAUDE.md (for other AI tools)
├── .algo-espresso.yaml        # User configuration (language, etc.)
├── templates/                 # READ-ONLY (do not modify)
│   ├── plan.md                # Problem checklist template
│   ├── mistakes.md            # Mistake log template
│   ├── progress.md            # Progress tracker template
│   └── problems/              # Problem stubs, tests, solutions, metadata
│       ├── 01_arrays_hashing/
│       │   ├── 01_contains_duplicate/
│       │   │   ├── contains_duplicate.go      # Function stub
│       │   │   ├── contains_duplicate_test.go  # Tests
│       │   │   ├── solution.go                 # Optimal reference solution
│       │   │   └── README.md                   # Difficulty, topics, approach
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
│   ├── plan.md                # Your problem checklist with progress
│   ├── mistakes.md            # Your mistake log
│   ├── progress.md            # Your knowledge point tracker
│   └── problems/              # Your working copies (write solutions here)
├── tmp/                       # Variant problems (gitignored)
└── scripts/
    ├── init.sh                # Initialize my-progress/ from templates
    └── reset.sh               # Reset to fresh state
```

## Configuration

Edit `.algo-espresso.yaml`:

```yaml
# Language for agent interaction: "en" or "zh"
language: en
```

## License

MIT
