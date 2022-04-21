/*
 * @lc app=leetcode id=2325 lang=golang
 *
 * [2222] Number of Ways to Select Buildings
 *
 * https://leetcode.com/problems/number-of-ways-to-select-buildings/description/
 *
 * algorithms
 * Medium (44.55%)
 * Likes:    327
 * Dislikes: 11
 * Total Accepted:    9.8K
 * Total Submissions: 22K
 * Testcase Example:  '"001101"'
 *
 * You are given a 0-indexed binary string s which represents the types of
 * buildings along a street where:
 *
 *
 * s[i] = '0' denotes that the i^th building is an office and
 * s[i] = '1' denotes that the i^th building is a restaurant.
 *
 *
 * As a city official, you would like to select 3 buildings for random
 * inspection. However, to ensure variety, no two consecutive buildings out of
 * the selected buildings can be of the same type.
 *
 *
 * For example, given s = "001101", we cannot select the 1^st, 3^rd, and 5^th
 * buildings as that would form "011" which is not allowed due to having two
 * consecutive buildings of the same type.
 *
 *
 * Return the number of valid ways to select 3 buildings.
 *
 *
 * Example 1:
 *
 * Input: s = "001101"
 * Output: 6
 * Explanation:
 * The following sets of indices selected are valid:
 * - [0,2,4] from "001101" forms "010"
 * - [0,3,4] from "001101" forms "010"
 * - [1,2,4] from "001101" forms "010"
 * - [1,3,4] from "001101" forms "010"
 * - [2,4,5] from "001101" forms "101"
 * - [3,4,5] from "001101" forms "101"
 * No other selection is valid. Thus, there are 6 total ways.
 *
 *
 * Example 2:
 *
 * Input: s = "11100"
 * Output: 0
 * Explanation: It can be shown that there are no valid selections.
 *
 *
 *
 * Constraints:
 *
 *
 * 3 <= s.length <= 10^5
 * s[i] is either '0' or '1'.
 *
 *
 */
package leetcode

// @lc code=start
func numberOfWays(s string) int64 {
	// 超时了
	// if len(s) < 3 {
	// 	return 0
	// }
	// sum := make([]int, 0, len(s))
	// count := 0
	// for _, c := range s {
	// 	count += int(c - '0')
	// 	sum = append(sum, count)
	// }

	// res := int64(0)
	// for i := range s {
	// 	for j := i + 1; j < len(s); j++ {
	// 		if s[i] == s[j] {
	// 			continue
	// 		}
	// 		if s[j] == '0' {
	// 			res += int64(sum[len(s)-1] - sum[j])
	// 		} else {
	// 			res += int64(len(s) - j - 1 - (sum[len(s)-1] - sum[j]))
	// 		}
	// 	}
	// }

	// return res

	count_of_1 := 0
	left := 0
	res := int64(0)
	for _, c := range s {
		if c == '1' {
			count_of_1++
		}
	}

	for i, c := range s {
		if c == '1' {
			// 左边的0 * 右边的0
			res += int64((i - left) * (len(s) - i - 1 - (count_of_1 - left - 1)))
			left++
		} else {
			// 左边的1 * 右边的1
			res += int64(left * (count_of_1 - left))
		}
	}

	return res
}

// @lc code=end
