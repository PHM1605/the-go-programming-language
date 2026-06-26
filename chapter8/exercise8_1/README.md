## Write a clock server `clock2_tz` with PORT number; client `clockwall` displays multiple clocks from multiple countries

```sh
go mod init exercise8_1
go build -o clock2bin ./clock2_tz 
go build -o clockwallbin ./clockwall
```

NOTE: TZ=xxx is setting environment variable to affect `time.Now()` inside
```sh
TZ=US/Eastern ./clock2bin -port 8010 > /dev/null 2>&1 &
TZ=Asia/Tokyo ./clock2bin -port 8020 > /dev/null 2>&1 &
TZ=Europe/London ./clock2bin -port 8030 > /dev/null 2>&1 &
./clockwallbin NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
```

Terminate all pricess
```sh
killall clock2bin
```