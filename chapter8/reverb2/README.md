## Echo server with multiple goroutines per client
Reverberation: return what-being-received-from-client in a more-silent-voice gradually
- Client calls "Hello?"
- Server answers "HELLO?" and "Hello?" and "hello?"

Runs together with `netcat2` program
```sh
go run main.go
```

NEW: this version is better. With 1 Client, when it receives 1 input, echos not yet 3 times, it still can receive the next input from that Client and echoing