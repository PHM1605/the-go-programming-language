## Build a CLI similar to git, with command `issue`
Create issue (`POST` request)
```sh
go run . create phm1605/the-go-programming-language
```
Read issue `12345`
```sh
go run . read phm1605/the-go-programming-language 12345
```
Update issue `12345`
```sh
go run . update phm1605/the-go-programming-language 12345
```
Close issue `12345`
```sh
go run . close phm1605/the-go-programming-language 12345
```
Delete issue `12345` (currently not supported by Github)
```sh
issue delete phm1605/the-go-programming-language 12345
```
