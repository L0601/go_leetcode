/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

package leetcode

// @lc code=start
func twoSum(nums []int, target int) []int {
	mp := make(map[int]int, len(nums))
	for idx, num := range nums {
		t, ok := mp[target-num]
		if ok && t != idx {
			return []int{t, idx}
		}
		mp[num] = idx
	}

	return []int{}
}

// @lc code=end
