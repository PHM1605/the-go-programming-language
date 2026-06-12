## Read HTML file, print all elements found in that HTML with hierarchy

### NEW: Print comments, text, attributes `<a href='...'>`, short form of `<img />` instead of `<img>` and `</img>`

To install missing lib package
```sh
go mod init exercise5_7
go get golang.org/x/net/html
```

Run
```sh
go run main.go http://gopl.io
```