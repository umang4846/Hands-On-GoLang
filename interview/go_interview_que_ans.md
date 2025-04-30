### 1. What are the three uses of "..." in Go? (The first two are relatively familiar, but the third one might need a little more thought.)

In Go, `...` (ellipsis) has three main uses:

1. **Variadic Functions**  
   The most common use of `...` is to define functions that accept a variable number of arguments of a specific type.

   ```go
   func sum(nums ...int) int {
       total := 0
       for _, num := range nums {
           total += num
       }
       return total
   }

   func main() {
       fmt.Println(sum(1, 2, 3, 4)) // Output: 10
   }
   ```

2. **Expanding a Slice into Variadic Arguments**  
   When calling a variadic function, you can use `...` to expand a slice into individual arguments.

   ```go
   func main() {
       numbers := []int{1, 2, 3, 4}
       fmt.Println(sum(numbers...)) // Output: 10
   }
   ```

3. **Specifying an Array's Length Implicitly**  
   In array declarations, `...` can be used instead of specifying an explicit length, allowing the compiler to determine
   the size based on the provided elements.

   ```go
   func main() {
       arr := [...]int{1, 2, 3, 4, 5} // Compiler infers length as 5
       fmt.Println(len(arr))          // Output: 5
   }
   ```

---

### 2. What's the difference between a type definition and a type alias? in golang

In Go, **type definitions** and **type aliases** serve different purposes.

### **1. Type Definition (`type NewType OldType`)**

A type definition creates a completely new, distinct type from an existing type. The new type is **not interchangeable**
with the original type, even if they have the same underlying representation.

#### Example:

```go
package main

import "fmt"

type Age int // Age is a new, distinct type

func main() {
	var a Age = 30
	var b int = 30

	// fmt.Println(a == b) // Error: mismatched types Age and int

	fmt.Println(a) // Output: 30
}
```

- `Age` is a **new type** based on `int`, but it **cannot** be used interchangeably with `int` without explicit
  conversion.
- If you need to assign an `int` to `Age`, you must cast it:
  ```go
  var a Age = Age(b) // Explicit conversion required
  ```

---

### **2. Type Alias (`type NewType = OldType`)**

A type alias **does not create a new type** but instead provides an alternate name for an existing type. It is **fully
interchangeable** with the original type.

#### Example:

```go
package main

import "fmt"

type Age = int // Age is an alias for int

func main() {
	var a Age = 30
	var b int = 30

	fmt.Println(a == b) // Works fine since Age is just an alias for int
}
```

- `Age` is simply another name for `int`. The compiler treats `Age` and `int` as the same type.
- No explicit conversion is needed.

---

### **Key Differences**

| Feature                         | Type Definition (`type T OldType`)    | Type Alias (`type T = OldType`)             |
|---------------------------------|---------------------------------------|---------------------------------------------|
| Creates a new type?             | ✅ Yes                                 | ❌ No (just a different name)                |
| Requires explicit conversion?   | ✅ Yes                                 | ❌ No                                        |
| Can define new methods?         | ✅ Yes (methods can be attached to it) | ❌ No (inherits the original type’s methods) |
| Interchangeable with base type? | ❌ No                                  | ✅ Yes                                       |

---

### **When to Use Each?**

- **Use a type definition** when you want to create a distinct type with separate behavior (
  e.g., `UserID`, `Distance`, `Age`).
- **Use a type alias** when you want to rename a type without changing its behavior (e.g., migrating from `rune`
  to `char` while keeping compatibility).

---

### 3. Given a type parameter T, what is the constraint that requires values of T to support the == operator? What is the constraint for the < operator? What operators are usable with T values if the constraint is any?

### Constraints for Operators in Go Generics

#### **1. Constraint for `==` (Equality Operator)**

To use `==` (or `!=`) with a type parameter `T`, `T` must be a **comparable** type. In Go, this means `T` must be a type
that can be used in equality comparisons, such as:

- Built-in types like `int`, `float64`, `string`, `bool`
- Structs (if all fields are comparable)
- Pointers
- Arrays (if their elements are comparable)

👉 **Constraint:**

```go
type MyConstraint interface {
comparable
}
```

👉 **Example:**

```go
func IsEqual[T comparable](a, b T) bool {
return a == b
}

func main() {
fmt.Println(IsEqual(10, 10)) // ✅ Works (int is comparable)
fmt.Println(IsEqual("Go", "Go")) // ✅ Works (string is comparable)

// fmt.Println(IsEqual([]int{1,2}, []int{1,2})) ❌ ERROR: Slices are not comparable
}
```

---

#### **2. Constraint for `<` (Less Than Operator)**

The `<`, `>`, `<=`, and `>=` operators require a type that supports ordering, such as:

- `int`, `float64`, `string`

👉 **Constraint:**

```go
type Ordered interface {
~int | ~float64 | ~string // Custom constraint (Go 1.18+)
}
```

👉 **Example:**

```go
type Ordered interface {
~int | ~float64 | ~string
}

func Min[T Ordered](a, b T) T {
if a < b {
return a
}
return b
}

func main() {
fmt.Println(Min(10, 20)) // ✅ Works (int)
fmt.Println(Min(3.14, 2.71)) // ✅ Works (float64)
fmt.Println(Min("apple", "bat")) // ✅ Works (string)
}
```

**🚨 Note:** `bool` and `struct` types are not ordered, so they can't be used with `<`.

---

#### **3. What if `T` has no constraints (`any`)?**

