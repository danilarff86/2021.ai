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
	var operation set.IntegerSetOperationFunc

	if ee.index >= len(ee.tokens) {
		return nil, errors.New("Invalid expression")
	}

	switch t := ee.tokens[ee.index]; t {
	case "SUM":
		operation = set.Sum
	case "DIF":
		operation = set.Difference
	case "INT":
		operation = set.Intersection
	default:
		return nil, errors.New("Invalid expression")
	}

	ee.index++

	return ee.performOperation(operation)
}

func (ee *ExpressionEvaluator) performOperation(operation set.IntegerSetOperationFunc) (*set.IntegerSet, error) {
	var arg *set.IntegerSet
	var res *set.IntegerSet
	var err error
	for ee.index < len(ee.tokens) {
		t := ee.tokens[ee.index]
		ee.index++

		switch t {
		case "[":
			arg, err = ee.operationLevel()
		case "]":
			return res, nil
		default:
			arg, err = loader.ReadIntegerSetFromFile(t)
		}

		if err != nil {
			return nil, err
		}

		if res == nil {
			res = arg
		} else {
			operation(res, arg)
		}
	}

	return res, nil
}
