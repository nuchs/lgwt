package properties

import (
	"fmt"
	"testing"
	"testing/quick"
)

type TestCase struct {
	Arabic uint16
	Roman  string
}

var testCases = []TestCase{
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
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{47, "XLVII"},
	{49, "XLIX"},
	{90, "XC"},
	{100, "C"},
	{400, "CD"},
	{500, "D"},
	{798, "DCCXCVIII"},
	{900, "CM"},
	{1000, "M"},
	{1006, "MVI"},
	{1984, "MCMLXXXIV"},
	{2014, "MMXIV"},
	{3999, "MMMCMXCIX"},
}

func TestConvertToRoman(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%d converted to %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)

			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%q converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)

			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		t.Log("Testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
