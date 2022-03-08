package leetcode

/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */


// @lc code=start
func titleToNumber(columnTitle string) int {
	if len(columnTitle) == 0 {
		return 0
	}

	num := 0
	for i := range columnTitle {
		ch := columnTitle[i]
		num = num*26 + int(ch-'A') + 1
	}

	return num
}

// @lc code=end

