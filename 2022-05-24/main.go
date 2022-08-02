package main

import "fmt"

type fc func()
func main(){
	a := map[string]fc{
		"b":func(){
			fmt.Println("b")
		},
	}
	for _, fn := range a {
		fn();
	}
	c := map[string]interface{}{
		"d":func(){
			fmt.Println("d")
		},
	}

	c["d"].(func())()
	//a["b"]()
	
	
}

