package roman

import (
	"errors"
	"slices"
	"sort"
)

var romanInvCharacters = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func GetKeysSorted(m map[int]string) []int {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	slices.Reverse(keys)
	return keys
}

func IntToRoman(decimal int) (string, error) {
	if decimal < 1 || decimal > 3999 {
		return "", errors.New("number out of range")
	}
	var roman string = ""

	var keys = GetKeysSorted(romanInvCharacters)

	for decimal > 0 {
		for k := range keys {
			for decimal >= keys[k] {
				roman += romanInvCharacters[keys[k]]
				decimal -= keys[k]
			}
		}
		decimal = 0
	}
	return roman, nil
}
