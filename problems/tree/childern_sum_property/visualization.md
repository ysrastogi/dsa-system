# Visualization — Children Sum Property

## Input Tree

```
        2
       / \
     35   10
    / \   / \
   2   3 5   2
```

## Recursion Trace — Push Down, Pull Up

```mermaid
flowchart TD
    A["Node 2 (root)
━━━━━━━━━
childSum=45 ≥ 2
Set val=45
Recurse...
Pull up: 55+14=69"]:::modified
    B["Node 35
━━━━━━━━━
childSum=5 < 35
Push 35→children
Recurse...
Pull up: 35+20=55"]:::modified
    C["Node 10
━━━━━━━━━
childSum=7 < 10
Push 10→children
Recurse...
Pull up: 9+5=14"]:::modified
    D["Node 2→35 (leaf)
Val stays 35"]:::leaf
    E["Node 3→35 (leaf)
━━━━━━━━━
childSum=5<35, push
Pull up: stays 20"]:::leaf
    F["Node 5→10 (leaf)
━━━━━━━━━
childSum=2<10, push
Pull up: stays 9"]:::leaf
    G["Node 2→10 (leaf)
Val stays 5"]:::leaf

    A -->|left| B
    A -->|right| C
    B -->|left| D
    B -->|right| E
    C -->|left| F
    C -->|right| G

    classDef modified fill:#f59e0b,stroke:#b45309,color:#000
    classDef leaf fill:#10b981,stroke:#065f46,color:#fff
```

## Step-by-Step State Table

### Phase 1: Push Down (Preorder)

| Step | Node | childSum | Compare | Action |
|------|------|----------|---------|--------|
| 1 | **2** (root) | 35+10=45 | 45 ≥ 2 | `root.Val = 45` |
| 2 | **35** | 2+3=5 | 5 < 35 | Push: `left.Val=35, right.Val=35` |
| 3 | **2→35** (leaf) | 0 | leaf | no-op |
| 4 | **3→35** (leaf) | 0 | leaf | no-op |
| 5 | **10** | 5+2=7 | 7 < 10 | Push: `left.Val=10, right.Val=10` |
| 6 | **5→10** (leaf) | 0 | leaf | no-op |
| 7 | **2→10** (leaf) | 0 | leaf | no-op |

### Phase 2: Pull Up (Postorder)

| Step | Node | left.Val | right.Val | Pull Up |
|------|------|----------|-----------|---------|
| 8 | **35** | 35 | 35 | was pushed, but leaves → `35 + 20 = 55`* |
| 9 | **10** | 10 | 10 | `9 + 5 = 14`* |
| 10 | **root** | 55 | 14 | `55 + 14 = 69` |

> *Leaf values stabilize after push-down. Internal nodes get their final values from pull-up.

## Before vs After

```mermaid
flowchart LR
    subgraph BEFORE["Before"]
        A1["2"]
        B1["35"]
        C1["10"]
        D1["2"]
        E1["3"]
        F1["5"]
        G1["2"]
        A1 --> B1
        A1 --> C1
        B1 --> D1
        B1 --> E1
        C1 --> F1
        C1 --> G1
    end

    subgraph AFTER["After ✅"]
        A2["69"]
        B2["55"]
        C2["14"]
        D2["35"]
        E2["20"]
        F2["9"]
        G2["5"]
        A2 --> B2
        A2 --> C2
        B2 --> D2
        B2 --> E2
        C2 --> F2
        C2 --> G2
    end

    BEFORE -.->|"transform"| AFTER

    style A1 fill:#374151,stroke:#6b7280,color:#9ca3af
    style B1 fill:#374151,stroke:#6b7280,color:#9ca3af
    style C1 fill:#374151,stroke:#6b7280,color:#9ca3af
    style D1 fill:#374151,stroke:#6b7280,color:#9ca3af
    style E1 fill:#374151,stroke:#6b7280,color:#9ca3af
    style F1 fill:#374151,stroke:#6b7280,color:#9ca3af
    style G1 fill:#374151,stroke:#6b7280,color:#9ca3af
    style A2 fill:#10b981,stroke:#065f46,color:#fff
    style B2 fill:#10b981,stroke:#065f46,color:#fff
    style C2 fill:#10b981,stroke:#065f46,color:#fff
    style D2 fill:#10b981,stroke:#065f46,color:#fff
    style E2 fill:#10b981,stroke:#065f46,color:#fff
    style F2 fill:#10b981,stroke:#065f46,color:#fff
    style G2 fill:#10b981,stroke:#065f46,color:#fff
```

## Call Stack (Sequence View)

```mermaid
sequenceDiagram
    participant R as modify(2)
    participant L as modify(35)
    participant LL as modify(2→35)
    participant LR as modify(3→35)
    participant Ri as modify(10)
    participant RL as modify(5→10)
    participant RR as modify(2→10)

    Note over R: childSum=45≥2, set val=45
    R->>L: recurse left
    Note over L: childSum=5<35, push 35→children
    L->>LL: recurse left
    Note over LL: leaf, no-op
    LL-->>L: return
    L->>LR: recurse right
    Note over LR: leaf, no-op
    LR-->>L: return
    Note over L: Pull up: 35+20=55 ✅
    L-->>R: return
    R->>Ri: recurse right
    Note over Ri: childSum=7<10, push 10→children
    Ri->>RL: recurse left
    Note over RL: leaf, no-op
    RL-->>Ri: return
    Ri->>RR: recurse right
    Note over RR: leaf, no-op
    RR-->>Ri: return
    Note over Ri: Pull up: 9+5=14 ✅
    Ri-->>R: return
    Note over R: Pull up: 55+14=69 ✅
```

## Key Insight

> The push-down doesn't need to be "correct" — it just needs to ensure **enough budget** exists in the subtree. The pull-up phase will fix the parent to the exact sum of children. This is why we can naively overwrite both children with the parent's value.
