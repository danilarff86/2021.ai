package set

import (
	"2021.ai/set"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum(t *testing.T) {
	s1 := set.CreateIntegerSetFromSlice([]int{1, 2, 3})
	s2 := set.CreateIntegerSetFromSlice([]int{2, 3, 4})
	set.Sum(s1, s2)
	assert.Equal(t, []int{1, 2, 3, 4}, s1.ToSortedSlice())
}

func TestIntersection(t *testing.T) {
	s1 := set.CreateIntegerSetFromSlice([]int{1, 2, 3})
	s2 := set.CreateIntegerSetFromSlice([]int{2, 3, 4})
	set.Intersection(s1, s2)
	assert.Equal(t, []int{2, 3}, s1.ToSortedSlice())
}

func TestDifference(t *testing.T) {
	s1 := set.CreateIntegerSetFromSlice([]int{1, 2, 3})
	s2 := set.CreateIntegerSetFromSlice([]int{2, 3, 4})
	set.Difference(s1, s2)
	assert.Equal(t, []int{1}, s1.ToSortedSlice())
}
