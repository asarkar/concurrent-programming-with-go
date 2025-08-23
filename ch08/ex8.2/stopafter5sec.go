package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
2. In listing 8.16, we have a goroutine in the generateNumbers() function that outputs random
numbers. Can you write a main() function using a select statement that continuously consumes
from the output channel, printing the output on the console until 5 seconds have elapsed from
the start of the program? After 5 seconds, the function should stop consuming from the output
channel, and the program should terminate.
*/
func main() {
	fmt.Println(time.Now().Format("15:04:05"), ": Start")
	ticks := time.Tick(time.Second * 5)
	nums := generateNumbers()
	for nums != nil {
		select {
		case n := <-nums:
			fmt.Println(time.Now().Format("15:04:05"), ":", n)
		case now := <-ticks:
			fmt.Println(now.Format("15:04:05"), ": End")
			nums = nil
		}
	}
}

func generateNumbers() chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for {
			output <- rand.Intn(10)
			time.Sleep(200 * time.Millisecond)
		}
	}()
	return output
}
