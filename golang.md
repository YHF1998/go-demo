# 协程下的并发锁

## 针对同一个ID数据更改的并发限制

```
1.使用sync包的map来做同id并发限制处理
2.使用LoadOrStore添加值，如果已经存在的键，会返回false，这时当前goruntine应该进入休眠状态，等待其它goruntine对该recordid操作完成后，继续执行
3.在goruntine对recordid操作完成后，需要调用Delete方法释放这个recordid，让其它goruntine知道这个记录可以操作了
```



```go
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

```

