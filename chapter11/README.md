## word1
- check Palindrome naive by checking char by char

Run 
```sh
go test chapter11/word1
```

Test with time of each test
```sh
go test -v chapter11/word1
```

Test only a list of specific tests
```sh
go test -v -run="French|Canal" chapter11/word1
```

## word2
- check Palindrome using `rune` instead of `char`
- ignore spaces, punctuation and letter case
Run 
```sh
go test -v chapter11/word2
```
To check benchmark 
```sh
go test -bench=. chapter11/word2
```
To include `memory-allocations-per-operation` information in the benchmark
```sh
go test -bench=. -benchmem chapter11/word2
```

## exercise11_1
- test the `character count` program of the word `hello`

Run 
```sh
go test -v chapter11/exercise11_1
```

## exercise11_2
- test `IntSet` if it matches standard library `map[int]bool`

Run 
```sh
go test -v chapter11/exercise11_2
```

## exercise11_3
- test a random string input is palindrome or not

Run 
```sh
go test -v chapter11/exercise11_3
```

## exercise11_4
- test a random string is palindrome or not
- NEW: that string has spaces or punctuations too
Run 
```sh
go test -v chapter11/exercise11_4
```

## echo
- test `echo` COMMAND with command-line arguments
Normal run 
```sh
go run echo.go -n -s "," hello world
```
Test
```sh
go test -v chapter11/echo
```

## storage1
- web service like Google drive
- shows quota-checking logic (sends email to user when he exceeds 90% quote)
NOTE: we can't test this because we can't send email to Client to test

## storage2
- same as above BUT use a dummy `notifiedUser()` function when test to NOT send anything to the real user
Run
```sh
go test -v chapter11/storage2
```

## exercise11_5
- test splitting of long string 
Run 
```sh
go test -v chapter11/exercise11_5
```

## eval 
- test evaluation an expression with `Coverage`
Run
```sh
go test -v -coverprofile=eval.out chapter11/eval
go tool cover -html=eval.out
```

## exercise11_6
- benchmark among different approaches of `popcount` program
- remind: `popcount` counts number of 1 bits from an `uint64`
```sh
go test -bench=. -benchmem chapter11/exercise11_6
```

## exercise11_7
- bench mark our `IntSet` that replicates the `Set` of integers
- compare with standard library Set `map[int]bool`
```sh
go test -bench=. chapter11/exercise11_7
```

## profiling
Standard way
```sh
go test -cpuprofile=cpu.out
go test -cpuprofile=block.out
go test -memprofile=mem.out
```

Real case: \
We move to `profiling` folder first (this folder is only to save .exe and log)
```sh
cd profiling
```
- run benchmark test `BenchmarkClientServerParallel` in folder package `http` 
- we don't have test cases in this folder hence `-run=NONE`
- log of profiling is stored in `cpu.log`
- in profiling the test .exe is stored inside `http.test` instead of discarding (like normal `go test...` without profiling flags)
```sh
go test -run=NONE -bench=BenchmarkClientServerParallel -cpuprofile=cpu.log net/http
```
Analyze the result with `pprof`
- `nodecount` means this `text` table has 10 rows
```sh
go tool pprof -text -nodecount=10 ./http.test cpu.log
```
- we can open in web too
- notice that we must install `graphviz`: `brew install graphviz`
```sh
go tool pprof -web -nodecount=10 ./http.test cpu.log
```