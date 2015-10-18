package stemmer

import "strings"

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
	sCondition bool
}

func makeStem(m int) stem {
	return stem{
		m,
		false,
		false,
		false,
		false,
	}
}

func makeStemWithConditions(m int, o, v, d, s bool) stem {
	return stem{
		m,
		o,
		v,
		d,
		s,
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
func processStem(input string, sTarget byte) (s stem) {
	previousWasVowel := false
	uppedString := strings.ToUpper(input)
	length := len(uppedString)
	for i := 0; i < length; i++ {
		char := uppedString[i]
		if shouldIncrementMeasure(char, previousWasVowel) {
			s.measure++
		}

		previousWasVowel = currentIsVowel(char, previousWasVowel, &s)
	}

	if uppedString[length-1] == sTarget {
		s.sCondition = true
	}

	return
}

// replace replaces the last instance of target in input with replacement
func replace(input, target, replacement string) string {
	if idx := strings.LastIndex(input, target); idx >= 0 {
		return input[0:idx] + replacement
	}

	return input
}

//	TODO:
//	in the initial stem entry function filter out strings of
//	length 1 and 2
func step1A(input string) string {
	length := len(input)
	if input[length-1] == 'S' {
		if input[length-2] == 'E' {
			if input[length-3] == 'I' {
				return replace(input, "IES", "I")
			} else if input[length-3] == 'S' && input[length-4] == 'S' {
				return replace(input, "SSES", "SS")
			}
		}

		if input[length-2] != 'S' {
			return replace(input, "S", "")
		}
	}

	return input
}

func step1B(input string) string {
	return ""
}

//*********************
//	Utility Functions
//*********************

func shouldIncrementMeasure(char byte, prev bool) bool {
	if char == 'Y' {
		return prev
	}

	return letterToConsonantMap[char] && prev
}

func currentIsVowel(char byte, prev bool, s *stem) bool {
	var result bool
	if char == 'Y' {
		result = !prev
	} else {
		result = !letterToConsonantMap[char]
	}

	if !s.vCondition {
		s.vCondition = result
	}

	return result
}
