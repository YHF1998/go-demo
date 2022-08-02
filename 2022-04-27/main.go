package main

import (
	"fmt"
	"time"
)

func main()  {
	//打点器实现
	
	ticker := time.NewTicker(time.Millisecond * 500)
	
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	
	
	time.Sleep(time.Second * 5)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
