package RomanNumerals

import "testing"

type test struct {
	roman  string
	arabic int
}

var underTenTests = []test{
	test{
		roman:  "I",
		arabic: 1,
	},
	test{
		roman:  "II",
		arabic: 2,
	},
	test{
		roman:  "III",
		arabic: 3,
	},
	test{
		roman:  "IV",
		arabic: 4,
	},
	test{
		roman:  "V",
		arabic: 5,
	},
	test{
		roman:  "VI",
		arabic: 6,
	},
	test{
		roman:  "VII",
		arabic: 7,
	},
	test{
		roman:  "VIII",
		arabic: 8,
	},
	test{
		roman:  "IX",
		arabic: 9,
	},
}

func TestUnderTen(t *testing.T) {
	for _, e := range underTenTests {
		if got := convertToRoman(e.arabic); got != e.roman {
			t.Errorf("convertToRoman(%d) = %q, expected %q", e.arabic, got, e.roman)
		}
	}
}
