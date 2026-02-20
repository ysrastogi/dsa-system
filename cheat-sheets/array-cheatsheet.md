# 🔥 Array Cheat Sheet — Interview Grade

> 60-second pre-problem orientation. If this isn't automatic, you're still slow.

---

## 1️⃣ Mental Model

```
Array = Contiguous Memory + O(1) Access + Index-Driven Reasoning
```

**Before writing a single line, ask:**

| Question | Why It Matters |
|----------|---------------|
| Is it contiguous? | Sliding window viable? |
| Do I need order? | Can I sort? |
| Is input sorted? | Two pointer / binary search |
| Is domain bounded? | Array bucket > hashmap |
| Can I preprocess? | Prefix sum / sort first |

If you don't ask these, you're guessing.

---

## 2️⃣ Pattern Trigger Table

| If Problem Says | Brain Should Trigger |
|----------------|---------------------|
| "Subarray" | Sliding Window / Prefix Sum |
| "Contiguous" | Sliding Window |
| "Range sum queries" | Prefix Sum |
| "Count subarrays with sum…" | Prefix Sum + HashMap |
| "Pair in sorted array" | Two Pointer |
| "Next greater / smaller" | Monotonic Stack |
| "Maximum subarray" | Kadane's |
| "Search in sorted space" | Binary Search |
| "Frequency (small domain)" | Counting Array |
| "Min/max value such that…" | Binary Search on Answer |
| "Duplicate detection" | Hashing / In-place Marking |

---

## 3️⃣ Core Patterns — Compressed

### 🧠 Two Pointer

```
WHEN: sorted array, pair/triplet, converging decision
INVARIANT: left ≤ right, search space shrinks every step
DANGER: moving pointer without justification
```

```go
for left < right {
    if condition { left++ } else { right-- }
}
```

---

### 🧠 Sliding Window

```
WHEN: contiguous, optimize length/sum
FLOW: expand → fix violation → update answer
INVARIANT: window always valid after shrink step
DANGER: not defining "valid"
```

```go
for right := 0; right < n; right++ {
    add(arr[right])
    for !valid() { remove(arr[left]); left++ }
    best = max(best, right-left+1)
}
```

---

### 🧠 Prefix Sum

```
WHEN: range queries, counting subarrays
IDENTITY: sum(l,r) = prefix[r+1] - prefix[l]
ADVANCED: prefix XOR, prefix freq map, difference array
TRIGGER: "how many subarrays…" → prefix + hashmap
```

```go
prefix[i+1] = prefix[i] + arr[i]
// "count subarrays with sum K":
count += freq[prefixSum - k]; freq[prefixSum]++
```

---

### 🧠 Kadane's

```
WHEN: maximum/minimum subarray, contiguous optimization
DECISION: extend previous OR start fresh
WHY: negative prefix is a burden — drop it
```

```go
current = max(arr[i], current + arr[i])
best = max(best, current)
```

---

### 🧠 Binary Search on Answer

```
WHEN: min/max value, feasibility check exists, monotonic property
FRAMEWORK: define range → check feasibility → halve space
DANGER: not proving monotonicity
```

```go
for low < high {
    mid := low + (high-low)/2
    if feasible(mid) { high = mid } else { low = mid+1 }
}
```

---

### 🧠 Monotonic Stack

```
WHEN: next greater/smaller, histogram, span problems
INVARIANT: stack always monotonic (increasing or decreasing)
KEY: pop until invariant holds — popped element found its answer
DANGER: memorizing template without understanding why popping is safe
```

```go
for i := 0; i < n; i++ {
    for len(stack) > 0 && arr[stack[top]] < arr[i] {
        result[stack[top]] = arr[i]; pop()
    }
    push(i)
}
```

---

### 🧠 Hashing + Array

```
WHEN: frequency, fast lookup, duplicate detection
RULE: domain ≤ 10⁵ → use array, not hashmap
ADVANCED: in-place marking when values ∈ [1,n] → O(1) space
```

---

## 4️⃣ Complexity Table

| Operation | Time |
|-----------|------|
| Access by index | O(1) |
| Search (unsorted) | O(n) |
| Search (sorted) | O(log n) |
| Insert middle | O(n) |
| Delete middle | O(n) |
| Append (dynamic) | Amortized O(1) |
| Sort | O(n log n) |

---

## 5️⃣ Edge Case Checklist ⚠️

Before submitting, **always check:**

- [ ] `n = 0` — empty array?
- [ ] `n = 1` — single element?
- [ ] All negative values?
- [ ] All same elements?
- [ ] Integer overflow possible?
- [ ] Index out of bounds?
- [ ] Duplicate handling correct?
- [ ] Off-by-one in window/prefix boundaries?

> You lose **easy** interviews here.

---

## 6️⃣ Advanced Layer (Medium-Hard Differentiator)

| Technique | When |
|-----------|------|
| **Difference Array** | Batch range updates in O(1) each |
| **Coordinate Compression** | Map large sparse values to dense [0, k] |
| **In-place Marking** | Values in [1,n] → negate arr[val-1] for O(1) space |
| **Cyclic Sort** | Find missing/duplicate in [1,n] range |
| **Boyer–Moore Voting** | Find majority element in O(1) space |
| **Dutch National Flag** | 3-way partition (sort 0s, 1s, 2s) |

### Boyer–Moore Voting

```go
func majorityElement(nums []int) int {
    candidate, count := 0, 0
    for _, num := range nums {
        if count == 0 { candidate = num }
        if num == candidate { count++ } else { count-- }
    }
    return candidate
}
```

### Dutch National Flag

```go
func sortColors(nums []int) {
    low, mid, high := 0, 0, len(nums)-1
    for mid <= high {
        switch nums[mid] {
        case 0: nums[low], nums[mid] = nums[mid], nums[low]; low++; mid++
        case 1: mid++
        case 2: nums[mid], nums[high] = nums[high], nums[mid]; high--
        }
    }
}
```

### Cyclic Sort

```go
func findMissing(nums []int) int {
    for i := 0; i < len(nums); i++ {
        for nums[i] > 0 && nums[i] <= len(nums) && nums[nums[i]-1] != nums[i] {
            nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
        }
    }
    for i, num := range nums {
        if num != i+1 { return i + 1 }
    }
    return len(nums) + 1
}
```
