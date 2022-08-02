package main

import (
	"fmt"
	"reflect"
)

type MessageBox struct {
	Name    string `json:"JName" tag1:"TagName"` //tag是结构体的原信息, 可以在运行时通过反射的机制读取出来
	title   string `json:"JTitle" tag1:"TagTitle"`
	Age     int    `json:"JAge" tag1:"TagAge"`
	Content string `json:"JContent" tag1:"TagContent"`
}

func main() {
	//mbox := MessageBox{Name: "张三", title: "Hi", Age: 18, Content: "Hello World"}
	//fmt.Println("-----TODO1-----")
	//mbox.reflectMsg()
	//fmt.Println("-----TODO2-----")
	//mbox.reflectTag("tag1")
	//fmt.Println("-----TODO3-----")
	//mbox.reflectMethod()

	mbox := MessageBox{Name: "张三", title: "Hi", Age: 18, Content: "Hello World"}
	m := reflect.ValueOf(&mbox)
	f := m.Elem().FieldByName("title")
	f.SetString("就这？")
	fmt.Println(mbox)
	fmt.Println(f)
}




func (box MessageBox) Tag(name string) {
	reType := reflect.TypeOf(box)
	num := reType.NumField()

	for i := 0; i < num; i++ {
		fmt.Println(reType.Field(i).Tag.Get(name))
	}
}
func (box *MessageBox) reflectMsg() {
	reType := reflect.TypeOf(*box)
	reValue := reflect.ValueOf(*box)
	num := reType.NumField()

	for i := 0; i < num; i++ {
		fmt.Println(reType.Field(i), "---", reValue.Field(i))
	}
}

func (box *MessageBox) reflectTag(name string) {
	reType := reflect.TypeOf(box).Elem()
	num := reType.NumField()

	for i := 0; i < num; i++ {
		fmt.Println(reType.Field(i).Tag.Get(name))
	}
}

func (box *MessageBox) reflectMethod() {
	reValue := reflect.ValueOf(box).Elem()
	reMethod := reValue.MethodByName("Tag") //T类型的方法
	args := []reflect.Value{reflect.ValueOf("json")}
	//fmt.Println(reValue)
	reMethod.Call(args)
}