If `T` is unconstrained (`any`), **no operators** are guaranteed to work because `T` can be **any type** (even slices,
maps, or interfaces that don't support `==`, `<`, etc.).

👉 **Example:**

```go
func Print[T any](val T) {
fmt.Println(val)
}

func main() {
Print(42)      // ✅ Works
Print("Hello") // ✅ Works
Print([]int{1, 2}) // ✅ Works (but no operators allowed)
}
```

🚨 **You can't use `==` or `<` on `T` if it's just `any`**:

```go
func Compare[T any](a, b T) bool {
return a == b // ❌ ERROR: "invalid operation"
}
```

---

### **🔥 Summary (Easy to Remember)**

| Operator              | Constraint                                          |
|-----------------------|-----------------------------------------------------|
| `==`, `!=`            | `comparable`                                        |
| `<`, `>`, `<=`, `>=`  | `Ordered` (custom constraint for numbers & strings) |
| No constraint (`any`) | No operators are guaranteed                         |

💡 **Rule of thumb:**

- Use `comparable` for equality checks (`==`).
- Use a custom `Ordered` constraint for comparisons (`<`).
- Use `any` only if you don't need operators.

---

### 4. What are the 25 keywords of Go?

Go has **25 reserved keywords**, which are part of the language syntax and cannot be used as identifiers (e.g., variable
names, function names).

---

### **📌 The 25 Keywords in Go**

#### **1. Control Flow (5)**

- `if`
- `else`
- `switch`
- `case`
- `fallthrough`

🔹 **Example:**

```go
if x > 10 {
fmt.Println("Greater")
} else {
fmt.Println("Smaller")
}
```

---

#### **2. Looping (3)**

- `for`
- `range`
- `break`

🔹 **Example:**

```go
for i := 0; i < 5; i++ {
fmt.Println(i)
}
```

---

#### **3. Function & Defer (4)**

- `func`
- `return`
- `defer`
- `go`

🔹 **Example (defer and go routine):**

```go
func main() {
defer fmt.Println("Executed last") // Runs at the end of function
go fmt.Println("Runs in goroutine") // Runs asynchronously
}
```

---

#### **4. Variable Declarations (6)**

- `var`
- `const`
- `type`
- `struct`
- `interface`
- `map`

🔹 **Example:**

```go
type Person struct {
Name string
Age  int
}

const Pi = 3.14
var x int = 42
```

---

#### **5. Concurrency (2)**

- `chan`
- `select`

🔹 **Example:**

```go
ch := make(chan int)

go func () {
ch <- 10
}()

fmt.Println(<-ch) // Receive value from channel
```

---

#### **6. Import & Packages (2)**

- `import`
- `package`

🔹 **Example:**

```go
package main

import "fmt"
```

---

#### **7. Miscellaneous (3)**

- `default`
- `type`
- `nil`

🔹 **Example:**

```go
var m map[string]int = nil // nil is the zero value for maps
```

---

### **🔥 Trick to Remember These Keywords**

1. **Control Flow** (🚦) → `if`, `else`, `switch`, `case`, `fallthrough`
2. **Looping** (🔁) → `for`, `range`, `break`
3. **Function & Execution** (⚙️) → `func`, `return`, `defer`, `go`
4. **Variables & Types** (📦) → `var`, `const`, `type`, `struct`, `interface`, `map`
5. **Concurrency** (🔀) → `chan`, `select`
6. **Imports & Packages** (📂) → `import`, `package`
7. **Miscellaneous** (🔹) → `default`, `type`, `nil`

---

### **🚀 Final List (Quick Recap)**

```
break       default       func          interface   select
case        defer         go            map         struct
chan        else         goto          package     switch
const       fallthrough  if            range       type
continue    for          import        return      var
```

---

### 5. What happens when a goroutine tries to:

- Send to a nil channel
- Send to a closed channel
- Receive on a closed channel
- Close a closed channel
-

### **Behavior of Goroutines with Channels in Go**

Understanding how channels behave in different situations is crucial for handling concurrency in Go. Let's break down
each scenario:

---

### **a. Send to a `nil` channel**

🚨 **Deadlock! The goroutine will block forever.**

#### **Example:**

```go
var ch chan int // ch is nil

go func () {
ch <- 10 // 🚨 Blocks forever (deadlock)
}()
```

**Why?**

- A `nil` channel has no underlying buffer or receiver.
- Any send operation will **block indefinitely** because there's no way for the value to be received.

---

### **b. Send to a closed channel**

🚨 **Panic: "send on closed channel"**

#### **Example:**

```go
ch := make(chan int)
close(ch) // Close the channel

go func () {
ch <- 10 // 🚨 Panic: send on closed channel
}()
```

**Why?**

- Once a channel is closed, no more values can be sent.
- Any attempt to send **causes a runtime panic**.

---

### **c. Receive on a closed channel**

✅ **Returns the zero value of the channel’s type (if empty), or the remaining buffered values.**

#### **Example:**

```go
ch := make(chan int, 2)
ch <- 5
ch <- 10
close(ch) // Close the channel

fmt.Println(<-ch) // ✅ 5 (first value)
fmt.Println(<-ch) // ✅ 10 (second value)
fmt.Println(<-ch) // ✅ 0 (zero value of int, since channel is empty)
```

**Why?**

- If the channel has buffered values, they can still be received.
- If the channel is empty, further reads return the **zero value** of the channel’s type.
- **No panic occurs when receiving from a closed channel!**

---

### **d. Close a closed channel**

🚨 **Panic: "close of closed channel"**

#### **Example:**

```go
ch := make(chan int)
close(ch)

close(ch) // 🚨 Panic: close of closed channel
```

**Why?**

- A channel **can only be closed once**.
- If you try to close it again, **Go panics**.

---

### **🔥 Quick Summary**

| **Action**                          | **Effect**                                      |
|-------------------------------------|-------------------------------------------------|
| **Send to a `nil` channel**         | 🚨 **Deadlock (blocks forever)**                |
| **Send to a closed channel**        | 🚨 **Panic: "send on closed channel"**          |
| **Receive from a closed channel**   | ✅ **Returns remaining values, then zero value** |
| **Close an already closed channel** | 🚨 **Panic: "close of closed channel"**         |

---

### 6. If you have a struct type X which embeds a struct type Y, and both X and Y have a method Foo, what happens when you call Foo on the outer struct? Which Foo is actually called, X's or Y's? If X embeds two structs Y and Z, and both Y and Z have a Foo, but X doesn't, what happens when you call Foo on X now? Which Foo is called, Y's or Z's?

### **Method Resolution in Embedded Structs in Go**

In Go, **method resolution follows a depth-first, promotion-based approach** when struct embedding is used. Let's
analyze both cases:

---

## **Case 1: X and Y both have a method `Foo`**

👉 **If `X` defines its own `Foo()`, it overrides `Y`'s `Foo()`.**

### **Example:**

```go
package main

import "fmt"

type Y struct{}

func (y Y) Foo() {
	fmt.Println("Y's Foo")
}

type X struct {
	Y
}

func (x X) Foo() {
	fmt.Println("X's Foo")
}

func main() {
	x := X{}
	x.Foo() // ✅ "X's Foo"
}
```

### **Explanation:**

- Even though `X` embeds `Y`, **`X`'s method takes priority**.
- The method in the outer struct (`X`) **overrides** the inner struct's (`Y`) method.

---

## **Case 2: X embeds Y and Z, both of which have Foo, but X itself does not**

👉 **If `X` does not define `Foo()`, Go checks `Y` and `Z`. If both have `Foo()`, a compilation error occurs due to
ambiguity.**

### **Example (Ambiguous Call - Compilation Error)**

```go
package main

import "fmt"

type Y struct{}

func (y Y) Foo() {
	fmt.Println("Y's Foo")
}

type Z struct{}

func (z Z) Foo() {
	fmt.Println("Z's Foo")
}

type X struct {
	Y
	Z
}

func main() {
	x := X{}
	x.Foo() // ❌ Compilation error: ambiguous selector x.Foo
}
```

### **Explanation:**

- `X` has **two embedded structs (`Y` and `Z`)**.
- Both `Y` and `Z` have a method `Foo()`, but **Go doesn’t know which one to call**.
- **Solution:** You must explicitly specify `x.Y.Foo()` or `x.Z.Foo()` to avoid ambiguity.

---

### **🔥 Quick Summary**

| **Scenario**                                                         | **Which `Foo` is called?**               |
|----------------------------------------------------------------------|------------------------------------------|
| `X` and `Y` both have `Foo()`                                        | ✅ `X.Foo()` (overrides `Y.Foo()`)        |
| `X` embeds `Y` and `Z`, both of which have `Foo()`, but `X` does not | ❌ **Compilation error (ambiguous call)** |

💡 **Rule of Thumb:**

- **Outer struct methods override inner ones** (if defined).
- **If multiple embedded structs have the same method, an explicit call is required (`x.Y.Foo()`).**

**🔥 Go Struct Embedding Cheat Sheet :**
Here’s a **Go Struct Embedding Cheat Sheet** to help you quickly recall method resolution rules! 🚀

---

## **🔥 Go Struct Embedding Cheat Sheet**

### **🔹 Rule #1: Outer Struct Methods Override Embedded Struct Methods**

- If the **outer struct (`X`) has a method**, it **overrides** methods of embedded structs.

```go
type Y struct{}

func (y Y) Foo() {
fmt.Println("Y's Foo")
}

type X struct {
Y
}

func (x X) Foo() { // Overrides Y's Foo
fmt.Println("X's Foo")
}

func main() {
x := X{}
x.Foo() // ✅ Output: "X's Foo"
}
```

---

### **🔹 Rule #2: Methods from Embedded Structs are Promoted**

- If `X` **doesn't** define `Foo()`, Go **searches in embedded structs**.

```go
type Y struct{}

func (y Y) Foo() {
fmt.Println("Y's Foo")
}

type X struct {
Y // X embeds Y
}

func main() {
x := X{}
x.Foo() // ✅ Output: "Y's Foo"
}
```

---

### **🔹 Rule #3: Multiple Embeds with the Same Method Cause Ambiguity**

- If `X` embeds both `Y` and `Z`, **and both have `Foo()`**, a **compilation error** occurs.

```go
type Y struct{}
func (y Y) Foo() { fmt.Println("Y's Foo") }

type Z struct{}
func (z Z) Foo() { fmt.Println("Z's Foo") }

type X struct {
Y
Z
}

func main() {
x := X{}
x.Foo() // ❌ Compilation error: ambiguous selector x.Foo
}
```

✅ **Fix:** Explicitly specify `x.Y.Foo()` or `x.Z.Foo()`.

```go
func main() {
x := X{}
x.Y.Foo() // ✅ Output: "Y's Foo"
x.Z.Foo() // ✅ Output: "Z's Foo"
}
```

---

### **🔹 Rule #4: Methods Can Still Be Overridden Per Instance**

- You can override an embedded struct’s method dynamically.

```go
type Y struct{}

func (y Y) Foo() { fmt.Println("Y's Foo") }

type X struct {
Y
Foo func () // Overrides method with function field
}

func main() {
x := X{Foo: func () { fmt.Println("Custom Foo") }}
x.Foo() // ✅ Output: "Custom Foo"
x.Y.Foo() // ✅ Output: "Y's Foo"
}
```

---

## **📌 Final Summary Table**

| **Scenario**                                                                | **Behavior**                                |
|-----------------------------------------------------------------------------|---------------------------------------------|
| **Outer struct (`X`) has method `Foo()`**                                   | ✅ `X.Foo()` is called (overrides `Y.Foo()`) |
| **Outer struct (`X`) doesn’t have `Foo()`, but embedded struct (`Y`) does** | ✅ `Y.Foo()` is promoted and used            |
| **Multiple embedded structs (`Y`, `Z`) have the same method (`Foo()`)**     | ❌ Compilation error (ambiguous call)        |
| **Explicitly calling `x.Y.Foo()` or `x.Z.Foo()`**                           | ✅ Works without ambiguity                   |

---

### 7. What is the largest value a numeric constant can have in Go? (Be careful.)

### **🔹 The Largest Value a Numeric Constant Can Have in Go**

👉 **Numeric constants in Go are arbitrary-precision values** (i.e., they can be **much larger than any built-in type
like `int64` or `float64`**). However, when a constant is assigned to a variable, it must fit within the type’s range.

---

### **🚀 Theoretical Maximum Value of a Numeric Constant**

- Since **untyped constants** in Go **use arbitrary precision**, **they are only limited by available memory**.
- This means you can define extremely large numbers **without an overflow** (as long as they remain untyped).

🔹 **Example of a Huge Constant (No Error)**

```go
package main

import "fmt"

const Huge = 1e1000 // ✅ Allowed (arbitrary precision)

func main() {
	fmt.Println(Huge) // Prints a very large number
}
```

**💡 Key Insight:**

- `Huge` is an **untyped constant**, so Go doesn’t enforce any bit limits.

---

### **🔹 What Happens When Assigned to a Variable?**

- Once assigned to a **typed variable**, the value must fit within that type’s range.
- If the value exceeds the range, **Go throws a compilation error**.

🔹 **Example (Overflow Error)**

```go
const Huge = 1e1000
var x float64 = Huge // ❌ Compilation error: constant 1e+1000 overflows float64
```

- **`float64` has a max value of ≈ `1.8 × 10^308`**.
- `1e1000` exceeds this, so **Go throws an error**.

---

### **🔹 What About Integer Constants?**

- Integer constants **are also arbitrary-precision**, but when assigned to a **typed variable**, they must fit.

🔹 **Example (No Error - Arbitrary Precision)**

```go
const BigInt = 1 << 1000 // ✅ Allowed (untyped constant)
```

🔹 **Example (Overflow Error)**

```go
var i int64 = 1 << 100 // ❌ Compilation error: overflows int64
```

- `int64` max value: **`9,223,372,036,854,775,807` (≈ 2^63 - 1)**.
- `1 << 100` **exceeds this limit** → causes **overflow error**.

---

### **🔥 Final Answer:**

| **Type**             | **Max Value**                               |
|----------------------|---------------------------------------------|
| **Untyped Constant** | 🚀 **No fixed limit** (arbitrary precision) |
| **`int64`**          | `9,223,372,036,854,775,807` (`2^63 - 1`)    |
| **`uint64`**         | `18,446,744,073,709,551,615` (`2^64 - 1`)   |
| **`float64`**        | ≈ `1.8 × 10^308`                            |

**💡 Trick to Remember:**

- **Untyped constants** are **only limited by memory**.
- **Typed variables** must fit within their bit-size limits.

### 8. Can you explain the difference between a mock, a stub and a fake?

### **🛠 Mock vs Stub vs Fake in Testing (Go Edition)**

When writing **unit tests in Go**, we often need **test doubles** to replace real dependencies. The three most common
types are:

| **Type** | **What It Does**                                                                                  | **When to Use It?**                                                                             |
|----------|---------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------|
| **Stub** | Returns **predefined data**                                                                       | When you **only care about inputs/outputs** and not behavior.                                   |
| **Mock** | **Verifies interactions** (i.e., checks if a function was called, how many times, with what args) | When you need to **assert behavior** (e.g., did a function send an email?)                      |
| **Fake** | Implements a **simplified but real version** of a dependency                                      | When a **real implementation** is too slow/complex (e.g., an in-memory DB instead of Postgres). |

---

### **1️⃣ Stub → Predefined Data**

👉 **Stub replaces a dependency and returns hardcoded data** when called.

#### **Example (Stub in Go)**

```go
type UserRepositoryStub struct{}

func (u UserRepositoryStub) GetUser(id int) string {
return "John Doe" // Always returns the same data
}

func TestGetUser(t *testing.T) {
repo := UserRepositoryStub{}
user := repo.GetUser(1)

if user != "John Doe" {
t.Errorf("Expected John Doe, got %s", user)
}
}
```

**💡 When to Use:**

- When testing **functions that depend on external data** but don’t care how the data is fetched.
- **Example:** Stubbing an API response.

---

### **2️⃣ Mock → Checks Behavior (Calls & Arguments)**

👉 **Mocks don’t just return values; they verify interactions.**

- Useful when you need to check **if a function was called with the correct arguments**.

#### **Example (Mock in Go using testify/mock)**

```go
import (
"github.com/stretchr/testify/mock"
"testing"
)

type EmailServiceMock struct {
mock.Mock
}

func (e *EmailServiceMock) SendEmail(to, body string) error {
args := e.Called(to, body)
return args.Error(0)
}

func TestSendEmail(t *testing.T) {
emailMock := new(EmailServiceMock)
emailMock.On("SendEmail", "test@example.com", "Hello").Return(nil)

err := emailMock.SendEmail("test@example.com", "Hello")

emailMock.AssertCalled(t, "SendEmail", "test@example.com", "Hello") // ✅ Verifies the call
emailMock.AssertExpectations(t) // ✅ Ensures all expectations were met

if err != nil {
t.Errorf("Expected nil error, got %v", err)
}
}
```

**💡 When to Use:**

- When you want to **verify method calls** (e.g., "Was `SendEmail` called with the right parameters?").
- **Example:** Mocking a payment gateway call.

---

### **3️⃣ Fake → Lightweight Real Implementation**

👉 **A Fake is an actual working implementation, just simpler.**

#### **Example (Fake Database in Go)**

```go
type FakeUserRepo struct {
data map[int]string
}

func NewFakeUserRepo() *FakeUserRepo {
return &FakeUserRepo{data: make(map[int]string)}
}

func (f *FakeUserRepo) GetUser(id int) string {
return f.data[id]
}

func (f *FakeUserRepo) AddUser(id int, name string) {
f.data[id] = name
}

func TestFakeRepo(t *testing.T) {
repo := NewFakeUserRepo()
repo.AddUser(1, "Alice")

user := repo.GetUser(1)
if user != "Alice" {
t.Errorf("Expected Alice, got %s", user)
}
}
```

**💡 When to Use:**

- When a **real implementation is too slow** (e.g., real DB calls).
- **Example:** Using an **in-memory database** instead of a real SQL database.

---

### **🚀 Quick Recap**

| **Feature**              | **Stub**           | **Mock**                | **Fake**                    |
|--------------------------|--------------------|-------------------------|-----------------------------|
| Returns predefined data? | ✅ Yes              | ❌ No                    | ✅ Yes                       |
| Verifies method calls?   | ❌ No               | ✅ Yes                   | ❌ No                        |
| Implements real logic?   | ❌ No               | ❌ No                    | ✅ Yes (but simplified)      |
| Best for...              | Input/output tests | Checking function calls | Replacing slow dependencies |

---

### **🔥 TL;DR**

- **Stub** → Hardcoded response (like a canned API response).
- **Mock** → Checks function calls & arguments (used for behavior verification).
- **Fake** → A simple, working implementation (like an in-memory DB).

### 9. What are the benefits and drawbacks of modeling a set in Go via map[T]bool vs. map[T]struct{}?

### **🏆 Modeling a Set in Go: `map[T]bool` vs. `map[T]struct{}`**

Since Go **does not have a built-in `set` type**, we typically use **maps** (`map[T]bool` or `map[T]struct{}`) to
implement sets.

---

## **🔹 Option 1: `map[T]bool` (Using Boolean Flags)**

```go
set := make(map[string]bool)
set["apple"] = true
set["banana"] = true

if set["apple"] {
fmt.Println("Apple exists in the set")
}
```

### ✅ **Pros:**

1. **Readable & Explicit** → `set["apple"] == true` clearly conveys intent.
2. **Allows additional logic** → You can use `true` or `false` to indicate presence/state.

### ❌ **Cons:**

1. **Extra memory usage** → Each entry stores a `bool` (`1 byte` per element).
2. **Unnecessary value storage** → We don’t really need `true/false`, just existence.

---

## **🔹 Option 2: `map[T]struct{}` (Zero-Sized Struct)**

```go
set := make(map[string]struct{})
set["apple"] = struct{}{}
set["banana"] = struct{}{}

if _, exists := set["apple"]; exists {
fmt.Println("Apple exists in the set")
}
```

### ✅ **Pros:**

1. **Memory efficient** → `struct{}` takes **zero bytes** in memory, unlike `bool`.
2. **Slightly faster lookup** → Avoids loading an extra `bool` from memory.

### ❌ **Cons:**

1. **Less readable** → `if _, exists := set["apple"]; exists` is slightly more verbose.
2. **No extra metadata** → You cannot store `true/false` to track additional state.

---

## **🚀 Which One Should You Use?**

| **Criteria**              | **`map[T]bool`**                         | **`map[T]struct{}`**                         |
|---------------------------|------------------------------------------|----------------------------------------------|
| **Memory Efficiency**     | ❌ Uses extra `1 byte` per entry          | ✅ Uses **0 bytes** per entry                 |
| **Performance**           | 🔹 Slightly slower (extra `bool` lookup) | 🔥 Faster (`struct{}` is zero-sized)         |
| **Readability**           | ✅ More intuitive                         | ❌ Slightly verbose (`_, exists := set[...]`) |
| **Supports extra state?** | ✅ Yes (can use `true/false`)             | ❌ No                                         |

### **📌 Final Recommendation:**

- Use **`map[T]struct{}`** if you only need **existence checks** (best for large sets).
- Use **`map[T]bool`** if you need to store additional state (e.g., marking items as "visited" vs. "not visited").

### 10. How to handle errors in golang

### **🚀 How to Handle Errors in Go (Best Practices & Examples)**

Error handling in Go is **explicit**, meaning you check and handle errors manually rather than relying on exceptions.

---

## **1️⃣ Basic Error Handling (`error` interface)**

👉 **Every function that can fail should return an `error` as the last return value.**

🔹 **Example:**

```go
package main

import (
	"errors"
	"fmt"
)

// Function that returns an error
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err) // ✅ Explicit error handling
		return
	}
	fmt.Println("Result:", result)
}
```

**🔹 Key Takeaways:**

- **Always check `err != nil`** before using the result.
- Use **`errors.New("message")`** to create simple errors.

---

## **2️⃣ Using `fmt.Errorf` for Formatted Errors**

👉 **You can format errors with additional context using `fmt.Errorf()`.**

🔹 **Example:**

```go
import "fmt"

func openFile(filename string) error {
return fmt.Errorf("failed to open file %s: %w", filename, errors.New("file not found"))
}
```

**💡 `%w` allows error wrapping** (see next section).

---

## **3️⃣ Error Wrapping (`errors.Is` & `errors.As`)**

👉 **Since Go 1.13, use `errors.Is` & `errors.As` for better error checking.**

🔹 **Example:**

```go
import (
"errors"
"fmt"
"os"
)

var ErrFileNotFound = errors.New("file not found")

func readFile(filename string) error {
return fmt.Errorf("read error: %w", ErrFileNotFound)
}

func main() {
err := readFile("data.txt")
if errors.Is(err, ErrFileNotFound) {
fmt.Println("Handle missing file case.")
}
}
```

**🔹 Why Use `errors.Is`?**

- Checks if `err` wraps `ErrFileNotFound` **without comparing strings**.

---

## **4️⃣ Custom Error Types (Implementing `Error()` Method)**

👉 **Use a custom struct when errors need extra context.**

🔹 **Example:**

```go
type MyError struct {
Code    int
Message string
}

func (e *MyError) Error() string {
return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func doSomething() error {
return &MyError{Code: 404, Message: "Not Found"}
}

func main() {
err := doSomething()
if e, ok := err.(*MyError); ok {
fmt.Println("Custom Error:", e.Code, e.Message)
}
}
```

**🔹 Why Use Custom Errors?**

- **Encapsulates structured error data** (like HTTP status codes).

---

## **5️⃣ Panic vs. Error Handling**

👉 **Use `panic` for truly unrecoverable situations** (not normal errors).

🔹 **Example (Bad Approach)** ❌

```go
func divide(a, b int) int {
if b == 0 {
panic("division by zero") // ❌ Don't panic for expected errors
}
return a / b
}
```

🔹 **Better Approach (Returning an `error`)** ✅

```go
func divide(a, b int) (int, error) {
if b == 0 {
return 0, errors.New("division by zero")
}
return a / b, nil
}
```

**🔹 When to Use `panic`?**

- When **program cannot continue safely** (e.g., corrupted memory, invalid assumptions).

---

## **6️⃣ Recovering from `panic` (Using `defer` & `recover`)**

👉 **Use `recover()` inside `defer` to catch `panic`.**

🔹 **Example:**

```go
func safeFunction() {
defer func () {
if r := recover(); r != nil {
fmt.Println("Recovered from panic:", r)
}
}()
panic("Something went wrong!") // ❌ Will be caught
}

func main() {
safeFunction()
fmt.Println("Program continues...")
}
```

**🔹 Key Takeaways:**

- **Use `recover()` only in critical sections.**
- **Don’t recover everything; let normal errors propagate.**

---

## **7️⃣ Best Practices for Error Handling**

✔ **Return `error` instead of panicking.**  
✔ **Use `errors.Is` & `errors.As` for structured error checking.**  
✔ **Use `fmt.Errorf("%w", err)` to wrap errors.**  
✔ **Avoid `panic` in business logic (only for fatal conditions).**  
✔ **Log errors at the boundary (not deep inside functions).**

---

## **🚀 TL;DR (Cheat Sheet)**

| **Scenario**                | **Solution**                                          |
|-----------------------------|-------------------------------------------------------|
| **Simple Error**            | `return errors.New("message")`                        |
| **Formatted Error**         | `fmt.Errorf("something went wrong: %w", err)`         |
| **Check Specific Error**    | `errors.Is(err, ErrType)`                             |
| **Extract More Error Info** | `errors.As(err, &target)`                             |
| **Custom Error Type**       | Define `type MyError struct {}` & implement `Error()` |
| **Avoiding `panic`**        | Use `error` return values instead                     |
| **Handling `panic` safely** | Use `defer` + `recover()`                             |

---

### 11. How to design the sockets in golang.

### **🚀 Designing Sockets in Golang (TCP & UDP)**

Sockets allow network communication between processes (either on the same machine or across the internet). In Go, we use
the **`net`** package to work with **TCP & UDP sockets**.

---

## **1️⃣ TCP Sockets (Connection-Oriented)**

TCP (Transmission Control Protocol) is **reliable** and **ensures ordered delivery** of data.

### **🔹 TCP Server (Listening for Clients)**

```go
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Start a TCP server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()
	fmt.Println("TCP Server started on port 8080")

	for {
		conn, err := listener.Accept() // Accept client connections
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn) // Handle client in a separate goroutine
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected:", conn.RemoteAddr())
			break
		}
		fmt.Print("Received:", message)
		conn.Write([]byte("Message received\n")) // Send response
	}
}
```

### **🔹 TCP Client (Connecting to Server)**

```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // Connect to TCP server
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server. Type messages:")
	for {
		reader := bufio.NewReader(os.Stdin)
		message, _ := reader.ReadString('\n')

		_, err := conn.Write([]byte(message)) // Send message to server
		if err != nil {
			fmt.Println("Write error:", err)
			return
		}

		response, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Server:", response)
	}
}
```

**🔹 How It Works:**  
✅ The **server** listens on port **8080**.  
✅ Clients **connect**, send a message, and the **server responds**.  
✅ **Multiple clients** can connect since we use **goroutines**.

---

## **2️⃣ UDP Sockets (Connectionless)**

UDP (User Datagram Protocol) is **faster but unreliable** (no guarantee of delivery or order).

### **🔹 UDP Server (Receiving Messages)**

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	addr := net.UDPAddr{
		Port: 8081,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer conn.Close()
	fmt.Println("UDP Server started on port 8081")

	buffer := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer) // Read data
		if err != nil {
			fmt.Println("Read error:", err)
			continue
		}
		fmt.Printf("Received %s from %s\n", string(buffer[:n]), remoteAddr)
		conn.WriteToUDP([]byte("Message received"), remoteAddr) // Send response
	}
}
```

### **🔹 UDP Client (Sending Messages)**

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	serverAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8081")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("Connection failed:", err)
		return
	}
	defer conn.Close()

	message := []byte("Hello UDP Server")
	_, err = conn.Write(message) // Send UDP message
	if err != nil {
		fmt.Println("Write error:", err)
		return
	}

	buffer := make([]byte, 1024)
	n, _, err := conn.ReadFromUDP(buffer) // Read response
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}
	fmt.Println("Server:", string(buffer[:n]))
}
```

