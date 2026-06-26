## Web crawling in parallel

Install missing libraries
```sh
go mod init crawl1
go get golang.org/x/net/html
```

Run 
```sh
go run main.go http://gopl.io/
```

PROBLEMS: 
- program always wait on channel `worklist`, never terminates
- too much parallelism => network blocks
