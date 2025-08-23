package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

/*
1. In listings 6.5 and 6.6, we developed a recursive concurrent file search. When a goroutine finds
a file match, it outputs it on the console. Can you change the implementation of this file search
so that it prints all the file matches, sorted into alphabetical order, after the search completes?
Hint: try collecting the results in a shared data structure instead of printing them on the console
from the goroutine.
*/

func fileSearch(dir string, filename string, wg *sync.WaitGroup, mutex *sync.Mutex, names *[]string) {
	files, _ := os.ReadDir(dir)
	for _, file := range files {
		fpath := filepath.Join(dir, file.Name())
		if strings.Contains(file.Name(), filename) {
			mutex.Lock()
			*names = append(*names, fpath)
			mutex.Unlock()
		}
		if file.IsDir() {
			wg.Add(1)
			go fileSearch(fpath, filename, wg, mutex, names)
		}
	}
	wg.Done()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	mutex := sync.Mutex{}
	var names []string
	go fileSearch(os.Args[1], os.Args[2], &wg, &mutex, &names)
	wg.Wait()
	mutex.Lock()
	sort.Strings(names)
	fmt.Println(names)
	mutex.Unlock()
}
