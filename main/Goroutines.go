// A _goroutine_ is a lightweight thread of execution.

package main

import "fmt"

func F(from string) {
	for i := 0; i < 100; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	// Suppose we have a function call `f(s)`. Here's how
	// we'd call that in the usual way, running it
	// synchronously.
	//f("direct")

	// To invoke this function in a goroutine, use
	// `go f(s)`. This new goroutine will execute
	// concurrently with the calling one.
	go F("goroutine")

	// You can also start a goroutine for an anonymous
	// function call.
	go func(msg string) {
		//fmt.Println(msg)
		for i := 0; i < 100; i++ {
			fmt.Println(msg, ":", i)
		}
	}("going")

	// Our two function calls are running asynchronously in 就是会交错
	// separate goroutines now, so execution falls through
	// to here. This `Scanln` code requires we press a key
	// before the program exits.
	var input string//他这个程序如果没有fmt.Scanln(&input)的话也不会运行go的方法,不知道为啥有了fmt.Scanln(&input)就能运行
	fmt.Scanln(&input)
	fmt.Println("done")
}


