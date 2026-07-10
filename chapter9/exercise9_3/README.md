## Cache (memoization) a list of links
NEW: each request has additional `done` channel to cancel that request.

```sh
go mod init exercise9_3
go test -run=TestCancellation -race -v exercise9_3
```
