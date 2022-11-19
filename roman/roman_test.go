package roman

import (
	"fmt"
	"log"
	"testing"
	"testing/quick"
)

func TestConvertToRoman(t *testing.T) {
	cases := []struct {
		Arabic uint16
		Roman  string
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
		{39, "XXXIX"},
		{40, "XL"},
		{1984, "MCMLXXXIV"},
	}
	for _, test := range cases {
		t.Run(fmt.Sprintf("%d -> %q", test.Arabic, test.Roman), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
		t.Run(fmt.Sprintf("%q -> %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}

func TestPropertiesOfConversion(t *testing.T) {
  assertion := func(arabic uint16) bool {
    log.Println(arabic)
    roman := ConvertToRoman(arabic)
    fromRoman := ConvertToArabic(roman)
    return fromRoman == arabic
  }
  if err := quick.Check(assertion, &quick.Config{
    MaxCount: 1000,
  }); err != nil {
    t.Error("failed quick checks", err)
  }
}