**🔹 How It Works:**  
✅ **No connection setup** → Just send and receive messages.  
✅ **Faster** than TCP but **no delivery guarantee**.  
✅ **Good for real-time applications** (e.g., live streaming, online gaming).

---

## **3️⃣ WebSockets (For Persistent Connections)**

WebSockets allow **full-duplex communication** between the client & server.

### **🔹 WebSocket Server (Using `github.com/gorilla/websocket`)**

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil) // Upgrade HTTP to WebSocket
	if err != nil {
		fmt.Println("Upgrade failed:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected:", conn.RemoteAddr())

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Client disconnected")
			break
		}
		fmt.Println("Received:", string(message))

		conn.WriteMessage(messageType, []byte("Message received"))
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	fmt.Println("WebSocket server started on :8082")
	http.ListenAndServe(":8082", nil)
}
```

### **🔹 WebSocket Client (Using `github.com/gorilla/websocket`)**

```go
package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

func main() {
	conn, _, err := websocket.DefaultDialer.Dial("ws://localhost:8082/ws", nil)
	if err != nil {
		log.Fatal("Connection failed:", err)
	}
	defer conn.Close()

	message := "Hello WebSocket Server"
	conn.WriteMessage(websocket.TextMessage, []byte(message))

	_, response, _ := conn.ReadMessage()
	fmt.Println("Server:", string(response))
}
```

**🔹 Why WebSockets?**

- ✅ **Persistent bidirectional communication**.
- ✅ **Efficient** for **chat apps, notifications, & real-time updates**.
- ❌ **More complex** than TCP/UDP.

---

## **🚀 Choosing the Right Socket Type**

| **Use Case**                                            | **Recommended Socket**                   |
|---------------------------------------------------------|------------------------------------------|
| Reliable file transfer, API, general client-server apps | **TCP (`net.Listen`, `net.Dial`)**       |
| Fast, real-time apps (gaming, video streaming, IoT)     | **UDP (`net.ListenUDP`, `net.DialUDP`)** |
| Web-based real-time chat, notifications, live data      | **WebSockets (`gorilla/websocket`)**     |

---

### 12. How to handle the DB load in golang.

| **Technique**                                       | **Benefit**                               |
|-----------------------------------------------------|-------------------------------------------|
| ✅ **Use Connection Pooling (`sql.DB`)**             | Reuses connections, reduces DB overhead   |
| ✅ **Use Read Replicas**                             | Offload read queries from primary DB      |
| ✅ **Implement Caching (Redis, Memcached)**          | Reduces redundant DB queries              |
| ✅ **Use Bulk Inserts/Updates**                      | Faster writes & fewer transactions        |
| ✅ **Optimize Queries (Indexes, `EXPLAIN ANALYZE`)** | Avoids slow queries                       |
| ✅ **Use Goroutines for Concurrent Queries**         | Improves parallel query execution         |
| ✅ **Use Message Queues (Kafka, RabbitMQ)**          | Offloads heavy writes to async processing |

---

### 13. How to share resources in golang.

### **🚀 Sharing Resources in Golang (Concurrency-Safe Approaches)**

When multiple Goroutines need to **share resources (memory, files, DB connections, etc.)**, we must ensure **safe access
** to avoid **race conditions** and **data corruption**. Here are the best ways to share resources safely in Golang:

---

## **1️⃣ Use Mutex for Safe Access to Shared Data**

👉 **Use `sync.Mutex` when only one Goroutine should access a resource at a time.**

🔹 **Example: Protecting a Counter with Mutex**

```go
package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex // Mutex for safe access

