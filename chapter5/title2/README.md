## Fetch a link; returns `title` in `<header>` if HTML, returns `error` if not (e.g. `image`)
NEW: use `defer` \
Install missing libs
```sh
go mod init title2
go get golang.org/x/net/html
```

Run 
```sh
go run main.go http://gopl.io
go run main.go https://golang.org/doc/effective_go.html
go run main.go https://golang.org/doc/gopher/frontpage.png 
```