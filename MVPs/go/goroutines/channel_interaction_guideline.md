# Go Channels: Data Interaction Guide

This guide will help you learn how to interact with data inside Go channels, from the basics to advanced usage patterns.

---

## Table of Contents
1. [Introduction to Channels](#introduction-to-channels)
2. [Creating Channels](#creating-channels)
3. [Sending Data](#sending-data)
4. [Receiving Data](#receiving-data)
5. [Buffered vs Unbuffered Channels](#buffered-vs-unbuffered-channels)
6. [Closing Channels](#closing-channels)
7. [Range and Channels](#range-and-channels)
8. [Select Statement](#select-statement)
9. [Best Practices](#best-practices)
10. [Further Reading](#further-reading)

---

## Introduction to Channels
Channels are Go's way of communicating between goroutines. They provide a safe way to send and receive data.

## Creating Channels
```go
ch := make(chan int) // Unbuffered channel
chBuf := make(chan int, 5) // Buffered channel
```

## Sending Data
Use the `<-` operator to send data:
```go
ch <- 42
```

## Receiving Data
Use `<-` to receive data:
```go
value := <-ch
```
Or, receive and check if the channel is open:
```go
value, ok := <-ch
if !ok {
    // channel is closed
}
```

## Buffered vs Unbuffered Channels
- **Unbuffered:** Send blocks until another goroutine receives.
- **Buffered:** Send only blocks when buffer is full.

## Closing Channels
Close a channel to signal no more data will be sent:
```go
close(ch)
```
Never send on a closed channel (causes panic). Receives from a closed channel yield zero value.

## Range and Channels
Use `range` to receive values until the channel is closed:
```go
for v := range ch {
    fmt.Println(v)
}
```

## Select Statement
`select` lets you work with multiple channels:
```go
select {
case v := <-ch1:
    fmt.Println(v)
case ch2 <- 99:
    fmt.Println("sent 99")
default:
    fmt.Println("no communication")
}
```

## Best Practices
- Always close channels when done sending.
- Never close a channel from the receiver side.
- Use `ok` to check for closed channels.
- Avoid using channels as a replacement for all data sharing; use them for synchronization and communication.

## Further Reading
- [Go Blog: Concurrency Patterns](https://go.dev/blog/pipelines)
- [Effective Go: Channels](https://go.dev/doc/effective_go#channels)
- [Go Tour: Channels](https://go.dev/tour/concurrency/2)
