package main

import (
	"sync"
	"sync/atomic"
)

/*
题目 ：编写一个程序，
使用 sync.Mutex 来保护一个共享的计数器。
启动10个协程，
每个协程对计数器进行1000次递增操作，最后输出计数器的值。
*/
var counter int32
var mu = sync.Mutex{}

func lockCounter() {
	var wg = sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	println("lockCounter result =", counter)
}

func atomicCounter() {
	var wg = sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&counter, 1)
			}
		}()
	}
	wg.Wait()
	println("lockCounter result =", counter)
}

func main() {
	lockCounter()
	atomicCounter()
}
