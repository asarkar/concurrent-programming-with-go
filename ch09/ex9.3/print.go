package ex9_3

import "fmt"

/*
3. Write a goroutine that prints to the console the contents of any message it receives on a
channel and then forwards the message to the output channel. Again, use generics so that the
function can be reused in many situations:
func Print[T any](quit <-chan int, input <-chan T) <-chan T
*/
func Print[T any](quit <-chan int, input <-chan T) <-chan T {
	output := make(chan T)
	go func() {
		defer close(output)
		moreData := true
		var msg T
		for moreData {
			select {
			case msg, moreData = <-input:
				if moreData {
					fmt.Println(msg)
					output <- msg
				}
			case <-quit:
				return
			}
		}
	}()
	return output
}
