package lc

// @lc code=start
func convert(s string, numRows int) string {
	// 简单的计算题
	/*
	* L0      D6      RD
	* E2   O5 E8   IC IE
	* E2 C4   I9 HB   NF
	* T4      SA      GG
	* 很容易总结出第一排为i%(2*numRows) == 0的情况
	* 下面一排依次是 i%(2*numRows) == iota或i(2*numRows) == (2*numRows) - i的情况
	* 新建切片遍历取一遍即可得到结果
	 */

	lenS := len(s)
	resB := make([]byte, 0, lenS)
	b := []byte(s)
	mod := 2 * (numRows - 1)
	if numRows == 1 {
		return s
	}
	for i := 0; i <= mod/2; i++ {
		for j := 0; j < lenS; j++ {
			if j%mod == i || j%mod == (mod-i) {
				resB = append(resB, b[j])
			}
		}
	}
	return string(resB)
}

// @lc code=end
