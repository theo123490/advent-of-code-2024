package reader

import (
	"bufio"
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

func ConvertToInts(line string) ([]int, error) {
	var ints []int
	stringElements := strings.Fields(line)
	for _, e := range stringElements {
		i, err := strconv.Atoi(e)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}
	return ints, nil
}
