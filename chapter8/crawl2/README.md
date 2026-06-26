## Web crawling in parallel
NEW: limit number of parallelism calls by using `tokens channel`  

Install missing libraries
```sh
go mod init crawl2
go get golang.org/x/net/html
```

Run 
```sh
go run main.go http://gopl.io/
```
