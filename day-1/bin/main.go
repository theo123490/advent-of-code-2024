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
	rowA, rowB, err := reader.GetTwoRow(lines)
	if err != nil {
		panic(err)
	}
	return tools.CalculateDistance(rowA, rowB)
}

func SecondSolution(fileName string) int {
	lines, err := reader.ReadPerLine(fileName)
	if err != nil {
		panic(err)
	}
	rowA, rowB, err := reader.GetTwoRow(lines)
	if err != nil {
		panic(err)
	}
	return tools.CalculateOccurance(rowA, rowB)
}
