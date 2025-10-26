package main

import "fmt"

//实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2

// 方法1
func receive_slice_multi_two_1(int_slice *[]int) { //这里需要注意整数切片指针是 int_slice *[] int，指针跟在括号之后
	if (*int_slice) == nil { //有时会忘记空值处理 记得处理
		return
	}

	for i := 0; i < len(*int_slice); i++ {
		(*int_slice)[i] = (*int_slice)[i] * 2
	}
	//[] 运算符的优先级高于 *（解引用） ,所以对于切片的指针来说 ，先从切片地址解引用获得数字  再将数字乘法处理
}

// 方法2 另一种循环切片的方法
func receive_slice_multi_two_2(int_slice *[]int) {
	//init_num := 1
	for i := range *int_slice {
		(*int_slice)[i] = (*int_slice)[i] * 2
	}

}

func gotestfunc() {
	fmt.Println("before test")
	test_elem := []int{3, 7, 6, 9, 25}
	fmt.Println(test_elem)

	receive_slice_multi_two_2(&test_elem)
	fmt.Println("after test")
	fmt.Println(test_elem)

}

func main() {
	fmt.Println("******Begin test******")
	gotestfunc()

}
