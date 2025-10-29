package main

import (
	"fmt"
	"sync/atomic"
	"sync"
)

//使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//原子操作、并发数据安全。

var G_num uint64

//尝试两种方法 1.全局变量 2.传参

func go_counter(add_times int64) {

	for add_times > 0 {
		atomic.AddUint64(&G_num, 1)     //原子的 给需要的数字加一
		atomic.AddInt64(&add_times, -1) //循环中加完一 之后给计数的数字 减一
	}

}

func gotestfunc() {

	fmt.Printf("before num is:%d\n" , G_num)
	
	var dotime int64 = 1000//循环次数

	var wg sync.WaitGroup
	goroutine_num := 10 //创建协程的个数 防止“魔法数字的出现”
	wg.Add(goroutine_num)
	

	for i := 0; i < goroutine_num; i++ {
		fmt.Println("this is time:", i)
		go func() {
			defer wg.Done()
			go_counter(dotime)
		}()
	}

	wg.Wait()

	fmt.Printf("finally num is:%d\n" , G_num)

}

func main() {
	fmt.Println("begin test:")
	gotestfunc()

	//fmt.Printf("finally num is:%d\n" , G_num)

}
//done