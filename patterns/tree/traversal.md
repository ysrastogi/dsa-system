# Traversal Pattern (DFS / BFS)

> This is the foundation. Everything else builds on this.

---

## Mental Model

You are **visiting nodes in a specific order**, not computing DP. The traversal itself is the goal — the order determines what you can do at each node.

---

## Traversal Orders

### DFS — Three Orders

```go
func dfs(node *TreeNode) {
    if node == nil {
        return
    }

    // PREORDER — process before children
    // Use: serialize, copy tree, prefix expressions
    process(node)
    dfs(node.Left)
    dfs(node.Right)
}
```

```go
func dfs(node *TreeNode) {
    if node == nil {
        return
    }

    dfs(node.Left)
    // INORDER — process between children
    // Use: BST sorted order, infix expressions
    process(node)
    dfs(node.Right)
}
```

```go
func dfs(node *TreeNode) {
    if node == nil {
        return
    }

    dfs(node.Left)
    dfs(node.Right)
    // POSTORDER — process after children
    // Use: delete tree, aggregate subtrees, evaluate expressions
    process(node)
}
```

### BFS — Level Order

```go
func bfs(root *TreeNode) {
    if root == nil {
        return
    }

    q := []*TreeNode{root}
    for len(q) > 0 {
        size := len(q)
        for i := 0; i < size; i++ {
            node := q[0]
            q = q[1:]

            process(node)

            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
    }
}
```

---

## DFS vs BFS — When to Use Which

| Use DFS when | Use BFS when |
|---|---|
| Need to reach leaves first (postorder) | Need level-by-level processing |
| Problem involves paths root → leaf | Need shortest path in unweighted tree |
| Space is O(h) — fine for balanced trees | Need to process by depth |
| Recursion is natural fit | Need positional/index info per level |

---

## Recognition Signals

Keywords that point to traversal:

- "visit", "print", "collect values"
- "serialize / deserialize"
- "convert to list"
- "level order", "zigzag"
- "right side view", "left side view"
- "flatten tree"

---

## Common Uses by Order

| Order | Typical Problems |
|---|---|
| **Preorder** | Serialize tree, clone tree, construct string from tree |
| **Inorder** | BST → sorted array, kth smallest in BST, validate BST |
| **Postorder** | Delete tree, evaluate expression tree, subtree problems |
| **Level order** | Level averages, right/left view, max width, zigzag |

---

## Common Mistakes

1. **Using recursion for BFS** — BFS is iterative with a queue. Don't force recursion.
2. **Forgetting `size := len(q)`** in BFS — without snapshotting level size, you can't distinguish levels.
3. **Modifying tree during inorder** — if collecting sorted values from BST, don't mutate while traversing.
4. **Using BFS when DFS suffices** — BFS uses O(w) space (width of tree). For deep narrow trees, DFS with O(h) is better.

---

## Iterative DFS (Stack-Based)

When recursion depth could overflow (deep skewed trees):

```go
func iterativeInorder(root *TreeNode) []int {
    var result []int
    var stack []*TreeNode
    curr := root

    for curr != nil || len(stack) > 0 {
        for curr != nil {
            stack = append(stack, curr)
            curr = curr.Left
        }
        curr = stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        result = append(result, curr.Val)
        curr = curr.Right
    }

    return result
}
```

---

## Complexity

| Variant | Time | Space |
|---|---|---|
| DFS (recursive) | O(n) | O(h) — call stack |
| DFS (iterative) | O(n) | O(h) — explicit stack |
| BFS | O(n) | O(w) — queue width |

Where `h` = height, `w` = max width of tree.

---

## When NOT to Use Traversal

- Problem says "combine left and right" → **Tree DP**
- Problem needs height/depth calculation → **Tree Height**
- Problem needs root-to-leaf accumulation → **Path Tracking**
- Traversal is necessary but not sufficient — if you need to **compute** something from children's results, you've moved beyond pure traversal
