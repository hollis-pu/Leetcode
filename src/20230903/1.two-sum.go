package main

import "fmt"

/*
*
  - Description:
    https://leetcode.cn/problems/two-sum/
  - @Author Hollis
  - @Create 2023/9/3 9:53
*/
func main() {
	nums := []int{3, 3, 9, 8}
	target := 6

	r1 := twoSum01(nums, target)
	fmt.Println(r1)

	r2 := twoSum02(nums, target)
	fmt.Println(r2)

}

// 暴力破解
func twoSum01(nums []int, target int) []int {
	var result []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(append(result, i), j)
			}
		}
	}
	return result
}

// 2023.09.03 哈希表
/* 思路：遍历题目给出的切片，将每次遍历到的元素存入到hashmap中。（key存元素值，value存元素的下标）
   在遍历时，先查看hashmap中是否有满足条件的key，如果有，则直接返回key的value（元素下标）和当前遍历的元素下标i。
   需要注意的是，一定要先检查hashmap中的元素，然后将遍历的元素加入到hashmap中。
*/
func twoSum02(nums []int, target int) []int {
	var hashMap = make(map[int]int)  // map01存放出现过的元素，注意hashmap的key存的存的是元素值，value存的是元素的下标，这样方便查找
	for i := 0; i < len(nums); i++ { // 遍历目标数组
		targetSub := target - nums[i]                // 目标差值
		if p, exists := hashMap[targetSub]; exists { // 如果map中存在以目标差值为key的元素，则直接返回结果
			return []int{p, i}
		}
		hashMap[nums[i]] = i // key:元素值，value:元素下标
		// 先判断是否满足条件，再将值存入hashmap中，否者会存在如下的情况：
		// nums={3,3},target=6   如果一上来就加入到hashmap中的话，后面的3会把前面的3给覆盖掉。
	}
	return nil
}
