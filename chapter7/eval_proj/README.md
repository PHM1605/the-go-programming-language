## Evaluator for arithmetic expressions
Expressions are 
```sh
floating-point literals
binary +, -, *, /
unary +, -
function calls e.g. pow(x,y)
variables like x, pi
parentheses
```

Run test (with `-v` means `verbose`)
```sh
go mod init eval_proj
go test -v ./eval
```

Run `Check()` several terms in our `main` function 
```sh
go run main.go
```