package leetcode
/*
 * @lc app=leetcode id=2150 lang=golang
 *
 * [2150] Find All Lonely Numbers in the Array
 */

// @lc code=start
func findLonely(nums []int) []int {
	mp := make(map[int]int, len(nums))
	result := make([]int, 0, len(nums))

	for _, num := range nums {
		mp[num] = mp[ num] + 1
	}

	for k, v := range mp {
		if v == 1 && mp[k - 1] == 0 && mp[k + 1] == 0 {
			result = append(result, k)
		}
	}

	return result
}
// @lc code=end

