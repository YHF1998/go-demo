package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//str := "Hello,世界"
	str := "Hello1111111"
	//n := strings.IndexRune(str, rune('界'))
	//if strings.Contains(str, "世") {
	//	fmt.Println("包含")
	//}
	//fmt.Println(n)
	//fmt.Println([]rune(str)[n:])
	//fmt.Println(string([]byte(str)[n:]))
	fmt.Println(len([]rune(str)))
	fmt.Println(utf8.RuneCountInString(str))
}

//快速排序函数
func quickSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	left, right := 0, len(data)-1
	pivot := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] < pivot {
			left++
		} else {
			right--
		}
	}
	
	data[0], data[left] = data[left], data[0]
	return append(quickSort(data[:left]), append([]int{pivot}, quickSort(data[left+1:])...)...)
}

















