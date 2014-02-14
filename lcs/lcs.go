package lcs

import ()

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
