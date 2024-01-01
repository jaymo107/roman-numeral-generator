package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	HUNDREDS = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	TENS     = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	ONES     = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
)

func GetNumeral(amount int) string {
	thousands := strings.Repeat("M", amount/1000)
	hundreds := HUNDREDS[(amount%1000)/100]
	tens := TENS[(amount%100)/10]
	units := ONES[amount%10]

	return thousands + hundreds + tens + units
}

func CheckValue(amount int) {
	if amount > 9999 {
		panic(errors.New("value too high"))
	}
}

func main() {
	if amount, err := strconv.Atoi(os.Args[1]); err == nil {
		CheckValue(amount)
		numeral := GetNumeral(amount)
		fmt.Println(numeral)
	} else {
		panic(err)
	}
}