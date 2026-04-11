package belajar

import (
	"math"
	"strconv"
	"strings"
)

// IsAnagram mengecek apakah dua string merupakan anagram (case-insensitive).
// Anagram = memiliki huruf yang sama dengan jumlah yang sama (spasi diabaikan).
// Contoh: IsAnagram("listen", "silent") -> true
//
//	IsAnagram("Hello", "Olelh") -> true
//	IsAnagram("hello", "world") -> false
//	IsAnagram("Astronomer", "Moon starer") -> true
//
// Hint: gunakan strings.ToLower, abaikan spasi
func IsAnagram(s1, s2 string) bool {
	// TODO: implementasi di sini
	val1 := strings.ToLower(strings.Replace(s1, " ", "", len(s1)))
	val2 := strings.ToLower(strings.Replace(s2, " ", "", len(s2)))

	if len(val1) != len(val2) {
		return false
	}
	if val1 == val2 {
		return true
	}

	data1 := make(map[rune]int)
	data2 := make(map[rune]int)

	for _, v := range val1 {
		data1[v]++
	}
	for _, v := range val2 {
		data2[v]++
	}

	for k := range data1 {
		if data1[k] != data2[k] {
			return false
		}
	}
	for k := range data2 {
		if data1[k] != data2[k] {
			return false
		}
	}

	return true
}

// LongestCommonPrefix mengembalikan prefix (awalan) terpanjang yang sama dari semua string di slice.
// Contoh: LongestCommonPrefix([]string{"flower", "flow", "flight"}) -> "fl"
//
//	LongestCommonPrefix([]string{"dog", "racecar", "car"}) -> ""
//	LongestCommonPrefix([]string{"interspecies", "interstellar", "interstate"}) -> "inters"
//	LongestCommonPrefix([]string{}) -> ""
//	LongestCommonPrefix([]string{"alone"}) -> "alone"
func LongestCommonPrefix(strs []string) string {
	// TODO: implementasi di sini

	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	var lowLength float64 = math.Inf(1)
	var indx int

	for k, v1 := range strs {
		if lowLength >= float64(len(v1)) {
			lowLength = float64(len(v1))
			indx = k
		}
	}

	var word string
	for i := 0; i < int(lowLength); i++ {
		data := strs[indx][i]
		done := true
		for _, v := range strs {
			if data != v[i] {
				return v[:i]
			}
			if done {
				word = word + string(data)
				done = false
			}
		}
	}
	return word
}

// Compress melakukan Run-Length Encoding pada string.
// Setiap karakter berurutan yang sama diikuti jumlah kemunculannya.
// Contoh: Compress("aaabbc") -> "a3b2c1"
//
//	Compress("aabbcc") -> "a2b2c2"
//	Compress("abc") -> "a1b1c1"
//	Compress("") -> ""
func Compress(s string) string {
	// TODO: implementasi di sini

	data := make(map[rune]int)

	for _, v := range s {
		data[v]++
	}
	val := ""
	for k, v := range data {
		vale := strconv.Itoa(v)
		val = val + string(k) + vale
	}
	return val
}

// Decompress membalikkan Run-Length Encoding.
// Contoh: Decompress("a3b2c1") -> "aaabbc"
//
//	Decompress("x5y1") -> "xxxxxy"
//	Decompress("") -> ""
//
// Asumsi: format input selalu valid (huruf diikuti angka 1 digit atau lebih).
// Hint: gunakan strconv.Atoi untuk parsing angka lebih dari 1 digit
func Decompress(s string) string {
	// TODO: implementasi di sini

	word := make(map[rune]string)
	var lastWord rune
	for _, v := range s {
		switch v {
		case 48, 49, 50, 51, 52, 53, 54, 55, 56, 57:
			word[lastWord] = string(word[lastWord]) + string(v)
		default:
			lastWord = v
			word[lastWord] = ""
		}
	}
	scope := ""
	for k, v := range word {
		val, _ := strconv.Atoi(v)
		for i := 0; i < val; i++ {
			scope = scope + string(k)
		}
	}
	return scope
}

