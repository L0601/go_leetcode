package leetcode

/*
 * @lc app=leetcode id=2001 lang=golang
 *
 * [2001] Number of Pairs of Interchangeable Rectangles
 */

// @lc code=start
func interchangeableRectangles(rectangles [][]int) int64 {
	offset := 100000
	mp := make(map[int64]int64, len(rectangles))

	for _, r := range rectangles {
		a, b := r[0], r[1]
		a, b = transform(a, b)
		key := int64(a*offset + b)
		mp[key] += 1
	}

	count := int64(0)
	for _, v := range mp {
		count += (v - 1) * v / 2
	}

	return count
}

func gcd(a, b int) int {
	mod := a % b
	if mod == 0 {
		return b
	}

	return gcd(b, mod)
}

func transform(a, b int) (int, int) {
	g := gcd(a, b)
	return a / g, b / g
}


// @lc code=end
