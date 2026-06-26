## Echo server (from `reverb2`) and `netcat2` client
NEW: control Clients. When Client doesn't type anything in 10s then disconnect it.

Open 3 terminals, cd into `netcat` (twice) and `reverb`

Run 3 times `go run main.go` to start 1 Server and 2 Clients.

Then type something in 1 Client, the other will close after 10s of no activity.