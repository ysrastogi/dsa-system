# Binary Tree Maximum Path Sum

## Problem

A **path** in a binary tree is a sequence of nodes where each pair of adjacent nodes has an edge connecting them. A node can only appear in the sequence **at most once**. The path does **not** need to pass through the root.

The **path sum** is the sum of the node values in the path.

Given the `root` of a binary tree, return the **maximum path sum** of any non-empty path.

## Constraints

- Number of nodes: `[1, 3 × 10⁴]`
- Node values: `[-1000, 1000]`

## Examples

### Example 1
```
Input:  [1, 2, 3]
Tree:
    1
   / \
  2   3

Output: 6
Path: 2 → 1 → 3
```

### Example 2
```
Input:  [-10, 9, 20, null, null, 15, 7]
Tree:
    -10
    / \
   9  20
      / \
     15  7

Output: 42
Path: 15 → 20 → 7
```

## Key Observations

1. A path **forks at most once** — it goes up from one subtree, through a node, and down into another
2. The return value to parent must be a **single chain** (can't fork twice)
3. Negative contributions should be **clamped to 0** (don't extend a path if it hurts)

## Difficulty

**Hard** (LeetCode 124)
