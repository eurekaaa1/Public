package main
import (
	"fmt"
	"sync"
	//"time"
)
//实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
//考察点 ：通道的缓冲机制

func producer(ch_ele chan int, wg *sync.WaitGroup){

	go func() {
		defer wg.Done()
		defer close(ch_ele) //协程结束之后关闭通道
		for i := 0; i < 100; i++ {
			ch_ele <- i
			// fmt.Printf("[Producer] Sent: %d, Buffer remaining: %d\n", i, len(ch_ele))
			// time.Sleep(10 * time.Millisecond) // 生产者稍慢，方便观察缓冲区变化
		}
	}()

}

// func producer(ch_ele chan int, wg *sync.WaitGroup) {
//     defer wg.Done() // 直接在当前协程中处理，无需嵌套
//     for i := 0; i < 100; i++ {
//         ch_ele <- i
//     }
//     close(ch_ele) // 发送完成后关闭通道
// }

func consumer(ch_ele chan int, wg *sync.WaitGroup){
	go func() {
		defer wg.Done()
		for i := range ch_ele{// for range语法专门遍历通道的语法
			//i := <-ch_ele

			fmt.Println("num is:" , i)
			// fmt.Printf("[Consumer] Received: %d, Buffer remaining: %d\n", i, len(ch_ele))
			// time.Sleep(50 * time.Millisecond)
		}
		
	}()
}

func gotestfunc(){

	var ch chan int = make(chan int, 10)
	var wg sync.WaitGroup

	wg.Add(2)

	producer(ch , &wg)

	consumer(ch , &wg)
	//go consumer(ch, &wg)

	wg.Wait()
	fmt.Println("All tasks completed")

}

func main(){
	
	fmt.Println("begin test 2")

	gotestfunc()
}