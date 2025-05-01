# 10 - Safe Counter Using Mutex

## Problem
Launch 100 goroutines that each increment a shared counter. Use `sync.Mutex` to avoid race conditions.

## Requirements
- Use `sync.Mutex` to protect counter access.

## Example Output
```
Final Counter: 100
```