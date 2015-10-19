package stemmer

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
