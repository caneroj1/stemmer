package stemmer

// testStemmer holds a slice of strings
// as input and has a corresponding slice
// of stem structs for what the test should
// should output.
type testStemmer struct {
	in   string
	want stem
}

// testFindCase tests the function to find a substring
// at the range given by start and end.
type testFindCase struct {
	in     string
	target string
	start  int
	end    int
	want   bool
}

// testSub holds a few string values. in is the input
// string, target is the substring we want to replace with
// replacement, and want is what the result should be.
type testSub struct {
	in          string
	target      string
	replacement string
	want        string
}

// testRule holds an in string and a want string, which
// is what the the result of the test should be.
type testRule struct {
	in   string
	want string
}

// step holds a function that will be used to map all of the inputs
// in rules.
type step struct {
	f     func(string) string
	rules []testRule
}

var findTests = []testFindCase{
	testFindCase{
		"DOG",
		"DO",
		0,
		2,
		true,
	},

	testFindCase{
		"DOG",
		"DO",
		0,
		3,
		false,
	},

	testFindCase{
		"D",
		"DO",
		0,
		1,
		false,
	},

	testFindCase{
		"D",
		"D",
		0,
		1,
		true,
	},

	testFindCase{
		"THAT",
		"T",
		-2,
		4,
		false,
	},
}

var stemCases = []testStemmer{

	testStemmer{
		"TR",
		makeStem(0, false, false, true, 'R'),
	},

	testStemmer{
		"EE",
		makeStem(0, false, true, false, 'E'),
	},

	testStemmer{
		"TREE",
		makeStem(0, false, true, false, 'E'),
	},

	testStemmer{
		"Y",
		makeStem(0, false, true, false, 'Y'),
	},

	testStemmer{
		"BY",
		makeStem(0, false, true, false, 'Y'),
	},

	testStemmer{
		"TROUBLE",
		makeStem(1, false, true, false, 'E'),
	},

	testStemmer{
		"TRoublE",
		makeStem(1, false, true, false, 'E'),
	},

	testStemmer{
		"OATS",
		makeStem(1, false, true, true, 'S'),
	},

	testStemmer{
		"TREES",
		makeStem(1, false, true, false, 'S'),
	},

	testStemmer{
		"IVY",
		makeStem(1, false, true, false, 'Y'),
	},

	testStemmer{
		"ORRERY",
		makeStem(2, false, true, false, 'Y'),
	},

	testStemmer{
		"OATEN",
		makeStem(2, true, true, false, 'N'),
	},

	testStemmer{
		"PRIVATE",
		makeStem(2, false, true, false, 'E'),
	},

	testStemmer{
		"TROUBLES",
		makeStem(2, true, true, false, 'S'),
	},

	testStemmer{
		"troubles",
		makeStem(2, true, true, false, 'S'),
	},

	testStemmer{
		"relational",
		makeStem(4, true, true, false, 'L'),
	},

	testStemmer{
		"rational",
		makeStem(3, true, true, false, 'L'),
	},
}

var subCases = []testSub{

	testSub{
		"caresses",
		"sses",
		"ss",
		"caress",
	},

	testSub{
		"ponies",
		"ies",
		"i",
		"poni",
	},

	testSub{
		"caress",
		"ss",
		"ss",
		"caress",
	},

	testSub{
		"cats",
		"s",
		"",
		"cat",
	},

	testSub{
		"relational",
		"ational",
		"ate",
		"relate",
	},

	testSub{
		"conditional",
		"tional",
		"tion",
		"condition",
	},
}

