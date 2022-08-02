package main

import (
	"fmt"
	"time"
)

func worker(id int, work <-chan int, result chan<- int) {
	for n := range work {
		fmt.Println("worker", id, "started", "num=", n)
		//time.Sleep(time.Millisecond * 100)
		time.Sleep(time.Second)
		result <- n * 2
	}
}

func main() {
	//协程的demo

	//建两个数据通道
	workChan := make(chan int, 100)
	resultChan := make(chan int, 100)

	//启动10个协程
	for i := 1; i <= 10; i++ {
		go worker(i, workChan, resultChan)
	}

	//发送数据
	for j := 1; j <= 90; j++ {
		workChan <- j
	}
	//关闭通道
	//close(workChan)
	

	//接收数据，打印结果
	// 最后，我们收集所有这些任务的返回值。
	count := 0
	check :=0
	TEST:
	for {
		select {
		case r := <-resultChan:
			fmt.Println("result:", r)
			count++
		case <-time.After(time.Second * 2):
			fmt.Println("好像没数据了",count)
			check++
			if check == 5 {
				break TEST
			}
		}
	}
	//for a := 1; a <= 9; a++ {
	//	fmt.Println(<-resultChan)
	//}
	//close(resultChan)
	fmt.Println(count)
	//time.Sleep(time.Second*5)
}
