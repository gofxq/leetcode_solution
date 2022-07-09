package lc

import "sort"

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start

func nextPermutation(nums []int) {
	// 查找不满足逆序上升的第一个数
	l := len(nums)
	brokeI := -1
	for i := l - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			brokeI = i
			break
		}
	}
	if brokeI == -1 {
		// 从右往左一直是升序
		sort.Ints(nums)
	} else {
		j := len(nums) - 1
		// 查找大于borker的最小数
		for j >= 0 && nums[j] <= nums[brokeI] {
			j--
		}
		//交换位置
		nums[brokeI], nums[j] = nums[j], nums[brokeI]
		// 重新排序
		sort.Ints(nums[brokeI+1:])

	}
}

// @lc code=end
