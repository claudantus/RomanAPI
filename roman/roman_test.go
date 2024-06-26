package roman

import "testing"

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		input int
		want  string
		err   bool
	}{
		{1, "I", false},        // Lower limit
		{3999, "MMMCMXCIX", false}, // Upper limit
		{0, "", true},          // Outside lower limit
		{4000, "", true},       // Outside upper limit
		{4, "IV", false},       // Edge case
		{9, "IX", false},       // Edge case
		{40, "XL", false},      // Edge case
		{90, "XC", false},      // Edge case
		{400, "CD", false},     // Edge case
		{900, "CM", false},     // Edge case
		{123, "CXXIII", false}, // Normal case
		{-1, "", true},         // Negative input
	}

	// perform all tests as defined above
	for _, tt := range tests {
		got, err := DecimalToRoman(tt.input)
		// case where an error was expected but not returned
		if tt.err && err == nil {
			t.Errorf("DecimalToRoman(%d): expected error, got nil", tt.input)
		}
		// case where no error was expected but an error was returned
		if !tt.err && err != nil {
			t.Errorf("DecimalToRoman(%d): unexpected error: %v", tt.input, err)
		}
		// case where the returned value does not match the expected value
		if got != tt.want {
			t.Errorf("DecimalToRoman(%d): got %q, want %q", tt.input, got, tt.want)
		}
	}
}



