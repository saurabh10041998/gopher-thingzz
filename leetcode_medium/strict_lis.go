package main

import "fmt"

func lengthOfLIS(nums []int) int {
	n := len(nums)
	var dp []int = make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[j] < nums[i] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}
	var max_ele int = dp[0]
	for _, ele := range dp[1:] {
		if ele > max_ele {
			max_ele = ele
		}
	}
	return max_ele
}

// == Tests ==
package main

import "testing"

func Test_tc1(t *testing.T) {
	var nums []int = []int{10, 9, 2, 5, 3, 7, 101, 18}
	result := lengthOfLIS(nums)
	if result != 4 {
		t.Errorf("Testcase 1: FAIL")
	}
}

func Test_tc2(t *testing.T) {
	var nums []int = []int{0, 1, 0, 3, 2, 3}
	result := lengthOfLIS(nums)
	if result != 4 {
		t.Errorf("Testcase 2: FAIL")
	}
}

func Test_tc3(t *testing.T) {
	var nums []int = []int{7, 7, 7, 7, 7, 7, 7}
	result := lengthOfLIS(nums)
	if result != 1 {
		t.Errorf("Testcase 3: FAIL")
	}
}

func main() {
	fmt.Println("vim-go")
}
