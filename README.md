# Ringbuffer
A simple ring buffer implementation in go using generics.
The ring buffer will reuse elements that have been already read, or will auto-expand when the length of unread elements exceeds the buffer.

For ring buffer uses see: [https://en.wikipedia.org/wiki/Circular_buffer](https://en.wikipedia.org/wiki/Circular_buffer)

## Example

```go
package main

import (
    "fmt"
    "github.com/clarktravism/ringbuffer"
)

func main() {
	buffer := ringbuffer.NewRingBuffer[int](4)
	buffer.Append(1)
	buffer.Append(2)
	buffer.Append(3)
	
	for !buffer.Empty() {
		value := buffer.Next()
		fmt.Println(value)
	}
	
	buffer.Append(4)
	buffer.Append(5)
	buffer.Append(6)
	
	for !buffer.Empty() {
		value := buffer.Next()
		fmt.Println(value)
	}
}
```