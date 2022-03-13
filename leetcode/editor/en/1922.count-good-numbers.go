/*
 * @lc app=leetcode id=1922 lang=golang
 *
 * [1922] Count Good Numbers
 */

package leetcode

// @lc code=start
func countGoodNumbers(n int64) int {
	mod := int64(1000000007)

	return int(modPower(5, (n+1)/2, mod) * modPower(4, n/2, mod) % mod)
}

func modPower(base, power, mod int64) int64 {
	if power == 0 {
		return 1
	}

	half := modPower(base, power/2, mod)
	if power%2 == 0 {
		return half * half % mod
	}

	return half * half * base % mod
}

// @lc code=end
