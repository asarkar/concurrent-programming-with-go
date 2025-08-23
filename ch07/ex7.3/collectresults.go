package main

import (
	"fmt"
	"math/rand"
)

/*
3. In listing 7.13, we use a child goroutine to calculate the factors of one number and the main()
goroutine to work out the factors of the other. Modify this listing so that, using multiple
goroutines, we collect the factors of 10 random numbers.
*/
func main() {
	factors := make(chan []int, 10)
	for i := 1; i <= 10; i++ {
		n := rand.Intn(100)
		go findFactors(n, factors)
	}
	for i := 1; i <= 10; i++ {
		res := <-factors
		fmt.Printf("%2d : %v\n", res[len(res)-1], res)
	}
}

func findFactors(number int, factors chan []int) {
	var res []int
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			res = append(res, i)
		}
	}
	factors <- res
}
