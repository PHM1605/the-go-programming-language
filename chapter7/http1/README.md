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
go mod init http1
go build http1
./http1
```
Then visit `http://localhost:8000`