// By default channels are _unbuffered_, meaning that they
// will only accept sends (`chan <-`) if there is a
// corresponding receive (`<- chan`) ready to receive the
// sent value. 这里是关键,说的是没有指定大小的channel的话,如果没有receieve chan ready的话,程序运行到send那里就阻塞了,连send都send不出去,就会error
// _Buffered channels_ accept a limited
// number of  values without a corresponding receiver for
// those values.  有了buffer之后send就不会阻塞了

package main

import "fmt"

func main() {

	// Here we `make` a channel of strings buffering up to
	// 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these
	// values into the channel without a corresponding
	// concurrent receive.
	messages <- "buffered"
	messages <- "channel" //如果再send一个就会报错,因为channel里装不下三个,但是如果装了两个再receive有一个send一个也行,是先进先出

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

