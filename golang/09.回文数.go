package lc

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 *
 * https://leetcode-cn.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (57.30%)
 * Likes:    997
 * Dislikes: 0
 * Total Accepted:    293.8K
 * Total Submissions: 512.5K
 * Testcase Example:  '121'
 *
 * 判断一个整数是否是回文数。回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。
 *
 * 示例 1:
 *
 * 输入: 121
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: -121
 * 输出: false
 * 解释: 从左向右读, 为 -121 。 从右向左读, 为 121- 。因此它不是一个回文数。
 *
 *
 * 示例 3:
 *
 * 输入: 10
 * 输出: false
 * 解释: 从右向左读, 为 01 。因此它不是一个回文数。
 *
 *
 * 进阶:
 *
 * 你能不将整数转为字符串来解决这个问题吗？
 *
 */
// @lc code=start
func isPalindrome(x int) bool {
	// 直接用NO.7的题解判断是否相等即可
	if x < 0 {
		return false
	}
	return x == reverse1(x)
}

func reverse1(x int) int {
	tX := x
	var res int
	for ; tX != 0; tX = tX / 10 {
		res = res*10 + tX%10
	}
	return res
}

// @lc code=end
