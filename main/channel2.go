package main

import "strconv"
import (
	"fmt"
	"time"
)

func main() {
	taskChan := make(chan string, 3)
	doneChan := make(chan int,1)

	for i := 0; i < 3; i++ {
		taskChan <- strconv.Itoa(i)
		fmt.Println("send: ", i)
	}

	go func() {
		for i := 0; i < 3; i++ {
			task := <-taskChan
			fmt.Println("received: ", task)
			time.Sleep(time.Second*2)
		}
		doneChan <- 1//如果这不用doneChan的话,那么这个这个taskChan就接受不到msg,这个doneChan给不给他一个容量都可以
	}()

	<-doneChan
}
