package main

import "fmt"

/*
题目1 ：编写一个Go程序，定义一个函数，
该函数接收一个整数指针作为参数，
在函数内部将该指针指向的值增加10，
然后在主函数中调用该函数并输出修改后的值。
*/
func pointAdd(numPtr *int) {
	*numPtr += 10
}

/*
题目 ：实现一个函数，
接收一个整数切片的指针，
将切片中的每个元素乘以2。
*/
func pointMultiply(numPtrs *[]int) {
	for i := 0; i < len(*numPtrs); i++ {
		(*numPtrs)[i] *= 2
	}
}

func main() {
	var numPtr *int = new(int)
	*numPtr = 5
	pointAdd(numPtr)
	fmt.Println(*numPtr)

	nums := []int{1, 2, 3, 4, 5}
	pointMultiply(&nums)
	fmt.Println(nums)

}
