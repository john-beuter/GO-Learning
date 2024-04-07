package main

/*


NOTES: Light weight threads. Using the syntax go function(x, y, z)
x, y, z are evaluated in the current routine but the execution of this
happens in the new routine (think new thread but smaller)


*/

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
