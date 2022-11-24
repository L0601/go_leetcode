package leetcode

//You are given a 0-indexed array of positive integers nums. Find the number of
//triplets (i, j, k) that meet the following conditions:
//
//
// 0 <= i < j < k < nums.length
// nums[i], nums[j], and nums[k] are pairwise distinct.
//
// In other words, nums[i] != nums[j], nums[i] != nums[k], and nums[j] != nums[
//k].
//
//
//
// Return the number of triplets that meet the conditions.
//
//
// Example 1:
//
//
//Input: nums = [4,4,2,4,3]
//Output: 3
//Explanation: The following triplets meet the conditions:
//- (0, 2, 4) because 4 != 2 != 3
//- (1, 2, 4) because 4 != 2 != 3
//- (2, 3, 4) because 2 != 4 != 3
//Since there are 3 triplets, we return 3.
//Note that (2, 0, 4) is not a valid triplet because 2 > 0.
//
//
// Example 2:
//
//
//Input: nums = [1,1,1,1,1]
//Output: 0
//Explanation: No triplets meet the conditions so we return 0.
//
//
//
// Constraints:
//
//
// 3 <= nums.length <= 100
// 1 <= nums[i] <= 1000
//
//
// Related Topics Array Hash Table 👍 114 👎 10

// leetcode submit region begin(Prohibit modification and deletion)
func unequalTriplets(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	cnt := 0
	cntMap := make(map[int]int, len(nums))
	cntArr := make([]int, 0, len(nums))

	for _, val := range nums {
		cntMap[val]++
	}

	for key := range cntMap {
		cntArr = append(cntArr, key)
	}

	for i := range cntArr {
		a := cntMap[cntArr[i]]
		for j := i + 1; j < len(cntArr); j++ {
			b := a * cntMap[cntArr[j]]
			for k := j + 1; k < len(cntArr); k++ {
				cnt += b * cntMap[cntArr[k]]
			}
		}
	}

	return cnt
}

//leetcode submit region end(Prohibit modification and deletion)
