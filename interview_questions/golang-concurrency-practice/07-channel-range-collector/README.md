# 07 - Channel Range Collector

## Problem
Send numbers 1 to 10 from a goroutine to a channel and use `range` in the main goroutine to collect and print them.

## Requirements
- Use `close` to end the channel.

## Example Output
```
1
2
...
10
```