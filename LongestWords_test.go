package codeTest

import (
	"github.com/raspberrypython/testify/assert"
	"testing"
)

func TestLongestWord(t *testing.T) {
	assert := assert.New(t)
	answers := LongestWords("You are just an old antidisestablishmentarian")
	assert.Equal(len(answers), 1, "Should be only 1 word")

	assert.Equal(answers[0], "antidisestablishmentarian", "Should be antidisestablishmentarian")
}

func TestLongestWords(t *testing.T) {
	assert := assert.New(t)
	answers := LongestWords("I gave a present to my parents")
	assert.Equal(len(answers), 2, "Should be 2 words")

	assert.True(answers[0] == "present" && answers[1] == "parents", "Should be [\"present\", \"parents\"]")
}

func TestLongestWordNoDupe(t *testing.T) {
	assert := assert.New(t)
	answers := LongestWords("Buffalo buffalo Buffalo buffalo buffalo buffalo Buffalo buffalo")
	assert.Equal(len(answers), 1, "Should be only 1 word")

	assert.True(answers[0] == "Buffalo" || answers[0] == "buffalo", "Should be buffalo or Buffalo")
}
