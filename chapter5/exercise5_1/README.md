## Read HTML file, parse it, and print all links found in that HTML
### NEW: traverse `n.FirstChild` linked list using recursive to `visit` instead of a loop
To install missing lib package
```sh
go mod init exercise5_1
go get golang.org/x/net/html
```

Run
```sh
go build -o fetch ./fetch_lib
go build -o exercise5_1 
./fetch https://golang.org | ./exercise5_1
```