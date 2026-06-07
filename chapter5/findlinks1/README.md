## Read HTML file, parse it, and print all links found in that HTML
To install missing lib package
```sh
go mod init findlinks1
go get golang.org/x/net/html
```

Run
```sh
go build -o fetch ./fetch_lib
go build -o findlinks1 
./fetch https://golang.org | ./findlinks1
```