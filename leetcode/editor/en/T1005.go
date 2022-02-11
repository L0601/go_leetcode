package leetcode

import "math"

//Given an integer array nums and an integer k, modify the array in the
//following way:
//
//
// choose an index i and replace nums[i] with -nums[i].
//
//
// You should apply this process exactly k times. You may choose the same index
//i multiple times.
//
// Return the largest possible sum of the array after modifying it in this way.
//
//
//
// Example 1:
//
//
//Input: nums = [4,2,3], k = 1
//Output: 5
//Explanation: Choose index 1 and nums becomes [4,-2,3].
//
//
// Example 2:
//
//
//Input: nums = [3,-1,0,2], k = 3
//Output: 6
//Explanation: Choose indices (1, 2, 2) and nums becomes [3,1,0,2].
//
//
// Example 3:
//
//
//Input: nums = [2,-3,-1,5,-4], k = 2
//Output: 13
//Explanation: Choose indices (1, 4) and nums becomes [2,3,-1,5,4].
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -100 <= nums[i] <= 100
// 1 <= k <= 10⁴
//
// Related Topics Array Greedy Sorting 👍 875 👎 75

//leetcode submit region begin(Prohibit modification and deletion)
func largestSumAfterKNegations(nums []int, k int) int {
	const (
		LEN   = 201
		LEVEL = 100
	)

	dp := [LEN]int{}
	for _, val := range nums {
		dp[val+LEVEL]++
	}

	negativeCount := 0
	for idx, count := range dp {
		for idx < LEVEL && count > 0 {
			negativeCount++
			count--
		}
	}

	if negativeCount <= k {
		return sumMax(nums, k-negativeCount)
	}

	// 将最小的k个取反
	sum, cnt := 0, 0
	for i := 0; i < LEN; i++ {
		for dp[i] != 0 {
			if cnt < k {
				sum += LEVEL - i
				cnt++
			} else {
				sum += i - LEVEL
			}
			dp[i]--
		}
	}

	return sum
}

func sumMax(nums []int, k int) int {
	min, sum := math.MaxInt32, 0
	for _, val := range nums {
		if val < 0 {
			val = -val
		}

		sum += val
		if val < min {
			min = val
		}
	}

	if k%2 == 0 {
		return sum
	}

	// 将绝对值 最小的取反
	return sum - 2*min
}

//leetcode submit region end(Prohibit modification and deletion)
