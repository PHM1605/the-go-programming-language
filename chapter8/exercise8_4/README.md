## `netcat` client and `reverb` echo server (with reducing tone)
NEW: using `sync.WaitGroup` to monitor the number of goroutines

Init
```sh
go mod init exercise8_4
```

Run server
```sh
cd reverb
go run main.go
```

Run client
```sh
cd netcat
go run main.go
```
