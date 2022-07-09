package lc

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 *
 * https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/description/
 *
 * algorithms
 * Medium (53.14%)
 * Likes:    662
 * Dislikes: 0
 * Total Accepted:    100.3K
 * Total Submissions: 188.5K
 * Testcase Example:  '"23"'
 *
 * 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
 *
 * 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
 *
 *
 *
 * 示例:
 *
 * 输入："23"
 * 输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 * 说明:
 * 尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
 *
 */

// @lc code=start
func letterCombinations(digits string) []string {
	// BFS 广度优先搜索
	ans := []string{}
	dMap := map[byte]string{}
	dMap['2'] = "abc"
	dMap['3'] = "def"
	dMap['4'] = "ghi"
	dMap['5'] = "jkl"
	dMap['6'] = "mno"
	dMap['7'] = "pqrs"
	dMap['8'] = "tuv"
	dMap['9'] = "wxyz"
	printLetters(digits, dMap, 0, 0, []byte{}, &ans)
	return ans

}

func printLetters(digits string, dMap map[byte]string, dcur int, dmcur int, tmp []byte, ans *[]string) {

	ld := len(digits)
	if (dcur >= ld) || (dmcur >= len(dMap[digits[dcur]])) {
		return
	}

	if dmcur != 0 {
		tmp[dcur] = dMap[digits[dcur]][dmcur]
	} else {
		tmp = append(tmp, dMap[digits[dcur]][dmcur])
	}

	if len(tmp) == len(digits) {
		*ans = append(*ans, string(tmp))
	}
	printLetters(digits, dMap, dcur+1, 0, tmp, ans)
	printLetters(digits, dMap, dcur, dmcur+1, tmp, ans)
}

// @lc code=end
