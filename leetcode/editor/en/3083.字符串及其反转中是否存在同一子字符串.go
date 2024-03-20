/*
 * @lc app=leetcode.cn id=3083 lang=golang
 *
 * [3083] 字符串及其反转中是否存在同一子字符串
 *
 * https://leetcode.cn/problems/existence-of-a-substring-in-a-string-and-its-reverse/description/
 *
 * algorithms
 * Easy (71.73%)
 * Likes:    3
 * Dislikes: 0
 * Total Accepted:    5.6K
 * Total Submissions: 7.7K
 * Testcase Example:  '"leetcode"'
 *
 * 给你一个字符串 s ，请你判断字符串 s 是否存在一个长度为 2 的子字符串，在其反转后的字符串中也出现。
 *
 * 如果存在这样的子字符串，返回 true；如果不存在，返回 false 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：s = "leetcode"
 *
 * 输出：true
 *
 * 解释：子字符串 "ee" 的长度为 2，它也出现在 reverse(s) == "edocteel" 中。
 *
 *
 * 示例 2：
 *
 *
 * 输入：s = "abcba"
 *
 * 输出：true
 *
 * 解释：所有长度为 2 的子字符串 "ab"、"bc"、"cb"、"ba" 也都出现在 reverse(s) == "abcba" 中。
 *
 *
 * 示例 3：
 *
 *
 * 输入：s = "abcd"
 *
 * 输出：false
 *
 * 解释：字符串 s 中不存在满足「在其反转后的字符串中也出现」且长度为 2 的子字符串。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= s.length <= 100
 * 字符串 s 仅由小写英文字母组成。
 *
 *
 */
package leetcode

// @lc code=start

func isSubstringPresent(s string) bool {
	if len(s) < 2 {
		return false
	}

	mp := make(map[string]interface{}, len(s))
	for i := 0; i < len(s)-1; i++ {
		mp[s[i:i+2]] = 
	}

	for i := 0; i < len(s)-1; i++ {
		if _, ok := mp[string(s[i+1])+string(s[i])]; ok {
			return true
		}
	}

	return false
}

// @lc code=end
