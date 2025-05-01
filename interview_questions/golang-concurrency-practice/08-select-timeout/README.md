# 08 - Select With Timeout

## Problem
Simulate a blocking operation with a channel and use `select` with `time.After` to timeout if it takes too long.

## Requirements
- Use `select` to handle both cases.

## Example Output
```
Operation timed out
```