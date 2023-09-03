package main

import "fmt"

/*
*
  - Description:
    https://leetcode.cn/problems/squares-of-a-sorted-array/
  - @Author Hollis
  - @Create 2023/9/3 14:18
*/
func main() {
	nums := []int{-4, -1, 0, 3, 10}
	r := sortedSquares(nums)
	fmt.Println(r)
}

// 2023.09.03 左右双指针
// 参考：// 参考：https://www.bilibili.com/video/BV1tZ4y1q7XE/?p=3（代码随想录）
func sortedSquares(nums []int) []int {
	lenNum := len(nums)
	left := 0
	right := lenNum - 1
	numResult := make([]int, lenNum)
	for i := 0; i < lenNum; i++ {
		if nums[left]*nums[left] > nums[right]*nums[right] { // 当左边大时，取左边的值，并移动左指针
			numResult[lenNum-i-1] = nums[left] * nums[left]
			left++
		} else { // 否者，移动右指针
			numResult[lenNum-i-1] = nums[right] * nums[right] // 目标值从后往前填充
			right--
		}
	}
	return numResult
}
