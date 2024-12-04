package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMul(t *testing.T) {
	testString := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	expected := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}
	result, err := FindMul(testString)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.Equal(t, expected, result)

	emptyMul := "sdasagfua0iafhafoaobjoaoas"
	expected = []string(nil)
	result, err = FindMul(emptyMul)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.Equal(t, expected, result)
}

func TestMultiplyMul(t *testing.T) {
	testCase := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}
	result, err := MultiplyMul(testCase)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	assert.Equal(t, 161, result)
}

func TestGetDo(t *testing.T) {
	testString := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	result := GetDo(testString)
	expected := []string{"xmul(2,4)&mul[3,7]!^", "?mul(8,5))"}
	assert.Equal(t, expected, result)
}
