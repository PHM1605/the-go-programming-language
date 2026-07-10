## Cache (memoization) for concurrent goroutines

In this version, DATA RACE when we gets from same URL without mutex

```sh
go mod init memo1
go test -run=TestConcurrent -race -v memo1
```
