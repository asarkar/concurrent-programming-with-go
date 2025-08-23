package main

import "fmt"

/*
1. In listings 7.1 and 7.2, the receiver doesn't output the last message STOP. This is because
the main() goroutine terminates before the receiver() goroutine gets the chance to print out
the last message. Can you change the logic, without using extra concurrency tools and without
using the sleep function, so that the last message is printed?
*/
func main() {
	messages := make(chan string)
	go receiver(messages)
	fmt.Println("Sending HELLO...")
	messages <- "HELLO"
	fmt.Println("Sending THERE...")
	messages <- "THERE"
	fmt.Println("Sending STOP...")
	messages <- "STOP"
	<-messages // Block for a message
	fmt.Println("main() done")
}

func receiver(messages chan string) {
	for msg := range messages {
		fmt.Println("Received:", msg)
		if msg == "STOP" {
			close(messages)
		}
	}
	fmt.Println("receiver() done")
}
