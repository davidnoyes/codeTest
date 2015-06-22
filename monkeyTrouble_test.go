package codeTest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBothSmiling(t *testing.T) {
	assert.True(t, MonkeyTrouble(true, true), "We should be in trouble")
}

func TestNoSmiling(t *testing.T) {
	assert.True(t, MonkeyTrouble(false, false), "We should be in trouble")
}

func TestASmiling(t *testing.T) {
	assert.False(t, MonkeyTrouble(true, false), "Monkey A smiling")
}
