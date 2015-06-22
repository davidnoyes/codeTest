package codeTest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWeekendNoVacation(t *testing.T) {
	assert.True(t, SleepIn(false, false), "Should be true")
}

func TestWeekdayNoVacation(t *testing.T) {
	assert.False(t, SleepIn(true, false), "Should be false")
}

func TestWeekendVacation(t *testing.T) {
	assert.True(t, SleepIn(false, true), "Should be true")
}
