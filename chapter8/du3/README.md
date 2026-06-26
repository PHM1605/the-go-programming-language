## Reports disk usage of 1 or more directories (from command line)
We can type `verbose -v` flag to print details during calculating process

NEW: concurrent in `walkDir()`

Run
```sh
go run main.go -v /bin /opt
```