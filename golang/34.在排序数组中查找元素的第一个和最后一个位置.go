package lc

import "sort"

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	leftIndex :=
		//binarySearchLeftFirst(nums, target)
		sort.SearchInts(nums, target)

	if leftIndex > -1 {
		rightIndex := leftIndex
		for rightIndex+1 < len(nums) &&
			nums[rightIndex+1] == target {
			rightIndex += 1
		}
		return []int{leftIndex, rightIndex}
	}
	return []int{-1, -1}
}

func binarySearchLeftFirst(nums []int, target int) (index int) {
	if len(nums) == 1 && nums[0] == target {
		return 0
	}
	ori, dst := 0, len(nums)-1
	for ori < dst {
		if ori+1 == dst {
			if nums[ori] == target {
				return ori
			} else if nums[dst] == target {
				return dst
			} else {
				return -1
			}
		}

		index = (ori + dst) / 2
		if nums[index] == target {
			for index > 0 && nums[index-1] == target {
				index -= 1
			}
			return index
		} else if nums[index] > target {
			dst = index
		} else {
			ori = index
		}

	}
	return -1
}

// GetMiddle 编译优化导致没差别
func GetMiddle(i, j int) int {
	return (i + j) / 2
}

func GetMiddleB(i, j int) int {
	return int(uint(i+j) >> 1)
}

// @lc code=end
