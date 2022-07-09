package lc

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 *
 * https://leetcode-cn.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (33.99%)
 * Likes:    1815
 * Dislikes: 0
 * Total Accepted:    323K
 * Total Submissions: 950.1K
 * Testcase Example:  '123'
 *
 * 给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。
 *
 * 示例 1:
 *
 * 输入: 123
 * 输出: 321
 *
 *
 * 示例 2:
 *
 * 输入: -123
 * 输出: -321
 *
 *
 * 示例 3:
 *
 * 输入: 120
 * 输出: 21
 *
 *
 * 注意:
 *
 * 假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回
 * 0。
 *
 */

// @lc code=start

func reverse(x int) int {
	// 简单计算题
	tX := x
	var res int
	for ; tX != 0; tX = tX / 10 {
		res = res*10 + tX%10
	}
	if res > 1<<31 || res < -(1<<31) {
		return 0
	}
	return res
}

// @lc code=end
