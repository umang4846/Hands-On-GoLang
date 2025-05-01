# 04 - Buffered vs Unbuffered Channel

## Problem
Demonstrate the difference between buffered and unbuffered channels.

## Requirements
- Show that unbuffered channels block until the value is received.
- Show that buffered channels allow sending up to their capacity.

## Example Output
```
Sending to buffered channel
Sent to buffered channel
Sending to unbuffered channel
(Unblocks only after receive)
Received from unbuffered channel
```