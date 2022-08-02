package main

import (
	"chatbox/zdbmodel"
	"chatbox/zlogs"
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

var list map[int]map[*websocket.Conn]*websocket.Conn
var ch map[int]chan []byte
var msg map[int][][]byte
var mu sync.Mutex

func svrConnHandler(conn *websocket.Conn) {

	r := conn.Request()
	r.ParseForm()
	sgroupid := r.Form["groupid"][0]
	suserid := r.Form["userid"][0]
	name := r.Form["name"][0]

	groupid, _ := strconv.Atoi(sgroupid)
	userid, _ := strconv.Atoi(suserid)

	_, ok := list[groupid]

	if !ok {
		list[groupid] = make(map[*websocket.Conn]*websocket.Conn, 500)
		ch[groupid] = make(chan []byte, 500)
		msg[groupid] = make([][]byte, 0, 500)
	}

	//TODO :  获取历史数据
	go func() {
		mu.Lock()
		list[groupid][conn] = conn
		for _, v := range msg[groupid] {
			conn.Write(v)
		}
		mu.Unlock()
	}()

	//TODO : 接收客户端消息
	go func() {
		var bb [1024]byte
		var content []byte
		for {
			n, err := conn.Read(bb[:])

			if err != nil {
				delete(list[groupid], conn)
				conn.Close()
				zlogs.ErrorLog(err)
				break
			}
			content = append(content, bb[:n]...)
			content2 := []byte(fmt.Sprint(userid, ":::", name, ":::"))
			content3 := append(content2, content...)

			if n < 4096 {
				ch[groupid] <- content3
				content = []byte{}
			}
		}
	}()

	//TODO : 向客户端发送消息
	if !ok {
		for {
			select {
			case content := <-ch[groupid]:
				mu.Lock()

				dd := strings.Split(string(content), ":::")

				msg[groupid] = append(msg[groupid], []byte("<span class='name'>"+dd[1]+"："+"</span>"+dd[2]))
				for _, v := range list[groupid] {
					v.Write([]byte("<span class='name'>" + dd[1] + "：" + "</span>" + dd[2]))
				}
				mu.Unlock()
				go func() {
					nn, _ := strconv.Atoi(dd[0])
					zdbmodel.InsertMessage(groupid, nn, dd[1], dd[2])
				}()
			}
		} //<--向客户端发送消息-->
	} else {
		for {
			time.Sleep(time.Second * 60 * 60)
		}
	}
}

func main() {
	list = make(map[int]map[*websocket.Conn]*websocket.Conn, 10)
	ch = make(map[int]chan []byte, 10)
	msg = make(map[int][][]byte)

	zlogs.InitLog()
	zdbmodel.InitDb()

	http.Handle("/message", websocket.Handler(svrConnHandler))
	err := http.ListenAndServe(":8095", nil)
	zlogs.ErrorLog(err)
}
