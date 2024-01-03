package main

import (
	"fmt"
)

func numberOfBeams(bank []string) int {
	var curr, prev, ans int
	prev, ans = 0, 0
	for _, row := range bank {
		curr = 0
		for _, c := range row {
			if c == '1' {
				curr++
			}
		}
		if curr != 0 {
			ans += (curr * prev)
			prev = curr
		}
	}
	return ans
}


// -- Tests  ---
package main

import (
	"testing"
)

func Test_tc1(t *testing.T) {
	var bank []string = []string{"011001", "000000", "010100", "001000"}
	result := numberOfBeams(bank)
	if result != 8 {
		t.Errorf("Testcase 1: FAIL")
	}
}

func Test_tc2(t *testing.T) {
	var bank []string = []string{"000", "111", "000"}
	result := numberOfBeams(bank)
	if result != 0 {
		t.Errorf("Testcase 2: FAIL")
	}
}

func main() {
	fmt.Println("test 123...")
}
