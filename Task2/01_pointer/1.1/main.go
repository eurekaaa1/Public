package main
import "fmt"
//编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。

func addfunc(addnum *int)int{
	after_add := *addnum + 10
	return after_add
}

func gotestfunc(){

	test_elem := 888
	ret := addfunc(&test_elem)
	fmt.Println(ret)
}

func main(){
	fmt.Printf("******Begin test******\n")

	gotestfunc()
}