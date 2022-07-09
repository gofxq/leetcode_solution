package lc

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return Combine(candidates, []int{}, target)
}

func Combine(candidates, resCandidates []int, target int) [][]int {
	res := make(SSlice, 0, 127)
	// TODO 谨记golang 切片浅拷贝
	for i, candidate := range candidates {
		resNew := make([]int, len(resCandidates))
		copy(resNew, resCandidates)
		resNew = append(resNew, candidate)
		if SumSlice(resNew) < target {
			res = append(res, Combine(candidates[i:], resNew, target)...)
		} else if SumSlice(resNew) == target {
			res = append(res, resNew)
		}
	}
	//sort.Sort(res)
	return res
}

type SSlice [][]int

func (x SSlice) Len() int { return len(x) }
func (x SSlice) Less(i, j int) bool {
	minLen := len(x[i])
	if minLen > len(x[j]) {
		minLen = len(x[j])
	}
	for idx := 0; idx < minLen; idx++ {
		//log.Println(i, j, len(x[i]), len(x[j]), idx, minLen)
		if x[i][idx] < x[j][idx] {
			return true
		} else if x[i][idx] == x[j][idx] {
			continue
		} else {
			return false
		}
	}
	return false

}
func (x SSlice) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// [2,7,6,3,5,1] 9
// [[1,1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,2],[1,1,1,1,1,1,3],[1,1,1,1,1,2,2],[1,1,1,1,2,3],[1,1,1,1,5],[1,1,1,2,2,2],[1,1,1,3,3],[1,1,1,6],[1,1,2,2,3],[1,1,2,5],[1,1,7],[1,2,2,2,2],[1,2,3,3],[1,2,6],[1,3,5],[2,2,2,1],[2,2,5],[2,7],[3,3,3],[3,6]]
// [[1,1,1,1,1,1,1,1,1],[1,1,1,1,1,1,1,2],[1,1,1,1,1,1,3],[1,1,1,1,1,2,2],[1,1,1,1,2,3],[1,1,1,1,5],[1,1,1,2,2,2],[1,1,1,3,3],[1,1,1,6],[1,1,2,2,3],[1,1,2,5],[1,1,7],[1,2,2,2,2],[1,2,3,3],[1,2,6],[1,3,5],[2,2,2,3],[2,2,5],[2,7],[3,3,3],[3,6]]
// @lc code=end
