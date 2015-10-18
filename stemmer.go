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
	measure int
}

// A rule represents a mapping from an old suffix to a
// new suffix when a condition is met.
type rule struct {
	measure   int
	oldSuffix string
	newSuffix string
}

var letterToConsonantMap = map[int32]bool{
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

// measure computes how many times a string
// switches from a sequence of vowels to a sequence of
// consonants.
func measure(input string) (s stem) {
	previousWasVowel := false
	for _, char := range strings.ToUpper(input) {
		if shouldIncrementMeasure(char, previousWasVowel) {
			s.measure++
		}

		previousWasVowel = currentIsVowel(char, previousWasVowel)
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

func shouldIncrementMeasure(char int32, prev bool) bool {
	if char == 'Y' {
		return prev
	}

	return letterToConsonantMap[char] && prev
}

func currentIsVowel(char int32, prev bool) bool {
	if char == 'Y' {
		return !prev
	}
	return !letterToConsonantMap[char]
}
