/*
 * @lc app=leetcode id=391 lang=golang
 *
 * [391] Perfect Rectangle
 *
 * https://leetcode.com/problems/perfect-rectangle/description/
 *
 * algorithms
 * Hard (32.06%)
 * Likes:    591
 * Dislikes: 96
 * Total Accepted:    34K
 * Total Submissions: 106.1K
 * Testcase Example:  '[[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]'
 *
 * Given an array rectangles where rectangles[i] = [xi, yi, ai, bi] represents
 * an axis-aligned rectangle. The bottom-left point of the rectangle is (xi,
 * yi) and the top-right point of it is (ai, bi).
 *
 * Return true if all the rectangles together form an exact cover of a
 * rectangular region.
 *
 *
 * Example 1:
 *
 * Input: rectangles = [[1,1,3,3],[3,1,4,2],[3,2,4,4],[1,3,2,4],[2,3,3,4]]
 * Output: true
 * Explanation: All 5 rectangles together form an exact cover of a rectangular
 * region.
 *
 *
 * Example 2:
 *
 * Input: rectangles = [[1,1,2,3],[1,3,2,4],[3,1,4,2],[3,2,4,4]]
 * Output: false
 * Explanation: Because there is a gap between the two rectangular regions.
 *
 *
 * Example 3:
 *
 * Input: rectangles = [[1,1,3,3],[3,1,4,2],[1,3,2,4],[2,2,4,4]]
 * Output: false
 * Explanation: Because two of the rectangles overlap with each other.
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= rectangles.length <= 2 * 10^4
 * rectangles[i].length == 4
 * -10^5 <= xi, yi, ai, bi <= 10^5
 *
 *
 */

package leetcode

import "fmt"

// @lc code=start
func isRectangleCover(rectangles [][]int) bool {
	left, right, top, bottom := 100001, -100001, -100001, 100001
	mp := make(map[string]bool, len(rectangles))
	area := int64(0)
	for _, rec := range rectangles {
		area += int64((rec[2] - rec[0]) * (rec[3] - rec[1]))
		if rec[0] < left {
			left = rec[0]
		}
		if rec[1] < bottom {
			bottom = rec[1]
		}
		if rec[2] > right {
			right = rec[2]
		}
		if rec[3] > top {
			top = rec[3]
		}
		mp[fmt.Sprintf("%v_%v", rec[0], rec[1])] = true
		mp[fmt.Sprintf("%v_%v", rec[2], rec[3])] = true

	}

	return area == int64((top-bottom)*(right-left)) && mp[fmt.Sprintf("%v_%v", left, bottom)] && mp[fmt.Sprintf("%v_%v", left, top)] && mp[fmt.Sprintf("%v_%v", right, bottom)] && mp[fmt.Sprintf("%v_%v", right, top)]
}

// @lc code=end
