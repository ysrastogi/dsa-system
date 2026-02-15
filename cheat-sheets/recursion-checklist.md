# 🧠 Recursion Checklist

Use this before any recursion problem in an interview. 60 seconds to orient yourself.

---

## Step 1: Identify the Pattern

| Signal | Pattern |
|--------|---------|
| "Generate all…" | Backtracking |
| "Count ways to…" | DP / Memoized recursion |
| "Max/min path in tree" | Tree DP (postorder) |
| "Shortest path" | BFS, not recursion |
| "Explore connected" | DFS |
| "Divide into halves" | Divide & conquer |

---

## Step 2: Define the Recursion

Ask these **three questions**:

1. **What does this function return?** (the "contract" with its caller)
2. **What is the base case?** (when to stop recursing)
3. **What is the recursive step?** (how to break into subproblems)

---

## Step 3: Track State

| Question | Answer → Goes Where |
|----------|-------------------|
| What flows **UP** (return)? | `return` statement |
| What flows **DOWN** (parameter)? | Function arguments |
| What is **GLOBAL** (side effect)? | Outer variable / pointer |

---

## Step 4: Mental Simulation

Before coding, trace through the **smallest non-trivial example**:

1. Draw the recursion tree (3-4 nodes max)
2. Write the return value at each leaf
3. Propagate values upward
4. Check: does the root's return value match expected output?

---

## Step 5: Watch for Traps

- [ ] **Off-by-one**: Is base case `nil` or leaf? (They're different!)
- [ ] **Negative values**: Are you clamping to 0 when needed?
- [ ] **Global vs Return**: Are you confusing what goes up with what's global?
- [ ] **Visited set**: For graph recursion, are you marking visited **before** recursing?
- [ ] **Copy vs Reference**: Are you modifying shared state? Need `copy()`?
- [ ] **Memoization**: Are there overlapping subproblems? Add a cache.

---

## Quick Templates

### Tree DFS
```go
func dfs(node *TreeNode) int {
    if node == nil { return 0 }
    left := dfs(node.Left)
    right := dfs(node.Right)
    // combine and return
}
```

### Backtracking
```go
func backtrack(state, choices) {
    if done { record; return }
    for choice in choices {
        if valid { add; recurse; remove }
    }
}
```

### Memoized DP
```go
func dp(i int, memo map[int]int) int {
    if val, ok := memo[i]; ok { return val }
    if base { return baseVal }
    result := /* recurrence */
    memo[i] = result
    return result
}
```
