## Surface program that parses an `Expression string`
Run 
```sh
go mod init surface
go run . (as we have many files of package main)
```
The visit the following URL
```sh
http://localhost:8000/plot?expr=sin(-x)*pow(1.5,-r)
http://localhost:8000/plot?expr=pow(2,sin(y))*pow(2,sin(x))/12
http://localhost:8000/plot?expr=sin(x*y/10)/10
```