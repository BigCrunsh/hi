package lcs

import (
	"sync"
)

type eqFunc func(int, int) bool

func LCSLength(s1, s2 []int, equal eqFunc) [][]int {
	n, m := len(s1), len(s2)
	C := make([][]int, n)

	for i := 0; i < n; i++ {
		C[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if equal(s1[i], s2[j]) {
				if i == 0 || j == 0 {
					C[i][j] = 1
				} else {
					C[i][j] = C[i-1][j-1] + 1
				}
			} else {
				if i == 0 || j == 0 {
					C[i][j] = 1
				} else {
					if C[i][j-1] > C[i-1][j] {
						C[i][j] = C[i][j-1]
					} else {
						C[i][j] = C[i-1][j]
					}
				}

			}
		}
	}

	return C
}

type setElement []int
type set []setElement

func backtrackAll(C [][]int, s1, s2 []int, equal eqFunc) set {
	return backtrackIJ(C, s1, s2, equal, len(s1)-1, len(s2)-1)
}

func (s *set) addPostfix(postfix int) {
	for i, e := range *s {
		(*s)[i] = append(e, postfix)
	}
}

func geq(a, b setElement) bool {
	if len(a) > len(b) {
		return true
	}

	if len(a) < len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return true
		}
		if a[i] < b[i] {
			return false
		}
	}

	// equal
	return true
}

// assumed that this and that are sorted
// TODO(cs): make set unique
func (this *set) union(that set) {
	for _, v := range that {
		*this = append(*this, v)
	}
}

// add positions to match set
func backtrackIJ(C [][]int, s1, s2 []int, equal eqFunc, i, j int) set {
	if equal(s1[i], s2[j]) {
		var res set
		if i == 0 || j == 0 {
			res = make(set, 1)
		} else {
			res = backtrackIJ(C, s1, s2, equal, i-1, j-1)
		}
		res.addPostfix(s1[i])
		return res
	}

	res := make(set, 0)
	if C[i][j-1] >= C[i-1][j] {
		res = backtrackIJ(C, s1, s2, equal, i, j-1)
	}

	if C[i-1][j] >= C[i][j-1] {
		res.union(backtrackIJ(C, s1, s2, equal, i-1, j))
	}

	return res
}

//
// ------------------
//

type subsequence struct {
	Items []matchedItem
}

type matchedItem struct {
	Value     int
	PositionX int
	PositionY int
}

type extractSubSeqs func(x, y []int, equal eqFunc, minLength, maxError int) []subsequence

// Find one subsequence starting at given positions.
// Advance the position for x on a match and for y always.
// Break when maxError is reached and only return a
// subsequence if the match is longer than minLength. 
func singleMatch(x, y []int, equal eqFunc, minLength, maxError, xPos, yPos int) *subsequence {
	var match *subsequence

	var buffer []matchedItem
	matchErrors := 0
	for (xPos < len(x)) && (yPos < len(y)) && (matchErrors <= maxError) {
		if equal(x[xPos], y[yPos]) {
			buffer = append(buffer, matchedItem{x[xPos], xPos, yPos})
			xPos++
		} else {
			matchErrors++
		}

		if matchErrors <= maxError && len(buffer) >= minLength {
			match = &subsequence{buffer}
		}

		yPos++ 
	}

	return match
}

// Find all subsequences
func GetSeqs(x, y []int, equal eqFunc, minLength, maxError int) []subsequence {
	result := []subsequence{}

	for xPos := 0; xPos < len(x); xPos++ {
		m := singleMatch(x, y, equal, minLength, maxError, xPos, 0)
		if m != nil {
			result = append(result, *m)
		}

		// call mirrored
		m = singleMatch(y, x, equal, minLength, maxError, 0, xPos)
		if m != nil {
			result = append(result, *m)
		}
	}

	return result
}

// Find all subsequences by spawning a go routine for each item in x
func GetSeqsConcurrently(x, y []int, equal eqFunc, minLength, maxError int) []subsequence {
	result := make(chan *subsequence)

	var wg sync.WaitGroup
	for xPos := 0; xPos < len(x); xPos++ {
		wg.Add(1)
		go func(result chan *subsequence, x, y []int, equal eqFunc, minLength, maxError, xPos int) {
			defer wg.Done()

			m := singleMatch(x, y, equal, minLength, maxError, xPos, 0)
			if m != nil {
				result <- m
			}

			// call mirrored
			m = singleMatch(y, x, equal, minLength, maxError, 0, xPos)
			if m != nil {
				result <- m
			}
		}(result, x, y, equal, minLength, maxError, xPos)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	subseqs := []subsequence{}
	for res := range result {
		subseqs = append(subseqs, *res)
	}

	return subseqs
}

// Go in with one sequence, retrieve the rest from an index including positions.
func GetSeqsFromIndex(x []int, equal eqFunc, minLength, maxError int) []subsequence {
	result := []subsequence{}

	for xPos := 0; xPos < len(x); xPos++ {
		allYs := indexLookup(x[xPos])

		for _, y := range(allYs) {
			m := singleMatch(x, y.Seq, equal, minLength, maxError, xPos, y.Pos)
			if m != nil {
				result = append(result, *m)
			}

			// call mirrored
			m = singleMatch(y.Seq, x, equal, minLength, maxError, y.Pos, xPos)
			if m != nil {
				result = append(result, *m)
			}
		}
	}

	return result
}

type lookupResult struct {
	Seq []int
	Pos int
}

var sequences [][]int  // data in the pseudo index

// Replace this in production by a large and efficient index lookup.
// FLAW: this assumes that it is possible to do a lookup, which means
// that the equal function is just '==' after some normalization.
func indexLookup(itemValue int) []lookupResult {
	result := []lookupResult{}
	for _, seq := range(sequences) {
		for i, v := range(seq) {
			if v == itemValue {
				result = append(result, lookupResult{seq, i})
			}
		}
	}
	return result
}
