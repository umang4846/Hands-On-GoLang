# 13 - Fan-Out Pattern

## Problem
Distribute work across multiple goroutines reading from a single input channel.

## Requirements
- Use 3 worker goroutines to read from a common job channel.
- Print processed job info.
