package main

import (
	"math/rand"
	"sync"
	"time"
)

/*
编写一个程序，
使用 go 关键字启动两个协程，
一个协程打印从1到10的奇数，
另一个协程打印从2到10的偶数。
*/
func printOddEven() {
	var wg = sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			if i%2 != 0 {
				println("f1:", i)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 1; i < 10; i++ {
			if i%2 == 0 {
				println("f2:", i)
			}
		}
	}()
	wg.Wait()
}

/*
设计一个任务调度器，
接收一组任务（可以用函数表示），
并使用协程并发执行这些任务，
同时统计每个任务的执行时间。
*/
type TaskFunc func(num int)

// 任务调度器
func taskScheduler(tasks []TaskFunc) {
	var wg = sync.WaitGroup{}
	for index, task := range tasks {
		wg.Add(1)
		go func(t TaskFunc) {
			defer wg.Done()
			startTime := time.Now()
			t(index)
			endTime := time.Now()
			executionTime := endTime.Sub(startTime) / time.Second
			println("Task-", index, " spend time = ", executionTime)
		}(task)
	}
	wg.Wait()
}

// 模拟随机执行时间的任务
func randomDuration() {
	pauseDuration := time.Duration(rand.Intn(int(10))) * time.Second
	time.Sleep(pauseDuration)
}

func main() {
	//打印奇数偶数
	printOddEven()
	//批量执行任务
	taskScheduler([]TaskFunc{
		func(num int) {
			println("Task ", num, " is running")
			randomDuration()
		},
		func(num int) {
			println("Task ", num, " is running")
			randomDuration()
		},
		func(num int) {
			println("Task ", num, " is running")
			randomDuration()
		},
	})

}
