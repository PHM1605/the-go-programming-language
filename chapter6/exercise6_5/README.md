## Demonstrate BITVECTOR type; using a list of bits (divided to many rows) to represent >=0 numbers
NEW: aware that machine can be 32 BIT ONLY too (so `uint64` might crash) \
=> Basic idea: add a `wordSize` to be calculated for each 32- or 64-bit system
```sh
const wordSize = 32 << (^uint(0) >> 63)
```
and use it anywhere `64` is coded \

Run
```sh
go run main.go
```
