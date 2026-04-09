package belajar

import (
	"reflect"
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - RomanToInt
// =============================================================

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"I", "I", 1},
		{"III", "III", 3},
		{"IV", "IV", 4},
		{"V", "V", 5},
		{"IX", "IX", 9},
		{"X", "X", 10},
		{"XL", "XL", 40},
		{"XLII", "XLII", 42},
		{"L", "L", 50},
		{"XC", "XC", 90},
		{"C", "C", 100},
		{"CD", "CD", 400},
		{"D", "D", 500},
		{"CM", "CM", 900},
		{"M", "M", 1000},
		{"MCMXCIV", "MCMXCIV", 1994},
		{"MMXXVI", "MMXXVI", 2026},
		{"MMMCMXCIX", "MMMCMXCIX", 3999},
		{"LVIII", "LVIII", 58},
		{"DCCCXC", "DCCCXC", 890},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RomanToInt(tt.input)
			if result != tt.expected {
				t.Errorf("RomanToInt(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - IntToRoman
// =============================================================

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{"1", 1, "I"},
		{"3", 3, "III"},
		{"4", 4, "IV"},
		{"5", 5, "V"},
		{"9", 9, "IX"},
		{"10", 10, "X"},
		{"40", 40, "XL"},
		{"42", 42, "XLII"},
		{"50", 50, "L"},
		{"90", 90, "XC"},
		{"100", 100, "C"},
		{"400", 400, "CD"},
		{"500", 500, "D"},
		{"900", 900, "CM"},
		{"1000", 1000, "M"},
		{"1994", 1994, "MCMXCIV"},
		{"2026", 2026, "MMXXVI"},
		{"3999", 3999, "MMMCMXCIX"},
		{"58", 58, "LVIII"},
		{"890", 890, "DCCCXC"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IntToRoman(tt.input)
			if result != tt.expected {
				t.Errorf("IntToRoman(%d) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. INTEGRATION TEST - Roman Numeral Round Trip
// =============================================================

func TestRomanRoundTrip(t *testing.T) {
	for i := 1; i <= 3999; i++ {
		roman := IntToRoman(i)
		back := RomanToInt(roman)
		if back != i {
			t.Errorf("IntToRoman(%d) = %q, RomanToInt(%q) = %d; want %d", i, roman, roman, back, i)
		}
	}
}

// =============================================================
// 4. TABLE-DRIVEN TEST - IsPalindromeNumber
// =============================================================

func TestIsPalindromeNumber(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"121", 121, true},
		{"negatif", -121, false},
		{"10", 10, false},
		{"0", 0, true},
		{"12321", 12321, true},
		{"satu digit", 7, true},
		{"1001", 1001, true},
		{"1234", 1234, false},
		{"11", 11, true},
		{"1000", 1000, false},
		{"12344321", 12344321, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsPalindromeNumber(tt.input)
			if result != tt.expected {
				t.Errorf("IsPalindromeNumber(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 5. TABLE-DRIVEN TEST - DecimalToBase
// =============================================================

func TestDecimalToBase(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		base     int
		expected string
	}{
		{"255 ke hex", 255, 16, "FF"},
		{"10 ke binary", 10, 2, "1010"},
		{"8 ke octal", 8, 8, "10"},
		{"0 ke binary", 0, 2, "0"},
		{"0 ke hex", 0, 16, "0"},
		{"255 ke binary", 255, 2, "11111111"},
		{"16 ke hex", 16, 16, "10"},
		{"100 ke desimal", 100, 10, "100"},
		{"15 ke hex", 15, 16, "F"},
		{"26 ke hex", 26, 16, "1A"},
		{"7 ke octal", 7, 8, "7"},
		{"1 ke binary", 1, 2, "1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DecimalToBase(tt.n, tt.base)
			if result != tt.expected {
				t.Errorf("DecimalToBase(%d, %d) = %q; want %q", tt.n, tt.base, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 6. TABLE-DRIVEN TEST - BaseToDecimal
// =============================================================

func TestBaseToDecimal(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		base      int
		expected  int
		expectErr bool
	}{
		{"hex FF", "FF", 16, 255, false},
		{"hex lowercase ff", "ff", 16, 255, false},
		{"binary 1010", "1010", 2, 10, false},
		{"octal 10", "10", 8, 8, false},
		{"desimal 100", "100", 10, 100, false},
		{"hex 1A", "1A", 16, 26, false},
		{"invalid hex GG", "GG", 16, 0, true},
		{"string kosong", "", 10, 0, true},
		{"binary invalid", "102", 2, 0, true},
		{"octal invalid", "89", 8, 0, true},
		{"single digit", "5", 10, 5, false},
		{"nol", "0", 10, 0, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, hasErr := BaseToDecimal(tt.s, tt.base)

			if tt.expectErr {
				if !hasErr {
					t.Errorf("BaseToDecimal(%q, %d) expected error, got (%d, false)", tt.s, tt.base, result)
				}
				return
			}

			if hasErr {
				t.Errorf("BaseToDecimal(%q, %d) unexpected error", tt.s, tt.base)
				return
			}

			if result != tt.expected {
				t.Errorf("BaseToDecimal(%q, %d) = %d; want %d", tt.s, tt.base, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 7. TABLE-DRIVEN TEST - PrimeFactors
// =============================================================

func TestPrimeFactors(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []int
	}{
		{"12", 12, []int{2, 2, 3}},
		{"60", 60, []int{2, 2, 3, 5}},
		{"17 prima", 17, []int{17}},
		{"100", 100, []int{2, 2, 5, 5}},
		{"2", 2, []int{2}},
		{"8", 8, []int{2, 2, 2}},
		{"30", 30, []int{2, 3, 5}},
		{"49", 49, []int{7, 7}},
		{"97 prima", 97, []int{97}},
		{"360", 360, []int{2, 2, 2, 3, 3, 5}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PrimeFactors(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PrimeFactors(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 8. TABLE-DRIVEN TEST - IsArmstrong
// =============================================================

func TestIsArmstrong(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected bool
	}{
		{"0", 0, true},
		{"1", 1, true},
		{"153", 153, true},
		{"370", 370, true},
		{"371", 371, true},
		{"407", 407, true},
		{"9474", 9474, true},
		{"123", 123, false},
		{"10", 10, false},
		{"100", 100, false},
		{"satu digit 5", 5, true},
		{"satu digit 9", 9, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsArmstrong(tt.input)
			if result != tt.expected {
				t.Errorf("IsArmstrong(%d) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 9. TABLE-DRIVEN TEST - NextPrime
// =============================================================

func TestNextPrime(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"setelah 0", 0, 2},
		{"setelah 1", 1, 2},
		{"setelah 2", 2, 3},
		{"setelah 3", 3, 5},
		{"setelah 10", 10, 11},
		{"setelah 13", 13, 17},
		{"setelah 20", 20, 23},
		{"setelah 96", 96, 97},
		{"setelah 97", 97, 101},
		{"setelah 100", 100, 101},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NextPrime(tt.input)
			if result != tt.expected {
				t.Errorf("NextPrime(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 10. TABLE-DRIVEN TEST - PrimesBetween
// =============================================================

func TestPrimesBetween(t *testing.T) {
	tests := []struct {
		name     string
		a, b     int
		expected []int
	}{
		{"10 sampai 20", 10, 20, []int{11, 13, 17, 19}},
		{"1 sampai 10", 1, 10, []int{2, 3, 5, 7}},
		{"terbalik", 20, 10, []int{}},
		{"tidak ada prima", 14, 16, []int{}},
		{"range sempit ada", 2, 2, []int{2}},
		{"range sempit tidak ada", 4, 4, []int{}},
		{"range besar", 1, 30, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}},
		{"range 50-60", 50, 60, []int{53, 59}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := PrimesBetween(tt.a, tt.b)
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("PrimesBetween(%d, %d) = %v; want %v", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 11. TABLE-DRIVEN TEST - DigitalRoot
// =============================================================

func TestDigitalRoot(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected int
	}{
		{"942", 942, 6},
		{"132189", 132189, 6},
		{"0", 0, 0},
		{"satu digit", 5, 5},
		{"10", 10, 1},
		{"99", 99, 9},
		{"999", 999, 9},
		{"123456789", 123456789, 9},
		{"38", 38, 2},
		{"1", 1, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DigitalRoot(tt.input)
			if result != tt.expected {
				t.Errorf("DigitalRoot(%d) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 12. INTEGRATION TEST - DecimalToBase & BaseToDecimal Round Trip
// =============================================================

func TestBaseConversion_RoundTrip(t *testing.T) {
	bases := []int{2, 8, 10, 16}
	for _, base := range bases {
		for n := 0; n <= 255; n++ {
			str := DecimalToBase(n, base)
			back, hasErr := BaseToDecimal(str, base)
			if hasErr {
				t.Errorf("BaseToDecimal(%q, %d) returned error for n=%d", str, base, n)
				continue
			}
			if back != n {
				t.Errorf("DecimalToBase(%d, %d) = %q, BaseToDecimal(%q, %d) = %d; want %d", n, base, str, str, base, back, n)
			}
		}
	}
}

// =============================================================
// 13. PARALLEL TEST
// =============================================================

func TestRomanToInt_Parallel(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"case1", "XIV", 14},
		{"case2", "XCIX", 99},
		{"case3", "MMMCMXCIX", 3999},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := RomanToInt(tt.input)
			if result != tt.expected {
				t.Errorf("RomanToInt(%q) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 14. BENCHMARK TEST
// =============================================================

func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RomanToInt("MCMXCIV")
	}
}

func BenchmarkIntToRoman(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntToRoman(1994)
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrimeFactors(360)
	}
}

func BenchmarkIsArmstrong(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsArmstrong(9474)
	}
}

func BenchmarkDigitalRoot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DigitalRoot(123456789)
	}
}

func BenchmarkDecimalToBase(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DecimalToBase(255, 16)
	}
}
