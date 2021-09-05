package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPalindromeChecker(t *testing.T) {
	trueInputs := []string{"ada", "kodok", "katak"}
	falseInputs := []string{"tidak", "palindrome"}

	for _, trueInput := range trueInputs {
		result := PalindromeChecker(trueInput)
		assert.Equal(t, result, true, "the result is shoud be true")
	}

	for _, falseInput := range falseInputs {
		result := PalindromeChecker(falseInput)
		assert.Equal(t, result, false, "the result is shoud be false")
	}
}
