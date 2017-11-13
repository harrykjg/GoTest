package main
import "fmt"

// We can use channels to synchronize execution
// across goroutines. Here's an example of using a
// blocking receive to wait for a goroutine to finish.



func main() {

	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.
	messages := make(chan string)



	// _Send_ a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	go func(){//如果这个输入channel的不放在go func里的话就会error,解释见http://www.tuicool.com/articles/vAZvYr的2的例子,而
		messages <- "ping" //后面这个接受的channel就不用放在go func里,但是!我放进go func里反而不行了!
	}()

	go func(){
		msg := <-messages
		fmt.Println(msg)
	}()

	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.

	fmt.Println("xxxx")

}
