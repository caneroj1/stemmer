package stemmer

import "testing"

// TestStemmer holds a slice of strings
// as input and has a corresponding slice
// of stem structs for what the test should
// should output.
type TestStemmer struct {
	in      string
	sTarget byte
	want    stem
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
		' ',
		makeStem(0),
	},

	TestStemmer{
		"EE",
		' ',
		makeStemWithConditions(0, false, true, false, false),
	},

	TestStemmer{
		"TREE",
		' ',
		makeStemWithConditions(0, false, true, false, false),
	},

	TestStemmer{
		"Y",
		' ',
		makeStemWithConditions(0, false, true, false, false),
	},

	TestStemmer{
		"BY",
		' ',
		makeStemWithConditions(0, false, true, false, false),
	},

	TestStemmer{
		"TROUBLE",
		' ',
		makeStemWithConditions(1, false, true, false, false),
	},

	TestStemmer{
		"TRoublE",
		' ',
		makeStemWithConditions(1, false, true, false, false),
	},

	TestStemmer{
		"OATS",
		' ',
		makeStemWithConditions(1, false, true, false, false),
	},

	TestStemmer{
		"TREES",
		' ',
		makeStemWithConditions(1, false, true, false, false),
	},

	TestStemmer{
		"IVY",
		' ',
		makeStemWithConditions(1, false, true, false, false),
	},

	TestStemmer{
		"ORRERY",
		' ',
		makeStemWithConditions(2, false, true, false, false),
	},

	TestStemmer{
		"OATEN",
		' ',
		makeStemWithConditions(2, false, true, false, false),
	},

	TestStemmer{
		"PRIVATE",
		' ',
		makeStemWithConditions(2, false, true, false, false),
	},

	TestStemmer{
		"TROUBLES",
		' ',
		makeStemWithConditions(2, false, true, false, false),
	},

	TestStemmer{
		"troubles",
		' ',
		makeStemWithConditions(2, false, true, false, false),
	},

	TestStemmer{
		"relational",
		' ',
		makeStemWithConditions(4, false, true, false, false),
	},

	TestStemmer{
		"rational",
		' ',
		makeStemWithConditions(3, false, true, false, false),
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

	"step2B": Step{
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
		},
	},
}

// TestMeasure tests functionality to compute
// the 'measure' of a stem.
func TestMeasure(t *testing.T) {
	for _, test := range stemCases {
		got := processStem(test.in, test.sTarget)
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
