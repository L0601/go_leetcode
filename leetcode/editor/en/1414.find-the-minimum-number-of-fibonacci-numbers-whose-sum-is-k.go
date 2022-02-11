/*
 * @lc app=leetcode id=1414 lang=golang
 *
 * [1414] Find the Minimum Number of Fibonacci Numbers Whose Sum Is K
 */

package leetcode

// @lc code=start
func findMinFibonacciNumbers(k int) int {
	fibo := []int{1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711, 28657, 46368, 75025, 121393, 196418, 317811, 514229, 832040, 1346269, 2178309, 3524578, 5702887, 9227465, 14930352, 24157817, 39088169, 63245986, 102334155, 165580141, 267914296, 433494437, 701408733}

	cnt := 0
	endIdx := len(fibo) - 1
	for k > 0 {
		maxIdx := find(fibo, k, endIdx)
		k -= fibo[maxIdx]
		cnt += 1
		endIdx = maxIdx
	}

	return cnt
}

func find(array []int, target int, endIdx int) int {
	for endIdx >= 0 {
		if target >= array[endIdx] {
			return endIdx
		}
		endIdx -= 1
	}

	return 0
}

func generateFibo(max int) []int {
	a, b := 1, 1
	result := make([]int, 0, 40)
	for a < max {
		result = append(result, a)
		a, b = b, a+b
	}

	result = append(result, a)
	return result
}

// @lc code=end
