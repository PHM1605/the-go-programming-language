## Temperature conversion app; store `Celsius` regardless what is passed in by `--temp` flag

```sh
go mod init tempflag
go build -o tempflag
```

Usage
```sh
go build tempflag
./tempflag
./tempflag -temp -18C
./tempflag -temp 212°F
./tempflag -temp 272.15K
./tempflag -help
```