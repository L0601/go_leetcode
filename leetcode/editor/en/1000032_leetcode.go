/*
 * @lc app=leetcode.cn id=1000032 lang=golang
 *
 * [面试题 17.04] 消失的数字
 *
 * https://leetcode.cn/problems/missing-number-lcci/description/
 *
 * lcci
 * Easy (58.31%)
 * Likes:    124
 * Dislikes: 0
 * Total Accepted:    61.5K
 * Total Submissions: 105.4K
 * Testcase Example:  '[3,0,1]'
 *
 * 数组nums包含从0到n的所有整数，但其中缺了一个。请编写代码找出那个缺失的整数。你有办法在O(n)时间内完成吗？
 *
 * 注意：本题相对书上原题稍作改动
 *
 * 示例 1：
 *
 * 输入：[3,0,1]
 * 输出：2
 *
 *
 *
 * 示例 2：
 *
 * 输入：[9,6,4,2,3,5,7,0,1]
 * 输出：8
 *
 *
 */
package leetcode

// @lc code=start
func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	ret := len(nums)
	for i, v := range nums {
		ret -= v
		ret += i
	}

	return ret
}

// @lc code=end
