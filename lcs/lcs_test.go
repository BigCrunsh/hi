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
	expected := [][]int{
		[]int{1, 2, 3, 5},
	}
	execTest(t, GetSeqs, x, y, 3, 1, expected)
}

// FIXME!!! this test fails currently; remove prefix 'x' to run
func xTestGetSeqsFlipped(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	y := []int{1, 1, 2, 3, 5, 8, 13}
	expected := [][]int{
		[]int{1, 2, 3, 5},
	}
	execTest(t, GetSeqs, x, y, 3, 1, expected)
}

func TestGetSeqsConcurrently(t *testing.T) {
	x := []int{1, 1, 2, 3, 5, 8, 13}
	y := []int{1, 2, 3, 4, 5, 6, 7}
	expected := [][]int{
		[]int{1, 2, 3, 5},
	}
	execTest(t, GetSeqsConcurrently, x, y, 3, 1, expected)
}

// FIXME!!! this test fails currently; remove prefix 'x' to run
func xTestGetSeqsConcurrentlyFlipped(t *testing.T) {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	y := []int{1, 1, 2, 3, 5, 8, 13}
	expected := [][]int{
		[]int{1, 2, 3, 5},
	}
	execTest(t, GetSeqsConcurrently, x, y, 3, 1, expected)
}

// ---

func execTest(t *testing.T, fun seqFun, x, y []int, minLength, maxErrors int, expected [][]int) {
	results := fun(x, y, equal, minLength, maxErrors)

	if len(expected) != len(results) {
		t.Fatalf("Number of subsequences incorrect; expected %d, got %d", len(expected), len(results))
	}

	for i, result := range results {
		exp := expected[i] 

		if len(exp) != len(result) {
			fmt.Println("x:", x, "y:", y, "result:", result)
			t.Fatalf("Number of items in subsequence incorrect; expected %d, got %d", len(exp), len(result))
		}

		for j, item := range result {
			if item != exp[j] {
				fmt.Println("x:", x, "y:", y, "result:", result)
				t.Fatalf("int at index %d was not %d but was %d", i, exp[j], item)
			}
		}
	}
}
