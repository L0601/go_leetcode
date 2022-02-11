package leetcode

import (
	"sort"
)

/*
 * @lc app=leetcode id=2148 lang=golang
 *
 * [2148] Count Elements With Strictly Smaller and Greater Elements
 */

// @lc code=start
func countElements(nums []int) int {
	if len(nums) <= 2 {
		return 0
	}
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	for l < r {
		anyHit := false
		if l+1 <= r && nums[l] == nums[l+1] {
			l += 1
			anyHit = true
		}
		if r-1 >= l && nums[r] == nums[r-1] {
			r -= 1
			anyHit = true
		}

		if !anyHit {
			break
		}
	}

	r = r - l - 1
	if r < 0 {
		return 0
	}

	return r
}

// @lc code=end
