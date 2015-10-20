package stemmer

// TestStemmer holds a slice of strings
// as input and has a corresponding slice
// of stem structs for what the test should
// should output.
type TestStemmer struct {
	in   string
	want stem
}

// TestFindCase tests the function to find a substring
// at the range given by start and end.
type TestFindCase struct {
	in     string
	target string
	start  int
	end    int
	want   bool
}

// TestSub holds a few string values. in is the input
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

var findTests = []TestFindCase{
	TestFindCase{
		"DOG",
		"DO",
		0,
		2,
		true,
	},

	TestFindCase{
		"DOG",
		"DO",
		0,
		3,
		false,
	},

	TestFindCase{
		"D",
		"DO",
		0,
		1,
		false,
	},

	TestFindCase{
		"D",
		"D",
		0,
		1,
		true,
	},
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

	"step3": Step{
		step3,
		[]TestRule{
			TestRule{
				"TRIPLICATE",
				"TRIPLIC",
			},

			TestRule{
				"FORMATIVE",
				"FORM",
			},

			TestRule{
				"FORMALIZE",
				"FORMAL",
			},

			TestRule{
				"ELECTRICITI",
				"ELECTRIC",
			},

			TestRule{
				"ELECTRICAL",
				"ELECTRIC",
			},

			TestRule{
				"HOPEFUL",
				"HOPE",
			},

			TestRule{
				"GOODNESS",
				"GOOD",
			},
		},
	},

	"step4": Step{
		step4,
		[]TestRule{
			TestRule{
				"REVIVAL",
				"REVIV",
			},

			TestRule{
				"ALLOWANCE",
				"ALLOW",
			},

			TestRule{
				"INFERENCE",
				"INFER",
			},

			TestRule{
				"AIRLINER",
				"AIRLIN",
			},

			TestRule{
				"GYROSCOPIC",
				"GYROSCOP",
			},

			TestRule{
				"ADJUSTABLE",
				"ADJUST",
			},

			TestRule{
				"DEFENSIBLE",
				"DEFENS",
			},

			TestRule{
				"IRRITANT",
				"IRRIT",
			},

			TestRule{
				"REPLACEMENT",
				"REPLAC",
			},

			TestRule{
				"ADJUSTMENT",
				"ADJUST",
			},

			TestRule{
				"DEPENDENT",
				"DEPEND",
			},

			TestRule{
				"ADOPTION",
				"ADOPT",
			},

			TestRule{
				"HOMOLOGOU",
				"HOMOLOG",
			},

			TestRule{
				"COMMUNISM",
				"COMMUN",
			},

			TestRule{
				"ACTIVATE",
				"ACTIV",
			},

			TestRule{
				"HOMOLOGOUS",
				"HOMOLOG",
			},

			TestRule{
				"ANGULARITI",
				"ANGULAR",
			},

			TestRule{
				"EFFECTIVE",
				"EFFECT",
			},

			TestRule{
				"BOWDLERIZE",
				"BOWDLER",
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
