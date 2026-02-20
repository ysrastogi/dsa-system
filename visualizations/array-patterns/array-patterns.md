# Array Patterns — Visual Reference

> Mermaid diagrams for each core array pattern. Use for rapid revision.

---

## 🗺️ Pattern Selection Flowchart

```mermaid
flowchart TD
    Start["Array Problem"] --> Q1{"Is input sorted?"}

    Q1 -->|Yes| Q2{"Looking for pair/triplet?"}
    Q1 -->|No| Q3{"Is it about subarrays?"}

    Q2 -->|Yes| TP["🎯 Two Pointer"]
    Q2 -->|No| BS["🎯 Binary Search"]

    Q3 -->|Yes| Q4{"Contiguous constraint?"}
    Q3 -->|No| Q5{"Next greater/smaller?"}

    Q4 -->|"Optimize length/sum"| SW["🎯 Sliding Window"]
    Q4 -->|"Range sum / count subarrays"| PS["🎯 Prefix Sum"]
    Q4 -->|"Max subarray sum"| KD["🎯 Kadane's"]

    Q5 -->|Yes| MS["🎯 Monotonic Stack"]
    Q5 -->|No| Q6{"Frequency / lookup?"}

    Q6 -->|Yes| HA["🎯 Hashing + Array"]
    Q6 -->|No| AD["🎯 Advanced Pattern"]

    style TP fill:#1b3d2d,color:#9effa5,stroke:#4caf50
    style SW fill:#1b3d2d,color:#9effa5,stroke:#4caf50
    style PS fill:#1b3d2d,color:#9effa5,stroke:#4caf50
    style KD fill:#1b3d2d,color:#9effa5,stroke:#4caf50
    style BS fill:#1b3d2d,color:#9effa5,stroke:#4caf50
    style MS fill:#1b3d2d,color:#9effa5,stroke:#4caf50
    style HA fill:#1b3d2d,color:#9effa5,stroke:#4caf50
    style AD fill:#1e3a5f,color:#90caf9,stroke:#42a5f5
```

---

## 1. Two Pointer — Converging Search

### How Search Space Shrinks

```mermaid
flowchart LR
    subgraph "Iteration 0"
        A["[① ② ③ ④ ⑤ ⑥ ⑦ ⑧]
   L→            ←R"]
    end
    subgraph "Iteration 1"
        B["[× ② ③ ④ ⑤ ⑥ ⑦ ×]
     L→        ←R"]
    end
    subgraph "Iteration 2"
        C["[× × ③ ④ ⑤ ⑥ × ×]
       L→    ←R"]
    end
    subgraph "Found!"
        D["[× × × ④ ⑤ × × ×]
         L  R ✅"]
    end

    A --> B --> C --> D
```

### Invariant Diagram

```mermaid
stateDiagram-v2
    [*] --> Evaluate: compare arr[L] + arr[R]
    Evaluate --> MoveLeft: sum < target
    Evaluate --> MoveRight: sum > target
    Evaluate --> Found: sum == target
    MoveLeft --> Evaluate: L++
    MoveRight --> Evaluate: R--
    Found --> [*]
```

---

## 2. Sliding Window — Expand-Shrink Cycle

### Window Lifecycle

```mermaid
stateDiagram-v2
    [*] --> Expand: right++, add element
    Expand --> CheckValid
    CheckValid --> Shrink: INVALID
    CheckValid --> UpdateAnswer: VALID
    Shrink --> CheckValid: left++, remove element
    UpdateAnswer --> Expand: continue
    UpdateAnswer --> [*]: right == n
```

### Fixed vs Variable Window

```mermaid
flowchart LR
    subgraph "Fixed Window (k=3)"
        F1["[②  ①  ⑤] · · ·"]
        F2["· [①  ⑤  ①] · ·"]
        F3["· · [⑤  ①  ③] ·"]
        F4["· · · [①  ③  ②]"]
        F1 --> F2 --> F3 --> F4
    end

    subgraph "Variable Window"
        V1["[a] ✅ grow"]
        V2["[a,b] ✅ grow"]
        V3["[a,b,c,a] ❌ shrink"]
        V4["[b,c,a] ✅ update"]
        V1 --> V2 --> V3 --> V4
    end
```

