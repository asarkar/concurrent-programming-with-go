package main

import (
	ex9_1 "concurrent-programming-with-go/ch09/ex9.1"
	ex9_2 "concurrent-programming-with-go/ch09/ex9.2"
	ex9_3 "concurrent-programming-with-go/ch09/ex9.3"
	ex9_4 "concurrent-programming-with-go/ch09/ex9.4"
)

/*
5. Connect the components developed in exercises 1 to 4 together in a main() function using the
following pseudocode:
Create quit channel
Drain(quitChannel,

	Print(quitChannel,
	    TakeUntil({ s <= 1000000 }, quitChannel,
	        GenerateSquares(quitChannel))))

Wait on quit channel
*/
func main() {
	quit := make(chan int)
	ex9_4.Drain(quit,
		ex9_3.Print(quit,
			ex9_2.TakeUntil(func(s int) bool { return s <= 1000000 }, quit,
				ex9_1.GenerateSquares(quit))))
	<-quit
}