func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()   // Lock before modifying shared resource
	counter++   // Critical section
	mu.Unlock() // Unlock after modification
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increment(&wg)
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter)
}
```

✅ **Why?**

- Prevents **race conditions** when multiple Goroutines modify `counter`.
- Ensures **only one Goroutine accesses the critical section** at a time.

---

## **2️⃣ Use RWMutex for Read-Mostly Data**

👉 **Use `sync.RWMutex` when many Goroutines read but few write.**

🔹 **Example: Allowing Multiple Readers, One Writer**

```go
var rwMu sync.RWMutex
var sharedData string

func readData(wg *sync.WaitGroup) {
defer wg.Done()
rwMu.RLock() // Allow multiple readers
fmt.Println("Reading:", sharedData)
rwMu.RUnlock()
}

func writeData(wg *sync.WaitGroup, data string) {
defer wg.Done()
rwMu.Lock() // Only one writer allowed
sharedData = data
fmt.Println("Written:", data)
rwMu.Unlock()
}
```

✅ **Why?**

- **Multiple readers** can access data **simultaneously**.
- **Only one writer** at a time → prevents corruption.

---

## **3️⃣ Use Channels for Safe Goroutine Communication**

👉 **Instead of sharing memory, use channels to pass data between Goroutines.**

🔹 **Example: Using Channels Instead of Locks**

```go
package main

