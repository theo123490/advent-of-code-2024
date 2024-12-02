package main

import (
	"fmt"

	"aoc.2024/day-1/pkg/reader"
	"aoc.2024/day-1/src/tools.go"
)

func main() {
	fmt.Println(FirstSolution("input/first-question.txt"))
	fmt.Println("********")
	fmt.Println(SecondSolution("input/second-question.txt"))
}

func FirstSolution(fileName string) int {
	lines, err := reader.ReadPerLine(fileName)
	if err != nil {
		panic(err)
	}

	counter := 0

	for _, line := range lines {
		ints, err := reader.ConvertToInts(line)
		if err != nil {
			panic(err)
		}

		if tools.IsStable(ints) {
			counter++
		}
	}
	return counter
}

func SecondSolution(fileName string) int {
	lines, err := reader.ReadPerLine(fileName)
	if err != nil {
		panic(err)
	}

	counter := 0

	for _, line := range lines {
		ints, err := reader.ConvertToInts(line)
		if err != nil {
			panic(err)
		}

		if tools.IsStableWithDampener(ints) {
			counter++
		}
	}
	return counter
}
