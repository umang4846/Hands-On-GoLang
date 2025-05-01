# 19 - Broadcast Using Channel Close

## Problem
Broadcast a shutdown signal to multiple goroutines by closing a channel.

## Requirements
- All goroutines should exit on channel close.

## Example Output
```
Worker 1 exiting
Worker 2 exiting
...
```