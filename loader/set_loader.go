package loader

import (
	"2021.ai/set"
	"bufio"
	"log"
	"os"
	"strconv"
)

type IntegerSetLoader interface {
	ReadIntegerSetFromFile(filename string) (*set.IntegerSet, error)
}

type IntegerSetLoaderFromFile struct {
}

func (*IntegerSetLoaderFromFile) ReadIntegerSetFromFile(filename string) (*set.IntegerSet, error) {
	readFile, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
		return nil, err
	}

	res := set.CreateIntegerSet()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		x, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			return res, err
		}
		res.AddElement(x)
	}

	return res, fileScanner.Err()
}
