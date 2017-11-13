// Go's _select_ lets you wait on multiple channel
// operations. Combining goroutines and channels with
// select is a powerful feature of Go.

package main

import "time"
import "fmt"

func main() {

	// For our example we'll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.
	go func() {
		time.Sleep(time.Second * 1 )//这两个go func就是多线程同时启动
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {    //他实际上就是等这个select里哪个channel先来值就把它打印出来然后就退出select了,然后去下一层for循环又开始等select里谁来就打印谁
		case msg1 := <-c1:   //,因此需要for循环运行2次.如果是下面这个注释的写法(不用select的话),就是先运行msg1 := <-c1,而这个会被block知道c1
			fmt.Println("received", msg1) //来了msg,然后程序才会往下走到msg2 := <-c2. 比如c1 要sleep5秒而c2要sleep1秒,则还是先
		case msg2 := <-c2:    //得等5秒之后c1才会打印出来,尽管c2早就来了但还是得等c1来了之后才会运行c2. 而如果是c1 sleep1秒而c2 sleep3秒,
			fmt.Println("received", msg2) //则c1来了之后2秒后c2也会打印出来
		}
		//time.Sleep(time.Millisecond * 1)
		//msg1 := <-c1
		//fmt.Println("received", msg1)
		//msg2 := <-c2
		//fmt.Println("received", msg2)

	}
}

