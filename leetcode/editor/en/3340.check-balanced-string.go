/*
 * @lc app=leetcode id=3340 lang=golang
 *
 * [3340] Check Balanced String
 *
 * https://leetcode.com/problems/check-balanced-string/description/
 *
 * algorithms
 * Easy (82.07%)
 * Likes:    45
 * Dislikes: 0
 * Total Accepted:    28.9K
 * Total Submissions: 35.2K
 * Testcase Example:  '"1234"'
 *
 * You are given a string num consisting of only digits. A string of digits is
 * called balanced if the sum of the digits at even indices is equal to the sum
 * of digits at odd indices.
 *
 * Return true if num is balanced, otherwise return false.
 *
 *
 * Example 1:
 *
 *
 * Input: num = "1234"
 *
 * Output: false
 *
 * Explanation:
 *
 *
 * The sum of digits at even indices is 1 + 3 == 4, and the sum of digits at
 * odd indices is 2 + 4 == 6.
 * Since 4 is not equal to 6, num is not balanced.
 *
 *
 *
 * Example 2:
 *
 *
 * Input: num = "24123"
 *
 * Output: true
 *
 * Explanation:
 *
 *
 * The sum of digits at even indices is 2 + 1 + 3 == 6, and the sum of digits
 * at odd indices is 4 + 2 == 6.
 * Since both are equal the num is balanced.
 *
 *
 *
 *
 * Constraints:
 *
 *
 * 2 <= num.length <= 100
 * num consists of digits only
 *
 *
 */

package leetcode

// @lc code=start
func isBalanced(num string) bool {
	result := 0
	sign := true
	for _, element := range num {
		r := int(element) - '0'
		if !sign {
			r = -r
		}
		result += r
		sign = !sign
	}

	return result == 0

}

// @lc code=end
