package main

import (
	"chapter13/exercise13_3/bzip"
	"os"
	"sync"
)

func main() {
	w := bzip.NewWriter(os.Stdout)
	defer w.Close()

	var wg sync.WaitGroup
	inputs := [][]byte{
		[]byte("Hello "),
		[]byte("from "),
		[]byte("multiple "),
		[]byte("goroutines "),
	}
	for _, data := range inputs {
		wg.Add(1)
		go func(d []byte) {
			defer wg.Done()
			if _, err := w.Write(d); err != nil {
				panic(err)
			}
		}(data)
	}
	wg.Wait()
}
