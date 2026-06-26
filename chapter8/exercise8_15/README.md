## Chat server
NEW: improve BLOCKING if one of clients broadcasting is slow or not through here
```sh
for cli := range clients {
  cli.ch <- msg
}
```

- go to `server`, then `go run main.go`
- go to `client`, then `go run main.go Ken` 
- go to `client`, then `go run main.go Alice` to open 2nd Client
Start chatting