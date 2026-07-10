## Cache (memoization) for concurrent goroutines

Make sure that mutex has safeguarded the `cache` => but it makes program slows as we block every 2nd goroutine

```sh
go mod init memo2
go test -run=TestConcurrent -race -v memo2
```
