package main

import "fmt"

func removeDuplicates(nums []int) int {
	length := len(nums)

	if length == 0 {
		return 0
	}

	slow := 0
	for fast:= 1 ; fast < length; fast++{//快指针
		
		if nums[fast] != nums[slow] {//如果这俩数字指向的地方不一样了 就把快指针的数放前面
			nums[slow + 1] = nums[fast]
			slow++
		}
	}
	return (slow + 1)
}

func main(){
	fmt.Println("begin test")
	nums := []int {0,0,1,1,1,2,2,3,3,4}
	ret := removeDuplicates(nums)
	fmt.Println(ret)
}
