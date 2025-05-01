# 14 - Context Cancellation

## Problem
Use `context.WithCancel` to gracefully stop a goroutine doing long-running work.

## Requirements
- Use a cancellable context.
- Stop goroutine using `<-ctx.Done()`.
