## Make thumbnails for images in folder `images/` in parallel
NEW: 
- using a separate goroutine to parse file name
- use `sync.WaitGroup` to countdown goroutines to track

Run
```sh
go run main.go
```