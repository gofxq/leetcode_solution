package lc

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=40 lang=golang
 *
 * [40] 组合总和 II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return Combine2(candidates, []int{}, target)
}

func Combine2(candidates, resCandidates []int, target int) [][]int {
	res := make([][]int, 0, 127)
	// TODO 谨记golang 切片浅拷贝

	for i, candidate := range candidates {
		if i > 0 && candidates[i-1] == candidate {
			continue
		}
		resNew := make([]int, len(resCandidates))
		copy(resNew, resCandidates)
		resNew = append(resNew, candidate)
		if SumSlice(resNew) < target {
			if i < len(candidates)-1 {
				res = append(res, Combine2(candidates[i+1:], resNew, target)...)
			}
		} else if SumSlice(resNew) == target {
			res = append(res, resNew)
		}
	}
	//sort.Sort(res)
	return res
}

func SumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

// @lc code=end
