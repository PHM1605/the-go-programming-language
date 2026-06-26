## Build a simple version of `netcat` (TCP client)
Run together with the TCP server `clock1` in this same folder
```sh
go mod init netcat1
go build
./netcat1
```
This will send to server `localhost:8000`