var testCases = map[string]step{

	"step1A": step{
		step1A,
		[]testRule{
			testRule{
				"CARESSES",
				"CARESS",
			},

			testRule{
				"PONIES",
				"PONI",
			},

			testRule{
				"TIES",
				"TI",
			},

			testRule{
				"CARESS",
				"CARESS",
			},

			testRule{
				"CATS",
				"CAT",
			},
		},
	},

	"step1B": step{
		step1B,
		[]testRule{
			testRule{
				"FEED",
				"FEED",
			},

			testRule{
				"AGREED",
				"AGREE",
			},

			testRule{
				"PLASTERED",
				"PLASTER",
			},

			testRule{
				"BLED",
				"BLED",
			},

			testRule{
				"MOTORING",
				"MOTOR",
			},

			testRule{
				"SING",
				"SING",
			},

			testRule{
				"CONFLATED",
				"CONFLATE",
			},

			testRule{
				"TROUBLED",
				"TROUBLE",
			},

			testRule{
				"SIZED",
				"SIZE",
			},

			testRule{
				"HOPPING",
				"HOP",
			},

			testRule{
				"TANNED",
				"TAN",
			},

			testRule{
				"TANNING",
				"TAN",
			},

			testRule{
				"FALLING",
				"FALL",
			},

			testRule{
				"FALLED",
				"FALL",
			},

			testRule{
				"HISSING",
				"HISS",
			},

			testRule{
				"FIZZING",
				"FIZZ",
			},

			testRule{
				"FIZZED",
				"FIZZ",
			},

			testRule{
				"FAILING",
				"FAIL",
			},

			testRule{
				"FILING",
				"FILE",
			},
		},
	},

	"step1C": step{
		step1C,
		[]testRule{
			testRule{
				"HAPPY",
				"HAPPI",
			},
			testRule{
				"SKY",
				"SKY",
			},
			testRule{
				"GUY",
				"GUI",
			},
		},
	},

	"step2": step{
		step2,
		[]testRule{
			testRule{
				"RELATIONAL",
				"RELATE",
			},

			testRule{
				"CONDITIONAL",
				"CONDITION",
			},

			testRule{
				"RATIONAL",
				"RATIONAL",
			},

			testRule{
				"VALENCI",
				"VALENCE",
			},

			testRule{
				"HESITANCI",
				"HESITANCE",
			},

			testRule{
				"DIGITIZER",
				"DIGITIZE",
			},

			testRule{
				"CONFORMABLI",
				"CONFORMABLE",
			},

			testRule{
				"RADICALLI",
				"RADICAL",
			},

			testRule{
				"DIFFERENTLI",
				"DIFFERENT",
			},

			testRule{
				"VILELI",
				"VILE",
			},

			testRule{
				"ANALOGOUSLI",
				"ANALOGOUS",
			},

			testRule{
				"VIETNAMIZATION",
				"VIETNAMIZE",
			},

			testRule{
				"PREDICATION",
				"PREDICATE",
			},

			testRule{
				"OPERATOR",
				"OPERATE",
			},

			testRule{
				"FEUDALISM",
				"FEUDAL",
			},

			testRule{
				"DECISIVENESS",
				"DECISIVE",
			},

			testRule{
				"HOPEFULNESS",
				"HOPEFUL",
			},

			testRule{
				"CALLOUSNESS",
				"CALLOUS",
			},

			testRule{
				"FORMALITI",
				"FORMAL",
			},

			testRule{
				"SENSITIVITI",
				"SENSITIVE",
			},

			testRule{
				"SENSIBILITI",
				"SENSIBLE",
			},
		},
	},

	"step3": step{
		step3,
		[]testRule{
			testRule{
				"TRIPLICATE",
				"TRIPLIC",
			},

			testRule{
				"FORMATIVE",
				"FORM",
			},

			testRule{
				"FORMALIZE",
				"FORMAL",
			},

			testRule{
				"ELECTRICITI",
				"ELECTRIC",
			},

			testRule{
				"ELECTRICAL",
				"ELECTRIC",
			},

			testRule{
				"HOPEFUL",
				"HOPE",
			},

			testRule{
				"GOODNESS",
				"GOOD",
			},
		},
	},

	"step4": step{
		step4,
		[]testRule{
			testRule{
				"REVIVAL",
				"REVIV",
			},

			testRule{
				"ALLOWANCE",
				"ALLOW",
			},

			testRule{
				"INFERENCE",
				"INFER",
			},

			testRule{
				"AIRLINER",
				"AIRLIN",
			},

			testRule{
				"GYROSCOPIC",
				"GYROSCOP",
			},

			testRule{
				"ADJUSTABLE",
				"ADJUST",
			},

			testRule{
				"DEFENSIBLE",
				"DEFENS",
			},

			testRule{
				"IRRITANT",
				"IRRIT",
			},

			testRule{
				"REPLACEMENT",
				"REPLAC",
			},

			testRule{
				"ADJUSTMENT",
				"ADJUST",
			},

			testRule{
				"DEPENDENT",
				"DEPEND",
			},

			testRule{
				"ADOPTION",
				"ADOPT",
			},

			testRule{
				"HOMOLOGOU",
				"HOMOLOG",
			},

			testRule{
				"COMMUNISM",
				"COMMUN",
			},

			testRule{
				"ACTIVATE",
				"ACTIV",
			},

			testRule{
				"HOMOLOGOUS",
				"HOMOLOG",
			},

			testRule{
				"ANGULARITI",
				"ANGULAR",
			},

			testRule{
				"EFFECTIVE",
				"EFFECT",
			},

			testRule{
				"BOWDLERIZE",
				"BOWDLER",
			},
		},
	},

	"step5A": step{
		step5A,
		[]testRule{
			testRule{
				"PROBATE",
				"PROBAT",
			},

			testRule{
				"RATE",
				"RATE",
			},

			testRule{
				"CEASE",
				"CEAS",
			},
		},
	},

	"step5B": step{
		step5B,
		[]testRule{
			testRule{
				"CONTROLL",
				"CONTROL",
			},

			testRule{
				"ROLL",
				"ROLL",
			},
		},
	},
}
