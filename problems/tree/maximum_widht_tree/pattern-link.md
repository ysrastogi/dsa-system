# Pattern Link — Maximum Width of Binary Tree

## Pattern

**BFS Level-Order — Index Tracking**

→ [patterns/tree/bfs-level-order.md](../../../patterns/tree/bfs-level-order.md)

## How This Problem Uses the Pattern

| Aspect | In This Problem |
|--------|----------------|
| **Queue item** | `{node, positional index}` |
| **Index formula** | Left: `2*idx + 1`, Right: `2*idx + 2` |
| **Per-level computation** | `width = lastIdx - firstIdx + 1` |
| **Global tracking** | `maxWidth = max(maxWidth, width)` |
| **Normalization** | Subtract `start` index each level to prevent overflow |
| **Null handling** | Nulls are never enqueued — gaps are implicit via index arithmetic |

## Why This Pattern Fits

The width of a level is defined by the **positions** of the leftmost and rightmost non-null nodes, including gaps. This requires positional awareness that raw BFS doesn't have — hence the index pairing. Processing level-by-level (BFS, not DFS) is natural because width is a **per-level** property.

## Why Alternatives Fail

| Approach | Problem |
|----------|---------|
| **DFS** | No natural concept of "level width" — you'd need to track min/max index per depth in a map, losing the clean level-by-level flow |
| **Array-based BFS** | Storing nulls to compute positions uses O(2ⁿ) space and TLE on deep skewed trees |
| **No normalization** | Index 2⁵⁰ on a left-skewed tree with depth 50 overflows `int` |

## Related Problems Using Same Pattern

- **Binary Tree Right Side View** (LC 199) — BFS, last node per level
- **Find Bottom Left Tree Value** (LC 513) — BFS, first node of last level
- **Binary Tree Level Order Traversal** (LC 102) — BFS, collect all nodes per level
- **Cousins in Binary Tree** (LC 993) — BFS with parent+depth tracking
