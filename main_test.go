package main

import (
	"fmt"
	"testing"
)

var cases = []RomanNumeral{
	{3, "III"},
	{14, "XIV"},
	{111, "CXI"},
	{47, "XLVII"},
	{1984, "MCMLXXXIV"},
	{69, "LXIX"},
}

func TestRomanize(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("romanize %d", test.Arabic), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Roman {
				t.Errorf("want %q, but got %q", test.Roman, got)
			}
		})
	}
}

func TestArabize(t *testing.T) {
	for _, test := range cases {
		t.Run(fmt.Sprintf("arabize %s", test.Roman), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			if got != test.Arabic {
				t.Errorf("want %d, but got %d", test.Arabic, got)
			}
		})
	}
}
