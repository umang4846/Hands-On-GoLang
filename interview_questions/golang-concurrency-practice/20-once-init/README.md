# 20 - One-Time Initialization with sync.Once

## Problem
Ensure a function runs only once using `sync.Once`.

## Requirements
- Show that even with multiple goroutines, the init function runs once.

## Example Output
```
init called
```