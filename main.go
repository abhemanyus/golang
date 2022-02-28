package main

import "strings"

type RomanNumeral struct {
	Arabic int
	Roman  string
}

var allNumerals = []RomanNumeral{
	{1000, "M"},
	{900, "CM"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder
	for _, numeral := range allNumerals {
		for arabic >= numeral.Arabic {
			result.WriteString(numeral.Roman)
			arabic -= numeral.Arabic
		}
	}
	return result.String()
}

var numToArab = map[string]int{
	"M": 1000,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}

func ConvertToArabic(roman string) int {
	prev := 0
	arab := 0
	for i := len(roman) - 1; i >= 0; i-- {
		ar := numToArab[string(roman[i])]
		if ar < prev {
			arab -= ar
		} else {
			arab += ar
		}
		prev = ar
	}
	return arab
}
