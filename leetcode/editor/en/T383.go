package leetcode

//Given two strings ransomNote and magazine, return true if ransomNote can be
//constructed from magazine and false otherwise.
//
// Each letter in magazine can only be used once in ransomNote.
//
//
// Example 1:
// Input: ransomNote = "a", magazine = "b"
//Output: false
// Example 2:
// Input: ransomNote = "aa", magazine = "ab"
//Output: false
// Example 3:
// Input: ransomNote = "aa", magazine = "aab"
//Output: true
//
//
// Constraints:
//
//
// 1 <= ransomNote.length, magazine.length <= 10⁵
// ransomNote and magazine consist of lowercase English letters.
//
// Related Topics Hash Table String Counting 👍 1491 👎 287

//leetcode submit region begin(Prohibit modification and deletion)
func canConstruct(ransomNote string, magazine string) bool {
	dp := make([][26]int, 2)
	for i, s := range []string{ransomNote, magazine} {
		for _, val := range s {
			dp[i][val-'a']++
		}
	}

	for i := 0; i < 26; i++ {
		if dp[0][i] > dp[1][i] {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
