## Test the effect of `GOMAXPROCS`, using 2 goroutines (1 prints 0s, 1 prints 1s)
NOTE: `GOMAXPROCS` is the number of CPU cores that will execute the goroutines

Run
```sh
GOMAXPROCS=1 go run main.go
GOMAXPROCS=2 go run main.go
```