# AIgoCoach 🤖

[English](README.md) | 中文

> 你的 AI 算法教练 — 150 道经典算法题，用 Go 实现，AI 全程引导。

## 这是什么？

一个专为 AI 编程助手（Claude Code、Cursor、Windsurf 等）设计的面试刷题仓库。克隆下来，打开项目，就能开始和 AI 教练一起练习算法题：

- 引导你完成 150 道题目（基于 NeetCode 150），不会直接给答案
- 在 `my-progress/checklist.md` 中追踪你的刷题进度
- 在 `my-progress/progress.md` 中记录心得和错误，帮你发现薄弱环节
- 根据你的错误模式生成变体练习题
- 每道题都附带最优参考答案
- 支持中文和英文交互

## 快速开始

```bash
git clone https://github.com/xuezhaojun/aigocoach.git
cd aigocoach

# 初始化你的个人工作空间
./scripts/init.sh

# （可选）设置交互语言
# 编辑 .aigocoach.yaml → language: "zh" 或 "en"

# 用 VS Code 配合 Claude Code / Cursor / Windsurf 打开
code .

# 开始和 AI 教练对话：
# "从一道简单题开始吧"
# "我想练习二分查找"
# "帮我看看我的薄弱项"
```

## 使用流程

### 1. 选择题目

打开 `my-progress/checklist.md`，查看按分类和难度排列的 150 道题目（编号格式 `01.01`、`01.02`……）。每道题都链接到你的工作文件。

### 2. 编写解法

编辑 `my-progress/problems/` 下对应的 `.go` 文件 — 函数签名已经定义好，只需填写函数体。

### 3. 运行测试

```bash
# 测试单道题目
go test ./my-progress/problems/01_arrays_hashing/03_two_sum/... -v

# 测试整个分类
go test ./my-progress/problems/07_trees/... -v
```

### 4. 追踪进度

AI 教练会自动更新 `my-progress/checklist.md`（打勾）和 `my-progress/progress.md`（心得、错误记录、知识点掌握度）。

### 5. 查看参考答案

每道题的参考答案在 `templates/problems/<分类>/<题目>/solution.go`。卡住的时候可以让 AI 教练提示，或者直接查看。

### 6. 重置进度

```bash
# 清除所有进度和解法，恢复初始状态
./scripts/reset.sh
```

## 项目结构

```
aigocoach/
├── CLAUDE.md                  # AI 教练行为指令
├── AGENTS.md                  # CLAUDE.md 的符号链接（兼容其他 AI 工具）
├── .aigocoach.yaml        # 用户配置（语言等）
├── templates/                 # 只读模板（不要修改）
│   ├── checklist.md           # 题目清单模板
│   ├── progress.md            # 进度追踪模板
│   └── problems/              # 题目桩、测试、参考答案、元数据
│       ├── 01_arrays_hashing/
│       │   ├── 01_contains_duplicate/
│       │   │   ├── contains_duplicate.go      # 函数桩
│       │   │   ├── contains_duplicate_test.go  # 测试
│       │   │   ├── solution.go                 # 最优参考解法
│       │   │   └── README.md                   # 难度、知识点、思路
│       │   └── ...
│       ├── 02_two_pointers/       # 双指针 — 5 题
│       ├── 03_sliding_window/     # 滑动窗口 — 6 题
│       ├── 04_stack/              # 栈 — 7 题
│       ├── 05_binary_search/      # 二分查找 — 7 题
│       ├── 06_linked_list/        # 链表 — 11 题
│       ├── 07_trees/              # 树 — 15 题
│       ├── 08_tries/              # 字典树 — 3 题
│       ├── 09_heap/               # 堆 — 7 题
│       ├── 10_backtracking/       # 回溯 — 9 题
│       ├── 11_graphs/             # 图 — 13 题
│       ├── 12_advanced_graphs/    # 高级图算法 — 6 题
│       ├── 13_dp_1d/              # 一维动态规划 — 12 题
│       ├── 14_dp_2d/              # 二维动态规划 — 11 题
│       ├── 15_greedy/             # 贪心 — 8 题
│       ├── 16_intervals/          # 区间 — 6 题
│       ├── 17_math_geometry/      # 数学与几何 — 8 题
│       └── 18_bit_manipulation/   # 位运算 — 7 题
├── my-progress/               # 你的个人工作空间（不纳入 git）
│   ├── checklist.md           # 你的题目清单和进度
│   ├── progress.md            # 你的进度追踪（心得、错误、统计）
│   └── problems/              # 你的工作副本（在这里写解法）
├── tmp/                       # 变体练习题（不纳入 git）
└── scripts/
    ├── init.sh                # 从模板初始化 my-progress/
    └── reset.sh               # 重置脚本
```

## 配置

编辑 `.aigocoach.yaml`：

```yaml
# AI 交互语言："en"（英文）或 "zh"（中文）
language: en
```

## 许可证

MIT
