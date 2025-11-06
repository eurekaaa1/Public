package main

import "fmt"

// func plusOne(digits []int) []int {
// 	length := len(digits)

// 	if length == 0 { //空的
// 		return []int{1}
// 	} else if length == 1 { //不超过10的
// 		res_num := digits[0] + 1
// 		return []int{res_num}
// 	} else {
// 		if digits[length-1] != 9 {
// 			num := digits[length-1] + 1
// 			digits = digits[:length-1]
// 			digits = append(digits, num)
// 			return digits
// 		} else {
// 			for digits[length-1] == 9 && digits[0] != 9 {
// 				digits = digits[:length-1]
// 				digits = append(digits, 0)
// 				length--
// 			}
// 			the_num := digits[0] + 1
// 			digits[0] = the_num
// 			return digits
// 		}
// 	}

// }

func plusOne(digits []int) []int {
    length := len(digits)
    for i := length - 1; i >= 0; i-- {
        if digits[i] != 9 {
            digits[i]++
            for index := i + 1; index < length; index++ {
                digits[index] = 0
            }
            return digits
        }
    }
    // digits 中所有的元素均为 9

    digits = make([]int, length+1)
    digits[0] = 1
    return digits
}

func main() {
	fmt.Println("begin test")
	num := []int{ 4,9,9,9}
	//num := []int{ 3,2,1}
	ret := plusOne(num)
	fmt.Println(ret)

}
