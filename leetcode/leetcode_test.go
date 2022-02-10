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