## format
- using `reflect.ValueOf()` to wrap a value of any type => print its string representation
```sh
go run chapter12/format
```
## display
- improve printing of composite types
```sh
go run chapter12/display
```

## exercise12_1
- improve `map` case of the `display` program; so that its key can be `struct` or `array`
```sh
go run chapter12/exercise12_1
```

## exercise12_1
- solve cyclic display by setting `maxDepth` parameters
```sh
go run chapter12/exercise12_2
```

## sexpr
- an encoding/decoding of an object mechanism called `S-expression` (like JSON or XML)
```sh
go run chapter12/sexpr
```

## exercise12_3
- an encoding of an object mechanism called `S-expression` (like JSON or XML)
- NEW: cover `float`, `complex`, `interface`
```sh
go run chapter12/exercise12_3
```

## exercise12_4
- an encoding of an object mechanism called `S-expression` (like JSON or XML)
- NEW: print it prettier with `indent` and `newline` instead of a very long line
```sh
go run chapter12/exercise12_4
```

## exercise12_5
- an encoding of an object to `JSON` format
```sh
go run chapter12/exercise12_5
```

## exercise12_6
- an encoding of an object mechanism called `S-expression` (like JSON or XML)
- NEW: `nil` or `boolean false` values are not encoded
```sh
go run chapter12/exercise12_6
```

## sexpr
- an encoding/decoding of an object mechanism called `S-expression` (like JSON or XML)
```sh
go run chapter12/sexpr
```

## exercise12_7
- an encoding/decoding of an object mechanism called `S-expression` (like JSON or XML)
- NEW: `streaming` instead of `Unmarshal`; we can use `conn` instead of `[]byte`
```sh
go run chapter12/exercise12_7
```

## exercise12_8
- an encoding/decoding of an object mechanism called `S-expression` (like JSON or XML)
- NEW: decode a list of Movie from file `movies.sexpr`
```sh
go run chapter12/exercise12_8
```

## exercise12_9
- `S-expression` decode to stream of `Token`s like `xml.Decoder`
```sh
go run chapter12/exercise12_9
```

## exercise12_10
- an encoding/decoding of an object mechanism called `S-expression` (like JSON or XML)
- NEW: handle `boolean`, `float`, `interface` as well
```sh
go run chapter12/exercise12_10
```

## params
- like an HTTP server to extract parameters in the HTTP request
```sh
go run chapter12/params 
```
- test with `curl`
```sh
curl "http://localhost:12345/search?l=golang&l=books&max=2&x=true"
curl "http://localhost:12345/search"
curl "http://localhost:12345/search?x=true&l=golang&l=programming"
curl "http://localhost:12345/search?q=hello&x=123"
```

## exercise12_11
- write a `Pack` function that convert a `struct` to HTTP request's parameters
```sh
go run chapter12/exercise12_11
```

## exercise12_12
- like an HTTP server to extract parameters in the HTTP request
- NEW: check if Client send valid request with `email`, `credit-card-number`, `US ZIP code` info
```sh
go run chapter12/exercise12_12 
```
valid request
```sh
curl "http://localhost:12345/search?email=bob@example.com&zip=90210&card=4111111111111111"
```
invalid request
```sh
curl "http://localhost:12345/search?email=abc&zip=123&card=xyz"
```

## exercise12_13
- an encoding/decoding of an object mechanism called `S-expression` (like JSON or XML)
- NEW: use `tag` in `Movie struct` to change the way that field is encoded into `S-expression`
```sh
go run chapter12/exercise12_13
```

## methods
- display Methods of a Type
```sh
go run chapter12/methods
```