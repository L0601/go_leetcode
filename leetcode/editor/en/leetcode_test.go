package leetcode

import (
	"fmt"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	fmt.Print(twoSum([]int{1, 2, 3, 4, 5, 6}, 3))
}

func TestGCD(t *testing.T) {
	fmt.Println(gcd(3, 4))
	fmt.Println(gcd(4, 6))
	fmt.Println(gcd(6, 4))

	fmt.Println(transform(4, 6))
}

func TestFibo(t *testing.T) {
	ar := generateFibo(1000000000)
	for _, i := range ar {
		fmt.Print(i, ",")
	}
}

func TestAdd(t *testing.T) {
	fmt.Println(add("1234", 4))
}

func TestSegement(t *testing.T) {
	fmt.Println(countSegments("  a dfja fwfew  we "))
	fmt.Println(countSegments("Hello, my name is John"))
}

func TestIsland(t *testing.T) {
	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
	fmt.Println(numIslands([][]byte{{'1', '0', '1', '1', '1'}, {'1', '0', '1', '0', '1'}, {'1', '1', '1', '0', '1'}}))
}
