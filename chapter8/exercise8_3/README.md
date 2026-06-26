## `netcat`: hanging to receive response, or send to Server
Use `channel` to synchronize 2 goroutines in this Client, to properly close
NEW: Close only `Write` half of the connection so that it can still receives echos
Run 
```sh
go run main.go
```
Run together with the TCP server `reverb1` in this same folder

When main goroutine is pressed `Ctrl+D`, it ends but still receive the rest of echos

