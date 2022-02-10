package leetcode

/*
 * @lc app=leetcode id=1601 lang=golang
 *
 * [1601] Maximum Number of Achievable Transfer Requests
 */

// @lc code=start
func maximumRequests(n int, requests [][]int) int {
	cnt := 1 << len(requests)
	max := 0
	for i := 0; i < cnt; i++ {
		tmp := test(i, &requests)
		if tmp > max {
			max = tmp
		}
	}

	return max
}

func test(cnt int, requests *[][]int) int {
	mp := make(map[int]int)
	success := 0
	for i := 0; i < len(*requests); i++ {
		if ((1 << i) & cnt) != 0 {
			mp[(*requests)[i][0]] -= 1
			mp[(*requests)[i][1]] += 1
			success++
		}
	}

	for _, v := range mp {
		if v != 0 {
			return 0
		}
	}

	return success
}

// @lc code=end
