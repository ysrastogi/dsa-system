# AVL Tree Rotations — Visual Guide

## Why Rotations?

After an insertion or deletion, the tree may become **unbalanced** (height difference > 1). Rotations restore balance in O(1) while preserving BST ordering.

## Right Rotation (Left-Heavy / LL Case)

```
Before:         After:
    z               y
   / \             / \
  y   T4          x   z
 / \             /   / \
x   T3          T1  T3  T4
/
T1
```

```mermaid
flowchart LR
    subgraph Before
        Z1[z] --> Y1[y]
        Z1 --> T4a[T4]
        Y1 --> X1[x]
        Y1 --> T3a[T3]
    end
    subgraph After
        Y2[y] --> X2[x]
        Y2 --> Z2[z]
        Z2 --> T3b[T3]
        Z2 --> T4b[T4]
    end

    Before -->|"Right Rotate(z)"| After
```

## Left Rotation (Right-Heavy / RR Case)

```
Before:         After:
  z                y
 / \              / \
T1   y           z   x
    / \         / \   \
   T2  x      T1  T2  T3
        \
        T3
```

```mermaid
flowchart LR
    subgraph Before
        Z3[z] --> T1a[T1]
        Z3 --> Y3[y]
        Y3 --> T2a[T2]
        Y3 --> X3[x]
    end
    subgraph After
        Y4[y] --> Z4[z]
        Y4 --> X4[x]
        Z4 --> T1b[T1]
        Z4 --> T2b[T2]
    end

    Before -->|"Left Rotate(z)"| After
```

## Left-Right (LR Case) — Double Rotation

```mermaid
flowchart LR
    subgraph "Step 1: Before"
        Za[z] --> Ya[y]
        Ya --> Xa[x]
    end
    subgraph "Step 2: Left Rotate y"
        Zb[z] --> Xb[x]
        Xb --> Yb[y]
    end
    subgraph "Step 3: Right Rotate z"
        Xc[x] --> Yc[y]
        Xc --> Zc[z]
    end

    Za -->|"Left Rotate(y)"| Zb
    Zb -->|"Right Rotate(z)"| Xc
```

## Right-Left (RL Case) — Double Rotation

```mermaid
flowchart LR
    subgraph "Step 1: Before"
        Zd[z] --> Yd[y]
        Yd --> Xd[x]
    end
    subgraph "Step 2: Right Rotate y"
        Ze[z] --> Xe[x]
        Xe --> Ye[y]
    end
    subgraph "Step 3: Left Rotate z"
        Xf[x] --> Zf[z]
        Xf --> Yf[y]
    end

    Zd -->|"Right Rotate(y)"| Ze
    Ze -->|"Left Rotate(z)"| Xf
```

## Quick Reference

| Case | Balance Factor | Fix |
|------|---------------|-----|
| LL | bf(z)=2, bf(y)=1 | Right rotate z |
| RR | bf(z)=-2, bf(y)=-1 | Left rotate z |
| LR | bf(z)=2, bf(y)=-1 | Left rotate y, then right rotate z |
| RL | bf(z)=-2, bf(y)=1 | Right rotate y, then left rotate z |