---

## 3. Prefix Sum — Precompute Once, Query O(1)

### Construction

```mermaid
flowchart TD
    subgraph "Array"
        A["[3, 1, 4, 1, 5]"]
    end
    subgraph "Prefix Sum"
        P["[0, 3, 4, 8, 9, 14]"]
    end
    subgraph "Query: sum(1,3)"
        Q["prefix[4] - prefix[1]
= 9 - 3 = 6 ✅"]
    end

    A -->|"cumulative add"| P -->|"one subtraction"| Q

    style Q fill:#1b3d2d,color:#9effa5,stroke:#4caf50
```

### Prefix + HashMap for "Count subarrays with sum K"

```mermaid
flowchart TD
    I["For each index: prefixSum += arr[i]"]
    L["Look up freq[prefixSum - K]"]
    C["count += freq[prefixSum - K]"]
    U["freq[prefixSum]++"]
    N["Next index"]

    I --> L --> C --> U --> N --> I

    style L fill:#1e3a5f,color:#90caf9,stroke:#42a5f5
```

### Difference Array for Range Updates

```mermaid
flowchart LR
    subgraph "Add +5 to [1,3]"
        D1["diff[1] += 5"]
        D2["diff[4] -= 5"]
    end
    subgraph "Difference Array"
        D["[0, +5, 0, 0, -5, 0]"]
    end
    subgraph "Running Sum → Result"
        R["[0, 5, 5, 5, 0, 0]"]
    end

    D1 --> D
    D2 --> D
    D -->|"prefix sum"| R
```

---

## 4. Kadane's — Extend vs Restart Decision

### Decision Tree at Each Index

```mermaid
flowchart TD
    D["At index i, arr[i] = x"]
    C{"current + x > x ?
(is prefix positive?)"}
    E["✅ EXTEND
current += x"]
    R["🔄 RESTART
current = x"]
    U["best = max(best, current)"]

    D --> C
    C -->|"Yes"| E
    C -->|"No"| R
    E --> U
    R --> U

    style E fill:#1b3d2d,color:#9effa5,stroke:#4caf50
    style R fill:#1e3a5f,color:#90caf9,stroke:#42a5f5
```

### Trace on `[-2, 1, -3, 4, -1, 2, 1, -5, 4]`

```mermaid
flowchart LR
    subgraph "Phase 1: Searching"
        A["-2 → cur=-2"] --> B["1 → 🔄 cur=1"]
        B --> C["-3 → cur=-2"]
    end
    subgraph "Phase 2: Building"
        D["4 → 🔄 cur=4"] --> E["-1 → cur=3"]
        E --> F["2 → cur=5"] --> G["1 → cur=6 ✅"]
    end
    subgraph "Phase 3: Declining"
        H["-5 → cur=1"] --> I["4 → cur=5"]
    end

    C --> D
    G --> H

    style G fill:#1b3d2d,color:#9effa5,stroke:#4caf50
```

---

## 5. Binary Search — Halving the Space

### Classic Search Visualization

```mermaid
flowchart TD
    S0["[1, 3, 5, 7, 9, 11, 13] target=9
low=0, high=6, mid=3 → 7 < 9"]
    S1["low=4, high=6, mid=5 → 11 > 9"]
    S2["low=4, high=4, mid=4 → 9 == 9 ✅"]

    S0 -->|"go right"| S1
    S1 -->|"go left"| S2

    style S2 fill:#1b3d2d,color:#9effa5,stroke:#4caf50
```

### Binary Search on Answer — Monotonicity

```mermaid
graph LR
    A1["val=1 ❌"] --> A2["val=2 ❌"] --> A3["val=3 ❌"]
    A3 --> A4["val=4 ❌"] --> A5["val=5 ✅ ← boundary"]
    A5 --> A6["val=6 ✅"] --> A7["val=7 ✅"]

    style A5 fill:#1b3d2d,color:#9effa5,stroke:#4caf50
    style A6 fill:#1b3d2d,color:#9effa5,stroke:#4caf50
    style A7 fill:#1b3d2d,color:#9effa5,stroke:#4caf50
```

