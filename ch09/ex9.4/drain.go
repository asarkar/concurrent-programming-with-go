package ex9_4

/*
4. Write a goroutine that drains the contents of its input channel without doing anything with
them. The goroutine simply reads a message and throws it away:
func Drain[T any](quit <-chan int, input <-chan T)
*/
func Drain[T any](quit <-chan int, input <-chan T) {
	go func() {
		for {
			select {
			case <-input:
			case <-quit:
				return
			}
		}
	}()
}
