## Simple e-commerce website

Run program
```sh
go mod init http3
go build http3
./http3
```
Then visit 
```sh
http://localhost:8000/list
http://localhost:8000/price?item=socks
http://localhost:8000/price?item=shoes
```