---

## 6. Monotonic Stack — Filtered History

### Stack State as Elements Arrive

Array: `[3, 7, 1, 4, 2]` — finding Next Greater

```mermaid
flowchart TD
    S0["Push 3
Stack: [3]"]
    S1["7 > 3 → Pop 3, NGE[0]=7
Push 7
Stack: [7]"]
    S2["1 < 7 → Push 1
Stack: [7, 1]"]
    S3["4 > 1 → Pop 1, NGE[2]=4
4 < 7 → Push 4
Stack: [7, 4]"]
    S4["2 < 4 → Push 2
Stack: [7, 4, 2]
Done → NGE[1,3,4] = -1"]

    S0 --> S1 --> S2 --> S3 --> S4

    style S1 fill:#1b3d2d,color:#9effa5,stroke:#4caf50
    style S3 fill:#1b3d2d,color:#9effa5,stroke:#4caf50
```

### Why Popping is Safe

```mermaid
flowchart TD
    Q["Stack has [A, B] where A > B
New element C arrives, C > B"]
    P["Pop B → B's answer is C"]
    W["Why safe? Any future element D:"]
    C1["If D > B → D would also > B, but C is closer
So C is still the right answer for B ✅"]
    C2["B can never be useful again → remove it"]

    Q --> P --> W --> C1 --> C2

    style C1 fill:#1b3d2d,color:#9effa5,stroke:#4caf50
```

---

## 7. Hashing + Array — Domain-Aware Lookup

### Decision: HashMap vs Array Bucket

```mermaid
flowchart TD
    Q{"What is the value domain?"}
    S1{"≤ 26?"} -->|Yes| A1["[26]int ← fastest"]
    S2{"≤ 128?"} -->|Yes| A2["[128]int"]
    S3{"≤ 10⁵?"} -->|Yes| A3["[]int slice"]
    S4{"Unbounded"} --> A4["map[int]int"]

    Q --> S1
    Q --> S2
    Q --> S3
    Q --> S4

    style A1 fill:#1b3d2d,color:#9effa5,stroke:#4caf50
    style A4 fill:#2d1b36,color:#ff9e9e,stroke:#ff5252
```

### In-Place Marking: Values in [1,n]

```mermaid
flowchart LR
    subgraph "Original"
        O["[3, 1, 3, 4, 2]"]
    end
    subgraph "After marking"
        M["[-3, -1, -3, -4, 2]
idx 4 still positive → 5 is missing!"]
    end

    O -->|"negate arr[val-1]"| M

    style M fill:#1e3a5f,color:#90caf9,stroke:#42a5f5
```

---

## 🧩 Advanced Patterns

### Boyer-Moore Voting

```mermaid
flowchart LR
    subgraph "Array: [2, 2, 1, 1, 2, 2, 1]"
        S1["cand=2, cnt=1"] --> S2["cand=2, cnt=2"]
        S2 --> S3["cnt=1 (1≠2)"] --> S4["cnt=0 (1≠2)"]
        S4 --> S5["cand=2, cnt=1"] --> S6["cand=2, cnt=2"]
        S6 --> S7["cnt=1 (1≠2)"]
    end

    S7 --> R["Answer: 2 ✅"]

    style R fill:#1b3d2d,color:#9effa5,stroke:#4caf50
```

### Dutch National Flag (3-Way Partition)

```mermaid
flowchart TD
    subgraph "Before"
        B["[2, 0, 2, 1, 1, 0]
low=0, mid=0, high=5"]
    end
    subgraph "After"
        A["[0, 0, 1, 1, 2, 2]
All 0s left, 1s middle, 2s right"]
    end

    B -->|"3 pointers: low, mid, high"| A

    style A fill:#1b3d2d,color:#9effa5,stroke:#4caf50
```
