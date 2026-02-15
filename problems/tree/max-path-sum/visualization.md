# Visualization — Max Path Sum

## Input Tree

```
        -10
        / \
       9   20
           / \
          15   7
```

## Recursion Tree with Return Values

```mermaid
flowchart TD
    A["Node -10\n───────────\nL=9, R=42\nglobal = max(_, 9+42+(-10)) = 41\nReturn: max(0, max(9,42)+(-10)) = 32"]
    B["Node 9\n───────────\nLeaf\nglobal = max(_, 9) = 9\nReturn: max(0, 9) = 9"]
    C["Node 20\n───────────\nL=15, R=7\nglobal = max(_, 15+7+20) = 42 ✅\nReturn: max(0, max(15,7)+20) = 35"]
    D["Node 15\n───────────\nLeaf\nReturn: max(0, 15) = 15"]
    E["Node 7\n───────────\nLeaf\nReturn: max(0, 7) = 7"]

    A -->|left| B
    A -->|right| C
    C -->|left| D
    C -->|right| E
```

## Call Stack (Sequence Diagram)

```mermaid
sequenceDiagram
    participant R as dfs(-10)
    participant N9 as dfs(9)
    participant N20 as dfs(20)
    participant N15 as dfs(15)
    participant N7 as dfs(7)

    R->>N9: recurse left
    N9-->>R: return 9 (leaf)
    R->>N20: recurse right
    N20->>N15: recurse left
    N15-->>N20: return 15 (leaf)
    N20->>N7: recurse right
    N7-->>N20: return 7 (leaf)
    Note over N20: global = 15+7+20 = 42 ✅
    N20-->>R: return 35
    Note over R: global = max(42, 9+35-10) = 42
    Note over R: Answer: 42
```

## State at Each Node

| Node | Left Return | Right Return | Global Update | Return to Parent |
|------|------------|--------------|---------------|------------------|
| 15 | 0 | 0 | 15 | 15 |
| 7 | 0 | 0 | 7 | 7 |
| 20 | 15 | 7 | **15+7+20=42** | 35 |
| 9 | 0 | 0 | 9 | 9 |
| -10 | 9 | 35 | 9+35-10=34 | 32 |

**Answer**: 42 (path: 15 → 20 → 7)

> Node 20 is where the magic happens — it forks through both children. Its return to parent (35) drops one child, but the global already captured the full fork.
