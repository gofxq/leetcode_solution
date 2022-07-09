package lc

// @lc code=start
//
func threeSumT(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	initLen := len(nums)
	posiNums := make(map[int]struct{}, initLen)
	negaNums := make(map[int]struct{}, initLen)
	zeroNums := 0

	for _, v := range nums {
		if v > 0 {
			posiNums[v] = struct{}{}
		} else if v < 0 {
			negaNums[v] = struct{}{}
		} else {
			zeroNums++
		}
	}
	type Cell struct {
		p int // positive num
		t int // target num
		n int // negative num
	}
	resMap := make(map[Cell]struct{}, initLen)

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
				delete(posiNums, targetB)
				if _, ok := posiNums[targetB]; ok {
					resMap[Cell{p, targetB, n}] = struct{}{}
				}
				posiNums[targetB] = struct{}{}
			} else { // targetB < 0
				delete(negaNums, targetB)
				if _, ok := negaNums[targetB]; ok {
					resMap[Cell{p, targetB, n}] = struct{}{}
				}
				negaNums[targetB] = struct{}{}
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
