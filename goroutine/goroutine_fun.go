package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	flag = false
	str  string
)

func foo() {
	flag = true
	str = "setup complete!"
}

func Procs() {
	//单核执行如果for前面或者中间不延迟，主线程不会让出CPU，导致异步的线程无法执行，从而无法设置flag的值，从而出现死循环。
	runtime.GOMAXPROCS(2)
	go foo()
	for {
		if flag {
			break
		}
	}
	fmt.Println(str)
}

func WaitGroup() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calc(&wg, i)
	}

	wg.Wait()
	fmt.Println("all goroutine finish")
}

func calc(w *sync.WaitGroup, i int) {

	fmt.Println("calc:", i)
	time.Sleep(time.Second)
	w.Done()
}

func ChanPooled() {
	data := make(chan int, 3) // 缓冲区可以存储 3 个元素
	exit := make(chan bool)

	data <- 1 // 在缓冲区未满前，不会阻塞。
	data <- 2
	data <- 3

	go func() {
		for d := range data { // 在缓冲区未空前，不会阻塞。
			fmt.Println(d)
		}

		exit <- true
	}()

	data <- 4 // 如果缓冲区已满，阻塞。
	data <- 5
	close(data)

	<-exit
}

func ChanNotPooled() {
	data := make(chan int)  // 数据交换队列
	exit := make(chan bool) // 退出通知

	go func() {
		for d := range data { // 从队列迭代接收数据，直到 close 。
			fmt.Println(d)
		}

		fmt.Println("recv over.")
		exit <- true // 发出退出通知。
	}()

	data <- 1 // 发送数据。
	data <- 2
	data <- 3
	close(data) // 关闭队列。
	fmt.Println("send over.")
	<-exit // 等待退出通知。
}

func Timeout() {
	w := make(chan bool)
	c := make(chan int, 2)

	go func() {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-time.After(time.Second * 3):
			fmt.Println("timeout.")
		}

		w <- true
	}()

	// c <- 1 // 注释掉，引发 timeout。
	<-w
}

func main() {
	// Procs()
	// WaitGroup()
	// ChanPooled()
	// ChanNotPooled()
	// Timeout()
}
