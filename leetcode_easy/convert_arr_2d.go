package main

import "fmt"

func findMatrix(nums []int) [][]int {
	var mp map[int]int = map[int]int{}
	var ans [][]int
	for _, num := range nums {
		mp[num]++
		for mp[num]-1 >= len(ans) {
			ans = append(ans, []int{})
		}
		ans[mp[num]-1] = append(ans[mp[num]-1], num)
	}
	return ans
}

func main() {
	fmt.Println("vim-go")
}

//-- d2_test.go --

package main

import (
	"testing"
)

func verifyConfiguration(nums [][]int, expectedRows int) bool {
	for _, rowVec := range nums {
		var set map[int]bool = map[int]bool{}
		for _, num := range rowVec {
			set[num] = true
		}
		if len(set) != len(rowVec) {
			return false
		}
	}
	return true && len(nums) == expectedRows
}

func Test_tc1(t *testing.T) {
	var nums []int = []int{1, 3, 4, 1, 2, 3, 1}
	var result [][]int = findMatrix(nums)
	ok := verifyConfiguration(result, 3)
	if !ok {
		t.Errorf("Incorrect result for the testcase 1\n")
	}
}

func Test_tc2(t *testing.T) {
	var nums []int = []int{1, 2, 3, 4}
	var result [][]int = findMatrix(nums)
	ok := verifyConfiguration(result, 1)
	if !ok {
		t.Errorf("Incorrect result for the testcase 2\n")
	}
}
