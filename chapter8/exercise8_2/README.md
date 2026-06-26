## Create an `ftpserver` and `ftpclient`

Init 
```sh
go mod init exercise8_2
```
Commands that `server` interprets from `client`
- cd: to change directory
- ls: list a directory
- get: send the contents of a file
- close: close the connection

Run 
```sh
go build -o ftpserverbin ./ftpserver
go build -o ftpclientbin ./ftpclient
./ftpserverbin
./ftpclientbin => start typing
```