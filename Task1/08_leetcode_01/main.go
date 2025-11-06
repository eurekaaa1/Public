package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int { //暴力双循环直接求解 类似双指针吧 一个指向第一个 另一个指向后一个
	length := len(nums)
	result := []int{}
	for slow := 0; slow < length; slow++ {
		for fast := slow + 1; fast < length; fast++ {
			if nums[fast]+nums[slow] != target {
				//todo
			} else {
				result = append(result, fast)
				result = append(result, slow)
				return result
			}
		}
	}
	return result
}

func main() {
	fmt.Println("begin test")
	//num := []int{2, 7, 11, 15}

	// num := []int{2, 5, 5, 11}
	// target_num := 10

	num := []int{3, 2, 3}
	target_num := 6
	ret := twoSum(num, target_num)
	fmt.Println(ret)
}
