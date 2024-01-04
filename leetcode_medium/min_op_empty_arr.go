package main

import "fmt"

func minOperations(nums []int) int {
	var mp map[int]int = map[int]int{}
	for _, n := range nums {
		mp[n]++
	}
	var res int = 0
	for _, v := range mp {
		if v == 1 {
			return -1
		}
		if mod := v % 3; mod == 0 {
			res += v / 3
		} else if mod == 1 {
			res += 2 + (v-4)/3
		} else {
			res += 1 + (v-2)/3
		}
	}
	return res
}


//=== TESTS ===
package main

import (
	"testing"
)

func Test_tc1(t *testing.T) {
	var nums []int = []int{2, 3, 3, 2, 2, 4, 2, 3, 4}
	result := minOperations(nums)
	if result != 4 {
		t.Errorf("Testcase 1: FAIL")
	}
}

func Test_tc2(t *testing.T) {
	var nums []int = []int{2, 1, 2, 2, 3, 3}
	result := minOperations(nums)
	if result != -1 {
		t.Errorf("Testcase 2: FAIL")
	}
}


func main() {
	fmt.Println("vim-go")
}
