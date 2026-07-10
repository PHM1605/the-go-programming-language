## Cache (memoization) for concurrent goroutines

Mutex has safeguarded the `cache` TWICE
- 1st time when lookup
- 2nd time when update
Problem: `golang.org` is fetched TWICE due to reaching the mutex lock inside `if` nearly at the same time

```sh
go mod init memo3
go test -run=TestConcurrent -race -v memo3
```
