package stemmer

import "testing"

// TestStemmer holds a slice of strings
// as input and has a corresponding slice
// of stem structs for what the test should
// should output.
type TestStemmer struct {
	in   string
	want stem
}

// TestSubs holds a few string values. in is the input
// string, target is the substring we want to replace with
// replacement, and want is what the result should be.
type TestSub struct {
	in          string
	target      string
	replacement string
	want        string
}

// TestRule holds an in string and a want string, which
// is what the the result of the test should be.
type TestRule struct {
	in   string
	want string
}

// Step holds a function that will be used to map all of the inputs
// in rules.
type Step struct {
	f     func(string) string
	rules []TestRule
}

var stemCases = []TestStemmer{

	TestStemmer{
		"TR",
		makeStem(0, false, false, true, 'R'),
	},

	TestStemmer{
		"EE",
		makeStem(0, false, true, false, 'E'),
	},

	TestStemmer{
		"TREE",
		makeStem(0, false, true, false, 'E'),
	},

	TestStemmer{
		"Y",
		makeStem(0, false, true, false, 'Y'),
	},

	TestStemmer{
		"BY",
		makeStem(0, false, true, false, 'Y'),
	},

	TestStemmer{
		"TROUBLE",
		makeStem(1, false, true, false, 'E'),
	},

	TestStemmer{
		"TRoublE",
		makeStem(1, false, true, false, 'E'),
	},

	TestStemmer{
		"OATS",
		makeStem(1, false, true, true, 'S'),
	},

	TestStemmer{
		"TREES",
		makeStem(1, false, true, false, 'S'),
	},

	TestStemmer{
		"IVY",
		makeStem(1, false, true, false, 'Y'),
	},

	TestStemmer{
		"ORRERY",
		makeStem(2, false, true, false, 'Y'),
	},

	TestStemmer{
		"OATEN",
		makeStem(2, true, true, false, 'N'),
	},

	TestStemmer{
		"PRIVATE",
		makeStem(2, false, true, false, 'E'),
	},

	TestStemmer{
		"TROUBLES",
		makeStem(2, true, true, false, 'S'),
	},

	TestStemmer{
		"troubles",
		makeStem(2, true, true, false, 'S'),
	},

	TestStemmer{
		"relational",
		makeStem(4, true, true, false, 'L'),
	},

	TestStemmer{
		"rational",
		makeStem(3, true, true, false, 'L'),
	},
}

var subCases = []TestSub{

	TestSub{
		"caresses",
		"sses",
		"ss",
		"caress",
	},

	TestSub{
		"ponies",
		"ies",
		"i",
		"poni",
	},

	TestSub{
		"caress",
		"ss",
		"ss",
		"caress",
	},

	TestSub{
		"cats",
		"s",
		"",
		"cat",
	},

	TestSub{
		"relational",
		"ational",
		"ate",
		"relate",
	},

	TestSub{
		"conditional",
		"tional",
		"tion",
		"condition",
	},
}

var testCases = map[string]Step{

	"step1A": Step{
		step1A,
		[]TestRule{
			TestRule{
				"CARESSES",
				"CARESS",
			},

			TestRule{
				"PONIES",
				"PONI",
			},

			TestRule{
				"TIES",
				"TI",
			},

			TestRule{
				"CARESS",
				"CARESS",
			},

			TestRule{
				"CATS",
				"CAT",
			},
		},
	},

	"step1B": Step{
		step1B,
		[]TestRule{
			TestRule{
				"FEED",
				"FEED",
			},

			TestRule{
				"AGREED",
				"AGREE",
			},

			TestRule{
				"PLASTERED",
				"PLASTER",
			},

			TestRule{
				"BLED",
				"BLED",
			},

			TestRule{
				"MOTORING",
				"MOTOR",
			},

			TestRule{
				"SING",
				"SING",
			},

			TestRule{
				"CONFLATED",
				"CONFLATE",
			},

			TestRule{
				"TROUBLED",
				"TROUBLE",
			},

			TestRule{
				"SIZED",
				"SIZE",
			},

			TestRule{
				"HOPPING",
				"HOP",
			},

			TestRule{
				"TANNED",
				"TAN",
			},

			TestRule{
				"TANNING",
				"TAN",
			},

			TestRule{
				"FALLING",
				"FALL",
			},

			TestRule{
				"FALLED",
				"FALL",
			},

			TestRule{
				"HISSING",
				"HISS",
			},

			TestRule{
				"FIZZING",
				"FIZZ",
			},

			TestRule{
				"FIZZED",
				"FIZZ",
			},

			TestRule{
				"FAILING",
				"FAIL",
			},

			TestRule{
				"FILING",
				"FILE",
			},
		},
	},

	"step1C": Step{
		step1C,
		[]TestRule{
			TestRule{
				"HAPPY",
				"HAPPI",
			},
			TestRule{
				"SKY",
				"SKY",
			},
			TestRule{
				"GUY",
				"GUI",
			},
		},
	},

	"step2": Step{
		step2,
		[]TestRule{
			TestRule{
				"RELATIONAL",
				"RELATE",
			},

			TestRule{
				"CONDITIONAL",
				"CONDITION",
			},

			TestRule{
				"RATIONAL",
				"RATIONAL",
			},

			TestRule{
				"VALENCI",
				"VALENCE",
			},

			TestRule{
				"HESITANCI",
				"HESITANCE",
			},

			TestRule{
				"DIGITIZER",
				"DIGITIZE",
			},

			TestRule{
				"CONFORMABLI",
				"CONFORMABLE",
			},

			TestRule{
				"RADICALLI",
				"RADICAL",
			},

			TestRule{
				"DIFFERENTLI",
				"DIFFERENT",
			},

			TestRule{
				"VILELI",
				"VILE",
			},

			TestRule{
				"ANALOGOUSLI",
				"ANALOGOUS",
			},

			TestRule{
				"VIETNAMIZATION",
				"VIETNAMIZE",
			},

			TestRule{
				"PREDICATION",
				"PREDICATE",
			},

			TestRule{
				"OPERATOR",
				"OPERATE",
			},

			TestRule{
				"FEUDALISM",
				"FEUDAL",
			},

			TestRule{
				"DECISIVENESS",
				"DECISIVE",
			},

			TestRule{
				"HOPEFULNESS",
				"HOPEFUL",
			},

			TestRule{
				"CALLOUSNESS",
				"CALLOUS",
			},

			TestRule{
				"FORMALITI",
				"FORMAL",
			},

			TestRule{
				"SENSITIVITI",
				"SENSITIVE",
			},

			TestRule{
				"SENSIBILITI",
				"SENSIBLE",
			},
		},
	},

	"step5A": Step{
		step5A,
		[]TestRule{
			TestRule{
				"PROBATE",
				"PROBAT",
			},

			TestRule{
				"RATE",
				"RATE",
			},

			TestRule{
				"CEASE",
				"CEAS",
			},
		},
	},

	"step5B": Step{
		step5B,
		[]TestRule{
			TestRule{
				"CONTROLL",
				"CONTROL",
			},

			TestRule{
				"ROLL",
				"ROLL",
			},
		},
	},
}

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
