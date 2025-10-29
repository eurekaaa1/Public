package main
import (
	"fmt"
	"sync"
)
//编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
//sync.Mutex 的使用、并发数据安全。

var G_sum = 0

func add(mu *sync.Mutex , sum *int)(ret int){
	for i := 0 ; i < 1000 ; i++{
		mu.Lock()
		(*sum)++
		mu.Unlock()
	}
	fmt.Printf("sum is:%d\n", *sum)
	return *sum
}

func gotestfunc(runtime int , wg *sync.WaitGroup , mu *sync.Mutex){
	
	wg.Add(runtime)
	for ; runtime > 0 ; runtime--{//外层循环十次 这样协程就不会并发的取次数导致错乱了
		go func()(int){//匿名函数执行协程 无输入参数 输出参数为 执行加法的结果
			defer wg.Done()
			ret := add(mu , &G_sum)
			return ret
		}()
		//wg.Wait()   //这里好像也行
	}

	wg.Wait()
	//return 

}

func main(){

	fmt.Println("begin test ->")
	var wg sync.WaitGroup
	var mu sync.Mutex
	
	gotestfunc(10 , &wg , &mu)

}

//done