## Reports disk usage of 1 or more directories (from command line)
We can type `verbose -v` flag to print details during calculating process

NEW: add `Cancellation` i.e. using a `done` channel and `close(done)` \
to broadcast to ALL goroutines that are running "hey User has done his job, exit yourself"

Run
```sh
go run main.go -v /bin /opt
```