package main



//waitgroup 是用来等待所有go routines运行完再取运行某些代码, 否则主线程再go routine运行完之前就退出了.channel也能达到这个效果
//http://www.liguosong.com/2014/05/06/golang-sync-waitgroup/
//http://studygolang.com/articles/2027
import(
	"fmt"
	//"time"
	"time"
	//"sync"
)

//func main(){ //这段代码也可以用
//	//ch:=make(chan int)
//	go func(){//假如我有一个taks要启动线程,会运行3秒才完成,如果没有channel的话,主线程进入启动了这个routine之后,这个线程再运作3秒,
//		//但没等3秒主线程就取执行println("x")然后就完事退出了,那么finish就永远不会打印出来了,所以我们可以建一个channel来,等
//		//print finish之后再传入信号,在main之后接受这个信号,那么channel就能阻塞main线程的print x了
//		//不用channel的话也可以用waitgroup来阻塞
//		time.Sleep(3e9)
//		fmt.Println("finish")
//		//ch<-2
//	}()
//	//fmt.Println(<-ch)
//	fmt.Println("x")
//}

//func main(){
//	var wg sync.WaitGroup
//	wg.Add(1)//add 1 代表这个wg要等待1个线程完成之后wg.wait才会解除阻塞,go func里的wg.Done就想动与wg.Add(-1).当wg里面是0的时候才解除阻塞
//	go func(){
//		//不用channel的话也可以用waitgroup来阻塞
//		time.Sleep(3e9)
//		fmt.Println("finish1")
//		wg.Done()
//
//	}()
// 	wg.Add(1)//再启动一个线程,wg加一,不写的话wg就只要等待某一个线程调用wg.Done就会变成0了,wg就不阻塞了,导致另一个运行时间长的线程运行不完,main就退出了
//	go func(){
//		//不用channel的话也可以用waitgroup来阻塞
//		time.Sleep(2e9)
//		fmt.Println("finish2")
//		wg.Done()
//
//	}()
//
//	wg.Wait()
//	fmt.Println("x")
//}

//如果用channel也想实现等待多个线程运行完之后才执行主线程的话,可以用select
//https://gobyexample.com/select

func main() {  //他这个例子就是select里面列出所有要等的线程,谁先来就解除阻塞谁,要全部都解除阻塞之后才会运行后面的

	// For our example we'll select across two channels.
	c1 := make(chan string)
	c2 := make(chan string)

	// Each channel will receive a value after some amount
	// of time, to simulate e.g. blocking RPC operations
	// executing in concurrent goroutines.
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// We'll use `select` to await both of these values
	// simultaneously, printing each one as it arrives.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	fmt.Println("x")
}
