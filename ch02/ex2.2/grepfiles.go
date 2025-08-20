package main

import (
    "fmt"
    "log"
    "os"
    "time"
    "strings"
)

/*
2. Expand the program you wrote in the first exercise so that instead of printing the contents 
of the text files, it searches for a string match. The string to search for is the first 
argument on the command line. When you spawn a new goroutine, instead of printing the file's 
contents, it should read the file and search for a match. If the goroutine finds a match, 
it should output a message saying that the filename contains a match. Call the program 
grepfiles.go. Here's how you can execute this Go program ("bubbles" is the search string in 
this example):

go run grepfiles.go bubbles ../../files/file1.txt ../../files/file2.txt ../../files/file3.txt
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
    filenames := os.Args[2:]
    for _, filename := range filenames {
        go grep(filename, word)
    }
    time.Sleep(2 * time.Second)
}