package tools

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func FindMul(line string) ([]string, error) {
	r, err := regexp.Compile(`mul\([0-9]+,[0-9]+\)`)
	if err != nil {
		return nil, fmt.Errorf("Error compiling regex: %v", err)
	}

	return r.FindAllString(line, -1), nil
}

func MultiplyMul(muls []string) (int, error) {
	total := 0
	for _, mul := range muls {
		trimMul := strings.TrimRight(strings.TrimLeft(mul, "mul("), ")")
		nums := strings.Split(trimMul, ",")
		a, err := strconv.Atoi(nums[0])
		if err != nil {
			return 0, fmt.Errorf("error converting string to int: %v", err)
		}
		b, err := strconv.Atoi(nums[1])
		if err != nil {
			return 0, fmt.Errorf("error converting string to int: %v", err)
		}
		total += a * b
	}
	return total, nil
}

func GetDo(line string) []string {
	dos := strings.Split(line, "do()")
	listOfDos := []string{}
	for _, do := range dos {
		removedDonts := strings.Split(do, "don't()")[0]
		listOfDos = append(listOfDos, removedDonts)
	}
	return listOfDos
}
