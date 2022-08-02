package main

import (
	"fmt"
	"sync"
	"time"
)

var m sync.Map

func main() {

	//userid
	for userid := 1; userid < 3; userid++ {
		go func(userid int) {
			//recordId 记录id
			for recordId := 1; recordId < 5; recordId++ {
				go action(userid, recordId)
			}
		}(userid)
	}

	time.Sleep(time.Second * 30)
}

//
//  action
//  @Description: 操作记录
//  @param recordId
//
func action(userid, recordId int) {
	//查询是否存在map
	//同一时间，一条记录只能只能有一个用户操作
	for {
		_, ok := m.LoadOrStore(recordId, 1)
		if !ok {
			break
		}
		fmt.Println(userid, "==", recordId, "有协程在操作，休眠")
		time.Sleep(time.Second * 1)

	}

	fmt.Println(userid, "==", recordId)
	time.Sleep(time.Second * 5)
	m.Delete(recordId)
}
