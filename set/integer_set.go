package set

import (
	"sort"
)

type void struct{}

var member void

type IntegerSetOperationFunc func(*IntegerSet, *IntegerSet)

type IntegerSet struct {
	data map[int]void
}

func CreateIntegerSet() *IntegerSet {
	return &IntegerSet{data: make(map[int]void)}
}

func (s *IntegerSet) AddElement(v int) {
	s.data[v] = member
}

func (s *IntegerSet) RemoveElement(v int) {
	delete(s.data, v)
}

func (s *IntegerSet) ToSortedSlice() []int {
	slice := []int{}
	for v := range s.data {
		slice = append(slice, v)
	}

	sort.Ints(slice)

	return slice
}

func Sum(lhs *IntegerSet, rhs *IntegerSet) {
	for v := range rhs.data {
		lhs.data[v] = member
	}
}

func Intersection(lhs *IntegerSet, rhs *IntegerSet) {
	for v := range lhs.data {
		_, exists := rhs.data[v]
		if !exists {
			delete(lhs.data, v)
		}
	}
}

func Difference(lhs *IntegerSet, rhs *IntegerSet) {
	for v := range rhs.data {
		_, exists := lhs.data[v]
		if exists {
			delete(lhs.data, v)
		}
	}
}
