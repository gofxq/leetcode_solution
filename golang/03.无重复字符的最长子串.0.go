package lc

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
// import "log"

// func main() {
// 	log.Println(lengthOfLongestSubstring("abcabcbb"))
// }

func lengthOfLongestSubstring0(s string) int {
	lenb := len(s)

	if lenb == 1 {
		return 1
	}
	b := []byte(s)
	var maxLen int

	// 滑动窗口法
	// 首层循环定义起始位置
	for i := 0; i < lenb; i++ {
		// 新建空map保存窗口中所有字符
		m := make(map[byte]int)
		m[b[i]] = i

		// 嵌套循环定义结束位置
		for j := i + 1; j < lenb; j++ {
			// log.Println(i, j, maxLen)
			// 如果i到达最后一个字符，则结束
			if charIndex, ok := m[b[j]]; ok {
				// map中已存在当前字符
				// 计算窗口长度，并决定是否更新最大长度
				if (j - i) > maxLen {
					maxLen = j - i
				}
				// 更新窗口起始位置为当前重复字节的出现位置
				i = charIndex
				break
			} else {
				// 窗口到达结尾
				if j == lenb-1 {
					if (j - i + 1) > maxLen {
						maxLen = j - i + 1
					}
					// 更新窗口起始位置为当前重复字节的出现位置
					return maxLen
				}

				// 更新map中的值
				m[b[j]] = j
				continue
			}

		}
	}
	return maxLen
}

// @lc code=end
