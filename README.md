# 🧠 DSA System

**A pattern extraction system. A mental simulation trainer. A recursion visualizer. A pre-interview rapid revision engine.**

> If it doesn't help you revise in 10 minutes before an interview → it failed.

---

## 📁 Structure

```
dsa-system/
├── patterns/          → Pattern files (6-section format)
│   ├── recursion/     → Backtracking, divide-and-conquer
│   ├── tree/          → Tree DP, LCA, traversals
│   ├── graph/         → BFS, DFS, topological sort
│   ├── dp/            → Knapsack, LIS, interval DP
│   ├── sliding-window/→ Fixed/variable window
│   └── two-pointers/  → Opposite/same direction
├── visualizations/    → Mermaid recursion trees, BFS/DFS, rotations
├── templates/         → Minimal Go code templates
├── problems/          → Solved problems linked to patterns
└── cheat-sheets/      → Rapid revision checklists
```

---

## 🚀 How to Use

### Before an Interview (10 min)
1. Open `cheat-sheets/` → skim the relevant checklist
2. Open `patterns/<category>/` → re-read the **Core Idea** and **Mermaid diagram**
3. Open `templates/` → glance at the clean Go skeleton

### Deep Practice
1. Pick a problem from `problems/`
2. Read `problem.md` → try solving it
3. Check `solution.go` → compare
4. Read `pattern-link.md` → connect it to the underlying pattern
5. Study `visualization.md` → simulate the recursion mentally

---

## 🧩 Pattern File Format

Every pattern file follows this **6-section** structure:

```
1️⃣ Pattern Name
2️⃣ Core Idea         → 1 paragraph max (what flows up/down/global?)
3️⃣ Template Code     → Minimal, clean Go
4️⃣ When To Use       → How to recognize this in interviews
5️⃣ Why Naive Fails   → Complexity pitfall
6️⃣ Mermaid Diagram   → Recursion tree or sequence diagram
```

---

## 📌 Problem Folder Rule

Every problem **must** contain:

| File | Purpose |
|------|---------|
| `problem.md` | Statement, constraints, examples |
| `solution.go` | Clean, annotated Go solution |
| `pattern-link.md` | Link back to the pattern file |
| `visualization.md` | Mermaid recursion tree with return values |

If you don't link a problem to a pattern, you're wasting time.

---

## ➕ Adding a New Pattern

1. Create `patterns/<category>/<pattern-name>.md`
2. Fill in all 6 sections
3. Add a Mermaid diagram — **labeled returns** are mandatory
4. Add at least one problem that uses this pattern in `problems/`

## ➕ Adding a New Problem

1. Create `problems/<category>/<problem-name>/`
2. Add all 4 files: `problem.md`, `solution.go`, `pattern-link.md`, `visualization.md`
3. Link back to the relevant pattern

---

## 🔧 Tech Stack

- **Language**: Go
- **Diagrams**: Mermaid (renders in GitHub, VS Code, any Markdown previewer)
- **No dependencies**: Pure markdown + Go. No frontend, no animation engine, no Notion sync.

---

## ⚡ Future Upgrade (Optional)

When ready, build a small script that:
- Parses an input tree
- Generates Mermaid automatically
- Prints recursion expansion

Until then: **Markdown + Mermaid + GitHub preview is enough.**
