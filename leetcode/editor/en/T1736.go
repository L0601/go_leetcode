package leetcode

//You are given a string time in the form of hh:mm, where some of the digits in
//the string are hidden (represented by ?).
//
// The valid times are those inclusively between 00:00 and 23:59.
//
// Return the latest valid time you can get from time by replacing the hidden
//digits.
//
//
// Example 1:
//
//
//Input: time = "2?:?0"
//Output: "23:50"
//Explanation: The latest hour beginning with the digit '2' is 23 and the
//latest minute ending with the digit '0' is 50.
//
//
// Example 2:
//
//
//Input: time = "0?:3?"
//Output: "09:39"
//
//
// Example 3:
//
//
//Input: time = "1?:22"
//Output: "19:22"
//
//
//
// Constraints:
//
//
// time is in the format hh:mm.
// It is guaranteed that you can produce a valid time from the given string.
//
//
// Related Topics String Greedy 👍 227 👎 129

//leetcode submit region begin(Prohibit modification and deletion)
func maximumTime(time string) string {
	res := make([]byte, 5, 5)
	for i, v := range time {
		switch v {
		case '?':
			switch i {
			case 0:
				next := time[1]
				if next == '?' {
					res[i] = '2'
				} else if next < '4' {
					res[i] = '2'
				} else {
					res[i] = '1'
				}
			case 1:
				pre := res[0]
				if pre == '2' {
					res[i] = '3'
				} else {
					res[i] = '9'
				}
			case 3:
				res[i] = '5'
			case 4:
				res[i] = '9'
			}
		default:
			res[i] = byte(v)
		}
	}
	return string(res)
}

//leetcode submit region end(Prohibit modification and deletion)
