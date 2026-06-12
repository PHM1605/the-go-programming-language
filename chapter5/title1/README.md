## Fetch a link; returns `title` if HTML, returns `error` if not (e.g. `image`)
Install missing libs
```sh
go mod init title1
go get golang.org/x/net/html
```

Run 
```sh
go run main.go http://gopl.io
go run main.go https://golang.org/doc/effective_go.html
go run main.go https://golang.org/doc/gopher/frontpage.png 
```