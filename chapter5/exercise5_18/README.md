## Improve the `fetch` program that writes HTTP response to a FILE (instead of `Stdout`)
NEW: use the correct `defer` way for closing the file opened by `os.Create()` \
Run 
```sh
go run main.go https://golang.org
```
