package numerals

import (
	"fmt"
	"testing"
	"testing/quick"
)

var tt []struct {
	arabic uint16
	roman  string
} = []struct {
	arabic uint16
	roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{19, "XIX"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{100, "C"},
	{90, "XC"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{2022, "MMXXII"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestArabicToRoman(t *testing.T) {

	for _, test := range tt {
		t.Run(fmt.Sprintf("%d gets converted to %q", test.arabic, test.roman), func(t *testing.T) {
			got := ConvertToRoman(test.arabic)

			if got != test.roman {
				t.Errorf("got %q, want %q", got, test.roman)
			}
		})
	}
}

func TestRomanToArabic(t *testing.T) {
	for _, test := range tt {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.roman, test.arabic), func(t *testing.T) {
			got := ConvertToArabic(test.roman)

			if got != test.arabic {
				t.Errorf("got %d, want %d", got, test.arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		// if arabic > 3999 {
		// 	t.Log("skipping", arabic)
		// 	return true
		// }
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
