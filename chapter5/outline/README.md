## Read HTML file, parse it, and print all links found in that HTML
To install missing lib package
```sh
go mod init outline
go get golang.org/x/net/html
```

Run
```sh
go build -o fetch ./fetch_lib
go build -o outline 
./fetch https://golang.org | ./outline
```