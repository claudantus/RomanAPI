package roman

import (
	"errors"
)

const (
	// lower limit of accepted input
	LOWER_LIMIT int = 1
	// upper limit of accepted input
	UPPER_LIMIT int = 3999
)

// ordered keys for the roman numerals
var romanKeys = [13]int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

// map for the roman characters
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

func IntToRoman(decimal int) (string, error) {
	// return error if the input integer is out of bounds
	if decimal < LOWER_LIMIT || decimal > UPPER_LIMIT {
		return "", errors.New("number out of range")
	}
	
	// initialize roman output
	var roman string

	for decimal > 0 { // while remainder from decimal is greater than 0
		for k := range romanKeys { // go through the sorted keys
			for decimal >= romanKeys[k] { // while the remainder of the decimal is greater than the key
				roman += romanInvCharacters[romanKeys[k]] // append the corresponding roman character
				decimal -= romanKeys[k] // and subtract the key from the remainder (deicmal)
			}
		}
	}
	return roman, nil
}
