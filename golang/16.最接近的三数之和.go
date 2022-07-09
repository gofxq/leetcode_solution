package lc

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
// 暴力解法
func threeSumClosest(nums []int, target int) int {
	diff := 10000

	var sum, tempDiff, resSum int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				sum = nums[i] + nums[j] + nums[k]
				if sum > target {
					tempDiff = sum - target
				} else {
					tempDiff = target - sum
				}
				if tempDiff < diff {
					diff = tempDiff
					resSum = sum
					if diff == 0 {
						return sum
					}
				}
			}
		}
	}
	return resSum
}

// @lc code=end
