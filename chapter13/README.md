## unsafeptr
- illustrate the use of `uintptr` as raw number
- should be converted to `unsafe.Pointer` IMMEDIATELY to satisfy Garbage Collector
```sh
go run chapter13/unsafeptr
```

## deepequal
- use `reflect.DeepEqual` to compare 2 `struct`s
```sh
go test chapter13/deepequal
```

## equal 
- write function `equal()` that compares 2 arbitrary values
- can compare 2 structs like `reflect.DeepEqual()`
- NEW: it considers `nil` and empty `[]string` or empty `map[string]int` as "equal"
```sh
go run chapter13/equal
```

## exercise13_1 
- write function `equal()` that compares 2 arbitrary values
- can compare 2 structs like `reflect.DeepEqual()`
- it considers `nil` and empty `[]string` or empty `map[string]int` as "equal"
- NEW: 1_000_000_000 and 1_000_000_001 are considered equal
```sh
go run chapter13/exercise13_1
```

## exercise13_2
- write a function that detects where a variable is a cyclic data structure
```sh
go run chapter13/exercise13_2
```

## bzip & bzipper
- data compression Go program
- for data compression: `libbzip2` package (in C) \
=> we make a `cgo` wrapper of a shorter file `bzip2.c` instead
- for data decompression: `compress/bzip2` package (in Go)
- `bzipper` reads `stdin`, compresses it, and writes to `stdout`

On MacOS, we install `bzip2` for Mac for data decompression
```sh
brew install bzip2
```
Tell the compiler "Please link to library `bz2`"
```sh
#cgo LDFLAGS: -lbz2
```
If we have C-error in Go code
```sh
Cmd+Shift+P
Go: Restart language server
```
Run
```sh
go build -o bzipper chapter13/bzipper
wc -c < /usr/share/dict/words # count how many words initially e.g. 2493000
sha256sum < /usr/share/dict/words # see sha256sum-representation of our original data
./bzipper/bzipper < /usr/share/dict/words | wc -c # compress, print data, pipeline to see data-size e.g. 857578
```
We can decompress to original with `bunzip2` and check the sha256sum-representation to see if it's the same
```sh
./bzipper/bzipper < /usr/share/dict/words | bunzip2 | sha256sum
```

## exercise13_3
- improve `bzip NewWriter` so that it is `concurrent safe`

Run 
```sh
go build -o exercise13_3/newbzipper chapter13/exercise13_3
./exercise13_3/newbzipper > exercise13_3/exercise13_3.bz2
```
To decompress to test
```sh
bunzip2 -c exercise13_3/exercise13_3.bz2 > exercise13_3/output.txt
```

## exercise13_4
- use pure Go for data compression and decompression
```sh
go build -o exercise13_4/bzippergo chapter13/exercise13_4
./exercise13_4/bzippergo < exercise13_4/exercise13_4.txt > exercise13_4/exercise13_4.bz2
bunzip2 -c exercise13_4/exercise13_4.bz2 > exercise13_4/decompressed.txt
```
