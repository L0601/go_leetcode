/*
 * @lc app=leetcode.cn id=3079 lang=golang
 *
 * [3079] 求出加密整数的和
 *
 * https://leetcode.cn/problems/find-the-sum-of-encrypted-integers/description/
 *
 * algorithms
 * Easy (80.96%)
 * Likes:    2
 * Dislikes: 0
 * Total Accepted:    3.4K
 * Total Submissions: 4.2K
 * Testcase Example:  '[1,2,3]'
 *
 * 给你一个整数数组 nums ，数组中的元素都是 正 整数。定义一个加密函数 encrypt ，encrypt(x) 将一个整数 x 中 每一个 数位都用
 * x 中的 最大 数位替换。比方说 encrypt(523) = 555 且 encrypt(213) = 333 。
 * 
 * 请你返回数组中所有元素加密后的 和 。
 * 
 * 
 * 
 * 示例 1：
 * 
 * 
 * 输入：nums = [1,2,3]
 * 
 * 输出：6
 * 
 * 解释：加密后的元素位 [1,2,3] 。加密元素的和为 1 + 2 + 3 == 6 。
 * 
 * 
 * 示例 2：
 * 
 * 
 * 输入：nums = [10,21,31]
 * 
 * 输出：66
 * 
 * 解释：加密后的元素为 [11,22,33] 。加密元素的和为 11 + 22 + 33 == 66 。
 * 
 * 
 * 
 * 
 * 提示：
 * 
 * 
 * 1 <= nums.length <= 50
 * 1 <= nums[i] <= 1000
 * 
 * 
 */
package leetcode
// @lc code=start
func encryptNum(n int) int {
	if n <= 9 {
		return n
	}

	max, base := 0, 0
	for n > 0 {
		r := n % 10
		n = n / 10
		base = base * 10 + 1
		if r > max {
			max = r
		}
	}

	return max * base
}
func sumOfEncryptedInt(nums []int) int {
	ret := 0
	for _, i := range nums {
		ret += encryptNum(i)
	}

	return ret
}
// @lc code=end

