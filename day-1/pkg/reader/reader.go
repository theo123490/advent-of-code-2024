package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ReadPerLine(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func GetTwoRow(lines []string) ([]int, []int, error) {
	firstRow := make([]int, len(lines))
	secondRow := make([]int, len(lines))
	var err error

	for i, line := range lines {
		elements := strings.Fields(line)
		if len(elements) != 2 {
			return nil, nil, fmt.Errorf("found 2 elements in element %v", line)
		}
		firstRow[i], err = strconv.Atoi(elements[0])
		if err != nil {
			return nil, nil, err
		}
		secondRow[i], err = strconv.Atoi(elements[1])
		if err != nil {
			return nil, nil, err
		}
	}
	return firstRow, secondRow, nil
}
