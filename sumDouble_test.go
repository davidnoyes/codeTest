package codeTest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSum1(t *testing.T) {
	assert.Equal(t, sumDouble(1, 2), 3, "Should equal 3")
}

func TestSum2(t *testing.T) {
	assert.Equal(t, sumDouble(3, 2), 5, "Should equal 5")
}

func TestDoubleSum(t *testing.T) {
	assert.Equal(t, sumDouble(2, 2), 8, "Should equal 8")
}
