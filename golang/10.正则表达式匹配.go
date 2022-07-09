package lc

import "regexp"

// @lc code=start
func isMatch(s string, p string) bool {

	res, _ := regexp.MatchString(p, "$"+s+"$")

	return res
}

// @lc code=end
