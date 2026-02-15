# Visualization — Maximum Width of Binary Tree

## Input Tree

```
         1
        / \
       3   2
      /     \
     5       9
    /         \
   6           7
```

## BFS Queue Walk-Through

```mermaid
flowchart TD
    A["Node 1\nidx: 0"]
    B["Node 3\nidx: 0 (norm)"]
    C["Node 2\nidx: 1"]
    D["Node 5\nidx: 0 (norm)"]
    G["Node 9\nidx: 3"]
    H["Node 6\nidx: 0 (norm)"]
    I["Node 7\nidx: 7"]

    A -->|"left: 2×0+1=1 → norm 0"| B
    A -->|"right: 2×0+2=2 → norm 1"| C
    B -->|"left: 2×0+1=1 → norm 0"| D
    C -->|"right: 2×1+2=4 → norm 3"| G
    D -->|"left: 2×0+1=1 → norm 0"| H
    G -->|"right: 2×3+2=8 → norm 7"| I

    style A fill:#2d6a4f,stroke:#1b4332,color:#d8f3dc
    style B fill:#2d6a4f,stroke:#1b4332,color:#d8f3dc
    style C fill:#2d6a4f,stroke:#1b4332,color:#d8f3dc
    style D fill:#e76f51,stroke:#9c4130,color:#fff
    style G fill:#e76f51,stroke:#9c4130,color:#fff
    style H fill:#264653,stroke:#1d3640,color:#a8dadc
    style I fill:#264653,stroke:#1d3640,color:#a8dadc
```

## Level-by-Level State

```mermaid
sequenceDiagram
    participant Q as Queue
    participant W as maxWidth

    Note over Q: Level 0
    Q->>Q: [{1, idx=0}]
    Note over Q: start=0, end=0
    Q->>W: width = 0 - 0 + 1 = 1
    Note over W: maxWidth = 1

    Note over Q: Level 1
    Q->>Q: [{3, idx=0}, {2, idx=1}]
    Note over Q: start=0, end=1
    Q->>W: width = 1 - 0 + 1 = 2
    Note over W: maxWidth = 2

    Note over Q: Level 2
    Q->>Q: [{5, idx=0}, {9, idx=3}]
    Note over Q: start=0, end=3
    Q->>W: width = 3 - 0 + 1 = 4
    Note over W: maxWidth = 4

    Note over Q: Level 3
    Q->>Q: [{6, idx=0}, {7, idx=7}]
    Note over Q: start=0, end=7
    Q->>W: width = 7 - 0 + 1 = 8 ✅
    Note over W: maxWidth = 8
```

## State Table

| Level | Queue Contents | Start Idx | End Idx | Width | maxWidth |
|-------|---------------|-----------|---------|-------|----------|
| 0 | `[{1, 0}]` | 0 | 0 | 1 | 1 |
| 1 | `[{3, 0}, {2, 1}]` | 0 | 1 | 2 | 2 |
| 2 | `[{5, 0}, {9, 3}]` | 0 | 3 | 4 | 4 |
| 3 | `[{6, 0}, {7, 7}]` | 0 | 7 | 8 | **8** |

**Answer**: 8 (the gap between node 6 and node 7 at level 3 spans 8 positions in a complete binary tree)

## Index Normalization Detail

```
Level 2 (before normalization):
  Node 5 raw idx = 2*0+1 = 1   →  minus start(1) → 0
  Node 9 raw idx = 2*1+2 = 4   →  minus start(1) → 3

Level 3 (children use normalized parents):
  Node 6: 2×0+1 = 1  →  minus start(1) → 0
  Node 7: 2×3+2 = 8  →  minus start(1) → 7

Width = 7 - 0 + 1 = 8
```

> **Key insight**: Without normalization, indices grow as 2ⁿ. By subtracting the level's start index each iteration, we keep numbers small while preserving the relative spacing.
