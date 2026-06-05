To import module from folder at same hierarchy: go to parent `myproject`, then
```sh
go mod init myproject
```

Then inside child folder
```sh
import "myproject/importedModule"
```

Some dependencies
```sh
go get github.com/joho/godotenv
```