## Clock server that writes current time to client once per second
NOTE: this is bad approach because server can push time to ONE client at one time

Run TCP server
```sh
go mod init clock1
go build
./clock1 &
```

Then use `netcat` or `telnet` to talk to that server
```sh
nc localhost 8000
```
Kill that process
```sh
killall clock1
```