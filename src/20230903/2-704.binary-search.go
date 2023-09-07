package main

import "fmt"

/*
*
  - Description:
    https://leetcode.cn/problems/binary-search/
  - @Author Hollis
  - @Create 2023/9/3 13:00
*/
func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9
	result := search(nums, target)
	fmt.Println(result)
}

// 二分查找，一定要注意区间的定义，是左闭右闭还是左闭右开。
// 参考：https://www.bilibili.com/video/BV1tZ4y1q7XE/?p=1
// 这里我们采用左闭右闭。
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := (left + right) / 2
		if target < nums[middle] {
			right = middle - 1
		} else if target > nums[middle] {
			left = middle + 1
		} else {
			return middle
		}
	}
	return -1
}

// 官方的写法：
//func search(nums []int, target int) int {
//	left, right := 0, len(nums)-1
//	for left <= right {
//		mid := (right-left)/2 + left  // 与 mid := (left + right) / 2 的写法相比，这种写法更有可能防止整数相加溢出的问题。
//		num := nums[mid]
//		if num == target {  // 把return的分支写在前面，可以减少不必要的条件判断，直接返回
//			return mid
//		} else if num > target {
//			right = mid - 1
//		} else {
//			left = mid + 1
//		}
//	}
//	return -1
//}
//
//作者：力扣官方题解
//链接：https://leetcode.cn/problems/binary-search/solutions/980494/er-fen-cha-zhao-by-leetcode-solution-f0xw/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
