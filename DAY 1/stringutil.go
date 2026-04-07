package belajar

import (
	"strings"
	"unicode"
)

// Reverse membalik urutan karakter dalam string
// Contoh: Reverse("hello") -> "olleh"
//
//	Reverse("golang") -> "gnalog"
func Reverse(s string) string {
	// TODO: implementasi di sini
	uniqcode := []rune(s)
	for a, b := 0, len(uniqcode)-1; a < b; a, b = a+1, b-1 {
		uniqcode[a], uniqcode[b] = uniqcode[b], uniqcode[a]
	}
	return string(uniqcode)
}

// IsPalindrome mengecek apakah string adalah palindrome (case-insensitive)
// Contoh: IsPalindrome("katak") -> true
//
//	IsPalindrome("Racecar") -> true
//	IsPalindrome("hello") -> false
func IsPalindrome(s string) bool {
	// TODO: implementasi di sini
	val := strings.ToLower(s)
	uniqcode := []rune(val)
	for a, b := 0, len(uniqcode)-1; a < b; a, b = a+1, b-1 {
		if !(uniqcode[a] == uniqcode[b]) {
			return false
		}
	}
	return true
}

// CountVowels menghitung jumlah huruf vokal (a, i, u, e, o) dalam string
// Case-insensitive
// Contoh: CountVowels("hello") -> 2
//
//	CountVowels("AEIOU") -> 5
func CountVowels(s string) int {
	// TODO: implementasi di sini
	val := strings.ToLower(s)
	count := 0
	for k, _ := range s {
		switch val[k] {
		case 'a':
			count += 1
		case 'i':
			count += 1
		case 'u':
			count += 1
		case 'e':
			count += 1
		case 'o':
			count += 1
		}
	}
	return count
}

// CamelToSnake mengubah camelCase / PascalCase ke snake_case
// Contoh: CamelToSnake("helloWorld") -> "hello_world"
//
//	CamelToSnake("MyFunc") -> "my_func"
func CamelToSnake(s string) string {
	// TODO: implementasi di sini
	val := ""
	for k, v := range s {
		if k == 0 && v == unicode.ToUpper(v) {
			val = val + string(unicode.ToLower(v))
		} else if v == unicode.ToUpper(v) {
			val = val + "_" + string(unicode.ToLower(v))
		} else {
			val = val + string(v)
		}
	}
	return val
}

// Truncate memotong string ke panjang maxLen
// Jika string lebih panjang dari maxLen, tambahkan "..." di akhir
// Total panjang hasil termasuk "..." tidak boleh lebih dari maxLen
// Contoh: Truncate("Hello World", 5) -> "He..."
//
//	Truncate("Hi", 10) -> "Hi"
func Truncate(s string, maxLen int) string {
	// TODO: implementasi di sini
	if len(s) == 0 || maxLen <= 0 {
		return ""
	}
	val := ""
	if len(s) > maxLen {
		for i := 0; i < maxLen; i++ {
			if string(s[i]) == " " {
				val = val + "..."
				break
			}
			val = val + string(s[i])
		}
		return val
	}
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			val = val + "..."
			break
		}
		val = val + string(s[i])
	}
	return val
}

// WordCount menghitung jumlah kata dalam string (dipisahkan spasi)
// Contoh: WordCount("hello world") -> 2
//
//	WordCount("  hello   world  ") -> 2
//	WordCount("") -> 0
func WordCount(s string) int {
	// TODO: implementasi di sini
	val := strings.ReplaceAll(strings.ReplaceAll(s, "\t", " "), "\n", " ")
	lastSpace := true
	count := 0

	for _, v := range val {
		if v == ' ' {
			lastSpace = true
		} else {
			if lastSpace {
				count++
			}
			lastSpace = false
		}
	}

	return count
}
