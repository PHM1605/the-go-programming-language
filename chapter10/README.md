## jpeg
- convert PNG from `mandelbrot` to JPEG

Run
```sh
go build -o mandelbrot.exe ./mandelbrot
go build -o jpeg.exe ./jpeg
./mandelbrot.exe | ./jpeg.exe > mandelbrot.jpg
```

## exercise10_1
- convert PNG from `mandelbrot` to a specific format
- use a flag to specify which format

Run
```sh
go build -o exercise10_1.exe ./exercise10_1
./mandelbrot.exe | ./exercise10_1.exe -format=gif > mandelbrot.gif
```
 
## exercise10_2
- design a package that reads `tar/zip` files 

Run 
```sh
go build -o exercise10_2.exe ./exercise10_2
cat test.zip | ./exercise10_2.exe -format=zip
```

## cross
- build for different operating systems `GOOS` and architectures `GOARCH`

Run 
```sh
GOARCH=arm64 GOOS=darwin go build -o cross.exe ./cross
./cross.exe
```

## exercise10_4
- report all packages in current workspace that transitively depend on packages specified on command line 

Run 
```sh
go build -o exercise10_4.exe ./exercise10_4
./exercise10_4.exe fmt strconv
```
