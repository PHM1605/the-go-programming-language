## 3 goroutines exchange data with each other
Send 100 numbers only through 2 channels \
NEW: define clearly `send-only` channel (`chan<- int`) and `receive-only` channel (`<-chan int`)

```sh
go run main.go
```