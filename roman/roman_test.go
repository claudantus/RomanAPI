package roman

import "testing"

func TestConvertToRomanI(t *testing.T) {
	roman, error := IntToRoman(1)
	if roman != "I" && error == nil {
		t.Errorf("I should return I, but %v", roman)
	}
}

func TestConvertToRomanZero(t *testing.T) {
	roman, error := IntToRoman(0)
	if roman != "" && error != nil {
		t.Errorf("Should return error out of range")
	}
}

func TestConvertToRomanMax(t *testing.T) {
	roman, error := IntToRoman(3999)
	if roman != "MMMCMXCIX" && error == nil {
		t.Errorf("I should return I, but %v", roman)
	}
}

func TestConvertToRoman4k(t *testing.T) {
	roman, error := IntToRoman(4000)
	if roman != "" && error != nil {
		t.Errorf("Should return error out of range")
	}
}
