package stemmer

import "testing"

// TestMeasure tests functionality to compute
// the 'measure' of a stem.
func TestMeasure(t *testing.T) {
	for _, test := range stemCases {
		got := processStem(test.in)
		if got != test.want {
			t.Errorf("measure(%v) = %v. Wanted %v", test.in, got, test.want)
		}
	}
}

// TestReplace tests functionality to
// replace the last matching substring of a word
// with a replacement string.
func TestReplace(t *testing.T) {
	for _, test := range subCases {
		got := replace(test.in, test.target, test.replacement)
		if got != test.want {
			t.Errorf("replace(%v, %v, %v) = %v. Wanted = %v", test.in, test.target, test.replacement, got, test.want)
		}
	}
}

// TestRules tests each of the steps in the Porter-Stemmer algorithm.
func TestRules(t *testing.T) {
	for stepName, step := range testCases {
		for _, test := range step.rules {
			got := step.f(test.in)
			if got != test.want {
				t.Errorf("%s(%v) = %v. Wanted = %v", stepName, test.in, got, test.want)
			}
		}
	}
}

// TestFind tests the function that checks a string
// for the presence of a substring.
func TestFind(t *testing.T) {
	for _, test := range findTests {
		got := find(test.in, test.target, test.start, test.end)
		if got != test.want {
			t.Errorf("find(%v, %d, %d) = %v. Wanted = %v", test.in, test.start, test.end, got, test.want)
		}
	}
}
