package set

import (
	"fmt"
	"reflect"
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

func (s *Set) Cardinality() int {
	return len((*s).data)
}

func (s *Set) Contains(a SetElement) bool {
	_, exists := (*s).data[a.Hash()]
	return exists
}

// TODO(cs): replace DeepEqual by comparisons of SetElements
func (s *Set) Equals(that Set) bool {
	return reflect.DeepEqual(*s, that)
}

func (s *Set) String() string {
	return fmt.Sprintf("%v\n", *s)
}
