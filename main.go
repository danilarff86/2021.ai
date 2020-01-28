package main

import (
	"fmt"
	"os"

	"2021.ai/calc"
	"2021.ai/loader"
)

func main() {
	argsWithoutProg := os.Args[1:]
	c := calc.CreateExpressionEvaluator(&loader.IntegerSetLoaderFromFile{}, argsWithoutProg)
	res, err := c.Evaluate()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range res.ToSortedSlice() {
		fmt.Println(v)
	}
}
