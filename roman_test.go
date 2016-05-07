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

func TestSix(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("VI"), 6, "The roman VI is the decimal 6")
}

func TestEight(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("VIII"), 8, "The roman VIII is the decimal 8")
}

func TestTen(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("X"), 10, "The roman X is the decimal 10")
}

func TestNine(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("IX"), 9, "The roman IX is the decimal 9")
}

func TestTwoelve(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("XII"), 12, "The roman XII is the decimal 12")
}

func TestFourteen(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("XIV"), 14, "The roman XIV is the decimal 14")
}

func TestFifteen(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("XV"), 15, "The roman XV is the decimal 15")
}

func TestSixteen(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("XVI"), 16, "The roman XVI is the decimal 16")
}

func TestNineteen(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("XIX"), 19, "The roman XIX is the decimal 19")
}

func TestTwenty(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("XX"), 20, "The roman XX is the decimal 20")
}

func TestTwentySeven(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("XXVII"), 27, "The roman XXVII is the decimal 27")
}

func TestThirtyThree(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("XXXIII"), 33, "The roman XXXIII is the decimal 33")
}

func TestThirtyNine(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("XXXIX"), 39, "The roman XXXIX is the decimal 39")
}

func TestFourtyNine(t *testing.T) {
	assert.Equal(t, roman.RomanToDecimal("XLIX"), 49, "The roman XLIX is the decimal 49")
}