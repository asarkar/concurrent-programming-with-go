package main

import (
	listing10_1 "concurrent-programming-with-go/ch10"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

/*
1. Implement the same directory hashing that we did in listing 10.4, but instead
of using channels to synchronize between iterations, try using waitgroups.
*/
func main() {
	dir := os.Args[1]
	files, _ := os.ReadDir(dir)
	sha := sha256.New()
	var prev, next *sync.WaitGroup
	for _, file := range files {
		if !file.IsDir() {
			next = &sync.WaitGroup{}
			next.Add(1)
			go func(filename string, prev, next *sync.WaitGroup) {
				fpath := filepath.Join(dir, filename)
				hashOnFile := listing10_1.FHash(fpath)
				// If not the first iteration
				if prev != nil {
					prev.Wait()
				}
				sha.Write(hashOnFile)
				next.Done()
			}(file.Name(), prev, next)
			prev = next
		}
	}
	next.Wait()
	fmt.Printf("%x\n", sha.Sum(nil))
}
