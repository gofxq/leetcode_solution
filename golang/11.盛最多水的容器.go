package lc

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 *
 * https://leetcode-cn.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (62.39%)
 * Likes:    1272
 * Dislikes: 0
 * Total Accepted:    172.6K
 * Total Submissions: 276.6K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为
 * (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 *
 * 说明：你不能倾斜容器，且 n 的值至少为 2。
 *
 *
 *
 *
 *
 * 图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
 *
 *
 *
 * 示例：
 *
 * 输入：[1,8,6,2,5,4,8,3,7]
 * 输出：49
 *
 */
// @lc code=start
func maxArea(height []int) int {
	// 双索引夹逼法
	// 证明过程可见题解
	// original
	o := 0
	// destination
	d := len(height) - 1

	areaTmp := 0
	for {
		if o == d {
			return areaTmp
		}
		if m(height[o], height[d]) {
			if areaTmp < (d-o)*height[d] {
				areaTmp = (d - o) * height[d]
			}
			d--
		} else {
			if areaTmp < (d-o)*height[o] {
				areaTmp = (d - o) * height[o]
			}
			o++
		}
	}
}
func m(a, b int) bool {
	if a > b {
		return true
	}
	return false
}

// @lc code=end
