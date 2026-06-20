## Simple e-commerce website

In order to be a valid HTTP handler, a type must comply with following interface
```sh
package http
type Handler interface {
  ServeHTTP(w ResponseWriter, r *Request)
}
```

Run program
```sh
go mod init http2
go build http2
./http2
```
Then visit 
```sh
http://localhost:8000/list
http://localhost:8000/price?item=socks
http://localhost:8000/price?item=shoes
```