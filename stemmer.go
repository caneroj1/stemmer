package stemmer

import (
	"runtime"
	"strings"
)

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

// find tries to match the substring given by target in int
// at the range specified by start and end.
func find(input, target string, start, end int) bool {
	length := len(input)
	if length < end || start < 0 {
		return false
	}

	return input[start:end] == target
}

// Stem is the entry function into the stemmer.
// We check to make sure the word isn't too short,
// and then we convert it to all uppercase
func Stem(input string) string {
	if len(input) < 3 {
		return input
	}

	upper := strings.ToUpper(input)
	stemmed := step1A(upper)
	stemmed = step1B(stemmed)
	stemmed = step1C(stemmed)
	stemmed = step2(stemmed)
	stemmed = step3(stemmed)
	stemmed = step4(stemmed)
	stemmed = step5A(stemmed)
	stemmed = step5B(stemmed)

	return stemmed
}

// StemMultiple accepts a slice of strings and stems each of them.
func StemMultiple(words []string) (output []string) {
	output = make([]string, len(words))
	for idx, word := range words {
		output[idx] = Stem(word)
	}

	return
}

// StemMultipleMutate accepts a pointer to a slice of strings and stems them in place.
// It modifies the original slice.
func StemMultipleMutate(words *[]string) {
	for idx, word := range *words {
		(*words)[idx] = Stem(word)
	}
}

// StemConcurrent accepts a pointer to a slice of strings and stems them in place.
// It tries to offload the work into multiple threads. It makes no guarantees about
// the order of the stems in the modified slice.
func StemConcurrent(words *[]string) {
	CPUs := runtime.NumCPU()
	length := len(*words)
	output := make(chan string)
	partition := length / CPUs

	var CPU int
	for CPU = 0; CPU < CPUs; CPU++ {
		go func(strs []string) {
			for _, word := range strs {
				output <- Stem(word)
			}
		}((*words)[CPU*partition : (CPU+1)*partition])
	}

	// if there are leftover words, stem them now
	if length-(CPU)*partition > 0 {
		go func(strs []string) {
			for _, word := range strs {
				output <- Stem(word)
			}
		}((*words)[(CPU)*partition : length])
	}

	for idx := 0; idx < length; idx++ {
		(*words)[idx] = <-output
	}
}
