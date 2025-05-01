package interview_questions

## Understanding `var` in Different Data Structures in Go

### **Arrays (`[N]T`)**
Arrays are **fixed-size** collections of the same type.

```go
var arr1 [3]int        // Zero-initialized: [0, 0, 0]
var arr2 = [3]int{1, 2, 3} // Explicit values
var arr3 = [...]int{1, 2, 3, 4} // Compiler determines size
```
💡 **Arrays are value types**, meaning assignments copy the entire array.

---

### **Slices (`[]T`)**
Slices are **dynamic-size** collections.

```go
var s1 []int         // Declares a nil slice (no allocation)
var s2 = []int{1, 2, 3} // Initializes with values
var s3 = make([]int, 5) // Allocates a slice of size 5 with zero values
```
💡 **Slices are reference types**, so they share underlying arrays.

---

### **Maps (`map[K]V`)**
Maps store key-value pairs.

```go
var m1 map[string]int // Declares a nil map (needs explicit allocation)
var m2 = map[string]int{"a": 1, "b": 2} // Initialized with values
var m3 = make(map[string]int) // Allocates an empty map
```
💡 **Maps are reference types**, so modifying a map inside a function affects the original.

---

### **Structs (`struct`)**
Structs group different fields together.

```go
type Person struct {
    Name string
    Age  int
}

var p1 Person       // Zero-initialized: {"" 0}
var p2 = Person{"Alice", 25} // Direct initialization
var p3 = Person{Age: 30} // Partial initialization
var p4 = new(Person) // Allocates struct, returns pointer
```
💡 **Structs are value types** but can be used with pointers.

---

### **Pointers (`*T`)**
Pointers store memory addresses.

```go
var x int = 10
var p *int = &x  // p holds address of x
fmt.Println(*p)  // Dereferencing prints 10
```
💡 **Use pointers when modifying variables inside functions.**

---

## **How to Remember This?**

| Data Structure | Default Value | `var` Behavior |
|---------------|--------------|---------------|
| **Array** (`[N]T`) | `[0, 0, 0]` | Fixed size, value type |
| **Slice** (`[]T`) | `nil` | Dynamic size, reference type |
| **Map** (`map[K]V`) | `nil` | Reference type, needs `make` |
| **Struct** (`struct`) | Zero-values | Value type, use `new()` for pointers |
| **Pointer** (`*T`) | `nil` | Stores memory addresses |


