package stemmer

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
	if input[last-6:last+1] == "TIONAL" {
		if input[last-7] == 'A' {
			if processStem(input[0:last-7]).measure > 0 {
				return replace(input, "ATIONAL", "ATE")
			}
		} else {
			if processStem(input[0:last-6]).measure > 0 {
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
			if processStem(input[0:last-3]).measure > 0 {
				return replace(input, "ANCI", "ANCE")
			}
		} else if input[last-3] == 'E' {
			if processStem(input[0:last-3]).measure > 0 {
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
					replace(input, "ENTLI", "ENT")
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
		if input[last-3:last-1] == "ATI" {
			if input[last-4] == 'Z' {
				if input[last-5] == 'I' {
					if processStem(input[0:last-5]).measure > 0 {
						replace(input, "IZATION", "IZE")
					}
				}
			} else {
				if processStem(input[0:last-4]).measure > 0 {
					replace(input, "ATION", "ATE")
				}
			}
		}
	} else if input[last] == 'R' {
		if input[last-3:last] == "ATO" {
			if processStem(input[0:last-3]).measure > 0 {
				replace(input, "ATOR", "ATE")
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