// CaesarEncrypt mengenkripsi string menggunakan Caesar cipher.
// Hanya huruf a-z dan A-Z yang digeser, karakter lain tetap.
// Shift bisa negatif (geser ke kiri) atau positif (geser ke kanan).
// Contoh: CaesarEncrypt("abc", 3) -> "def"
//
//	CaesarEncrypt("xyz", 3) -> "abc" (wrap around)
//	CaesarEncrypt("Hello, World!", 5) -> "Mjqqt, Btwqi!"
//	CaesarEncrypt("abc", -1) -> "zab"
func CaesarEncrypt(s string, shift int) string {
	result := ""

	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			offset := int(c - 'a')
			newPos := ((offset+shift)%26 + 26) % 26
			result += string(rune(newPos) + 'a')
		} else if c >= 'A' && c <= 'Z' {
			offset := int(c - 'A')
			newPos := ((offset+shift)%26 + 26) % 26
			result += string(rune(newPos) + 'A')
		} else {
			result += string(c)
		}
	}

	return result
}

// CaesarDecrypt mendekripsi string yang dienkripsi Caesar cipher.
// Contoh: CaesarDecrypt("def", 3) -> "abc"
//
//	CaesarDecrypt("Mjqqt, Btwqi!", 5) -> "Hello, World!"
//
// Hint: dekripsi = enkripsi dengan shift negatif
func CaesarDecrypt(s string, shift int) string {
	return CaesarEncrypt(s, -shift)
}

// IsValidBrackets mengecek apakah urutan bracket/kurung dalam string valid (balanced).
// Bracket yang didukung: (), [], {}
// Karakter selain bracket diabaikan.
// Contoh: IsValidBrackets("()[]{}") -> true
//
//	IsValidBrackets("([{}])") -> true
//	IsValidBrackets("(]") -> false
//	IsValidBrackets("([)]") -> false
//	IsValidBrackets("") -> true
//	IsValidBrackets("hello(world)") -> true
func IsValidBrackets(s string) bool {
	// TODO: implementasi di sini
	return false
}

// CountSubstring menghitung berapa kali substring muncul dalam string (tidak overlap).
// Case-sensitive.
// Contoh: CountSubstring("banana", "ana") -> 1 (tidak overlap: b[ana]na, bukan b[ana][na])
//
//	CountSubstring("hello hello hello", "hello") -> 3
//	CountSubstring("aaa", "aa") -> 1
//	CountSubstring("abc", "xyz") -> 0
//	CountSubstring("abc", "") -> 0
//
// Hint: gunakan strings.Count atau implementasi manual
func CountSubstring(s, sub string) int {
	// TODO: implementasi di sini
	return 0
}

// TitleCase mengubah huruf pertama setiap kata menjadi huruf kapital, sisanya huruf kecil.
// Kata dipisahkan oleh spasi.
// Contoh: TitleCase("hello world") -> "Hello World"
//
//	TitleCase("gO iS fUn") -> "Go Is Fun"
//	TitleCase("") -> ""
//	TitleCase("a") -> "A"
func TitleCase(s string) string {
	// TODO: implementasi di sini
	// Hint: gunakan strings.Fields dan strings.ToUpper/strings.ToLower
	return ""
}

// RemoveDuplicateChars menghapus karakter duplikat, hanya menyisakan kemunculan pertama.
// Case-sensitive.
// Contoh: RemoveDuplicateChars("abcabc") -> "abc"
//
//	RemoveDuplicateChars("hello") -> "helo"
//	RemoveDuplicateChars("aAbBcC") -> "aAbBcC" (case-sensitive, semua unik)
//	RemoveDuplicateChars("") -> ""
func RemoveDuplicateChars(s string) string {
	// TODO: implementasi di sini
	return ""
}
