package lc

// @lc code=start
func intToRoman(num int) string {
	// 简单的递归题
	s := ""
	if num >= 1000 {
		for i := 0; i < num/1000; i++ {
			s += "M"
		}
		return s + intToRoman(num%1000)
	} else if num >= 500 {
		if num/100 == 9 {
			return s + "CM" + intToRoman(num%100)
		}
		for i := 0; i < num/500; i++ {
			s += "D"
		}
		return s + intToRoman(num%500)
	} else if num >= 100 {
		if num/100 == 4 {
			return s + "CD" + intToRoman(num%100)
		}
		for i := 0; i < num/100; i++ {
			s += "C"
		}
		return s + intToRoman(num%100)
	} else if num >= 50 {
		if num/10 == 9 {
			return s + "XC" + intToRoman(num%10)
		}
		for i := 0; i < num/50; i++ {
			s += "L"
		}
		return s + intToRoman(num%50)
	} else if num >= 10 {
		if num/10 == 4 {
			return s + "XL" + intToRoman(num%10)
		}
		for i := 0; i < num/10; i++ {
			s += "X"
		}
		return s + intToRoman(num%10)
	} else if num >= 5 {
		if num == 9 {
			return s + "IX"
		}
		for i := 0; i < num/5; i++ {
			s += "V"
		}
		return s + intToRoman(num%5)
	} else {
		if num == 4 {
			return s + "IV"
		}
		for i := 0; i < num; i++ {
			s += "I"
		}
		return s
	}
}

// @lc code=end
