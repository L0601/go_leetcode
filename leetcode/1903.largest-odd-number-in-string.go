/*
 * @lc app=leetcode id=1903 lang=golang
 *
 * [1903] Largest Odd Number in String
 */
package leetcode

// @lc code=start
func largestOddNumber(num string) string {
	for idx := len(num) - 1; idx >= 0; idx-- {
		if (num[idx]-'0')%2 == 1 {
			return num[:idx+1]
		}
	}

	return ""
}

// @lc code=end
