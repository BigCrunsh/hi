package set

import (
	"fmt"
)

type SetElement interface {
	Hash() string
}

type Set struct {
	data map[string]SetElement
}

func NewSet() Set {
	var s Set
	s.data = make(map[string]SetElement)
	return s
}

func (s *Set) Add(a SetElement) bool {
	if (*s).Contains(a) {
		return false
	}

	(*s).data[a.Hash()] = a

	return true
}

func (s *Set) Union(that Set) {
	for _, v := range that.data {
		(*s).Add(v)
	}
}

func (s *Set) Cardinality() int {
	return len((*s).data)
}

func (s *Set) Contains(a SetElement) bool {
	_, exists := (*s).data[a.Hash()]
	return exists
}

// TODO(cs): replace DeepEqual by comparisons of SetElements
func (s *Set) Equals(that Set) bool {
	if len((*s).data) != len(that.data) {
		return false
	}

	for _, v := range (*s).data {
		if !that.Contains(v) {
			return false
		}
	}

	return true
}

func (s *Set) String() string {
	return fmt.Sprintf("%v\n", *s)
}
