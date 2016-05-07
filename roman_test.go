package roman

import (
	"testing"
	"github.com/stretchr/testify/assert"
	
	"github.com/julianlopezv/tdd-roman-numbers"
)


func TestOne(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("I"), 1, "The roman I is the decimal 1")
}

func TestTwo(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("II"), 2, "The roman II is the decimal 2")
}

func TestThree(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("III"), 3, "The roman III is the decimal 3")
}

func TestFour(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("IV"), 4, "The roman IV is the decimal 4")
}

func TestFive(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("V"), 5, "The roman V is the decimal 5")
}