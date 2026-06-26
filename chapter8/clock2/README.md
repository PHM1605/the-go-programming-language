## Clock server that writes current time to client once per second
NEW: handle multiple clients at the same time with `go function()` 

Run TCP server
```sh
go mod init clock2
go build
./clock2 &
```

Then use `netcat` or `telnet` to talk to that server
```sh
nc localhost 8000
```
Kill that process
```sh
killall clock2
```