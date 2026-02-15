# Graph Traversal Visualization — BFS vs DFS

## Sample Graph

```
    0
   / \
  1   2
  |   |
  3   4
   \ /
    5
```

Adjacency: `0→[1,2], 1→[0,3], 2→[0,4], 3→[1,5], 4→[2,5], 5→[3,4]`

## BFS — Level Order Expansion

```mermaid
flowchart TD
    subgraph "Level 0"
        A((0))
    end
    subgraph "Level 1"
        B((1))
        C((2))
    end
    subgraph "Level 2"
        D((3))
        E((4))
    end
    subgraph "Level 3"
        F((5))
    end

    A --> B
    A --> C
    B --> D
    C --> E
    D --> F
```

### BFS Queue Trace

| Step | Queue | Visited | Processing |
|------|-------|---------|------------|
| 0 | [0] | {0} | start |
| 1 | [1, 2] | {0,1,2} | dequeue 0 |
| 2 | [2, 3] | {0,1,2,3} | dequeue 1 |
| 3 | [3, 4] | {0,1,2,3,4} | dequeue 2 |
| 4 | [4, 5] | {0,1,2,3,4,5} | dequeue 3 |
| 5 | [5] | {0,1,2,3,4,5} | dequeue 4 |
| 6 | [] | {0,1,2,3,4,5} | dequeue 5 |

## DFS — Depth-First Stack Trace

```mermaid
sequenceDiagram
    participant S as Call Stack

    Note over S: push dfs(0)
    S->>S: visit 0 → neighbors [1, 2]
    Note over S: push dfs(1)
    S->>S: visit 1 → neighbors [3]
    Note over S: push dfs(3)
    S->>S: visit 3 → neighbors [5]
    Note over S: push dfs(5)
    S->>S: visit 5 → neighbors [4]
    Note over S: push dfs(4)
    S->>S: visit 4 → neighbors [2]
    Note over S: push dfs(2)
    S->>S: visit 2 → no unvisited
    Note over S: pop all — done
```

### DFS Visit Order: `0 → 1 → 3 → 5 → 4 → 2`
### BFS Visit Order: `0 → 1 → 2 → 3 → 4 → 5`

> **Takeaway**: BFS visits by proximity (level). DFS visits by depth (follows one path to the end before backtracking).
