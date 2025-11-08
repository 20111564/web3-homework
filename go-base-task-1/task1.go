package main

import (
	"fmt"
	"strconv"
)

// 只出现一次的数字：
// https://leetcode.cn/problems/single-number/
func singleNumber(nums []int) int {
	numCountMap := make(map[int]int)
	for _, v := range nums {
		count, flag := numCountMap[v]
		if flag {
			numCountMap[v] = count + 1
		} else {
			numCountMap[v] = 1
		}
	}

	for key, value := range numCountMap {
		if value == 1 {
			return key
		}
	}
	return 0
}

// 判断一个整数是否是回文数
// https://leetcode.cn/problems/palindrome-number/
func isPalindrome(num int) bool {
	numStr := strconv.Itoa(num)
	numlen := len(numStr)
	for i := 0; i < numlen/2; i++ {
		if numStr[i] != numStr[numlen-1-i] {
			return false
		}
	}
	return true
}

// 有效的括号
// https://leetcode-cn.com/problems/valid-parentheses/
func isValid(s string) bool {
	symbolStack := []string{}
	symbolStackMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	for _, value := range s {
		if string(value) == "(" || string(value) == "{" || string(value) == "[" {
			symbolStack = append(symbolStack, string(value))
		}
		if string(value) == ")" || string(value) == "]" || string(value) == "}" {
			//注意只输入右边的符号 ")]}"
			if len(symbolStack) > 0 && symbolStack[len(symbolStack)-1] == symbolStackMap[string(value)] {
				symbolStack = symbolStack[:len(symbolStack)-1]
			} else {
				return false
			}
		}

	}
	return len(symbolStack) == 0
}

// 最长公共前缀
// https://leetcode.cn/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var longestStr string = ""

	for index, value := range strs[0] {
		for _, value2 := range strs {
			//前面字符串比后面的长
			if len(value2) < (index+1) || string(value2[index]) != string(value) {
				goto Jump
			}
		}
		longestStr = longestStr + string(value)
	}
Jump:
	return longestStr
}

// 加一
// https://leetcode.cn/problems/plus-one/description/
func plusOne(digits []int) []int {
	inversionArr := inversionArray(digits)
	var flag = true
	for index, value := range inversionArr {
		if flag {
			//超过10，变成0
			if value+1 == 10 {
				inversionArr[index] = 0
			} else {
				inversionArr[index] = value + 1
				flag = false
			}
		}
	}
	//最后一位是0，则需要补充一位数
	if inversionArr[len(inversionArr)-1] == 0 {
		inversionArr = append(inversionArr, 1)
	}
	return inversionArray(inversionArr)
}

// 数组反转
func inversionArray(arr []int) []int {
	if len(arr) <= 0 {
		return arr
	}
	var result = []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		result = append(result, arr[i])
	}
	return result
}

// 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] < nums[j] {
				nums[i+1] = nums[j]
				break
			}
		}
	}
	//寻找不重复
	var count int = 1
	for i := 0; i < len(nums)-1; i++ {
		count++
		if nums[i] >= nums[i+1] {
			count = i + 1
			break
		}
	}
	return count
}

// 合并区间
// https://leetcode.cn/problems/merge-intervals/description/
func merge(intervals [][]int) [][]int {
	intervals = mergeArraySort(intervals)
	result := [][]int{}
	fmt.Println(len(result))
	for i := 0; i < len(intervals); i++ {
		//往result添加新元素,添加完退出进行下一轮循环
		if len(result) == 0 {
			result = append(result, intervals[i])
		} else {
			//新数组第二数字比待选第一个数字小 添加新元素
			if result[len(result)-1][1] < intervals[i][0] {
				result = append(result, intervals[i])
				continue
			}
		}
		//修改result最后元素
		if result[len(result)-1][1] < intervals[i][1] {
			result[len(result)-1][1] = intervals[i][1]
		}
	}
	return result
}

// 合并区间-数组排序
func mergeArraySort(intervals [][]int) [][]int {
	for i := 0; i < len(intervals)-1; i++ {
		for j := i + 1; j < len(intervals); j++ {
			if intervals[i][0] > intervals[j][0] {
				intervals[i], intervals[j] = intervals[j], intervals[i]
			}
		}
	}
	return intervals
}

// 两数之和
// https://leetcode.cn/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	sumArr := make([]int, 2)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				sumArr[0] = i
				sumArr[1] = j
			}
		}
	}
	return sumArr
}

func main() {
	//只出现一次的数字
	fmt.Println("=====只出现一次的数字=====")
	fmt.Println(singleNumber([]int{1, 2, 2, 3, 1}))
	//判断一个整数是否是回文数
	fmt.Println("========判断一个整数是否是回文数=====")
	fmt.Println(isPalindrome(10))
	//有效的括号
	fmt.Println("=====有效的括号=====")
	fmt.Println(isValid("(])"))
	//最长公共前缀
	fmt.Println("=====最长公共前缀=====")
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
	//加一
	fmt.Println("=====加一=====")
	fmt.Println(plusOne([]int{9, 9, 9, 9}))
	//删除有序数组中的重复项
	fmt.Println("=====删除有序数组中的重复项=====")
	fmt.Println(removeDuplicates([]int{1, 2, 2, 2}))
	//合并区间
	fmt.Println("=====合并区间=====")
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	//两数之和
	fmt.Println("=====两数之和=====")
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

}