import "fmt"

func main() {
	data := make(chan int)

	go func() {
		data <- 42 // Send data
	}()

	value := <-data // Receive data
	fmt.Println("Received:", value)
}
```

✅ **Why?**

- **No need for locks** (prevents deadlocks).
- **Encourages "Share by Communication"** instead of "Sharing Memory".

---

## **4️⃣ Use `sync.Once` for Resource Initialization (Singleton Pattern)**

👉 **Ensure a resource is initialized only once (e.g., database connection).**

🔹 **Example: Creating a Singleton Database Connection**

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var db *sql.DB
var once sync.Once

func getDBInstance() *sql.DB {
	once.Do(func() { // Ensures this block runs only ONCE
		var err error
		db, err = sql.Open("postgres", "postgres://user:pass@localhost/dbname?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Database connection created")
	})
	return db
}

func main() {
	conn1 := getDBInstance()
	conn2 := getDBInstance()
	fmt.Println(conn1 == conn2) // true (same instance)
}
```

✅ **Why?**

- Ensures **only one DB connection** is created.
- Prevents **race conditions in initialization**.

---

## **5️⃣ Use `sync.Pool` for Object Reuse (Reduce GC Pressure)**

👉 **Use `sync.Pool` to reuse objects and reduce memory allocation overhead.**

