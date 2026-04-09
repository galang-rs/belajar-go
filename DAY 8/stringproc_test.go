package belajar

import (
	"testing"
)

// =============================================================
// 1. TABLE-DRIVEN TEST - IsAnagram
// =============================================================

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name     string
		s1, s2   string
		expected bool
	}{
		{"anagram klasik", "listen", "silent", true},
		{"case insensitive", "Hello", "Olelh", true},
		{"bukan anagram", "hello", "world", false},
		{"dengan spasi", "Astronomer", "Moon starer", true},
		{"panjang beda", "abc", "abcd", false},
		{"string kosong", "", "", true},
		{"satu kosong", "abc", "", false},
		{"huruf sama beda jumlah", "aab", "abb", false},
		{"kata sama", "test", "test", true},
		{"anagram panjang", "debit card", "bad credit", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsAnagram(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("IsAnagram(%q, %q) = %v; want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 2. TABLE-DRIVEN TEST - LongestCommonPrefix
// =============================================================

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{"normal", []string{"flower", "flow", "flight"}, "fl"},
		{"tidak ada prefix", []string{"dog", "racecar", "car"}, ""},
		{"prefix panjang", []string{"interspecies", "interstellar", "interstate"}, "inters"},
		{"slice kosong", []string{}, ""},
		{"satu string", []string{"alone"}, "alone"},
		{"semua sama", []string{"abc", "abc", "abc"}, "abc"},
		{"satu char sama", []string{"a", "ab", "ac"}, "a"},
		{"ada string kosong", []string{"", "abc", "abd"}, ""},
		{"dua string", []string{"hello", "help"}, "hel"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestCommonPrefix(tt.input)
			if result != tt.expected {
				t.Errorf("LongestCommonPrefix(%v) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 3. TABLE-DRIVEN TEST - Compress
// =============================================================

func TestCompress(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"normal", "aaabbc", "a3b2c1"},
		{"semua beda", "abc", "a1b1c1"},
		{"semua sama", "aaaa", "a4"},
		{"string kosong", "", ""},
		{"satu karakter", "x", "x1"},
		{"dua karakter sama", "aa", "a2"},
		{"pola berulang", "aabbcc", "a2b2c2"},
		{"karakter tunggal di antara", "aabccc", "a2b1c3"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Compress(tt.input)
			if result != tt.expected {
				t.Errorf("Compress(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 4. TABLE-DRIVEN TEST - Decompress
// =============================================================

func TestDecompress(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"normal", "a3b2c1", "aaabbc"},
		{"satu karakter", "x1", "x"},
		{"banyak pengulangan", "x5y1", "xxxxxy"},
		{"string kosong", "", ""},
		{"semua satu", "a1b1c1", "abc"},
		{"angka besar", "z10", "zzzzzzzzzz"},
		{"pola berulang", "a2b2c2", "aabbcc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Decompress(tt.input)
			if result != tt.expected {
				t.Errorf("Decompress(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 5. TABLE-DRIVEN TEST - CaesarEncrypt
// =============================================================

func TestCaesarEncrypt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		shift    int
		expected string
	}{
		{"shift 3", "abc", 3, "def"},
		{"wrap around", "xyz", 3, "abc"},
		{"dengan spesial char", "Hello, World!", 5, "Mjqqt, Btwqi!"},
		{"shift negatif", "def", -3, "abc"},
		{"shift 0", "hello", 0, "hello"},
		{"shift 26 (full cycle)", "hello", 26, "hello"},
		{"huruf besar", "ABC", 1, "BCD"},
		{"campuran", "Go 1.21!", 10, "Qy 1.21!"},
		{"string kosong", "", 5, ""},
		{"wrap besar negatif", "abc", -1, "zab"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CaesarEncrypt(tt.input, tt.shift)
			if result != tt.expected {
				t.Errorf("CaesarEncrypt(%q, %d) = %q; want %q", tt.input, tt.shift, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 6. TABLE-DRIVEN TEST - CaesarDecrypt
// =============================================================

func TestCaesarDecrypt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		shift    int
		expected string
	}{
		{"shift 3", "def", 3, "abc"},
		{"dengan spesial char", "Mjqqt, Btwqi!", 5, "Hello, World!"},
		{"shift 0", "hello", 0, "hello"},
		{"wrap around", "abc", 3, "xyz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CaesarDecrypt(tt.input, tt.shift)
			if result != tt.expected {
				t.Errorf("CaesarDecrypt(%q, %d) = %q; want %q", tt.input, tt.shift, result, tt.expected)
			}
		})
	}

	// Round-trip test: encrypt lalu decrypt harus kembali ke asli
	t.Run("round trip", func(t *testing.T) {
		original := "The Quick Brown Fox Jumps Over The Lazy Dog! 123"
		for shift := -30; shift <= 30; shift++ {
			encrypted := CaesarEncrypt(original, shift)
			decrypted := CaesarDecrypt(encrypted, shift)
			if decrypted != original {
				t.Errorf("Round trip failed for shift %d: got %q", shift, decrypted)
			}
		}
	})
}

// =============================================================
// 7. TABLE-DRIVEN TEST - IsValidBrackets
// =============================================================

func TestIsValidBrackets(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"semua jenis valid", "()[]{}",  true},
		{"nested valid", "([{}])", true},
		{"salah pasangan", "(]", false},
		{"interleaved salah", "([)]", false},
		{"string kosong", "", true},
		{"hanya buka", "(((", false},
		{"hanya tutup", ")))", false},
		{"dengan teks", "hello(world)", true},
		{"complex valid", "{[()]}([])", true},
		{"satu pasang", "()", true},
		{"kurung kurawal salah", "{]", false},
		{"campuran teks valid", "func() { arr[0] = (x + y) }", true},
		{"lebih tutup", "())", false},
		{"lebih buka", "(()", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsValidBrackets(tt.input)
			if result != tt.expected {
				t.Errorf("IsValidBrackets(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 8. TABLE-DRIVEN TEST - CountSubstring
// =============================================================

func TestCountSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s, sub   string
		expected int
	}{
		{"normal", "hello hello hello", "hello", 3},
		{"overlap handled", "banana", "ana", 1},
		{"tidak ditemukan", "abc", "xyz", 0},
		{"substring kosong", "abc", "", 0},
		{"string kosong", "", "abc", 0},
		{"keduanya kosong", "", "", 0},
		{"satu kemunculan", "hello world", "world", 1},
		{"case sensitive", "Hello hello", "hello", 1},
		{"substring = string", "abc", "abc", 1},
		{"aaa cari aa", "aaa", "aa", 1},
		{"banyak kemunculan", "abababab", "ab", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CountSubstring(tt.s, tt.sub)
			if result != tt.expected {
				t.Errorf("CountSubstring(%q, %q) = %d; want %d", tt.s, tt.sub, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 9. TABLE-DRIVEN TEST - TitleCase
// =============================================================

func TestTitleCase(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"normal", "hello world", "Hello World"},
		{"sudah title case", "Hello World", "Hello World"},
		{"semua kapital", "gO iS fUn", "Go Is Fun"},
		{"string kosong", "", ""},
		{"satu kata", "hello", "Hello"},
		{"satu huruf", "a", "A"},
		{"banyak spasi", "hello   world", "Hello World"},
		{"huruf kecil semua", "the quick brown fox", "The Quick Brown Fox"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TitleCase(tt.input)
			if result != tt.expected {
				t.Errorf("TitleCase(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 10. TABLE-DRIVEN TEST - RemoveDuplicateChars
// =============================================================

func TestRemoveDuplicateChars(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"normal", "abcabc", "abc"},
		{"hello", "hello", "helo"},
		{"case sensitive", "aAbBcC", "aAbBcC"},
		{"string kosong", "", ""},
		{"semua sama", "aaaa", "a"},
		{"sudah unik", "abcdef", "abcdef"},
		{"dengan spasi", "aba cdc", "ab cd"},
		{"satu karakter", "x", "x"},
		{"palindrome", "abcba", "abc"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RemoveDuplicateChars(tt.input)
			if result != tt.expected {
				t.Errorf("RemoveDuplicateChars(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 11. INTEGRATION TEST - Compress & Decompress Round Trip
// =============================================================

func TestCompressDecompress_RoundTrip(t *testing.T) {
	tests := []string{
		"aaabbc",
		"abc",
		"aaaaaa",
		"x",
		"aabbccdd",
	}

	for _, original := range tests {
		t.Run(original, func(t *testing.T) {
			compressed := Compress(original)
			decompressed := Decompress(compressed)
			if decompressed != original {
				t.Errorf("Compress->Decompress(%q): compressed=%q, decompressed=%q", original, compressed, decompressed)
			}
		})
	}
}

// =============================================================
// 12. PARALLEL TEST
// =============================================================

func TestIsAnagram_Parallel(t *testing.T) {
	tests := []struct {
		name     string
		s1, s2   string
		expected bool
	}{
		{"case1", "abc", "cba", true},
		{"case2", "hello", "world", false},
		{"case3", "listen", "silent", true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := IsAnagram(tt.s1, tt.s2)
			if result != tt.expected {
				t.Errorf("IsAnagram(%q, %q) = %v; want %v", tt.s1, tt.s2, result, tt.expected)
			}
		})
	}
}

// =============================================================
// 13. BENCHMARK TEST
// =============================================================

func BenchmarkIsAnagram(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsAnagram("astronomer", "moon starer")
	}
}

func BenchmarkCompress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Compress("aaabbbcccdddeee")
	}
}

func BenchmarkCaesarEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CaesarEncrypt("The Quick Brown Fox Jumps Over The Lazy Dog", 13)
	}
}

func BenchmarkIsValidBrackets(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValidBrackets("{[()]}([]){{}}")
	}
}
