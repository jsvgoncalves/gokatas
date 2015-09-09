package RomanNumerals

var romanTable = map[int]string{
	1: "I",
	5: "V",
	10: "X",
	50: "L",
	100: "C",
	500: "D",
	1000: "M",
}

// convertToRomanScale takes an arabic numeral and it's given magnitude
// order and converts it to it's roman equivalent
func convertToRomanScale(an int, scale int) string {
	rn := ""
	// Get the represation of the digits for the current scale
	unit := romanTable[1 * scale]
	nextUnit := romanTable[10 * scale]
	halfUnit := romanTable[5 * scale]

	// First the border cases (4 and 9)
	if an == 9 {
		rn += unit + nextUnit
	} else if an == 4 {
		rn += unit + halfUnit
	// Then the simplest case (1,2,3)
	} else if an < 4 {
		for i := 1; i <= an; i++ {
			rn += unit
		}
	// Then 5,6,7,8
	} else if an >= 5 {
		rn += halfUnit
		for i := 6; i <= an; i++ {
			rn += unit
		}
	}
	return rn
}

// convertToRoman converts a given arabic numeral to it's equivalent roman numeral
func convertToRoman(an int) string {
	rn := ""
	// Iterate over the "scale", starting from the biggest
	// i.e, 1000, 100, 10, 1
	for scale := 1000; scale >= 1; scale /= 10 {
		// if the number has the current scale magnitude
		if an >= scale {
			// Then we get the digit for the current magnitude
			digit := an/scale
			// Convert it to a roman number with the given magnitude
			rn += convertToRomanScale(digit, scale)
			// And subtract the current magnitude unit
			an -= digit * scale
		}
	}
	return rn
}