🔹 **Example: Reusing Structs with sync.Pool**

```go
package main

import (
	"fmt"
	"sync"
)

type User struct {
	Name string
	Age  int
}

var userPool = sync.Pool{
	New: func() interface{} {
		return &User{} // Create a new user if pool is empty
	},
}

func main() {
	u1 := userPool.Get().(*User) // Get from pool
	u1.Name = "Alice"
	u1.Age = 25

	fmt.Println("User:", u1)

	userPool.Put(u1) // Return to pool

	u2 := userPool.Get().(*User)    // Reuse object
	fmt.Println("Reused User:", u2) // u2 has same memory as u1
}
```

✅ **Why?**

- **Avoids frequent memory allocations** (reduces GC overhead).
- Useful for **temporary objects** (e.g., JSON buffers, DB connections).

---

## **6️⃣ Use Atomic Variables for Performance-Optimized Counters**

👉 **For simple shared counters, `sync/atomic` is faster than Mutex.**

🔹 **Example: Atomic Counter**

```go
package main

import (
	"fmt"
	"sync/atomic"
)

var counter int64

func increment() {
	atomic.AddInt64(&counter, 1) // Thread-safe increment
}

func main() {
	for i := 0; i < 10; i++ {
		go increment()
	}
	fmt.Println("Final Counter:", counter)
}
```

✅ **Why?**

- **Faster than Mutex** for simple counters.
- **Thread-safe** without locks.

---

## **🚀 Summary: Best Practices for Resource Sharing**

| **Method**                   | **Use Case**                   | **Pros**                           | **Cons**                          |
|------------------------------|--------------------------------|------------------------------------|-----------------------------------|
| **Mutex (`sync.Mutex`)**     | Protect shared variables       | Simple, prevents race conditions   | Blocks Goroutines                 |
| **RWMutex (`sync.RWMutex`)** | Many reads, few writes         | Allows multiple readers            | Writers must wait                 |
| **Channels (`chan`)**        | Goroutine communication        | No locks needed, structured design | Requires channel management       |
| **`sync.Once`**              | Singleton initialization       | Ensures one-time setup             | Limited use cases                 |
| **`sync.Pool`**              | Object reuse                   | Reduces memory allocations         | Only useful for temporary objects |
| **Atomic (`sync/atomic`)**   | Performance-optimized counters | Fastest for simple counters        | Only works for primitive types    |

---

## **🚀 Choosing the Right Approach**

✔ **Need to modify a shared variable?** → Use **Mutex**  
✔ **Mostly reading, rarely writing?** → Use **RWMutex**  
✔ **Want Goroutines to communicate?** → Use **Channels**  
✔ **Need a Singleton (e.g., DB connection)?** → Use **sync.Once**  
✔ **Reusing temporary objects?** → Use **sync.Pool**  
✔ **Simple atomic counters?** → Use **sync/atomic**


---

### 14. similarities and dissimilarities in java and golang

### **🚀 Java vs. Golang: Similarities & Differences (Quick Notes)**

---

### **✅ Similarities Between Java & Golang**

1️⃣ **Both are compiled languages** → Java compiles to **bytecode (JVM)**, Go compiles to **native binary**.  
2️⃣ **Both are strongly typed** → No implicit type conversions.  
3️⃣ **Both support concurrency** → Java uses **Threads**, Go uses **Goroutines** (lighter).  
4️⃣ **Garbage Collection (GC)** → Both manage memory automatically.  
5️⃣ **Cross-Platform** → Java runs on **JVM**, Go compiles for multiple platforms.  
6️⃣ **Support for Networking** → Java has **Netty**, Go has **built-in HTTP & net packages**.  
7️⃣ **Object-Oriented Features** → Java uses **classes**, Go uses **structs with methods**.  
8️⃣ **Standard Libraries** → Both have rich standard libraries for IO, HTTP, and networking.

---

### **❌ Differences Between Java & Golang**

| Feature                   | Java 🟡                            | Golang 🔵                                  |
|---------------------------|------------------------------------|--------------------------------------------|
| **Compilation**           | Bytecode (JVM)                     | Native binary                              |
| **Performance**           | Slower (JVM overhead)              | Faster (No VM)                             |
| **Concurrency**           | Threads (Expensive)                | Goroutines (Lightweight)                   |
| **Error Handling**        | Exceptions (try-catch)             | Error values (`if err != nil`)             |
| **Generics**              | Available (since Java 5)           | Introduced in Go 1.18                      |
| **Memory Management**     | GC + Manual (`finalize`)           | GC, but no manual finalization             |
| **Inheritance**           | Supports OOP (extends, implements) | No classical inheritance, uses composition |
| **Interfaces**            | Explicit implementation            | Implicit implementation                    |
| **Dependency Management** | Maven, Gradle                      | Go Modules (go.mod)                        |
| **Syntax Complexity**     | Verbose (class, interface)         | Simpler, lightweight syntax                |
| **Development Speed**     | Slower (boilerplate)               | Faster (concise code)                      |
| **Package System**        | `package.class.method`             | `package.method`                           |
| **Compilation Time**      | Slower (JVM bytecode)              | Faster (direct native compilation)         |
| **Binary Size**           | Small (JVM required)               | Larger (static linking)                    |

---

### **🚀 When to Use Java vs. Golang?**

✔ **Use Java when:**

- Enterprise applications (Spring, Hibernate).
- You need **JVM ecosystem** (Scala, Kotlin).
- Large-scale applications with **complex OOP needs**.

✔ **Use Golang when:**

- High-performance systems (networking, databases).
- **Microservices & Cloud applications** (Docker, Kubernetes).
- When you need **fast compilation & deployment**.

### 16. diff between method and function in golang

### **🚀 Difference Between Method & Function in Golang**

In Golang, **functions and methods** both perform actions, but **methods are functions with a receiver** (associated
with a type).

---

### **✅ Function in Golang**

- **A standalone block of code** that can be called independently.
- **Not tied to any struct or type.**

🔹 **Example: A Regular Function**

```go
package main

import "fmt"

// Function without a receiver
func Add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(Add(5, 3)) // Output: 8
}
```

---

### **✅ Method in Golang**

- **A function that is associated with a specific type (struct, int, etc.).**
- Uses a **receiver** to bind the function to a type.
- Can be called using an **instance of the type** (`object.Method()`).

🔹 **Example: A Method with a Struct Receiver**

```go
package main

import "fmt"

// Define a struct
type Rectangle struct {
	Width, Height int
}

// Method with receiver of type Rectangle
func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	fmt.Println(rect.Area()) // Output: 50
}
```

---

### **🔄 Function vs. Method (Comparison Table)**

| Feature                            | Function                   | Method                         |
|------------------------------------|----------------------------|--------------------------------|
| **Definition**                     | Defined without a receiver | Defined with a receiver        |
| **Tied to Type?**                  | ❌ No, standalone           | ✅ Yes, bound to a struct/type  |
| **Calling Syntax**                 | `Add(5, 3)`                | `rect.Area()`                  |
| **Receives Data?**                 | Uses parameters            | Uses **receiver + parameters** |
| **Can be used on built-in types?** | ✅ Yes                      | ✅ Yes (with type alias)        |

---

### **✅ Can Methods Work with Built-in Types?**

Yes! We can define methods on **custom type aliases of built-in types.**

🔹 **Example: Method on an Integer Type Alias**

