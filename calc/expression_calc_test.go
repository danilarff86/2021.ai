package calc

import (
	"errors"
	"strings"

	"2021.ai/calc"

	"2021.ai/set"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

type IntegerSetLoaderFromMemory struct {
	data map[string]*set.IntegerSet
}

func (this *IntegerSetLoaderFromMemory) AddSet(filename string, s *set.IntegerSet) {
	this.data[filename] = s
}

func (this *IntegerSetLoaderFromMemory) ReadIntegerSetFromFile(filename string) (*set.IntegerSet, error) {
	if nil == this.data[filename] {
		return nil, errors.New("Not found")
	}
	return this.data[filename], nil
}

func TestExpression(t *testing.T) {
	l := &IntegerSetLoaderFromMemory{data: make(map[string]*set.IntegerSet)}
	l.AddSet("a.txt", set.CreateIntegerSetFromSlice([]int{1, 2, 3}))
	l.AddSet("b.txt", set.CreateIntegerSetFromSlice([]int{2, 3, 4}))
	l.AddSet("c.txt", set.CreateIntegerSetFromSlice([]int{3, 4, 5}))
	c := calc.CreateExpressionEvaluator(l, strings.Fields("[ SUM [ DIF a.txt b.txt c.txt ] [ INT b.txt c.txt ] ]"))
	s, err := c.Evaluate()
	require.Nil(t, err)
	assert.Equal(t, []int{1, 3, 4}, s.ToSortedSlice())
}
