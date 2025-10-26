package main
import "fmt"


func singleNumber(nums []int) int {
	Ansmap := make(map[int]int)

	for _ , value_of_nums := range nums{
		Ansmap[value_of_nums]++//map 中的 key 就是 数组的具体的某一个位置的值，map的 value是出现的次数
	}

	for key_of_map , value_of_map := range Ansmap{
		if value_of_map == 1{
			return key_of_map
		}
	}

	return 0
}

func gotestfunc(){

	arry := [] int {3, 5, 3, 6, 6, 1, 1, 10, 10, 11, 5}
	ret := singleNumber(arry)
	fmt.Println(ret)
}

func main(){
	fmt.Println("hello world")

	gotestfunc()
}