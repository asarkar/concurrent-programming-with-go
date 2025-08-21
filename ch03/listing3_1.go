package main

import (
	"fmt"
	"time"
)

/*
2. Run Go's race detector on listing 3.1. Does the result contain a race condition? If it does,
can you explain why it happens?

ANSWER:
==================
WARNING: DATA RACE
Write at 0x00c00011a038 by goroutine 7:
  main.countdown()
      /path/to/concurrent-programming-with-go/ch03/listing31.go:20 +0x54
  main.main.gowrap1()
      /path/to/concurrent-programming-with-go/ch03/listing31.go:10 +0x20

Previous read at 0x00c00011a038 by main goroutine:
  main.main()
      /path/to/concurrent-programming-with-go/ch03/listing31.go:13 +0xb8

Goroutine 7 (running) created at:
  main.main()
      /path/to/concurrent-programming-with-go/ch03/listing31.go:10 +0x9c
==================
*/

func main() {
	count := 5
	go countdown(&count)
	for count > 0 {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(count)
	}
}

func countdown(seconds *int) {
	for *seconds > 0 {
		time.Sleep(1 * time.Second)
		*seconds -= 1
	}
}
