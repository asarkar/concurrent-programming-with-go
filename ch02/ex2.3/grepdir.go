package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "strings"
    "path/filepath"
)

/*
3. Change the program you wrote in the second exercise so that instead of passing a list of text
filenames, you pass a directory path. The program will look inside this directory and list the 
files. For each file, you can spawn a goroutine that will search for a string match (the same as 
before). Call the program grepdir.go. Here's how you can execute this Go program:

go run grepdir.go bubbles ../../files
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

func main() {
    word := os.Args[1]
    dir := os.Args[2]
    files, err := os.ReadDir(dir)
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        if !file.IsDir() {
            go grep(filepath.Join(dir, file.Name()), word)
        }
    }
    time.Sleep(2 * time.Second)
}