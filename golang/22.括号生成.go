package lc

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	// 依然是简单的递归
	// l，r分别为未使用的左右括号数量
	// 递归的结束条件为l == 0， 即未使用的左括号为0
	// l不为0时，递归genNext(l-1,r)
	// 同时附加genNext(l,r-1),即此处填右括号的情况
	l := n
	r := n
	tmpS := ""
	res := make([]string, 0)
	res = append(res, genNext(l, r, tmpS)...)
	return res
}

func genNext(l int, r int, s string) []string {
	// 结束条件为l==0
	if l == 0 {
		for i := 0; i < r; i++ {
			s += ")"
		}
		return []string{s}
	}
	res := make([]string, 0)
	// l > 1
	res = append(res, genNext(l-1, r, s+"(")...)
	if l < r {
		res = append(res, genNext(l, r-1, s+")")...)
	}
	return res
}

// @lc code=end
