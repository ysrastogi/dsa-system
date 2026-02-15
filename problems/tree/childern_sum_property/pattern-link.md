# Pattern Link — Children Sum Property

## Pattern

**Tree Modification — Push Down, Pull Up**

→ [patterns/tree/tree-modification.md](../../patterns/tree/tree-modification.md)

## How This Problem Uses the Pattern

| Aspect | In This Problem |
|--------|----------------|
| **What flows DOWN** | Parent's value → overwrite children when `childSum < parent` |
| **What flows UP** | Sum of children's final values → set parent |
| **Mutation type** | In-place node value modification |
| **Constraint** | Values can only increase, never decrease |
| **Base case** | `nil` node → return (no-op) |
| **Leaf behavior** | Value unchanged (no children to sum) |

## Why This Pattern Fits

The problem requires **two-pass modification**: first push surplus value downward (so it's not lost), then recurse (children may grow from their subtrees), then pull final children values back up. Neither pure preorder nor pure postorder alone can solve this — you need the **hybrid**.

## The Three Phases at Each Node

```
1. COMPARE  →  childSum vs node.Val
2. PUSH DOWN →  if childSum < node.Val, stamp parent's value onto children
3. RECURSE   →  left, right (children may grow)
4. PULL UP   →  node.Val = left.Val + right.Val
```

## Related Problems Using Same Pattern

- **Distribute Coins in Binary Tree** (LC 979) — tree modification with flow tracking
- **Convert BST to Greater Tree** (LC 538) — different modification (reverse inorder)
