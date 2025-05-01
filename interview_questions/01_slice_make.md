## What is the difference between make([]int, 10), make([]int, 0, 10), make([]int, 10, 10) in golang in Go


## Understanding `make([]T, len, cap)` in Go

In Go, `make([]T, length, capacity)` is used to create a slice with a specified length and capacity. Below are the differences between the three cases:

### 1. `make([]int, 10)`
- Creates a slice of `int` with **length = 10** and **capacity = 10**.
- The slice is initialized with zero values (i.e., all elements are `0`).
- Since length and capacity are the same, there is no extra room for appending elements beyond 10 without allocating a new array.

Example:
```go
s := make([]int, 10)
fmt.Println(len(s), cap(s)) // Output: 10 10
fmt.Println(s)              // Output: [0 0 0 0 0 0 0 0 0 0]
```

---

### 2. `make([]int, 0, 10)`
- Creates a slice with **length = 0** but **capacity = 10**.
- Since the length is `0`, the slice is empty.
- However, the slice has preallocated space for up to 10 elements before a new allocation is needed.

Example:
```go
s := make([]int, 0, 10)
fmt.Println(len(s), cap(s)) // Output: 0 10
fmt.Println(s)              // Output: []

s = append(s, 1, 2, 3)
fmt.Println(len(s), cap(s)) // Output: 3 10
fmt.Println(s)              // Output: [1 2 3]
```

---

### 3. `make([]int, 10, 10)`
- This is **identical to `make([]int, 10)`**.
- It creates a slice with **length = 10** and **capacity = 10**.
- The slice is initialized with zero values.

Example:
```go
s := make([]int, 10, 10)
fmt.Println(len(s), cap(s)) // Output: 10 10
fmt.Println(s)              // Output: [0 0 0 0 0 0 0 0 0 0]
```

---

## Key Differences

| Expression           | Length (`len`) | Capacity (`cap`) | Initial Content |
|----------------------|--------------|--------------|----------------|
| `make([]int, 10)`  | 10           | 10           | `[0, 0, 0, 0, 0, 0, 0, 0, 0, 0]` |
| `make([]int, 0, 10)` | 0            | 10           | `[]` (empty slice) |
| `make([]int, 10, 10)` | 10           | 10           | `[0, 0, 0, 0, 0, 0, 0, 0, 0, 0]` |

---

## When to Use Each?
- **`make([]int, 10)` or `make([]int, 10, 10)`**
    - Use when you need a slice of exactly 10 elements, initialized with zero values.

- **`make([]int, 0, 10)`**
    - Use when you need an empty slice but want to preallocate memory for future appends to avoid frequent reallocations.

