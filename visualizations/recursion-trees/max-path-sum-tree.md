# Max Path Sum — Recursion Tree Visualization

This visualization traces `maxPathSum` on the tree `[1, 2, 3, 4, 5]`.

## The Tree

```
        1
       / \
      2   3
     / \
    4   5
```

## Recursion Tree with Return Values

```mermaid
flowchart TD
    A["maxPathSum(1)\n───────────────\nL=7, R=3\nglobal = max(global, 7+3+1) = 11\nReturn: max(0, max(7,3)+1) = 8"]
    B["maxPathSum(2)\n───────────────\nL=4, R=5\nglobal = max(global, 4+5+2) = 11 ✅\nReturn: max(0, max(4,5)+2) = 7"]
    C["maxPathSum(3)\n───────────────\nL=0, R=0\nglobal = max(global, 0+0+3) = 3\nReturn: max(0, 3) = 3"]
    D["maxPathSum(4)\n───────────────\nLeaf\nReturn: max(0, 4) = 4"]
    E["maxPathSum(5)\n───────────────\nLeaf\nReturn: max(0, 5) = 5"]

    A -->|left| B
    A -->|right| C
    B -->|left| D
    B -->|right| E
```

## Call Stack Trace

```mermaid
sequenceDiagram
    participant N1 as Node 1
    participant N2 as Node 2
    participant N4 as Node 4
    participant N5 as Node 5
    participant N3 as Node 3

    N1->>N2: recurse left
    N2->>N4: recurse left
    N4-->>N2: return 4
    N2->>N5: recurse right
    N5-->>N2: return 5
    Note over N2: global = max(_, 4+5+2) = 11
    N2-->>N1: return 7
    N1->>N3: recurse right
    N3-->>N1: return 3
    Note over N1: global = max(11, 7+3+1) = 11
    Note over N1: Return 8 (not used by anyone)
```

## Level-by-Level Analysis

| Level | Node | Left Return | Right Return | Global Update | Return to Parent |
|-------|------|-------------|--------------|---------------|------------------|
| 2 | 4 | — | — | 4 | 4 |
| 2 | 5 | — | — | 5 | 5 |
| 1 | 2 | 4 | 5 | **4+5+2=11** | 7 |
| 1 | 3 | 0 | 0 | 3 | 3 |
| 0 | 1 | 7 | 3 | 7+3+1=11 | 8 |

**Answer**: Global max = **11** (path 4→2→5)

> The return value at root (8) is never used — it's only meaningful if this tree were a subtree of something larger.
