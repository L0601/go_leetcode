package leetcode

/*
 * @lc app=leetcode id=2006 lang=golang
 *
 * [2006] Count Number of Pairs With Absolute Difference K
 */

// @lc code=start
func countKDifference(nums []int, k int) int {
	mp := [101][]int{}
	for i, val := range nums {
		mp[val] = append(mp[val], i)
	}

	count := 0
	for i, val := range nums {
		target := [2]int{val + k, val - k}
		for _, j := range target {
			if j >= 0 && j < 101 {
				for _, idx := range mp[j] {
					if idx > i {
						count++
					}
				}
			}
		}
	}

	return count

}

// @lc code=end
