## Cache (memoization) for concurrent goroutines
NOTE: good solution here, use a `result with ready channel` to flag when result is ready (force other goroutines to wait for that channel)

```sh
go mod init memo4
go test -run=TestConcurrent -race -v memo4
```
