## Crawling all links from a website
NEW: make local copies of the pages it finds, create directories as necessary.\
Don't make copies of other pages that come from different domains

Install missing libraries 
```sh
go mod init exercise5_13
go get golang.org/x/net/html
```

Run
```sh
go run exercise5_13 https://go.dev
```
(Check `mirror/` folder to see downloaded files)