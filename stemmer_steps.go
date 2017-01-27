package stemmer

//	TODO:
//	in the initial stem entry function filter out strings of
//	length 1 and 2
func step1A(input string) string {
	last := len(input) - 1
	if find(input, "S", last, last+1) {
		if find(input, "E", last-1, last) {
			if find(input, "I", last-2, last-1) {
				return replace(input, "IES", "I")
			} else if find(input, "S", last-2, last-1) && find(input, "S", last-3, last-2) {
				return replace(input, "SSES", "SS")
			}
		}

		if !find(input, "S", last-1, last) {
			return replace(input, "S", "")
		}

		return replace(input, "SS", "SS")
	}

	return input
}

func step1B(input string) string {
	last := len(input) - 1
	if find(input, "D", last, last+1) {
		if find(input, "E", last-1, last) {
			if find(input, "E", last-2, last-1) {
				if processStem(input[0:last-2]).measure > 0 {
					return replace(input, "EED", "EE")
				}
			} else {
				if (last-1) >= 0 && processStem(input[0:last-1]).vCondition {
					return step1BAfter(replace(input, "ED", ""))
				}
			}
		}
	} else if find(input, "G", last, last+1) {
		if find(input, "N", last-1, last) {
			if find(input, "I", last-2, last-1) {
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
	if find(input, "T", last, last+1) {
		if find(input, "A", last-1, last) {
			return replace(input, "AT", "ATE")
		}
	}

	if find(input, "L", last, last+1) {
		if find(input, "B", last-1, last) {
			return replace(input, "BL", "BLE")
		}
	}

	if find(input, "Z", last, last+1) {
		if find(input, "I", last-1, last) {
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
	if find(input, "Y", last, last+1) {
		stem := processStem(input[0:last])
		if stem.vCondition {
			return replace(input, "Y", "I")
		}
	}

	return input
}

func step2(input string) string {
	last := len(input) - 1

	if last-1 < 0 {
		return input
	}

	switch input[last-1] {
	case 'A':
		return step2ProcessA(input)
	case 'C':
		return step2ProcessC(input)
	case 'E':
		return step2ProcessE(input)
	case 'L':
		return step2ProcessL(input)
	case 'O':
		return step2ProcessO(input)
	case 'S':
		return step2ProcessS(input)
	case 'T':
		return step2ProcessT(input)
	}

	return input
}

func step2ProcessA(input string) string {
	last := len(input) - 1
	if find(input, "TIONAL", last-5, last+1) {
		if find(input, "A", last-6, last-5) {
			if processStem(input[0:last-6]).measure > 0 {
				return replace(input, "ATIONAL", "ATE")
			}
		} else {
			if processStem(input[0:last-5]).measure > 0 {
				return replace(input, "TIONAL", "TION")
			}
		}
	}

	return input
}

func step2ProcessC(input string) string {
	last := len(input) - 1
	if find(input, "NCI", last-2, last+1) {
		if find(input, "A", last-3, last-2) {
			if processStem(input[0:last-2]).measure > 0 {
				return replace(input, "ANCI", "ANCE")
			}
		} else if find(input, "E", last-3, last-2) {
			if processStem(input[0:last-2]).measure > 0 {
				return replace(input, "ENCI", "ENCE")
			}
		}
	}

	return input
}

func step2ProcessE(input string) string {
	last := len(input) - 1
	if find(input, "R", last, last+1) {
		if find(input, "Z", last-2, last-1) && find(input, "I", last-3, last-2) {
			return replace(input, "IZER", "IZE")
		}
	}

	return input
}

func step2ProcessL(input string) string {
	last := len(input) - 1
	if find(input, "I", last, last+1) {
		if find(input, "B", last-2, last-1) && find(input, "A", last-3, last-2) {
			if processStem(input[0:last-3]).measure > 0 {
				return replace(input, "ABLI", "ABLE")
			}
		} else if find(input, "L", last-2, last-1) && find(input, "A", last-3, last-2) {
			if processStem(input[0:last-3]).measure > 0 {
				return replace(input, "ALLI", "AL")
			}
		} else if find(input, "T", last-2, last-1) {
			if find(input, "EN", last-4, last-2) {
				if processStem(input[0:last-4]).measure > 0 {
					return replace(input, "ENTLI", "ENT")
				}
			}
		} else if find(input, "E", last-2, last-1) {
			if processStem(input[0:last-2]).measure > 0 {
				return replace(input, "ELI", "E")
			}
		} else if find(input, "S", last-2, last-1) {
			if find(input, "OUSLI", last-4, last+1) {
				if processStem(input[0:last-4]).measure > 0 {
					return replace(input, "OUSLI", "OUS")
				}
			}
		}
	}

	return input
}

func step2ProcessO(input string) string {
	last := len(input) - 1
	if find(input, "N", last, last+1) {
		if find(input, "ATI", last-4, last-1) {
			if find(input, "Z", last-5, last-4) {
				if find(input, "I", last-6, last-5) {
					if processStem(input[0:last-6]).measure > 0 {
						return replace(input, "IZATION", "IZE")
					}
				}
			} else {
				if processStem(input[0:last-4]).measure > 0 {
					return replace(input, "ATION", "ATE")
				}
			}
		}
	} else if find(input, "R", last, last+1) {
		if find(input, "ATO", last-3, last) {
			if processStem(input[0:last-3]).measure > 0 {
				return replace(input, "ATOR", "ATE")
			}
		}
	}

	return input
}

func step2ProcessS(input string) string {
	last := len(input) - 1
	if find(input, "M", last, last+1) {
		if find(input, "ALI", last-4, last-1) {
			if processStem(input[0:last-4]).measure > 0 {
				return replace(input, "ALISM", "AL")
			}
		}
	} else if find(input, "S", last, last+1) {
		if find(input, "NE", last-3, last-1) {
			if last-6 < 0 {
				return input
			}

			str := input[last-6 : last-3]
			if str == "IVE" {
				if processStem(input[0:last-6]).measure > 0 {
					return replace(input, "IVENESS", "IVE")
				}
			} else if str == "FUL" {
				if processStem(input[0:last-6]).measure > 0 {
					return replace(input, "FULNESS", "FUL")
				}
			} else if str == "OUS" {
				if processStem(input[0:last-6]).measure > 0 {
					return replace(input, "OUSNESS", "OUS")
				}
			}
		}
	}

	return input
}

func step2ProcessT(input string) string {
	last := len(input) - 1
	if find(input, "I", last, last+1) {
		if find(input, "I", last-2, last-1) {
			if find(input, "L", last-3, last-2) {
				if find(input, "A", last-4, last-3) {
					if processStem(input[0:last-4]).measure > 0 {
						return replace(input, "ALITI", "AL")
					}
				} else if find(input, "BI", last-5, last-3) {
					if processStem(input[0:last-5]).measure > 0 {
						return replace(input, "BILITI", "BLE")
					}
				}
			} else if find(input, "V", last-3, last-2) {
				if find(input, "I", last-4, last-3) {
					if processStem(input[0:last-4]).measure > 0 {
						return replace(input, "IVITI", "IVE")
					}
				}
			}
		}
	}

	return input
}

func step3(input string) string {
	last := len(input) - 1

	if find(input, "E", last, last+1) {
		if last-4 < 0 {
			return input
		}

		str := input[last-4 : last]
		if str == "ICAT" {
			if processStem(input[0:last-4]).measure > 0 {
				return replace(input, "ICATE", "IC")
			}
		} else if str == "ATIV" {
			if processStem(input[0:last-4]).measure > 0 {
				return replace(input, "ATIVE", "")
			}
		} else if str == "ALIZ" {
			if processStem(input[0:last-4]).measure > 0 {
				return replace(input, "ALIZE", "AL")
			}
		}
	} else if find(input, "I", last, last+1) {
		if find(input, "ICIT", last-4, last) {
			if processStem(input[0:last-4]).measure > 0 {
				return replace(input, "ICITI", "IC")
			}
		}
	} else if find(input, "L", last, last+1) {
		if find(input, "ICA", last-3, last) {
			if processStem(input[0:last-3]).measure > 0 {
				return replace(input, "ICAL", "IC")
			}
		} else if find(input, "FU", last-2, last) {
			if processStem(input[0:last-2]).measure > 0 {
				return replace(input, "FUL", "")
			}
		}
	} else if find(input, "S", last, last+1) {
		if find(input, "NES", last-3, last) {
			if processStem(input[0:last-3]).measure > 0 {
				return replace(input, "NESS", "")
			}
		}
	}

	return input
}

func step4(input string) string {
	last := len(input) - 1

	if last-1 < 0 {
		return input
	}

	switch input[last-1] {
	case 'A':
		return step4ProcessA(input)
	case 'C':
		return step4ProcessC(input)
	case 'E':
		return step4ProcessE(input)
	case 'I':
		return step4ProcessI(input)
	case 'L':
		return step4ProcessL(input)
	case 'N':
		return step4ProcessN(input)
	case 'O':
		return step4ProcessO(input)
	case 'S':
		return step4ProcessS(input)
	case 'T':
		return step4ProcessT(input)
	case 'U':
		return step4ProcessU(input)
	case 'V':
		return step4ProcessV(input)
	case 'Z':
		return step4ProcessZ(input)
	}

	return input
}

func step4ProcessA(input string) string {
	last := len(input) - 1
	if find(input, "L", last, last+1) {
		if processStem(input[0:last-1]).measure > 1 {
			return replace(input, "AL", "")
		}
	}

	return input
}

func step4ProcessC(input string) string {
	last := len(input) - 1
	if find(input, "E", last, last+1) {
		if last-3 < 0 {
			return input
		}

		str := input[last-3 : last-1]
		if str == "AN" {
			if processStem(input[0:last-3]).measure > 1 {
				return replace(input, "ANCE", "")
			}
		} else if str == "EN" {
			if processStem(input[0:last-3]).measure > 1 {
				return replace(input, "ENCE", "")
			}
		}
	}

	return input
}

func step4ProcessE(input string) string {
	last := len(input) - 1
	if find(input, "R", last, last+1) {
		if processStem(input[0:last-1]).measure > 1 {
			return replace(input, "ER", "")
		}
	}

	return input
}

func step4ProcessI(input string) string {
	last := len(input) - 1
	if find(input, "C", last, last+1) {
		if processStem(input[0:last-1]).measure > 1 {
			return replace(input, "IC", "")
		}
	}

	return input
}

func step4ProcessL(input string) string {
	last := len(input) - 1
	if find(input, "E", last, last+1) {
		if last-3 < 0 {
			return input
		}

		str := input[last-3 : last-1]
		if str == "AB" {
			if processStem(input[0:last-3]).measure > 1 {
				return replace(input, "ABLE", "")
			}
		} else if str == "IB" {
			if processStem(input[0:last-3]).measure > 1 {
				return replace(input, "IBLE", "")
			}
		}
	}

	return input
}

func step4ProcessN(input string) string {
	last := len(input) - 1
	if find(input, "T", last, last+1) {
		if find(input, "A", last-2, last-1) {
			if processStem(input[0:last-2]).measure > 1 {
				return replace(input, "ANT", "")
			}
		} else if find(input, "E", last-2, last-1) {
			if find(input, "M", last-3, last-2) {
				if find(input, "E", last-4, last-3) {
					if processStem(input[0:last-4]).measure > 1 {
						return replace(input, "EMENT", "")
					}
				} else {
					if processStem(input[0:last-3]).measure > 1 {
						return replace(input, "MENT", "")
					}
				}
			} else {
				if processStem(input[0:last-2]).measure > 1 {
					return replace(input, "ENT", "")
				}
			}
		}
	}

	return input
}

func step4ProcessO(input string) string {
	last := len(input) - 1
	if find(input, "N", last, last+1) {
		if find(input, "I", last-2, last-1) {
			stem := processStem(input[0 : last-2])
			if stem.measure > 1 && (stem.lastChar == 'S' || stem.lastChar == 'T') {
				return replace(input, "ION", "")
			}
		}
	} else if find(input, "U", last, last+1) {
		if processStem(input[0:last-1]).measure > 1 {
			return replace(input, "OU", "")
		}
	}

	return input
}

func step4ProcessS(input string) string {
	last := len(input) - 1
	if find(input, "M", last, last+1) {
		if find(input, "I", last-2, last-1) {
			if processStem(input[0:last-2]).measure > 1 {
				return replace(input, "ISM", "")
			}
		}
	}

	return input
}

func step4ProcessT(input string) string {
	last := len(input) - 1
	if find(input, "E", last, last+1) {
		if find(input, "A", last-2, last-1) {
			if processStem(input[0:last-2]).measure > 1 {
				return replace(input, "ATE", "")
			}
		}
	} else if find(input, "I", last, last+1) {
		if find(input, "I", last-2, last-1) {
			if processStem(input[0:last-2]).measure > 1 {
				return replace(input, "ITI", "")
			}
		}
	}

	return input
}

func step4ProcessU(input string) string {
	last := len(input) - 1
	if find(input, "S", last, last+1) {
		if find(input, "O", last-2, last-1) {
			if processStem(input[0:last-2]).measure > 1 {
				return replace(input, "OUS", "")
			}
		}
	}

	return input
}

func step4ProcessV(input string) string {
	last := len(input) - 1
	if find(input, "E", last, last+1) {
		if find(input, "I", last-2, last-1) {
			if processStem(input[0:last-2]).measure > 1 {
				return replace(input, "IVE", "")
			}
		}
	}

	return input
}

func step4ProcessZ(input string) string {
	last := len(input) - 1
	if find(input, "E", last, last+1) {
		if find(input, "I", last-2, last-1) {
			if processStem(input[0:last-2]).measure > 1 {
				return replace(input, "IZE", "")
			}
		}
	}

	return input
}

func step5A(input string) string {
	last := len(input) - 1
	if find(input, "E", last, last+1) {
		stem := processStem(input[0:last])
		if stem.measure > 1 || (stem.measure == 1 && !stem.oCondition) {
			return input[0:last]
		}
	}

	return input
}

func step5B(input string) string {
	stem := processStem(input)
	if stem.measure > 1 && stem.dCondition && stem.lastChar == 'L' {
		return input[0 : len(input)-1]
	}

	return input
}
