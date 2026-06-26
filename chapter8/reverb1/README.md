## Echo server with multiple goroutines per client
Reverberation: return what-being-received-from-client in a more-silent-voice gradually
- Client calls "Hello?"
- Server answers "HELLO?" and "Hello?" and "hello?"

Runs together with `netcat2` program
```sh
go run main.go
```

This version is not good because, even with 1 Client only, when it receives an input from that Client it must echo THRICE before echoing the next input
