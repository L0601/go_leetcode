/*
 * @lc app=leetcode id=2345 lang=golang
 *
 * [2224] Minimum Number of Operations to Convert Time
 */

package leetcode

import (
	"strconv"
	"strings"
)

// @lc code=start
func convertTime(current string, correct string) int {
	diff := int(str2minute(correct) - str2minute(current))
	if diff > 24*3600-diff {
		// diff > 0
		diff = 24*3600 - diff
	}
	if diff < 0 {
		diff = 24*3600 + diff
	}

	return diff/60 + (diff%60)/15 + (diff%15)/5 + diff%5
}

func str2minute(s string) int64 {
	ss := strings.Split(s, ":")
	a, _ := strconv.ParseInt(ss[0], 10, 64)
	b, _ := strconv.ParseInt(ss[1], 10, 64)

	return a*60 + b
}

// @lc code=end
