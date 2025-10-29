package main

import (
	"fmt"
	"sync"
	"time"
)

//设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。

func add_ten_times(receive_num *int){

	for i := 0; i < 10; i++ {
		(*receive_num)++
		fmt.Printf("receive_num is:%d\n", (*receive_num))
	}
}

func dotask_func(){
	var wg sync.WaitGroup
	wg.Add(3)

	give_num := 0
	go func(*int) {
		defer wg.Done()
		start := time.Now()
		add_ten_times(&give_num)
		endtime := time.Since(start)
		fmt.Printf("task1 use time%v\n" , endtime)
	}((&give_num))

	go func(*int) {
		defer wg.Done()
		start := time.Now()
		add_ten_times(&give_num)
		endtime := time.Since(start)
		fmt.Printf("task2 use time%v\n" , endtime)
	}((&give_num))

	go func(*int) {
		defer wg.Done()
		start := time.Now()
		add_ten_times(&give_num)
		endtime := time.Since(start)
		fmt.Printf("task3 use time%v\n" , endtime)
	}((&give_num))

	wg.Wait()

}

func gotestfunc(){
	dotask_func()
}

func main(){
	fmt.Println("begin test->")
	gotestfunc()
}

//done