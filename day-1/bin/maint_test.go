package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var DAY string = "day-1"
var INPUT_DIR string = fmt.Sprintf("/Users/theodore.chandra/personal/code/advent-of-code-2024/%v/input/", DAY)

func TestFirstSolution(t *testing.T) {
	result := FirstSolution(fmt.Sprintf("%v/first-question-test.txt", INPUT_DIR))
	assert.Equal(t, 11, result)
}

func TestSecondSolution(t *testing.T) {
	result := SecondSolution(fmt.Sprintf("%v/second-question-test.txt", INPUT_DIR))
	assert.Equal(t, 31, result)
}
