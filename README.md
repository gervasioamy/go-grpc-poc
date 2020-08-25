# go-grpc-poc
Very simple PoC for gRPC in Go.
A server that allows notifications to be sent by the clients

# Start the server
`go run server.go`

# Start the client
`go run client.go`

## Client CLI usage
- `send {something}` sends a notification to the server
- `list` will return all the notifications the server holds
- `remove {id}` removes the notification with the specified id

```
$ go run client.go 
-> send hello
2020/08/25 15:18:10 Response from Server: ID = 0
-> send world      
2020/08/25 15:19:20 Response from Server: ID = 1
-> list
2020/08/25 15:19:23 Response from Server: 
2020/08/25 15:19:23   - [0] (2020-08-25 15:18:10) - hola
2020/08/25 15:19:23   - [1] (2020-08-25 15:19:20) - world
-> remove 1
2020/08/25 15:19:27 Response from Server: Was removed? true
-> list
2020/08/25 15:19:29 Response from Server: 
2020/08/25 15:19:29   - [0] (2020-08-25 15:18:10) - hola
```