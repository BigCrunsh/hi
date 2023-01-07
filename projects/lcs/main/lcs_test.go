package lcs

import (
	"fmt"
	"testing"
)

func equal(a, b int) bool {
	return a == b
}

func TestLCS(t *testing.T) {
	s1 := []int{1, 1, 2, 3, 5, 8, 13}
	s2 := []int{1, 2, 3, 4, 5, 6, 7}

	fmt.Println(s1)
	fmt.Println(s2)

	C := LCSLength(s1, s2, equal)
	fmt.Println(C)

	fmt.Println(backtrackAll(C, s1, s2, equal))
}


//
// ------------------
//

func TestGetSeqs(t *testing.T) {
	x := []int{1, 1, 2, 3, 5, 8, 13}
	y := []int{1, 2, 3, 4, 5, 6, 7}
	expected := []subsequence{
		subsequence{ []matchedItem{
			matchedItem{1, 0, 0},
			matchedItem{2, 2, 1},
			matchedItem{3, 3, 2},
			matchedItem{5, 4, 4},
		}},
	}
	execTest(t, GetSeqs, x, y, 3, 1, expected)
}

func execTest(t *testing.T, fun extractSubSeqs, x, y []int, minLength, maxErrors int, expected []subsequence) {
	results := fun(x, y, equal, minLength, maxErrors)

	if len(expected) != len(results) {
		fmt.Println("expected:", expected, "results:", results)
		t.Fatalf("Number of subsequences incorrect; expected %d, got %d", len(expected), len(results))
	}

	for i, result := range results {
		exp := expected[i] 

		if len(exp.Items) != len(result.Items) {
			fmt.Println("x:", x, "y:", y, "result:", result)
			t.Fatalf("Number of items in subsequence incorrect; expected %d, got %d", len(exp.Items), len(result.Items))
		}

		for j, item := range result.Items {
			if item.Value != exp.Items[j].Value {
				fmt.Println("x:", x, "y:", y, "result:", result)
				t.Fatalf("value at index %d was not %d but was %d", i, exp.Items[j], item)
			}
		}
	}
}
