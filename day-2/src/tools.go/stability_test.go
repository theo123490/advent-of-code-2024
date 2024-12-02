package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsStable(t *testing.T) {
	increasingSlice := []int{1, 2, 3, 4}
	decreasingSlice := []int{10, 9, 8, 7}
	increaseUnstableSlice := []int{1, 2, 1, 4}
	decreasingUnstableSlice := []int{10, 9, 12, 1}
	unchangingSliceFirst := []int{1, 1, 2, 9}
	decreaseSaturatedSlice := []int{10, 4, 4, 1}
	increaseSaturatedSlice := []int{1, 5, 5, 9}
	increaseJump := []int{1, 5, 12, 13}
	decreasingJump := []int{10, 9, 4, 2}

	assert.True(t, IsStable(increasingSlice))
	assert.True(t, IsStable(decreasingSlice))
	assert.False(t, IsStable(increaseUnstableSlice))
	assert.False(t, IsStable(decreasingUnstableSlice))
	assert.False(t, IsStable(unchangingSliceFirst))
	assert.False(t, IsStable(decreaseSaturatedSlice))
	assert.False(t, IsStable(increaseSaturatedSlice))
	assert.False(t, IsStable(increaseJump))
	assert.False(t, IsStable(decreasingJump))
}

func TestIsStableWithDampener(t *testing.T) {
	increasingSlice := []int{1, 2, 3, 4, 5}
	decreasingSlice := []int{10, 9, 8, 7, 6}
	increasingDampenedSlice := []int{1, 0, 3, 4, 5}
	decreasingDampenendSlice := []int{10, 11, 8, 7, 6}
	unchangingSliceFirst := []int{1, 1, 2, 3, 4}
	unchangingSliceDecreaseFirst := []int{10, 10, 9, 6, 4}
	increaseJump := []int{1, 2, 3, 12, 6}
	decreasingJump := []int{10, 9, 4, 6, 3}

	tripleFailure := []int{1, 2, 3, 2, 2, 6, 7, 8, 9, 10}
	badStart := []int{3, 2, 1, 2, 3}
	badStartTwo := []int{20, 16, 14, 12, 10, 8, 7, 6}
	decreaseSaturatedSlice := []int{10, 4, 4, 4, 1}
	increaseSaturatedSlice := []int{1, 5, 5, 5, 9}

	assert.True(t, IsStableWithDampener(increasingSlice))
	assert.True(t, IsStableWithDampener(decreasingSlice))
	assert.True(t, IsStableWithDampener(increasingDampenedSlice))
	assert.True(t, IsStableWithDampener(unchangingSliceFirst))
	assert.True(t, IsStableWithDampener(unchangingSliceDecreaseFirst))
	assert.True(t, IsStableWithDampener(decreasingDampenendSlice))
	assert.True(t, IsStableWithDampener(increaseJump))
	assert.True(t, IsStableWithDampener(decreasingJump))
	assert.True(t, IsStableWithDampener(badStartTwo))

	assert.False(t, IsStableWithDampener(tripleFailure))
	assert.False(t, IsStableWithDampener(badStart))
	assert.False(t, IsStableWithDampener(decreaseSaturatedSlice))
	assert.False(t, IsStableWithDampener(increaseSaturatedSlice))
}
