package stemmer

// import "fmt"

//	TODO:
//	in the initial stem entry function filter out strings of
//	length 1 and 2
func step1A(input string) string {
	last := len(input) - 1
	if input[last] == 'S' {
		if input[last-1] == 'E' {
			if input[last-2] == 'I' {
				return replace(input, "IES", "I")
			} else if input[last-2] == 'S' && input[last-3] == 'S' {
				return replace(input, "SSES", "SS")
			}
		}

		if input[last-1] != 'S' {
			return replace(input, "S", "")
		}
	}

	return input
}

func step1B(input string) string {
	last := len(input) - 1
	if input[last] == 'D' {
		if input[last-1] == 'E' {
			if input[last-2] == 'E' {
				if processStem(input[0:last-2]).measure > 0 {
					return replace(input, "EED", "EE")
				}
			} else {
				if processStem(input[0 : last-1]).vCondition {
					return step1BAfter(replace(input, "ED", ""))
				}
			}
		}
	} else if input[last] == 'G' {
		if input[last-1] == 'N' {
			if input[last-2] == 'I' {
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
	if input[last] == 'T' {
		if input[last-1] == 'A' {
			return replace(input, "AT", "ATE")
		}
	}

	if input[last] == 'L' {
		if input[last-1] == 'B' {
			return replace(input, "BL", "BLE")
		}
	}

	if input[last] == 'Z' {
		if input[last-1] == 'I' {
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
	if input[last] == 'Y' {
		stem := processStem(input[0:last])
		if stem.vCondition {
			return replace(input, "Y", "I")
		}
	}

	return input
}

func step2(input string) string {
	last := len(input) - 1

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
	if input[last-5:last+1] == "TIONAL" {
		if input[last-6] == 'A' {
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
	if input[last-2:last+1] == "NCI" {
		if input[last-3] == 'A' {
			if processStem(input[0:last-2]).measure > 0 {
				return replace(input, "ANCI", "ANCE")
			}
		} else if input[last-3] == 'E' {
			if processStem(input[0:last-2]).measure > 0 {
				return replace(input, "ENCI", "ENCE")
			}
		}
	}

	return input
}

func step2ProcessE(input string) string {
	last := len(input) - 1
	if input[last] == 'R' {
		if input[last-2] == 'Z' && input[last-3] == 'I' {
			return replace(input, "IZER", "IZE")
		}
	}

	return input
}

func step2ProcessL(input string) string {
	last := len(input) - 1
	if input[last] == 'I' {
		if input[last-2] == 'B' && input[last-3] == 'A' {
			if processStem(input[0:last-3]).measure > 0 {
				return replace(input, "ABLI", "ABLE")
			}
		} else if input[last-2] == 'L' && input[last-3] == 'A' {
			if processStem(input[0:last-3]).measure > 0 {
				return replace(input, "ALLI", "AL")
			}
		} else if input[last-2] == 'T' {
			if input[last-4:last-2] == "EN" {
				if processStem(input[0:last-4]).measure > 0 {
					return replace(input, "ENTLI", "ENT")
				}
			}
		} else if input[last-2] == 'E' {
			if processStem(input[0:last-2]).measure > 0 {
				return replace(input, "ELI", "E")
			}
		} else if input[last-2] == 'S' {
			if input[last-4:last+1] == "OUSLI" {
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
	if input[last] == 'N' {
		if input[last-4:last-1] == "ATI" {
			if input[last-5] == 'Z' {
				if input[last-6] == 'I' {
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
	} else if input[last] == 'R' {
		if input[last-3:last] == "ATO" {
			if processStem(input[0:last-3]).measure > 0 {
				return replace(input, "ATOR", "ATE")
			}
		}
	}

	return input
}

func step2ProcessS(input string) string {
	last := len(input) - 1
	if input[last] == 'M' {
		if input[last-4:last-1] == "ALI" {
			if processStem(input[0:last-4]).measure > 0 {
				return replace(input, "ALISM", "AL")
			}
		}
	} else if input[last] == 'S' {
		if input[last-3:last-1] == "NE" {
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
	if input[last] == 'I' {
		if input[last-2] == 'I' {
			if input[last-3] == 'L' {
				if input[last-4] == 'A' {
					if processStem(input[0:last-4]).measure > 0 {
						return replace(input, "ALITI", "AL")
					}
				} else if input[last-5:last-3] == "BI" {
					if processStem(input[0:last-5]).measure > 0 {
						return replace(input, "BILITI", "BLE")
					}
				}
			} else if input[last-3] == 'V' {
				if input[last-4] == 'I' {
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

	if input[last] == 'E' {
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
	} else if input[last] == 'I' {
		if input[last-4:last] == "ICIT" {
			if processStem(input[0:last-4]).measure > 0 {
				return replace(input, "ICITI", "IC")
			}
		}
	} else if input[last] == 'L' {
		if input[last-3:last] == "ICA" {
			if processStem(input[0:last-3]).measure > 0 {
				return replace(input, "ICAL", "IC")
			}
		} else if input[last-2:last] == "FU" {
			if processStem(input[0:last-2]).measure > 0 {
				return replace(input, "FUL", "")
			}
		}
	} else if input[last] == 'S' {
		if input[last-3:last] == "NES" {
			if processStem(input[0:last-3]).measure > 0 {
				return replace(input, "NESS", "")
			}
		}
	}

	return input
}

func step5A(input string) string {
	last := len(input) - 1
	if input[last] == 'E' {
		stem := processStem(input[0:last])
		if stem.measure > 1 || (stem.measure == 1 && !stem.oCondition) {
			return input[0:last]
		}
	}

	return input
}

func step5B(input string) string {
	last := len(input) - 1
	stem := processStem(input)
	if stem.measure > 1 && stem.dCondition && stem.lastChar == 'L' {
		return input[0:last]
	}

	return input
}
