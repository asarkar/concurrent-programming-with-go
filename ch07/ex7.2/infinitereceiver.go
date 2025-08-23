package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
2. In listing 7.8, the receiver reads a 0 when the channel is closed. Can you try it with different
data types? What happens if the channel is of type string? What if it is of type slice?

ANSWER: Default value for string is an empty string (”), default for a slice is an empty slice.
*/
func main() {
	intChannel := make(chan int)
	go receiverInt(intChannel)
	stringChannel := make(chan string)
	go receiverString(stringChannel)
	sliceChannel := make(chan []int)
	go receiverSlice(sliceChannel)
	for i := 1; i <= 3; i++ {
		fmt.Println(time.Now().Format("15:04:05"), "Sending:", i)
		intChannel <- i
		stringChannel <- strconv.Itoa(i)
		sliceChannel <- []int{i}
		time.Sleep(1 * time.Second)
	}
	close(intChannel)
	close(stringChannel)
	close(sliceChannel)
	time.Sleep(3 * time.Second)
}

func receiverInt(messages <-chan int) {
	for {
		msg := <-messages
		fmt.Println(time.Now().Format("15:04:05"), "Received int:", msg)
		time.Sleep(1 * time.Second)
	}
}

func receiverString(messages <-chan string) {
	for {
		msg := <-messages
		fmt.Printf("%s Received string: '%s'\n", time.Now().Format("15:04:05"), msg)
		time.Sleep(1 * time.Second)
	}
}

func receiverSlice(messages <-chan []int) {
	for {
		msg := <-messages
		fmt.Println(time.Now().Format("15:04:05"), "Received slice:", msg)
		time.Sleep(1 * time.Second)
	}
}
