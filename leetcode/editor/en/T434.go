package leetcode

//Given a string s, return the number of segments in the string.
//
// A segment is defined to be a contiguous sequence of non-space characters.
//
//
// Example 1:
//
//
//Input: s = "Hello, my name is John"
//Output: 5
//Explanation: The five segments are ["Hello,", "my", "name", "is", "John"]
//
//
// Example 2:
//
//
//Input: s = "Hello"
//Output: 1
//
//
//
// Constraints:
//
//
// 0 <= s.length <= 300
// s consists of lowercase and uppercase English letters, digits, or one of the
//following characters "!@#$%^&*()_+-=',.:".
// The only space character in s is ' '.
//
// Related Topics String 👍 409 👎 977

//leetcode submit region begin(Prohibit modification and deletion)
func countSegments(s string) int {
	start := 0
	count := 0
	for i, ch := range s {
		if ch == ' ' {
			if i > start {
				count++
			}
			start = i + 1
		}
	}

	if len(s) > start {
		count++
	}

	return count
}

//leetcode submit region end(Prohibit modification and deletion)
