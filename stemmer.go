package stemmer

import "strings"

// import "fmt"

// A stem represents the part of a word
// or whole word before a matching
// condition.
// For example:
// condition = (SSES -> SS)
// rule      = caresses  ->  caress
// The stem is "care"
type stem struct {
	measure    int
	oCondition bool
	vCondition bool
	dCondition bool
	lastChar   byte
}

func makeStem(m int, o, v, d bool, b byte) stem {
	return stem{
		m,
		o,
		v,
		d,
		b,
	}
}

var letterToConsonantMap = map[byte]bool{
	65: false,
	66: true,
	67: true,
	68: true,
	69: false,
	70: true,
	71: true,
	72: true,
	73: false,
	74: true,
	75: true,
	76: true,
	77: true,
	78: true,
	79: false,
	80: true,
	81: true,
	82: true,
	83: true,
	84: true,
	85: false,
	86: true,
	87: true,
	88: true,
	89: true,
	90: true,
}

// stem computes how many times a string
// switches from a sequence of vowels to a sequence of
// consonants as well as verifying the different stem conditions.
func processStem(input string) (s stem) {
	previousWasVowel := false
	uppedString := strings.ToUpper(input)
	length := len(uppedString)
	doubleConsonants := 0

	for i := 0; i < length; i++ {
		char := uppedString[i]
		if shouldIncrementMeasure(char, previousWasVowel) {
			s.measure++
		}

		previousWasVowel = currentIsVowel(char, previousWasVowel, &s)
		if i == (length - 3) {

			if !previousWasVowel {
				s.oCondition = true
			}
		}

		if i == (length - 2) {
			if !previousWasVowel {
				if s.oCondition {
					s.oCondition = false
				}

				doubleConsonants++
			}
		}

		if i == (length - 1) {
			s.lastChar = char
			if !previousWasVowel {
				if s.oCondition && charCheck(char) {
					s.oCondition = false
				}

				doubleConsonants++
			} else {
				s.oCondition = false
			}
		}
	}

	s.dCondition = (doubleConsonants == 2)

	return
}

// replace replaces the last instance of target in input with replacement
func replace(input, target, replacement string) string {
	if idx := strings.LastIndex(input, target); idx >= 0 {
		return input[0:idx] + replacement
	}

	return input
}
