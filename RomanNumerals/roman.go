package RomanNumerals

var romanTable = map[int]string{
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
}

func convertToRoman(arabic_no int) string {
	return romanTable[arabic_no]
}
