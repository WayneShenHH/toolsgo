package services

import (
	"fmt"
	"time"

	"github.com/WayneShenHH/toolsgo/tools"
)

// GoroutineWithWaitGroup example for goroutine with wait group
func GoroutineWithWaitGroup() {
	tools.MultiStartWithGroup(5, func(i int) {
		time.Sleep(2)
		fmt.Println(i, ":goroutine")
	})
}

// GoroutineExample example for goroutine
func GoroutineExample() {
	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.
	f("direct")

	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.
	go f("goroutine")

	// You can also start a goroutine for an anonymous
	// function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in
	// separate goroutines now, so execution falls through
	// to here. This `Scanln` code requires we press a key
	// before the program exits.
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}
