# Tree Common Mistakes

> These mistakes cost more time than not knowing the pattern at all.
> Read this before interviews. Recognizing the mistake is faster than debugging it.

---

## Mistake 1 — Returning the Forked Path

**Where**: Tree DP problems (diameter, max path sum)

```go
// ❌ WRONG — returns forked path to parent
func dfs(node *TreeNode) int {
    left := dfs(node.Left)
    right := dfs(node.Right)
    return left + right + node.Val  // parent can't use this — path forks twice
}

// ✅ CORRECT — return chain, update global with fork
func dfs(node *TreeNode) int {
    left := dfs(node.Left)
    right := dfs(node.Right)
    global = max(global, left+right+node.Val)      // fork stays here
    return max(left, right) + node.Val              // chain goes up
}
```

**Why it breaks**: A path cannot fork at two nodes. If you return `left + right + node.Val`, the parent node will try to fork again.

---

## Mistake 2 — Wrong Base Case

**Where**: Any recursive tree problem

```go
// ❌ WRONG — height base case
if node == nil { return -1 }  // off-by-one everywhere

// ✅ CORRECT
if node == nil { return 0 }
```

```go
// ❌ WRONG — max path sum base case
if node == nil { return math.MinInt }  // poisons parent computation

// ✅ CORRECT
if node == nil { return 0 }
// then clamp: return max(0, max(left, right) + node.Val)
```

**Rule**: Ask *"What does nil contribute?"* — usually `0`, `true`, or `nil`.

---

## Mistake 3 — Forgetting Global Reset

**Where**: Tree DP with package-level globals

```go
var globalMax int  // ← still holds value from previous test case!

func maxPathSum(root *TreeNode) int {
    // ❌ WRONG — forgot to reset
    dfs(root)
    return globalMax
}

// ✅ CORRECT
func maxPathSum(root *TreeNode) int {
    globalMax = math.MinInt32  // reset before each call
    dfs(root)
    return globalMax
}
```

**In interviews**: This bug doesn't show on single test cases but fails on judges running multiple inputs.

---

## Mistake 4 — Mixing Top-Down and Bottom-Up

**Where**: Confusing Path Tracking with Tree DP

```go
// ❌ CONFUSED — trying to pass state down AND aggregate up
func dfs(node *TreeNode, pathSum int) int {   // path sum flowing DOWN
    left := dfs(node.Left, pathSum + node.Val) // but also returning UP?
    right := dfs(node.Right, pathSum + node.Val)
    return max(left, right)  // what does this even mean?
}
```

**Fix**: Choose one direction:
- **Top-down** (Path Tracking): state flows in parameters, no meaningful return
- **Bottom-up** (Tree DP / Height): result flows in return value, no meaningful parameters

---

## Mistake 5 — Using Preorder When Aggregation Is Needed

**Where**: Problems requiring combine-left-and-right

```go
// ❌ WRONG — preorder can't aggregate
func dfs(node *TreeNode) {
    process(node)        // don't have children's results yet!
    dfs(node.Left)
    dfs(node.Right)
}

// ✅ CORRECT — postorder for aggregation
func dfs(node *TreeNode) int {
    left := dfs(node.Left)
    right := dfs(node.Right)
    return combine(left, right, node)  // now you have both
}
```

**Rule**: If the word "combine" appears in your thinking, use postorder.

---

## Mistake 6 — Not Clamping Negative Chains

**Where**: Max path sum and similar DP problems

```go
// ❌ WRONG — negative chain poisons parent
left := dfs(node.Left)   // returns -5
right := dfs(node.Right)  // returns 3
return max(left, right) + node.Val  // uses -5 if node.Val is large

// ✅ CORRECT — clamp to 0
left := max(0, dfs(node.Left))
right := max(0, dfs(node.Right))
```

**Why**: A negative subtree is worse than not taking it. Clamping to 0 means "don't extend into this subtree."

---

## Mistake 7 — Confusing Height and Depth

**Where**: Height vs depth questions

```
Height = distance from node to farthest leaf (bottom-up)
Depth  = distance from root to node (top-down)
```

| Property | Direction | Pattern |
|----------|-----------|---------|
| Height | Bottom-up (postorder) | Tree Height |
| Depth | Top-down (parameter) | Path Tracking |

```go
// Height — postorder
func height(node *TreeNode) int {
    if node == nil { return 0 }
    return max(height(node.Left), height(node.Right)) + 1
}

// Depth — top-down parameter
func dfs(node *TreeNode, depth int) {
    if node == nil { return }
    dfs(node.Left, depth+1)
    dfs(node.Right, depth+1)
}
```

---

## Mistake 8 — Building Tree Without Index Boundaries

**Where**: Tree construction problems

```go
// ❌ WRONG — creating new slices (O(n²) total)
func build(preorder []int, inorder []int) *TreeNode {
    root := preorder[0]
    leftInorder := inorder[:splitIdx]       // copies!
    rightInorder := inorder[splitIdx+1:]    // copies!
}

// ✅ CORRECT — use index boundaries
func build(preorder []int, inorder []int, preStart, inStart, inEnd int) *TreeNode {
    // O(1) per call, indices only
}
```

**Also**: Use a hashmap for inorder index lookup. Linear search makes construction O(n²).

---

## Quick Checklist (Pre-Interview)

- [ ] Am I returning chain or fork? (Tree DP)
- [ ] Did I reset globals?
- [ ] Is my base case correct for nil?
- [ ] Am I clamping negative values?
- [ ] Height or depth — which direction?
- [ ] Preorder or postorder — which do I need?
- [ ] Top-down or bottom-up — pick one, don't mix
