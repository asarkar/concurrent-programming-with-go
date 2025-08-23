package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
1. In listing 8.15, we have two goroutines. The generateTemp() function simulates reading and
sending the temperature on a channel every 200 ms. The outputTemp() function simply outputs a
message found on a channel every 2 seconds. Can you write a main() function, using a select
statement, that reads messages coming from the generateTemp() goroutine and sends only the
latest temperature to the outputTemp() channel? Since the generateTemp() function outputs
values faster than the outputTemp() function, you'll need to discard some values so that only
the most up-to-date temperature is displayed.
*/

func main() {
	tempChannel := make(chan int)
	outputTemp(tempChannel)
	temps := generateTemp()
	t := <-temps
	for {
		select {
		case t = <-temps:
		case tempChannel <- t: // blocks until previous message is read
		}
	}
}

func generateTemp() chan int {
	output := make(chan int)
	go func() {
		temp := 50 //fahrenheit
		for {
			output <- temp
			temp += rand.Intn(3) - 1
			time.Sleep(200 * time.Millisecond)
		}
	}()
	return output
}

func outputTemp(input chan int) {
	go func() {
		for {
			fmt.Println("Current temp:", <-input)
			time.Sleep(2 * time.Second)
		}
	}()
}
