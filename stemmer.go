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

//	TODO:
//	in the initial stem entry function filter out strings of
//	length 1 and 2
func step1A(input string) string {
	last := len(input) - 1
	if input[last] == 'S' {
		if input[last-1] == 'E' {
			if input[last-2] == 'I' {
				return replace(input, "IES", "I")
			} else if input[last-2] == 'S' && input[last-3] == 'S' {
				return replace(input, "SSES", "SS")
			}
		}

		if input[last-1] != 'S' {
			return replace(input, "S", "")
		}
	}

	return input
}

func step1B(input string) string {
	last := len(input) - 1
	if input[last] == 'D' {
		if input[last-1] == 'E' {
			if input[last-2] == 'E' {
				if processStem(input[0:last-2]).measure > 0 {
					return replace(input, "EED", "EE")
				}
			} else {
				if processStem(input[0 : last-1]).vCondition {
					return step1BAfter(replace(input, "ED", ""))
				}
			}
		}
	} else if input[last] == 'G' {
		if input[last-1] == 'N' {
			if input[last-2] == 'I' {
				if processStem(input[0 : last-2]).vCondition {
					return step1BAfter(replace(input, "ING", ""))
				}
			}
		}
	}

	return input
}

func step1BAfter(input string) string {
	last := len(input) - 1
	if input[last] == 'T' {
		if input[last-1] == 'A' {
			return replace(input, "AT", "ATE")
		}
	}

	if input[last] == 'L' {
		if input[last-1] == 'B' {
			return replace(input, "BL", "BLE")
		}
	}

	if input[last] == 'Z' {
		if input[last-1] == 'I' {
			return replace(input, "IZ", "IZE")
		}
	}

	stem := processStem(input)
	if stem.dCondition {
		if stem.lastChar != 'L' && stem.lastChar != 'S' && stem.lastChar != 'Z' {
			return input[0:last]
		}
	} else if stem.oCondition && stem.measure == 1 {
		return input + "E"
	}

	return input
}

func step1C(input string) string {
	last := len(input) - 1
	if input[last] == 'Y' {
		stem := processStem(input[0:last])
		if stem.vCondition {
			return replace(input, "Y", "I")
		}
	}

	return input
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

func doubleConsonantCheck(prev bool, s *stem) {
	if !prev && !s.dCondition {
		s.dCondition = true
	}
}

func charCheck(char byte) bool {
	return char == 'W' || char == 'X' || char == 'Y'
}
