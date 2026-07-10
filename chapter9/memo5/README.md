## Cache (memoization) for concurrent goroutines
NOTE: good solution here, use a `map` attaching to a `server goroutine`

```sh
go mod init memo5
go test -run=TestConcurrent -race -v memo5
```
