package main

import (
	"fmt"
	"time"
)

/*
3. Consider the following listing. Can you find the race condition in this program without running
the race detector? Hint: Try running the program several times to see if it results in a race
condition.
*/

func addNextNumber(nextNum *[101]int) {
	i := 0
	for nextNum[i] != 0 {
		i++
	}
	nextNum[i] = nextNum[i-1] + 1
}

func main() {
	// Create an array of 101 integers, set the first element to 1, and leave the other 100
	// elements as 0.
	nextNum := [101]int{1}
	for i := 1; i <= 100; i++ {
		go addNextNumber(&nextNum)
	}
	for nextNum[100] == 0 {
		println("Waiting for goroutines to complete")
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println(nextNum)
}
