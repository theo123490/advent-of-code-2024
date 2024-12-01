package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDistance(t *testing.T) {
	listA := []int{1, 2, 3, 4}
	listB := []int{6, 3, 2, 4}

	result := CalculateDistance(listA, listB)
	assert.Equal(t, 5, result)
}

func TestCalculateOccurance(t *testing.T) {
	listA := []int{1, 2, 3, 3}
	listB := []int{6, 3, 2, 4}
	result := CalculateOccurance(listA, listB)
	assert.Equal(t, 8, result)
}
