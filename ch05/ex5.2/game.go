package main

import (
	"fmt"
	"sync"
	"time"
)

/*
2. Change the game-sync listings 5.8 and 5.9 so that, still using condition variables, the
players wait for a fixed number of seconds. If the players haven't all joined within this time,
the goroutines should stop waiting and let the game start without all the players.
Hint: try using another goroutine with an expiry timer.
*/

func playerHandler(cond *sync.Cond, playersRemaining *int,
	playerId int, cancel *bool) {
	cond.L.Lock()
	fmt.Println(playerId, ": Connected")
	*playersRemaining--
	if *playersRemaining == 0 {
		cond.Broadcast()
	}
	for *playersRemaining > 0 && !*cancel {
		fmt.Println(playerId, ": Waiting for more players")
		cond.Wait()
	}
	cond.L.Unlock()
	if *cancel {
		fmt.Println(playerId, ": Game cancelled")
	} else {
		fmt.Println("All players connected. Ready player", playerId)
	}
}

func timeout(cond *sync.Cond, cancel *bool) {
	time.Sleep(10 * time.Second)
	cond.L.Lock()
	*cancel = true
	cond.Broadcast()
	cond.L.Unlock()
}

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	cancel := false
	go timeout(cond, &cancel)
	playersInGame := 5
	for i := 0; i < 4; i++ {
		go playerHandler(cond, &playersInGame, i, &cancel)
		time.Sleep(1 * time.Second)
	}
	time.Sleep(60 * time.Second)
}
