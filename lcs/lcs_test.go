package lcs

import (
	"fmt"
	"testing"
)

func TestLCS(t *testing.T) {
	s1 := []int{1, 1, 2, 3, 5, 8, 13}
	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	equal := func(a, b int) bool {
		return a == b
	}

	fmt.Println(s1)
	fmt.Println(s2)

	C := LCSLength(s1, s2, equal)
	fmt.Println(C)

	fmt.Println(backtrackAll(C, s1, s2, equal))
}
