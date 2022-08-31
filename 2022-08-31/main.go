package main

import (
	"fmt"
)

type aaa func() string
func main()  {

	var h aaa = hello;
	//var string h = hello();
	fmt.Println(h())
}


func hello() string {
	return "hello"
}