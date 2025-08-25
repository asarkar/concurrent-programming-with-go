package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

/*
1. Write a program similar to the one in listing 2.3 that accepts a list of text filenames as
arguments.
For each filename, the program should spawn a new goroutine that will output the contents of that
file to the console. You can use the time.Sleep() function to wait for the child goroutines to
complete (until you know how to do this better). Call the program catfiles.go. Here's how you can
execute this Go program:

go run catfiles.go ../../files/file1.txt ../../files/file2.txt ../../files/file3.txt
*/

func cat(filename string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}

func main() {
	filenames := os.Args[1:]
	for _, filename := range filenames {
		go cat(filename)
	}
	time.Sleep(2 * time.Second)
}
