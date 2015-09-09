package RomanNumerals

import "testing"

type test struct {
	roman  string
	arabic int
}

var tests = []test{
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
	test{
		roman:  "X",
		arabic: 10,
	},
	test{
		roman:  "XI",
		arabic: 11,
	},
	test{
		roman:  "XII",
		arabic: 12,
	},
	test{
		roman:  "XXII",
		arabic: 22,
	},
	test{
		roman:  "XXIV",
		arabic: 24,
	},
	test{
		roman:  "LV",
		arabic: 55,
	},
	test{
		roman:  "CDXLIV",
		arabic: 444,
	},
	test{
		roman: "MDLXXIX",
		arabic: 1579,
	},
	test{
		roman:  "MCMXC",
		arabic: 1990,
	},
	test{
		roman:  "MCMXCIX",
		arabic: 1999,
	},	
}

func TestConversion(t *testing.T) {
	for _, tt := range tests {
		if got := convertToRoman(tt.arabic); got != tt.roman {
			t.Errorf("convertToRoman(%d) = %q, expected %q", tt.arabic, got, tt.roman)
		}
	}
}
