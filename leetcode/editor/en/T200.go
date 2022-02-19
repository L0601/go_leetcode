package leetcode

//Given an m x n 2D binary grid grid which represents a map of '1's (land) and
//'0's (water), return the number of islands.
//
// An island is surrounded by water and is formed by connecting adjacent lands
//horizontally or vertically. You may assume all four edges of the grid are all
//surrounded by water.
//
//
// Example 1:
//
//
//Input: grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
//]
//Output: 1
//
//
// Example 2:
//
//
//Input: grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
//]
//Output: 3
//
//
//
// Constraints:
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] is '0' or '1'.
//
// Related Topics Array Depth-First Search Breadth-First Search Union Find
//Matrix 👍 12305 👎 308

//leetcode submit region begin(Prohibit modification and deletion)
func numIslands(grid [][]byte) int {
	//m, n := len(grid), len(grid[0])
	//visited := make(map[int]bool, m)
	//steps := [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	//count := 0
	//
	//for i, row := range grid {
	//	for j, ch := range row {
	//		if ch == '1' && !visited[i*n+j] {
	//			count++
	//			visited[i*n+j] = true
	//			queue := [][2]int{{i, j}}
	//			for len(queue) > 0 {
	//				head := queue[0]
	//				queue = queue[1:]
	//				for _, step := range steps {
	//					x, y := head[0]+step[0], head[1]+step[1]
	//					if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '1' && !visited[x*n+y] {
	//						queue = append(queue, [2]int{x, y})
	//						visited[x*n+y] = true
	//					}
	//				}
	//			}
	//		}
	//	}
	//}
	//
	//return count

	return useInputArray(grid)
}

func useInputArray(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	steps := [4][2]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	count := 0

	for i, row := range grid {
		for j, ch := range row {
			if ch == '1' {
				count++
				grid[i][j] = '0'
				queue := [][2]int{{i, j}}
				for len(queue) > 0 {
					head := queue[0]
					queue = queue[1:]
					for _, step := range steps {
						x, y := head[0]+step[0], head[1]+step[1]
						if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == '1' {
							queue = append(queue, [2]int{x, y})
							grid[x][y] = '0'
						}
					}
				}
			}
		}
	}

	return count
}

//leetcode submit region end(Prohibit modification and deletion)
