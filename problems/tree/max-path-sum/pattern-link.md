# Pattern Link — Max Path Sum

## Pattern

**Tree DP — Postorder Accumulation**

→ [patterns/tree/tree-dp.md](../../patterns/tree/tree-dp.md)

## How This Problem Uses the Pattern

| Aspect | In This Problem |
|--------|----------------|
| **What flows UP** | Best single chain: `max(left, right) + node.Val` |
| **What is GLOBAL** | Best forked path: `left + right + node.Val` |
| **What flows DOWN** | Nothing (no parent-to-child info needed) |
| **Base case** | `nil` node returns `0` |
| **Negative handling** | Clamp child returns to `max(0, child)` |

## Why This Pattern Fits

The answer at any node depends on children's results → postorder is required. The "best answer" may pass through any node (not just root) → global tracking is required. The return value differs from the global value → dual tracking is required.

## Related Problems Using Same Pattern

- **Diameter of Binary Tree** (LC 543) — UP: height, GLOBAL: left+right
- **Longest Univalue Path** (LC 687) — UP: matching chain, GLOBAL: left+right matching
- **Binary Tree Cameras** (LC 968) — UP: coverage state, GLOBAL: camera count
