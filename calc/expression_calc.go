package calc

import (
	"errors"

	"2021.ai/loader"
	"2021.ai/set"
)

type ExpressionEvaluator struct {
	tokens []string
	index  int
}

func CreateExpressionEvaluator(tokens []string) *ExpressionEvaluator {
	return &ExpressionEvaluator{tokens: tokens, index: 0}
}

func (ee *ExpressionEvaluator) Evaluate() (*set.IntegerSet, error) {
	if len(ee.tokens) == 0 || ee.tokens[ee.index] != "[" {
		return nil, errors.New("Invalid expression")
	}
	ee.index++
	return ee.operationLevel()
}

func (ee *ExpressionEvaluator) operationLevel() (*set.IntegerSet, error) {
	res := set.CreateIntegerSet()
	var op set.IntegerSetOperationFunc

	if ee.index >= len(ee.tokens) {
		return nil, errors.New("Invalid expression")
	}

	switch t := ee.tokens[ee.index]; t {
	case "SUM":
		op = res.Sum
	case "DIF":
		op = res.Difference
	case "INT":
		op = res.Intersection
	default:
		return nil, errors.New("Invalid expression")
	}

	ee.index++

	return res, ee.performOperation(op, res)
}

func (ee *ExpressionEvaluator) performOperation(op set.IntegerSetOperationFunc, res *set.IntegerSet) error {
	var arg *set.IntegerSet
	var err error
	firstArg := true
	for ee.index < len(ee.tokens) {
		t := ee.tokens[ee.index]
		ee.index++

		if t == "[" {
			arg, err = ee.operationLevel()
		} else if t == "]" {
			break
		} else {
			arg, err = loader.ReadIntegerSetFromFile(t)
		}

		if err != nil {
			return err
		}

		if firstArg {
			res.Sum(arg)
			firstArg = false
		} else {
			op(arg)
		}
	}

	return nil
}
