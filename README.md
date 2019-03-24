# SSH-Trap

A ssh server that accepts connections and only writes random data in 10 seconds interval to keep the connection alive.

```go
go run main.go -port 2324 -alsologtostderr
```