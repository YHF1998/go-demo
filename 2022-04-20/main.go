package main

import "fmt"

//import "fmt"

type  AAAA map[string]interface{}

func main() {
	 map1 := AAAA{"name": "yhf", "age": "18"}
	 
	 
	 slice1 := []string{"name", "age"}
	
	
	tp := fmt.Sprintf("%T", map1)
	tp1 := fmt.Sprintf("%T", slice1)
	
	fmt.Println(tp=="map[string]string")
	fmt.Println(tp=="main.AAAA")
	fmt.Println(tp=="map[string]interface {}")
	fmt.Println(tp1=="[]string")
	
}