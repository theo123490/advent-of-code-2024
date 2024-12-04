package main

import (
	"fmt"
	"strings"

	"aoc.2024/day-3/pkg/reader"
	"aoc.2024/day-3/src/tools.go"
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
	total := 0
	for _, line := range lines {
		muls, err := tools.FindMul(line)
		if err != nil {
			panic(err)
		}
		lineTotal, err := tools.MultiplyMul(muls)
		if err != nil {
			panic(err)
		}
		total += lineTotal
	}
	return total
}

func SecondSolution(fileName string) int {
	lines, err := reader.ReadPerLine(fileName)
	line := strings.Join(lines, "")
	if err != nil {
		panic(err)
	}
	total := 0
	dos := tools.GetDo(line)
	for _, do := range dos {
		muls, err := tools.FindMul(do)
		if err != nil {
			panic(err)
		}
		doTotal, err := tools.MultiplyMul(muls)
		if err != nil {
			panic(err)
		}
		total += doTotal
	}
	return total
}
