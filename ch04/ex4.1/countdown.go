package main

import (
	"fmt"
	"sync"
	"time"
)

/*
1. Listing 4.15 (originally from chapter 3) does not use any mutexes to protect access to its
shared variable. This is bad practice. Change this program so that access to the shared seconds
variable is protected by a mutex. Hint: you might need to copy a variable.

Listing 4.15 Goroutines sharing a variable without synchronization

package main

import (
    "fmt"
    "time"
)

func countdown(seconds *int) {
    for *seconds > 0 {
        time.Sleep(1 * time.Second)
        *seconds -= 1
    }
}

func main() {
    count := 5
    go countdown(&count)
    for count > 0 {
        time.Sleep(500 * time.Millisecond)
        fmt.Println(count)
    }
}

ANWSER:
The shared variable is copied to reduce lock contention and ensure each loop iteration uses a
consistent snapshot of the shared state.

* The condition is evaluated based on a stable snapshot (remaining).
* Reads/writes of the shared variable (count) are always protected by the mutex.
* Neither goroutine ever holds the lock across Sleep or Println.
* No races, no blocking issues.
*/

func countdown(seconds *int, mutex *sync.Mutex) {
	mutex.Lock()
	remaining := *seconds
	mutex.Unlock()
	for remaining > 0 {
		time.Sleep(1 * time.Second)
		mutex.Lock()
		*seconds -= 1
		remaining = *seconds
		mutex.Unlock()
	}
}

func main() {
	mutex := sync.Mutex{}
	count := 5
	go countdown(&count, &mutex)
	remaining := count
	for remaining > 0 {
		time.Sleep(500 * time.Millisecond)
		mutex.Lock()
		fmt.Println(count)
		remaining = count
		mutex.Unlock()
	}
}
