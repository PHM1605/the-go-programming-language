## `netcat`: hanging to receive response, or send to Server
NEW: use `channel` to synchronize 2 goroutines in this Client, to properly close

Run 
```sh
go run main.go
```
Run together with the TCP server `clock1` in this same folder

When main goroutine is pressed `Ctrl+D`, we see the print `Done` from the branch goroutine.

