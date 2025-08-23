package ex9_2

/*
2. In listing 9.18, we developed a take(n) goroutine. Extend the functionality of this goroutine to
implement TakeUntil(f), where f is a function returning a Boolean. The goroutine needs to continue
consuming and forwarding the messages on its input channel while the return value of f is true.
Using generics ensures that we can reuse the TakeUntil(f) function and plug it into many other
pipelines. Here's the function signature:
func TakeUntil[K any](f func(K) bool,quit chan int,input <-chan K) <-chan K
*/
func TakeUntil[K any](f func(K) bool, quit chan int, input <-chan K) <-chan K {
	output := make(chan K)
	go func() {
		defer close(output)
		moreData := true
		var msg K
		for moreData {
			select {
			case msg, moreData = <-input:
				if moreData {
					moreData = f(msg)
					if moreData {
						output <- msg
					}
				}
			case <-quit:
				return
			}
		}
		if !moreData {
			close(quit)
		}
	}()
	return output
}
