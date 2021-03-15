package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestInvalid(t *testing.T) {
	result := log(-1, 0)
	assert.Equal(t, -1.0, result)
}

func TestNoIterations(t *testing.T) {
	result := log(25.3, 0)
	assert.Equal(t, 0.0, result)
}


func TestOne(t *testing.T) {
	result := log(1.0, 1)
	assert.Equal(t, 0.0, result)
}

func TestTwo(t *testing.T) {
	result := log(2.0, 2)
	assert.Equal(t, 0.691358024691358, result)
}

func TestEOneIteration(t *testing.T) {
	result := log(2.718281828459045, 1)
	assert.Equal(t, 0.9242343145200195, result)
}

func TestETenIteration(t *testing.T) {
	result := log(2.718281828459045, 10)
	assert.Equal(t, 	0.999999989210841, result)
}

func TestEOneHundredIteration(t *testing.T) {
	result := log(2.718281828459045, 100)
	assert.Equal(t, 0.9999999999999997, result)
}

func TestZeroOneIteration(t *testing.T) {
	result := log(0.0, 1)
	assert.Equal(t, -1.0, result)
}

func Test3(t *testing.T) {
	result := log(3.0, 14)
	assert.Equal(t, 1.0986122885003737, result)
}

func Test25(t *testing.T) {
	result := log(2.5, 1)
	assert.Equal(t, 0.8571428571428571, result)
}

func Test100(t *testing.T) {
	result := log(100.0, 1)
	assert.Equal(t, 1.9603960396039604, result)
}