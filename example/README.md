## how to run

### run docker

- make prepare

```bash
make prepare
```

### run server

```go
go run server/main.go
```

result:

```go
2022/07/26 13:38:23.859717 transport.go:91: [Info] HERTZ: HTTP server listening on address = 127.0.0.1:8888
2022/07/26 13:38:23.859745 transport.go:91: [Info] HERTZ: HTTP server listening on address = 127.0.0.1:8889
```

- 8888 port return pong1 and 8889 return pong2

### run client

```go
go run client/main.go
```

```go
2022/07/26 13:52:47.310617 main.go:46: [Info] code =200, body ={"ping":"pong2"}
2022/07/26 13:52:47.311019 main.go:46: [Info] code = 200, body ={"ping":"pong1"}
2022/07/26 13:52:47.311186 main.go:46: [Info] code = 200, body ={"ping":"pong1"}
2022/07/26 13:52:47.311318 main.go:46: [Info] code = 200, body ={"ping":"pong1"}
2022/07/26 13:52:47.311445 main.go:46: [Info] code = 200, body ={"ping":"pong1"}
2022/07/26 13:52:47.311585 main.go:46: [Info] code = 200, body ={"ping":"pong2"}
2022/07/26 13:52:47.311728 main.go:46: [Info] code = 200, body ={"ping":"pong1"}
2022/07/26 13:52:47.311858 main.go:46: [Info] code = 200, body ={"ping":"pong1"}
2022/07/26 13:52:47.311977 main.go:46: [Info] code = 200, body ={"ping":"pong1"}
2022/07/26 13:52:47.312107 main.go:46: [Info] code = 200, body ={"ping":"pong1"}
```