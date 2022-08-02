package main

import (
	"fmt"
	"math"
)

func main()  {
	num := divide(-2147483648,-1)
	fmt.Println(num)
	
	fmt.Println((-2147483648/-1))
	
	//inta := int32(math.Pow(-2,31))
	checkMax := 1 << 31
	fmt.Println(checkMax)
	
}

func strStr(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return -1
	}
	
	hn := len(haystack)
	
	haystack1 := []rune(haystack)
	
	for i := 0; i <= hn - n; i++ {
		if string(haystack1[i:i+n]) == needle {
			return i
		}
		
	}

	return -1
	
}


func divide(dividend int, divisor int) int {
	n := 0
	if divisor == 0 {
		return n
	}

	//设定边界值
	checkMin := -1 << 31
	checkMax := 1 << 31
	if dividend == checkMin && divisor == -1 {
		return checkMax - 1
	}
	
	
	symbol := 1
	if dividend < 0 && divisor > 0 || dividend > 0 && divisor < 0 {
		symbol = -1
	}
	
	dividend = int(math.Abs(float64(dividend)))
	divisor = int(math.Abs(float64(divisor)))

	
	
	
	for  {
		if dividend < divisor {
			break
		}
		dividend -= divisor
		n++
	}

	if symbol == -1 {
		return -n
	}
		
	return n
}
