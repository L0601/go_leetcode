/*
 * @lc app=leetcode id=2186 lang=golang
 *
 * [2186] Minimum Number of Steps to Make Two Strings Anagram II
 */

package leetcode

// @lc code=start
func minSteps(s string, t string) int {
	x, y := trans(s), trans(t)
	count := 0
	for i := 0; i < 26; i++ {
		diff := x[i] - y[i]
		if diff > 0 {
			count += diff
		} else {
			count -= diff
		}
	}

	return count
}

func trans(s string) [26]int {
	array := [26]int{}
	for _, ch := range s {
		array[ch-'a']++
	}

	return array
}

// @lc code=end
