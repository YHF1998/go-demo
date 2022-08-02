package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Map

func main() {

	for i := 0; i < 5; i++ {
		go func(i int) {
			for j := 0; j < 3; j++ {
				//查询是否存在map
				for{
					_,ok:=m.Load(i)
					if !ok {
						break
					}
					fmt.Println(i,"有协程在操作，休眠")
					time.Sleep(time.Second*1)
			
				}
				
				m.Store(i,1)
				fmt.Println(i)
				time.Sleep(time.Second*10)
				m.Delete(i)
			}
			
		}(i)
	}
	
	time.Sleep(time.Minute)
}