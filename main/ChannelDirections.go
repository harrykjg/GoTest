package main


// When using channels as function parameters, you can
// specify if a channel is meant to only send or receive
// values. This specificity increases the type-safety of
// the program.

import (
	"fmt"

	"time"
)

// This `ping` function only accepts a channel for sending
// values. It would be a compile-time error to try to
// receive on this channel.
func ping(pings chan <- string, msg string) {

	pings <- msg
	fmt.Println("sent ping")
}

// The `pong` function accepts one channel for receives
// (`pings`) and a second for sends (`pongs`).
func pong(pings <-chan string, pongs chan <- string) {//就是信息从ping这个方法的msg传给pings(但这个pings居然叫用来send的channel),
	msg := <-pings             //然后pings作为参数传给pong这个方法再把pings带的value传给pongs这个channel.注意pings本来是出和入
	fmt.Println("something")   //都行了,只是在ping这个方法里规定成为用来send的而在pong这个方法里规定成为receive的
	pongs <- msg             //这里pongs就是接受了一个msg然后就啥都没干了,如果pongs这个channel没声明成大小的话就错了
}

func main() {
	pings := make(chan string,1)//注意这里channel设了大小为1,如果这两个channel不设大小的话那么这个main会运行不起来被block,除非用go ping gopong
	pongs := make(chan string,1)
	ping(pings, "passed message") //注意这个和channel.go不一样,那里是用了go func 接受一个channel作为输入参数,channel在go func里

	time.Sleep(time.Second*3)

	pong(pings,pongs)

}

