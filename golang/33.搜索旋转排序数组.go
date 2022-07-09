package lc

// @lc code=start
func search(nums []int, target int) int {
	o, d := 0, len(nums)-1
	m := 0
	for {
		if o < d {
			m = (o + d) / 2
			// 这里如果能用c++的异或才是真的牛批
			if !(!(nums[0] > target) == (nums[0] > nums[m])) == (target > nums[m]) {
				o = m + 1
			} else {
				d = m
			}
		} else {
			break
		}
	}
	if o == d && nums[o] == target {
		return o
	}

	return -1
}

// @lc code=end

/* class Solution {
 public:
 int search(vector<int>& nums, int target) {
	 int lo = 0, hi = nums.size() - 1;
	 while (lo < hi) {
		 int mid = (lo + hi) / 2;
		 if ((nums[0] > target) ^ (nums[0] > nums[mid]) ^ (target > nums[mid]))
			 lo = mid + 1;
		 else
			 hi = mid;
	 }
	 return lo == hi && nums[lo] == target ? lo : -1;
 }
 };*/
