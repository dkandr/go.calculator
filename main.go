package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	romanNumerals := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9,
		"X": 10, "XI": 11, "XII": 12, "XIII": 13, "XIV": 14, "XV": 15, "XVI": 16, "XVII": 17, "XVIII": 18, "XIX": 19,
		"XX": 20, "XXI": 21, "XXIV": 24, "XXV": 25, "XXVII": 27, "XXVIII": 28,
		"XXX": 30, "XXXII": 32, "XXXV": 35, "XXXVI": 36,
		"XL": 40, "XLII": 42, "XLV": 45, "XLVIII": 48, "XLIX": 49,
		"L": 50, "LIV": 54, "LVI": 56,
		"LX": 60, "LXIII": 63, "LXIV": 64,
		"LXX": 70, "LXXII": 72,
		"LXXX": 80, "LXXXI": 81,
		"XC": 90, "C": 100}

	fmt.Print("> ")
	read := bufio.NewReader(os.Stdin)
	str, err := read.ReadString('\n')
	str = strings.TrimSpace(str)

	if err != nil {
		panic("error read stdin")
	}

	regexArabicNumerals, _ := regexp.Compile(`^[\d]+\s[\+\-\*\/]\s[\d]+$`)
	regexRomanNumerals, _ := regexp.Compile(`^[IVX]+\s[\+\-\*\/]\s[IVX]+$`)

	var isRomanNumerals bool

	switch {
	case regexArabicNumerals.MatchString(str):
		isRomanNumerals = false
	case regexRomanNumerals.MatchString(str):
		isRomanNumerals = true
	default:
		panic("bad string (a + b)")
	}

	var a, b int
	params := strings.Split(str, " ")

	if isRomanNumerals {
		var isAExist, isBExist bool
		a, isAExist = romanNumerals[params[0]]
		b, isBExist = romanNumerals[params[2]]
		if !isAExist || !isBExist {
			panic("unknown roman numeral")
		}
	} else {
		a, _ = strconv.Atoi(params[0])
		b, _ = strconv.Atoi(params[2])
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("bad numerals (1..10)")
	}

	var result int
	switch params[1] {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	}

	if isRomanNumerals {
	Loop:
		for key, value := range romanNumerals {
			if result == value {
				fmt.Println("<", key)
				break Loop
			}
		}
	} else {
		fmt.Println("<", result)
	}
}
