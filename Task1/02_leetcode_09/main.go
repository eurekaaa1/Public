package main

import "fmt"

func isPalindrome(x int) bool { //想法 给数字取到 之后异或

	restore := x
	rev_num := 0

	for x > 0 {
		last_num := x % 10

		rev_num = rev_num*10 + last_num

		x = x / 10
	}

	// if rev_num ^ restore == 0{
	// 	return  true
	// }

	// return false

	return rev_num^restore == 0 //这样写更好

}

func gotestfunc() {
	num := 123321
	ret := isPalindrome(num)
	fmt.Println(ret)
}

func main() {
	fmt.Println("hello world")
	gotestfunc()
}