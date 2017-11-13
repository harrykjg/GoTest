// _Closing_ a channel indicates that no more values
// will be sent on it. This can be useful to communicate
// completion to the channel's receivers.

package main

import "fmt"
import "time"

// In this example we'll use a `jobs` channel to
// communicate work to be done from the `main()` goroutine
// to a worker goroutine. When we have no more jobs for
// the worker we'll `close` the `jobs` channel.
func main() {
	jobs := make(chan int, 5)
	done := make(chan string)

	// Here's the worker goroutine. It repeatedly receives
	// from `jobs` with `j, more := <-jobs`. In this
	// special 2-value form of receive, the `more` value
	// will be `false` if `jobs` has been `close`d and all
	// values in the channel have already been received.
	// We use this to notify on `done` when we've worked
	// all our jobs.
	go func(msg string) {                      //go func这个方法里边要有channel来sent或receive,否则这个方法就根本不执行,如果fmt.Println(<-done)
		//这句放在go func里的任何位置,这个go根本不进来也不会报错

		fmt.Println(msg)
		//for {
		//	j, more := <-jobs
		//	if more {
		//		fmt.Println("received job", j)
		//	} else {
		//		fmt.Println("received all jobs")
		//		done <- true
		//		return
		//	}
		//}

		//程序运行到后面的fmt.Println(<-done)的时候,才进入这个go func,然后先执行pring xxx,此时
		//fmt.Println(<-done)还没有执行,睡了2秒之后sb才才打印出来,但是不知道channel在print xxx的时候里面已经有了sb
		time.Sleep(time.Second * 2 )
		//如果sleep放到done<-"sb"之后的话,则所有东西,sent job1,2,3和xxx和sb会同时打印出来
		done<-"sb"

	}("ccc")

	// This sends 3 jobs to the worker over the `jobs`
	// channel, then closes it.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the
	// [synchronization](channel-synchronization) approach
	// we saw earlier.
	fmt.Println(<-done)
	fmt.Println("after")
}

