## Crawling all links from a website (only to print links & sublinks out)
Install missing libraries 
```sh
go mod init findlinks3
go get golang.org/x/net/html
```

Run
```sh
go run findlinks3 https://golang.org
```