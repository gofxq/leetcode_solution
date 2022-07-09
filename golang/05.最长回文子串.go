package lc

// @lc code=start
// 中心扩展算法
func longestPalindrome(s string) string {
	// 存储s长度，空间换时间
	lenS := len(s)
	// golang如果空[]byte转string会得到\u0000，与标准答案不符
	if lenS == 0 {
		return ""
	}

	// 初始化
	maxLen := 1
	b := []byte(s)
	start := 0
	end := 1

	// 字串长度为奇数或者偶数标记
	var flagOdd, flagDouble bool
	// 奇数个
	for i := 0; i < lenS-1; i++ {
		flagOdd = true
		flagDouble = true
		for j := 0; j <= min(i, lenS-i-2); j++ {
			if flagOdd {
				// 检查字串为奇数的情况
				if b[i-j] == b[i+j] {
					if j*2+1 > maxLen {
						maxLen = j*2 + 1
						start = i - j
						end = i + j + 1
					}
				} else {
					flagOdd = false
				}
			}

			if flagDouble {
				if b[i-j] == b[i+j+1] {
					if (j+1)*2 > maxLen {
						maxLen = (j + 1) * 2
						start = i - j
						end = i + j + 2
						//res = string(b[i-j : i+j+2])
					}
				} else {
					flagDouble = false
				}

			}

		}

		// 奇数情况下，由于数组索引限制的原因，可能出现上述循环无法遍历到的情况
		if min(i, lenS-i-1) == lenS-i-1 {
			j := lenS - i - 1
			if flagOdd {
				if b[i-j] == b[i+j] {
					if j*2+1 > maxLen {
						maxLen = j*2 + 1
						start = i - j
						end = i + j + 1
					}
				} else {
					flagOdd = false
				}
			}
		}
	}
	return string(b[start:end])

}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end
