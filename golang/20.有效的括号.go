/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */
package lc

// @lc code=start
func isValid(s string) bool {
	bracketStack := make([]rune, 0, len(s))
	for _, c := range s {
		if len(bracketStack) == 0 {
			bracketStack = append(bracketStack, c)
			continue
		}
		bracketStackLast := bracketStack[len(bracketStack)-1]
		switch c {
		case ')':
			if bracketStackLast == '(' {
				bracketStack = bracketStack[:len(bracketStack)-1]
				continue
			}
		case '}':
			if bracketStackLast == '{' {
				bracketStack = bracketStack[:len(bracketStack)-1]
				continue
			}
		case ']':
			if bracketStackLast == '[' {
				bracketStack = bracketStack[:len(bracketStack)-1]
				continue
			}
		}
		bracketStack = append(bracketStack, c)
	}
	return len(bracketStack) == 0
}

// @lc code=end
