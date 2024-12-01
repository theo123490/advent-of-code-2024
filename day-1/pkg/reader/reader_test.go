package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTwoRow(t *testing.T) {
	testInput := []string{
		"1   3",
		"2   4",
		"2   4",
		"5   3",
		"2   4",
		"1   2",
	}
	expectedResultFirstRow := []int{1, 2, 2, 5, 2, 1}
	expectedResultSecondRow := []int{3, 4, 4, 3, 4, 2}
	resultFistRow, resultSecondRow, err := GetTwoRow(testInput)
	assert.Equal(t, expectedResultFirstRow, resultFistRow)
	assert.Equal(t, expectedResultSecondRow, resultSecondRow)
	assert.NoError(t, err)
}
