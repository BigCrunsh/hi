package set

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	data := []int{1, 2, 3, 2, 4, 1, 1}
	expectedReturns := []bool{true, true, true, false, true, false, false}

	s := NewSet()
	for i, expected := range expectedReturns {
		el := new(intElement)
		el.data = data[i]

		if got := s.Add(el); got != expected {
			t.Fatalf("add: expected %v but got %v", expected, got)
		}

	}
}

func TestCardinality(t *testing.T) {
	data := []int{1, 2, 3, 2, 4, 1, 1}
	expectedSizes := []int{0, 1, 2, 3, 3, 4, 4, 4}

	s := NewSet()
	for i, expected := range expectedSizes {
		if got := s.Cardinality(); got != expected {
			t.Fatalf("size: expected %d but got %d", expected, got)
		}

		if i < len(data) {
			el := new(intElement)
			el.data = data[i]
			s.Add(el)
		}
	}
}

func TestEquals(t *testing.T) {
	data1 := []int{2, 1, 3}
	data2 := []int{1, 2, 3, 3, 4}
	expectedEquals := []bool{false, true, true, true, false}

	s1, s2 := NewSet(), NewSet()
	for i, expected := range expectedEquals {
		// add element to sets
		if i < len(data1) {
			el1 := new(intElement)
			el1.data = data1[i]
			s1.Add(el1)
		}

		if i < len(data2) {
			el2 := new(intElement)
			el2.data = data2[i]
			s2.Add(el2)
		}

		if got := s1.Equals(s2); got != expected {
			t.Fatalf("equals: expected %v but got %v", expected, got)
		}

		if got := s1.Equals(s1); got != true {
			t.Fatalf("equals: set should be equal to itself")
		}
	}
}

func TestStructuredSet(t *testing.T) {
	i, j := new(structElement), new(structElement)
	i.data = []int{1, 2, 3}
	j.data = []int{2, 2, 3}

	s := NewSet()
	s.Add(i)

	if !s.Contains(i) {
		t.Fatalf("set should contain %v", i)
	}

	if s.Contains(j) {
		t.Fatalf("set contains only %v but also %v found", i, j)
	}
}

type intElement struct {
	data int
}

func NewIntElement(i int) intElement {
	var el intElement
	el.data = i
	return el
}

func (s *intElement) Hash() string {
	return fmt.Sprintf("%v", (*s).data)
}

type structElement struct {
	data []int
}

func (s *structElement) Hash() string {
	return fmt.Sprintf("%v", (*s).data)
}
