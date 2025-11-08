package main

import (
	"sync"
)

/*
题目 ：编写一个程序，
使用通道实现两个协程之间的通信。
一个协程生成从1到10的整数，
并将这些整数发送到通道中，
另一个协程从通道中接收这些整数并打印出来。
*/
func channel1() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	//创建通道
	ch := make(chan int)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()
	go func() {
		defer close(ch)
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			num := <-ch
			println("接收到的数字 =", num)
		}
	}()
	//等待协程执行完毕
	wg.Wait()
}

/*
题目:实现一个带有缓冲的通道，
生产者协程向通道中发送100个整数，
消费者协程从通道中接收这些整数并打印。
*/
func channel2() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	ch := make(chan int, 3)

	go func() {
		defer wg.Done()
		defer close(ch)
		for i := 1; i <= 100; i++ {
			ch <- i
			println("product =", i)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			num, ok := <-ch
			if !ok {
				break
			}
			println("consume =", num)
		}
	}()
	wg.Wait()

}

func main() {
	// channel打印数字
	channel1()
	//带缓冲区channel
	channel2()

}
