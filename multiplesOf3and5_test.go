package codeTest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplesTo10(t *testing.T) {
	assert := assert.New(t)
	values := MultiplesOf3And5(10)

	assert.Equal(len(values), 4, "Should be 4")

	assert.Equal(values[0], 3, "Should be 3")
	assert.Equal(values[1], 5, "Should be 5")
	assert.Equal(values[2], 6, "Should be 6")
	assert.Equal(values[3], 9, "Should be 9")
}

func TestSumMultiplesTo10(t *testing.T) {
	values := MultiplesOf3And5(10)
	assert.Equal(t, SumMultiples(values), 23, "Should be equal to 23")
}
