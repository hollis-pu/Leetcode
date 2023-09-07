package main

import "fmt"

/*
*
  - Description:
    https://leetcode.cn/problems/remove-element/
  - @Author Hollis
  - @Create 2023/9/3 14:01
*/
func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	r := removeElement(nums, val)
	fmt.Println(r)
}

// 2023.09.03 快慢指针
// 参考：https://www.bilibili.com/video/BV1tZ4y1q7XE/?p=2（代码随想录）
func removeElement(nums []int, val int) int {
	fast, slow := 0, 0
	for fast = 0; fast < len(nums); fast++ { // 快指针遍历切片
		if nums[fast] != val {
			nums[slow] = nums[fast] // 慢指针保存目标结果
			slow++
		}
	}
	return slow
}
