package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	
	var expect float64 = 10
	result := add(5, 5)

	assert.Equal(t, expect, result)
}

func TestSub(t *testing.T) {

	var expect float64 = 10
	result := subtract(20, 10)

	assert.Equal(t, expect, result)
}

func TestMultiply(t *testing.T) {

	var expect float64 = 10
	result := multiply(5, 2)

	assert.Equal(t, expect, result)
}

func TestDivide(t *testing.T) {

	var expect float64 = 10

	result, err := divide(20, 2)

	assert.Nil(t, err)
	assert.Equal(t, expect, result)

	expect = 0

	_, err = divide(10, 0)

	assert.Error(t, err)

}