package main

import (
	"fmt"
	"sync"
)

//编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。

func print_odd_num(wg *sync.WaitGroup) {
	//var wg sync.WaitGroup
	defer wg.Done()
	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println("the odd num is:", i)
		}
	}

}

func print_even_num(wg *sync.WaitGroup) {
	//var wg sync.WaitGroup
	//这里之前的做法是错误的 新建的wg实例和之前的没关系，是一个局部变量 不是test_goroutine里面的，所以会上报计数器为负数的错误
	defer wg.Done()
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("the even num is:", i)
		}
	}

}

func test_goroutine() {

	var wg sync.WaitGroup
	wg.Add(2)

	go print_odd_num(&wg) //这里注意需要传入同一个 wg实例
	go print_even_num(&wg)

	wg.Wait()
}

func goruntest() {
	fmt.Println("*****begin test******")
	// print_odd_num()

	// print_even_num()

	test_goroutine()
}

func main() {
	fmt.Println("hello world")
	goruntest()
}

//done