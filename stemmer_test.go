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

var stemCases = []TestStemmer{

	TestStemmer{
		"TR",
		stem{
			0,
		},
	},

	TestStemmer{
		"EE",
		stem{
			0,
		},
	},

	TestStemmer{
		"TREE",
		stem{
			0,
		},
	},

	TestStemmer{
		"Y",
		stem{
			0,
		},
	},

	TestStemmer{
		"BY",
		stem{
			0,
		},
	},

	TestStemmer{
		"TROUBLE",
		stem{
			1,
		},
	},

	TestStemmer{
		"TRoublE",
		stem{
			1,
		},
	},

	TestStemmer{
		"OATS",
		stem{
			1,
		},
	},

	TestStemmer{
		"TREES",
		stem{
			1,
		},
	},

	TestStemmer{
		"IVY",
		stem{
			1,
		},
	},

	TestStemmer{
		"ORRERY",
		stem{
			2,
		},
	},

	TestStemmer{
		"OATEN",
		stem{
			2,
		},
	},

	TestStemmer{
		"PRIVATE",
		stem{
			2,
		},
	},

	TestStemmer{
		"TROUBLES",
		stem{
			2,
		},
	},

	TestStemmer{
		"troubles",
		stem{
			2,
		},
	},

	TestStemmer{
		"relational",
		stem{
			4,
		},
	},

	TestStemmer{
		"rational",
		stem{
			3,
		},
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

var testCases = map[string][]TestRule{

	"step1A": []TestRule{
		TestRule{
			"caresses",
			"caress",
		},

		TestRule{
			"ponies",
			"poni",
		},

		TestRule{
			"ties",
			"ti",
		},

		TestRule{
			"caress",
			"caress",
		},

		TestRule{
			"cats",
			"cat",
		},
	},
}

// TestMeasure tests functionality to compute
// the 'measure' of a stem.
func TestMeasure(t *testing.T) {
	for _, test := range stemCases {
		got := measure(test.in)
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

// TestStep1A tests Step1A of the Porter-stemmer algorithm.
func TestStep1A(t *testing.T) {
	for _, test := range testCases["step1A"] {
		got := step1A(test.in)
		if got != test.want {
			t.Errorf("step1A(%v) = %v. Wanted = %v", test.in, got, test.want)
		}
	}
}
