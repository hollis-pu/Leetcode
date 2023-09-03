package main

import "fmt"

/*
*
  - Description:
    https://leetcode.cn/problems/minimum-size-subarray-sum/
  - @Author Hollis
  - @Create 2023/9/3 14:43
*/
func main() {
	target := 213
	nums := []int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12}
	r := minSubArrayLen(target, nums)
	fmt.Println(r)
}

// 2023.09.03 双指针（滑动窗口）
// 参考：https://www.bilibili.com/video/BV1tZ4y1q7XE/?p=4
/* 思路：使用两个指针（i、j），表示满足条件的窗口的左右下标。
   遍历切片，在循环体里面对遍历的每个元素进行求和sum，当sum满足条件时（sum>=target），求出此时滑动窗口的大小。
   然后移动滑动窗口的起始位置，使sum-=nums[i]，然后使i++。循环这个过程，求得滑动窗口的最小值。
*/
func minSubArrayLen(target int, nums []int) int {
	i := 0
	sum := 0
	lenResult := len(nums)
	flag := false // flag表示是否有满足条件的滑动窗口
	for j := 0; j < len(nums); j++ {
		sum += nums[j]
		for sum >= target {
			flag = true            // 如果能进入到这个for循环，则一定存在满足条件的滑动窗口，则最后返回滑动窗口的大小，否则返回0
			lenWindow := j - i + 1 // 滑动窗口的大小，注意值为j - i + 1
			if lenWindow < lenResult {
				lenResult = lenWindow
			}
			sum -= nums[i]
			i++
		}
	}
	if !flag {
		return 0
	}
	return lenResult
}
