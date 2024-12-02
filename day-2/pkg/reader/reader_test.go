package reader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToInts(t *testing.T) {
	input := "1 2 3 4"
	expected := []int{1, 2, 3, 4}
	result, err := ConvertToInts(input)
	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}
