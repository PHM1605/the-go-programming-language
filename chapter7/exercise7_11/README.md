## Simple e-commerce website with CRUD
Run program
```sh
go mod init exercise7_11
go build exercise7_11
./exercise7_11
```
List of all items in HTML format
```sh
curl "http://localhost:8000/list"
```
Read 
```sh
curl "http://localhost:8000/price?item=shoes"
```
Create
```sh
curl "http://localhost:8000/create?item=hat&price=20.5"
```
Update
```sh
curl "http://localhost:8000/update?item=hat&price=16"
```
Delete
```sh
curl "http://localhost:8000/delete?item=shoes"
```