package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	assert.True(t, isValid([]int32{1, 3, 2}))
	assert.True(t, isValid([]int32{2, 1, 3}))
	assert.True(t, isValid([]int32{3, 2, 1, 5, 4, 6}))
	assert.True(t, isValid([]int32{5, 2, 1, 3}))
	assert.True(t, isValid([]int32{5, 3, 1, 2}))
	assert.False(t, isValid([]int32{1, 3, 4, 2}))
}

func Test2(t *testing.T) {
	assert.True(t, isValid([]int32{3, 2, 1}))
	assert.True(t, isValid([]int32{3}))
	assert.False(t, isValid([]int32{2, 3, 1}))
	assert.True(t, isValid([]int32{2, 1, 3}))
	assert.False(t, isValid([]int32{3, 1, 4, 2}))
}

func Test3(t *testing.T) {
	assert.False(t, isValid([]int32{17, 7, 2, 5, 11, 10, 3, 12, 1, 18, 15, 13, 8, 4, 16, 6, 19, 14, 9}))
	assert.True(t, isValid([]int32{6, 2, 1, 5, 4, 3, 29, 14, 8, 7, 11, 10, 9, 12, 13, 24, 17, 16, 15, 18, 23, 21, 20, 19, 22, 26, 25, 27, 28}))
}