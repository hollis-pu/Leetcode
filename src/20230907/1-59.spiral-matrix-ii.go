package main

import "fmt"

/*
*
  - Description:
    https://leetcode.cn/problems/spiral-matrix-ii/
  - @Author Hollis
  - @Create 2023/9/7 15:24
*/
func main() {
	n := 9
	r := generateMatrix(n)
	//fmt.Println(r)
	for i := range r {
		fmt.Println(r[i])
	}
}

// 这道题的难度还是挺大的，重要的是如何思考，如何将题目的描述转换成代码可解决的逻辑关系。
// 思路：将这个问题转换为正方形循环遍历四条边的问题，每次遍历时，不填充最后一个元素，留给下一条边来处理（统一规则）。
// 这里是分层的概念，每完成一次外层循环，就是完成了一个正方形层的数据填充。使用startX来表示现在填充的是第几层。
// 依次遍历四条边，注意每次处理的边界，这里是左闭右开的。
// 参考：https://www.bilibili.com/video/BV1SL4y1N7mV（代码随想录）
func generateMatrix(n int) [][]int {
	startX := 0 // 代表了第几层的循环
	count := 1  // 每个位置要填充的数字，从1开始
	offset := 1 // 每层的偏移量，每循环一层，偏移量就加1

	// 创建n*n的切片，学会这种写法
	a := make([][]int, n)
	for i := range a {
		a[i] = make([]int, n)
	}

	i, j := 0, 0                // i表示行，j表示列
	for m := 1; m <= n/2; m++ { // 控制总的循环层数
		// 正方形的上边：行值不变，列值增加
		for j = startX; j < n-offset; j++ { // 这里不包括最后一个元素，所以j < n-offset，offset从1开始
			a[m-1][j] = count
			count++
		}
		// 正方形的右边：列值不变，行值增加
		for i = startX; i < n-offset; i++ {
			a[i][j] = count
			count++
		}
		// 正方形的底边：行值不变，列值减小
		for ; j > startX; j-- {
			a[i][j] = count
			count++
		}
		// 正方形的左边：列值不变，行值减小
		for ; i > startX; i-- {
			a[i][j] = count
			count++
		}
		startX++
		offset++
	}
	if n%2 != 0 { // 如果n为奇数，则最中心的位置会遍历不到，所以需要手动赋值
		a[startX][startX] = count
	}
	return a
}
