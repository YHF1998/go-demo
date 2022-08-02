package main

import "fmt"

func main() {
	
	num := []int{1, 1, 2, 3, 4, 4, 4, 4, 4, 5, 6, 7, 8, 9, 10}
	//removeDuplicates(&num)
	//removeDuplicates(num)
	//num =  num[:removeDuplicates(num)]
	num2 := make([]int, 5)
	copy(num2,num[:5])
	num2[2] = 100
	//removeElement(num,4)
	fmt.Println(num)
	fmt.Println(num2)
}


//func removeDuplicates(nums *[]int) int {
//	n := len(*nums)
//	if n == 0 {
//		return 0
//	}
//	slow := 1
//	for fast := 1; fast < n; fast++ {
//		if (*nums)[fast] != (*nums)[fast-1] {
//			(*nums)[slow] = (*nums)[fast]
//			slow++
//		}
//	}
//
//	(*nums) = (*nums)[:slow]
//	
//	return slow
//}

//func removeDuplicates(nums []int) int {
//	if len(nums)==0{
//		return 0
//	}
//	left,right := 1,1
//	for right < len(nums){
//		if nums[right] != nums[right-1]{
//			nums[left] = nums[right]
//			left++
//		}
//		right++
//	}
//	return left
//
//}

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	slow := 1
	for fast := 1; fast < n; fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

func removeElement(nums []int, val int) int {
	n := len(nums)
	if n== 0 {
		return 0
	}
	count := 0
	for i := 0; i < n; i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count++
		}
	}
	
	return count
}