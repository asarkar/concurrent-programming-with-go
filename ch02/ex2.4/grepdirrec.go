package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

/*
4. Adapt the program in the third exercise to continue searching recursively in any subdirectories.
If you give your search goroutine a file, it should search for a string match in that file, just
like in the previous exercises. Otherwise, if you give it a directory, it should recursively spawn
a new goroutine for each file or directory found inside. Call the program grepdirrec.go, and
execute it by running this command:

go run grepdirrec.go bubbles ../../files
*/

func grep(filename string, word string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	found := strings.Contains(string(content), word)
	if found {
		fmt.Println(filename, "contains", word)
	}
}

func list(path string, word string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}
	if fileInfo.IsDir() {
		files, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}
		for _, file := range files {
			go list(filepath.Join(path, file.Name()), word)
		}
	} else {
		go grep(path, word)
	}
}

func main() {
	word := os.Args[1]
	path := os.Args[2]
	list(path, word)

	time.Sleep(2 * time.Second)
}