```go
package main

import "fmt"

// Define a custom type alias for int
type MyInt int

// Define a method on MyInt
func (m MyInt) Double() MyInt {
	return m * 2
}

func main() {
	num := MyInt(5)
	fmt.Println(num.Double()) // Output: 10
}
```

---

### **🚀 Key Takeaways**

✔ **Functions** → Standalone, general-purpose.  
✔ **Methods** → Bound to a type, allow OOP-like behavior.  
✔ **Methods on built-in types?** → Possible with **type aliases**.

---

### 17. How to achive polymorphism in golang

### **🚀 Achieving Polymorphism in Golang**

In Go, **polymorphism** is achieved through **interfaces** because Go **does not support traditional inheritance** like
Java or C++.

---

### **✅ What is Polymorphism?**

Polymorphism allows different types to be **treated uniformly** based on a shared behavior (methods).

In Golang, any type that **implements an interface** can be used interchangeably, enabling polymorphism.

---

## **🎯 Polymorphism with Interfaces (Example)**

### **1️⃣ Define an Interface (Common Behavior)**

```go
package main

import "fmt"

// Define an interface
type Speaker interface {
	Speak() string
}

// Define two structs implementing Speaker
type Dog struct{}
type Cat struct{}

// Implement Speak() for Dog
func (d Dog) Speak() string {
	return "Woof!"
}

// Implement Speak() for Cat
func (c Cat) Speak() string {
	return "Meow!"
}

// Polymorphic function
func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {
	d := Dog{}
	c := Cat{}

	MakeSound(d) // Output: Woof!
	MakeSound(c) // Output: Meow!
}
```

✔ **Dog and Cat both implement `Speak()`, so they satisfy `Speaker`.**  
✔ The function `MakeSound(s Speaker)` **accepts any type that implements `Speaker`.**

---

### **2️⃣ Polymorphism with Empty Interface (`interface{}`)**

If we don't know the exact type, we can use **`interface{}` (empty interface)** to handle **any type**.

```go
package main

import "fmt"

// Function accepting any type (empty interface)
func PrintValue(val interface{}) {
	fmt.Println(val)
}

func main() {
	PrintValue(42)       // Output: 42
	PrintValue("Golang") // Output: Golang
	PrintValue(3.14)     // Output: 3.14
}
```

✔ **`interface{}` is the universal type that can hold any value.**  
✔ We can use **type assertions** or **type switches** to determine the actual type inside the function.

---

### **3️⃣ Polymorphism via Struct Embedding**

Go does not support **class inheritance**, but it allows **composition (struct embedding)** for reuse.

```go
package main

import "fmt"

// Base struct
type Animal struct{}

// Method Speak for Animal
func (a Animal) Speak() string {
	return "Some sound"
}

// Derived struct embedding Animal
type Dog struct {
	Animal
}

// Overriding Speak for Dog
func (d Dog) Speak() string {
	return "Bark!"
}

func main() {
	d := Dog{}
	fmt.Println(d.Speak()) // Output: Bark!
}
```

✔ `Dog` **inherits** `Animal` methods but overrides `Speak()`.

---

### **🚀 Summary: How Go Achieves Polymorphism**

| Approach                            | Description                                                         | Example                            |
|-------------------------------------|---------------------------------------------------------------------|------------------------------------|
| **Interfaces**                      | Structs implementing the same interface can be used interchangeably | `Speaker` interface (`Dog`, `Cat`) |
| **Empty Interface (`interface{}`)** | Can accept **any** type                                             | `PrintValue(interface{})`          |
| **Struct Embedding**                | Allows **composition** and method overriding                        | `Dog` embeds `Animal`              |

---

### **🔥 Key Takeaways**

✔ **Go does not have classical OOP inheritance.**  
✔ **Interfaces enable polymorphism.**  
✔ **Struct embedding allows method reuse & overriding.**  
✔ **Empty interfaces (`interface{}`) can hold any type.**

---

### 17. How to determine no of go routines incase of worker pool design.

### **🚀 How to Determine the Number of Goroutines in a Worker Pool?**

Choosing the right number of Goroutines in a **worker pool** depends on **CPU-bound vs. I/O-bound tasks**.

---

### **1️⃣ CPU-Bound Workloads (Compute-Intensive Tasks)**

- When **each Goroutine does heavy computation**, the number of Goroutines should match the number of CPU cores.
- **Formula:**  
  \[
  \text{NumGoroutines} = \text{NumCPU}
  \]
- **Example:**

```go
package main

import (
	"fmt"
	"runtime"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * job // Heavy computation
	}
}

func main() {
	numWorkers := runtime.NumCPU() // Get number of CPU cores
	fmt.Println("Using", numWorkers, "workers")

	jobs := make(chan int, 10)
	results := make(chan int, 10)

	for i := 0; i < numWorkers; i++ {
		go worker(i, jobs, results)
	}

	// Send jobs
	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 0; r < 10; r++ {
		fmt.Println(<-results)
	}
}
```

✔ Uses **`runtime.NumCPU()`** to get the **optimal number of workers**.  
✔ This prevents **excessive context switching** on CPU-heavy tasks.

---

### **2️⃣ I/O-Bound Workloads (Network, Database, File I/O)**

- When tasks involve **waiting for external responses**, Goroutines can **exceed CPU cores**.
- **Formula:**  
  \[
  \text{NumGoroutines} = \text{NumCPU} \times k
  \]
  Where **k** is an experimentally determined factor (e.g., 2-10).
- **Example:**

```go
package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func worker(id int, jobs <-chan string) {
	for url := range jobs {
		resp, _ := http.Get(url) // Simulating an I/O-bound task
		fmt.Println("Worker", id, "fetched", url, "Status:", resp.Status)
	}
}

func main() {
	numWorkers := runtime.NumCPU() * 4 // More goroutines for I/O tasks
	jobs := make(chan string, 10)

	for i := 0; i < numWorkers; i++ {
		go worker(i, jobs)
	}

	urls := []string{"https://golang.org", "https://google.com", "https://github.com"}
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)
}
```

✔ Uses **`NumCPU() * 4`** to scale up workers for high-latency tasks.  
✔ Efficiently handles multiple network requests **without blocking CPU**.

---

### **3️⃣ General Formula for Worker Pool Size**

#### **🔹 If CPU-bound** (math, encoding, compression):

\[
\text{NumGoroutines} = \text{NumCPU}
\]

#### **🔹 If I/O-bound** (DB queries, HTTP requests, file I/O):

\[
\text{NumGoroutines} = \text{NumCPU} \times k
\]
Where **k is between 2 to 10**, depending on workload.

---

### **4️⃣ Experimentation (Benchmarking)**

To find the optimal number of workers:

1. **Start with `NumCPU()` and increase incrementally.**
2. **Measure latency, throughput, and CPU usage.**
3. **Adjust worker count based on real-world performance.**

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for workers := 1; workers <= 16; workers *= 2 {
		start := time.Now()
		// Run workload with 'workers' number of goroutines
		fmt.Println("Workers:", workers, "Time:", time.Since(start))
	}
}
```

✔ Helps **determine the best `k` factor** dynamically.

---

### **🚀 Key Takeaways**

| Workload Type                      | Worker Pool Size                  |
|------------------------------------|-----------------------------------|
| **CPU-Bound (Heavy Processing)**   | `NumCPU()`                        |
| **I/O-Bound (DB, Network, Files)** | `NumCPU() * k` (k=2 to 10)        |
| **Experimentation**                | Start small, benchmark, and scale |

 
---

### 18. what is meta programming

### **🚀 What is Meta-Programming?**

**Meta-programming** is writing code that **generates, modifies, or analyzes other code at runtime or compile time**.

👉 In simple terms, **"code that writes code."**

---

### **🔹 Meta-Programming in Go**

Unlike dynamic languages like Python or Ruby, Go is **statically typed** and **does not support reflection-based code
generation at compile time**.

However, Go supports **meta-programming** through:  
1️⃣ **Reflection (`reflect` package)** – Inspect and modify structs, fields, and methods at runtime.  
2️⃣ **Code Generation (`go generate`)** – Auto-generate boilerplate code before compilation.  
3️⃣ **Templates (`text/template`)** – Dynamically generate code or configuration files.  
4️⃣ **Generics (Go 1.18+)** – Reuse logic across multiple types.

---

## **1️⃣ Reflection (`reflect` Package)**

Reflection allows a program to inspect its **own structure** and dynamically interact with variables, structs, and
methods.

### **Example: Inspecting a Struct at Runtime**

```go
package main

