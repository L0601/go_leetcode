/*
 * @lc app=leetcode id=2085 lang=golang
 *
 * [2085] Count Common Words With One Occurrence
 */

package leetcode

// @lc code=start
func countWords(words1 []string, words2 []string) int {
	m, n := count(words1), count(words2)
	count := 0
	for k, v := range m {
		if v == 1 && n[k] == 1 {
			count++
		}
	}

	return count
}

func count(words []string) map[string]int {
	mp := make(map[string]int, len(words))
	for _, word := range words {
		mp[word]++
	}

	return mp
}

// @lc code=end
