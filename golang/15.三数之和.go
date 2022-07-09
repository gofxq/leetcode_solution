package lc

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
//
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	posiNums := make(map[int]bool, 0)
	negaNums := make(map[int]bool, 0)
	zeroNums := 0

	for _, k := range nums {
		if k > 0 {
			if _, ok := posiNums[k]; ok {
				// 存在多个相同值
				posiNums[k] = true
			} else {
				posiNums[k] = false
			}
		} else if k < 0 {
			if _, ok := negaNums[k]; ok {
				negaNums[k] = true
			} else {
				negaNums[k] = false
			}
		} else {
			zeroNums++
		}
	}
	type Cell struct {
		p int // positive num
		t int // target num
		n int // negative num
	}
	resMap := make(map[Cell]struct{}, 0)

	if zeroNums >= 3 {
		resMap[Cell{0, 0, 0}] = struct{}{}
	}

	for p := range posiNums {
		for n := range negaNums {
			targetB := -p - n
			if targetB == 0 {
				if zeroNums > 0 {
					resMap[Cell{p, targetB, n}] = struct{}{}
				}
			} else if targetB > 0 {
				if v, ok := posiNums[targetB]; ok {
					if targetB < p {
						resMap[Cell{p, targetB, n}] = struct{}{}
					} else if targetB > p {
						resMap[Cell{targetB, p, n}] = struct{}{}
					} else {
						if v == true {
							resMap[Cell{p, targetB, n}] = struct{}{}
						}
					}
				}
			} else { // targetB < 0
				if v, ok := negaNums[targetB]; ok {
					if targetB > n {
						resMap[Cell{p, targetB, n}] = struct{}{}
					} else if targetB < n {
						resMap[Cell{p, n, targetB}] = struct{}{}
					} else {
						if v == true {
							resMap[Cell{p, targetB, n}] = struct{}{}
						}
					}
				}
			}
		}
	}
	res := make([][]int, 0)
	for v := range resMap {
		res = append(res, []int{v.n, v.t, v.p})
	}
	return res
}

// @lc code=end
