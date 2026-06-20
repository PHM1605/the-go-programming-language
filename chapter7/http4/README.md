## Simple e-commerce website
NEW: use `DefaultServeMux`, global `http.HandleFunc` and `http.Handle`

Run program
```sh
go mod init http4
go build http4
./http4
```
Then visit 
```sh
http://localhost:8000/list
http://localhost:8000/price?item=socks
http://localhost:8000/price?item=shoes
```