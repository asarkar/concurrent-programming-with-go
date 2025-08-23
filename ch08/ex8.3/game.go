package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

/*
3. Consider listing 8.17 containing the player() function. This function creates a goroutine
simulating a player in a game moving along a two-dimensional plane. The goroutine returns
the movements at random times by writing UP, DOWN, LEFT, or RIGHT on an output channel.
Create a main() function that creates four player goroutines and outputs on the console all
movements from the four players. The main() function should terminate only when there is one
player left in the game. Here is an example of what the output should look like:

Player 1: DOWN
Player 0: LEFT
Player 3: DOWN
Player 2 left the game. Remaining players: 3
Player 1: UP
. . .
Player 0: LEFT
. . .
Player 0: LEFT
Player 3 left the game. Remaining players: 2
Player 1: RIGHT
. . .
Player 1: RIGHT
Player 0 left the game. Remaining players: 1
Game finished
*/

func main() {
	players := make([]chan string, 4)
	for i := range 4 {
		players[i] = player()
	}

	cases := make([]reflect.SelectCase, len(players))
	for i, ch := range players {
		cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
	}

	for len(cases) > 1 {
		i, move, ok := reflect.Select(cases)
		if ok {
			fmt.Printf("Player %d: %s\n", i, move)
		} else {
			fmt.Printf("Player %d left the game. Remaining players: %d\n", i, len(cases)-1)
			cases[i] = cases[0]
			cases = cases[1:]
		}
	}
	fmt.Println("Game finished")
}

// func filter(players []chan string) []chan string {
// 	nonNil := slices.Collect(func(yield func(chan string) bool) {
// 		for _, p := range players {
// 			if p != nil {
// 				if !yield(n) {
// 					return
// 				}
// 			}
// 		}
// 	})
// 	return nonNil
// }

func player() chan string {
	output := make(chan string)
	count := rand.Intn(100)
	move := []string{"UP", "DOWN", "LEFT", "RIGHT"}
	go func() {
		defer close(output)
		for i := 0; i < count; i++ {
			output <- move[rand.Intn(4)]
			d := time.Duration(rand.Intn(200))
			time.Sleep(d * time.Millisecond)
		}
	}()
	return output
}