import (
	"fmt"
	"reflect"
)

// Define a struct
type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 25}

	// Get type and value at runtime
	t := reflect.TypeOf(p)
	v := reflect.ValueOf(p)

	// Loop through fields dynamically
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("Field: %s, Value: %v\n", field.Name, value)
	}
}
```

✔ **Use case:** Logging, debugging, serialization (e.g., JSON marshalling).

---

## **2️⃣ Code Generation (`go generate`)**

Go provides a built-in tool **`go generate`** to **automate code generation** before compilation.

### **Example: Generating Boilerplate Code**

```go
//go:generate echo "Generating code..."
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```

✔ **Use case:** Generating mock files, gRPC stubs, or enum string conversions.

Run:

```sh
go generate
```

---

## **3️⃣ Templates (`text/template` & `html/template`)**

Go templates allow dynamic content generation for **HTML files, configuration files, and even code.**

### **Example: Generating Code with Templates**

```go
package main

import (
	"os"
	"text/template"
)

const structTemplate = `
type {{.Name}} struct {
	{{range .Fields}}{{.}} string
	{{end}}
}
`

func main() {
	tmpl := template.Must(template.New("struct").Parse(structTemplate))
	data := struct {
		Name   string
		Fields []string
	}{
		Name:   "User",
		Fields: []string{"Name", "Email", "Phone"},
	}

	tmpl.Execute(os.Stdout, data)
}
```

✔ **Use case:** Generating Go structs, config files, or API clients dynamically.

---

## **4️⃣ Generics (Go 1.18+)**

Generics allow **writing flexible code** that works with multiple types without reflection.

### **Example: Generic Function**

```go
package main

import "fmt"

// Generic function
func Print[T any](value T) {
	fmt.Println(value)
}

func main() {
	Print(42)       // Works with int
	Print("Hello!") // Works with string
}
```

✔ **Use case:** Avoid code duplication in **sorting, mapping, or mathematical operations**.

---

## **🚀 Summary: How Go Supports Meta-Programming**

| **Feature**                         | **Use Case**                                                |
|-------------------------------------|-------------------------------------------------------------|
| **Reflection (`reflect`)**          | Inspect structs at runtime (e.g., serialization, ORM, JSON) |
| **Code Generation (`go generate`)** | Pre-compile automation (e.g., mocks, API clients)           |
| **Templates (`text/template`)**     | Dynamically generate Go code or configuration files         |
| **Generics (`T any`)**              | Reuse logic across multiple types without reflection        |

---

### 19. go new release feature

### 20. Runtime package (Different operations, current memory, cpu usage etc.)

# **🔹 `runtime` Package in Go**

The `runtime` package in Go provides functions for interacting with the Go runtime, including **Goroutine management,
memory statistics, garbage collection, and CPU control**.

---

## **1️⃣ Key Features of `runtime` Package**

| Feature            | Function                 | Description                             |
|--------------------|--------------------------|-----------------------------------------|
| **Goroutine Info** | `runtime.NumGoroutine()` | Get the number of active Goroutines     |
| **Memory Usage**   | `runtime.MemStats`       | Retrieve heap, stack, and GC stats      |
| **Force GC**       | `runtime.GC()`           | Manually trigger garbage collection     |
| **CPU Usage**      | `runtime.NumCPU()`       | Get the number of available CPU cores   |
| **Set Max CPUs**   | `runtime.GOMAXPROCS(n)`  | Set the number of CPU cores used        |
| **Caller Info**    | `runtime.Caller()`       | Get function call details (stack trace) |
| **Stack Trace**    | `runtime.Stack()`        | Dump Goroutine stack traces             |

---

## **2️⃣ Goroutine Management**

🔹 **Get the number of currently running Goroutines**

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func worker() {
	time.Sleep(2 * time.Second)
}

func main() {
	fmt.Println("Goroutines before:", runtime.NumGoroutine())

	go worker()
	fmt.Println("Goroutines after:", runtime.NumGoroutine())

	time.Sleep(3 * time.Second) // Allow Goroutine to finish
	fmt.Println("Goroutines final:", runtime.NumGoroutine())
}
```

✅ **Output:**

```
Goroutines before: 1
Goroutines after: 2
Goroutines final: 1
```

---

## **3️⃣ Get Memory Usage**

🔹 **Check current memory allocation, GC stats, and heap usage**

```go
package main

import (
	"fmt"
	"runtime"
)

func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc: %v KB\n", m.Alloc/1024)
	fmt.Printf("TotalAlloc: %v KB\n", m.TotalAlloc/1024)
	fmt.Printf("Sys: %v KB\n", m.Sys/1024)
	fmt.Printf("NumGC: %v\n", m.NumGC)
}

func main() {
	printMemoryUsage()
}
```

✅ **Example Output:**

```
Alloc: 512 KB
TotalAlloc: 1024 KB
Sys: 8192 KB
NumGC: 5
```

✔ **`Alloc`** → Memory currently allocated  
✔ **`TotalAlloc`** → Total allocated memory (including freed)  
✔ **`Sys`** → Total memory obtained from OS  
✔ **`NumGC`** → Number of times Garbage Collector (GC) has run

---

## **4️⃣ Force Garbage Collection (`runtime.GC()`)**

🔹 **Trigger the Garbage Collector manually**

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Before GC")
	runtime.GC() // Force GC
	fmt.Println("After GC")
}
```

✔ Usually **not needed** (Go's GC is automatic)  
✔ Can be useful **before high-memory operations**

---

## **5️⃣ CPU Usage and Tuning**

🔹 **Get the number of CPU cores**

```go
fmt.Println("CPU Cores:", runtime.NumCPU())
```

🔹 **Set how many CPU cores Go should use**

```go
runtime.GOMAXPROCS(2) // Limit Go to 2 CPU cores
```

✔ Can improve **performance in CPU-intensive tasks**  
✔ Use **all cores** for **parallel processing**

---

## **6️⃣ Get Function Call Info (`runtime.Caller()`)**

🔹 **Find out which function was called**

```go
package main

import (
	"fmt"
	"runtime"
)

func printCaller() {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("Could not get caller info")
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fmt.Printf("Called from function: %s (%s:%d)\n", funcName, file, line)
}

func main() {
	printCaller()
}
```

✅ **Output Example:**

```
Called from function: main.main (main.go:15)
```

✔ **Useful for logging, debugging, and stack tracing**

---

## **7️⃣ Get Stack Trace (`runtime.Stack()`)**

🔹 **Capture all Goroutine stack traces**

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	buf := make([]byte, 1024)
	n := runtime.Stack(buf, true)
	fmt.Println(string(buf[:n]))
}
```

✔ **Useful for debugging deadlocks and crashes**

---

## **🚀 Summary**

| Feature           | Function                 | Use Case                   |
|-------------------|--------------------------|----------------------------|
| **Goroutines**    | `runtime.NumGoroutine()` | Count active Goroutines    |
| **Memory Stats**  | `runtime.MemStats`       | Check heap usage & GC info |
| **Manual GC**     | `runtime.GC()`           | Trigger garbage collection |
| **CPU Cores**     | `runtime.NumCPU()`       | Get CPU core count         |
| **Set CPU Usage** | `runtime.GOMAXPROCS(n)`  | Limit Go's CPU usage       |
| **Caller Info**   | `runtime.Caller()`       | Debugging and logging      |
| **Stack Trace**   | `runtime.Stack()`        | Debug Goroutine issues     |
