

package main

import "fmt"
import "time"

//Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and
//receive those values into another goroutine.

// This is the function we'll run in a goroutine. The
// `done` channel will be used to notify another
// goroutine that this function's work is done.
func worker(done chan bool) {//这个不是bool也行,任意一个类型,只要随便传一个值给他就行了,
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	// Send a value to notify that we're done.
	<-done//true和false都行,把这个值发给done这个channel就相当于send了,这里channel就是相当于一个标志,说明send准备好了,
	//下面main方法的<-done就是receive,例子的注释说了默认时只有send和receive都ready的时候这个worker方法才执行.即这个方法他接受这个channel类型的参数
	//就只是为了让他同步(这个方法会一直block知道receive ready时候)
	//这里有两个关键点,1,这个chan是不是buffer的,2,worker里的chan是recevier还是sender. 如果chan有buffer的话,并且woker里是sender那么,done<-false
	//就不会被block并且hhhh也会马上被print出来;如果chan有buffer但worker里是receiver的话则hhh只有在5秒后<-done接受到信息之后才print出来(妈的只有改过
	// 代码之后的第一次会print hhhh,再运行就不print hhhh了不知道什么问题)
	//如果没buffer,并且worker里的done是receiver或者sender都不会print hhhh(不知道为啥)
	fmt.Println("hhhhh")

}

func main() {

	// Start a worker goroutine, giving it the channel to
	// notify on.
	done := make(chan bool)
	//sb := make(chan bool, 1)
	go worker(done)
        fmt.Println("do something before worker")
	// Block until we receive a notification from the
	// worker on the channel.
	time.Sleep(time.Second*5)
	done<-false
	fmt.Println("fwq")

}

