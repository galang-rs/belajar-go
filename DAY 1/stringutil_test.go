package belajar

import (
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - Reverse
// =============================================================

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"kata biasa", "hello", "olleh"},
		{"string kosong", "", ""},
		{"satu huruf", "a", "a"},
		{"palindrome", "madam", "madam"},
		{"angka", "12345", "54321"},
		{"unicode/emoji", "go🚀", "🚀og"},
		{"spasi", "hello world", "dlrow olleh"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Reverse(tt.input)
			if result != tt.expected {
				t.Errorf("Reverse(%q) = %q; expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TEST BOOLEAN RESULT - IsPalindrome
// =============================================================

func TestIsPalindrome(t *testing.T) {
	t.Run("palindrome", func(t *testing.T) {
		palindromes := []string{"madam", "katak", "racecar", "Racecar", "aba", "a", ""}
		for _, s := range palindromes {
			if !IsPalindrome(s) {
				t.Errorf("IsPalindrome(%q) = false; expected true", s)
			}
		}
	})

	t.Run("bukan palindrome", func(t *testing.T) {
		notPalindromes := []string{"hello", "golang", "ab", "test"}
		for _, s := range notPalindromes {
			if IsPalindrome(s) {
				t.Errorf("IsPalindrome(%q) = true; expected false", s)
			}
		}
	})
}

// =============================================================
// 3. TABLE-DRIVEN TEST - CountVowels
// =============================================================

func TestCountVowels(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"hello", "hello", 2},
		{"semua vokal", "aeiou", 5},
		{"huruf besar", "AEIOU", 5},
		{"tanpa vokal", "rhythm", 0},
		{"string kosong", "", 0},
		{"campuran", "Golang Itu Asyik", 6},
		{"satu huruf vokal", "a", 1},
		{"satu huruf konsonan", "b", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountVowels(tt.input)
			if result != tt.expected {
				t.Errorf("CountVowels(%q) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 4. TEST STRING TRANSFORMATION - CamelToSnake
// =============================================================

func TestCamelToSnake(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"camelCase", "helloWorld", "hello_world"},
		{"PascalCase", "MyFunc", "my_func"},
		{"huruf kecil semua", "hello", "hello"},
		{"satu kata besar", "Hello", "hello"},
		{"banyak kata", "thisIsALongName", "this_is_a_long_name"},
		{"string kosong", "", ""},
		{"sudah lowercase", "alreadylower", "alreadylower"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CamelToSnake(tt.input)
			if result != tt.expected {
				t.Errorf("CamelToSnake(%q) = %q; expected %q", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 5. TEST DENGAN EDGE CASES - Truncate
// =============================================================

func TestTruncate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		maxLen   int
		expected string
	}{
		{"tidak dipotong", "Hi", 10, "Hi"},
		{"pas", "Hello", 5, "Hello"},
		{"dipotong", "Hello World", 8, "Hello..."},
		{"sangat pendek", "Hello World", 3, "Hel"},
		{"maxLen 0", "Hello", 0, ""},
		{"string kosong", "", 5, ""},
		{"maxLen negatif", "Hello", -1, ""},
		{"panjang 1", "Hello", 1, "H"},
		{"unicode", "Halo Dunia!", 7, "Halo..."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Truncate(tt.input, tt.maxLen)
			if result != tt.expected {
				t.Errorf("Truncate(%q, %d) = %q; expected %q", tt.input, tt.maxLen, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 6. TEST WORD COUNT - Test parsing string
// =============================================================

func TestWordCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"dua kata", "hello world", 2},
		{"satu kata", "hello", 1},
		{"string kosong", "", 0},
		{"spasi saja", "   ", 0},
		{"banyak spasi", "  hello   world  ", 2},
		{"tiga kata", "saya belajar golang", 3},
		{"tab dan newline", "hello\tworld\nfoo", 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := WordCount(tt.input)
			if result != tt.expected {
				t.Errorf("WordCount(%q) = %d; expected %d", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 7. BENCHMARK TEST - Performa string operations
// =============================================================

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("The quick brown fox jumps over the lazy dog")
	}
}

func BenchmarkCountVowels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountVowels("The quick brown fox jumps over the lazy dog")
	}
}

func BenchmarkCamelToSnake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CamelToSnake("thisIsAVeryLongCamelCaseString")
	}
}
