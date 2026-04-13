package belajar

import (
	"fmt"
	"math"
	"strconv"
)

// RomanToInt mengonversi string angka Romawi ke integer.
// Input dijamin valid dan berada di rentang 1-3999.
// Simbol Romawi: I=1, V=5, X=10, L=50, C=100, D=500, M=1000
// Aturan subtractive: IV=4, IX=9, XL=40, XC=90, CD=400, CM=900
// Contoh: RomanToInt("III") -> 3
//
//	RomanToInt("IV") -> 4
//	RomanToInt("IX") -> 9
//	RomanToInt("XLII") -> 42
//	RomanToInt("MCMXCIV") -> 1994
func RomanToInt(s string) int {
	// TODO: implementasi di sini
	var count int
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'I':
			if i+2 < len(s) {
				if s[i+1] == 'I' && s[i+2] == 'I' {
					count += 3
					i += 2
				} else if s[i+1] == 'I' {
					count += 2
					i += 1
				} else {
					switch s[i+1] {
					case 'V':
						count += 4
					case 'X':
						count += 9
					case 'L':
						count += 49
					case 'C':
						count += 99
					case 'D':
						count += 499
					case 'M':
						count += 999
					}
					i += 1
				}
			} else if i+1 < len(s) {
				if s[i+1] == 'I' {
					count += 2
					i += 1
				} else {
					switch s[i+1] {
					case 'V':
						count += 4
					case 'X':
						count += 9
					case 'L':
						count += 49
					case 'C':
						count += 99
					case 'D':
						count += 499
					case 'M':
						count += 999
					}
					i += 1
				}
			} else {
				count += 1
			}
		case 'V':
			count += 5
		case 'X':
			if i+2 < len(s) {
				if s[i+1] == 'X' && s[i+2] == 'X' {
					count += 30
					i += 2
				} else if s[i+1] == 'X' {
					count += 20
					i += 1
				} else {
					switch s[i+1] {
					case 'L':
						count += 40
					case 'C':
						count += 90
					case 'D':
						count += 490
					case 'M':
						count += 990
					}
					i += 1
				}
			} else if i+1 < len(s) {
				if s[i+1] == 'I' {
					count += 2
					i += 1
				} else {
					switch s[i+1] {
					case 'L':
						count += 40
					case 'C':
						count += 90
					case 'D':
						count += 490
					case 'M':
						count += 990
					}
					i += 1
				}
			} else {
				count += 10
			}
		case 'L':
			count += 50
		case 'C':
			if i+2 < len(s) {
				if s[i+1] == 'C' && s[i+2] == 'C' {
					count += 300
					i += 2
				} else if s[i+1] == 'C' {
					count += 200
					i += 1
				} else {
					switch s[i+1] {
					case 'D':
						count += 400
					case 'M':
						count += 900
					}
					i += 1
				}
			} else if i+1 < len(s) {
				if s[i+1] == 'C' {
					count += 2
					i += 1
				} else {
					switch s[i+1] {
					case 'D':
						count += 400
					case 'M':
						count += 900
					}
					i += 1
				}
			} else {
				count += 100
			}
		case 'D':
			count += 500
		case 'M':
			if i+2 < len(s) {
				if s[i+1] == 'M' && s[i+2] == 'M' {
					count += 3000
					i += 2
				} else if s[i+1] == 'M' {
					count += 2000
					i += 1
				} else {
					count += 1000
				}
			} else if i+1 < len(s) {
				if s[i+1] == 'M' {
					count += 2000
					i += 1
				} else {
					count += 2000
				}
			} else {
				count += 1000
			}
		}

	}
	return count
}

// IntToRoman mengonversi integer (1-3999) ke string angka Romawi.
// Contoh: IntToRoman(3) -> "III"
//
//	IntToRoman(4) -> "IV"
//	IntToRoman(9) -> "IX"
//	IntToRoman(42) -> "XLII"
//	IntToRoman(1994) -> "MCMXCIV"
//	IntToRoman(3999) -> "MMMCMXCIX"
func IntToRoman(n int) string {
	// TODO: implementasi di sini
	var word string
	val := strconv.Itoa(n)
	words := val
	for i := 0; i < len(words); i++ {
		if 4 == len(words)-i {
			switch string(words[i]) {
			case "3":
				word = word + "MMM"
			case "2":
				word = word + "MM"
			case "1":
				word = word + "M"
			}
			fmt.Println(word, "Area 1")
		} else if 3 == len(words)-i {
			switch string(words[i]) {
			case "9":
				word = word + "CM"
			case "8":
				word = word + "DCCC"
			case "7":
				word = word + "DCC"
			case "6":
				word = word + "DC"
			case "5":
				word = word + "D"
			case "4":
				word = word + "CD"
			case "3":
				word = word + "CCC"
			case "2":
				word = word + "CC"
			case "1":
				word = word + "C"
			}
			fmt.Println(word, "Area 2")
		} else if 2 == len(words)-i {
			switch string(words[i]) {
			case "9":
				word = word + "XC"
			case "8":
				word = word + "LXXX"
			case "7":
				word = word + "LXX"
			case "6":
				word = word + "LX"
			case "5":
				word = word + "L"
			case "4":
				word = word + "XL"
			case "3":
				word = word + "XXX"
			case "2":
				word = word + "XX"
			case "1":
				word = word + "X"
			}
			fmt.Println(word, "Area 3")
		} else if 1 == len(words)-i {
			switch string(words[i]) {
			case "9":
				word = word + "IX"
			case "8":
				word = word + "VIII"
			case "7":
				word = word + "VII"
			case "6":
				word = word + "VI"
			case "5":
				word = word + "V"
			case "4":
				word = word + "IV"
			case "3":
				word = word + "III"
			case "2":
				word = word + "II"
			case "1":
				word = word + "I"
			}
			fmt.Println(word, "Area 4")
		}
		fmt.Println(word, "masuk sini ")
	}
	return word
}

// IsPalindromeNumber mengecek apakah bilangan integer adalah palindrome.
// Bilangan negatif bukan palindrome.
// Contoh: IsPalindromeNumber(121) -> true
//
//	IsPalindromeNumber(-121) -> false
//	IsPalindromeNumber(10) -> false
//	IsPalindromeNumber(0) -> true
//	IsPalindromeNumber(12321) -> true
func IsPalindromeNumber(n int) bool {
	// TODO: implementasi di sini

	return false
}

// DecimalToBase mengonversi bilangan desimal (non-negatif) ke string representasi di base tertentu (2-16).
// Gunakan huruf kapital untuk digit > 9 (A=10, B=11, ... F=15).
// Contoh: DecimalToBase(255, 16) -> "FF"
//
//	DecimalToBase(10, 2) -> "1010"
//	DecimalToBase(8, 8) -> "10"
//	DecimalToBase(0, 2) -> "0"
//	DecimalToBase(255, 2) -> "11111111"
func DecimalToBase(n int, base int) string {
	// TODO: implementasi di sini
	return ""
}

// BaseToDecimal mengonversi string bilangan dari base tertentu (2-16) ke desimal.
// Huruf case-insensitive (a-f dan A-F sama).
// Kembalikan hasil dan false jika valid, 0 dan true jika ada karakter tidak valid.
// Contoh: BaseToDecimal("FF", 16) -> 255, false
//
//	BaseToDecimal("1010", 2) -> 10, false
//	BaseToDecimal("10", 8) -> 8, false
//	BaseToDecimal("GG", 16) -> 0, true (G bukan digit hex)
//	BaseToDecimal("", 10) -> 0, true
func BaseToDecimal(s string, base int) (int, bool) {
	// TODO: implementasi di sini
	return 0, true
}

// PrimeFactors mengembalikan faktor-faktor prima dari bilangan n (n >= 2).
// Faktor diurutkan dari kecil ke besar, boleh ada duplikat.
// Contoh: PrimeFactors(12) -> []int{2, 2, 3}
//
//	PrimeFactors(60) -> []int{2, 2, 3, 5}
//	PrimeFactors(17) -> []int{17}
//	PrimeFactors(100) -> []int{2, 2, 5, 5}
func PrimeFactors(n int) []int {
	// TODO: implementasi di sini
	var hasilBagi int = n
	// var sisaBagi int = n

	var result []int
	for i := 2; i <= hasilBagi; {
		if hasilBagi%i == 0 {
			hasilBagi = hasilBagi / i
			result = append(result, i)
			fmt.Println(i, hasilBagi)
			i = 2
		} else {
			i++
		}
	}
	return result
}

// IsArmstrong mengecek apakah bilangan merupakan Armstrong number (Narcissistic number).
// Armstrong number: jumlah setiap digit dipangkatkan jumlah digit = bilangan itu sendiri.
// Contoh: IsArmstrong(153) -> true (1³ + 5³ + 3³ = 1 + 125 + 27 = 153)
//
//	IsArmstrong(370) -> true (3³ + 7³ + 0³ = 27 + 343 + 0 = 370)
//	IsArmstrong(9474) -> true (9⁴ + 4⁴ + 7⁴ + 4⁴ = 6561 + 256 + 2401 + 256 = 9474)
//	IsArmstrong(123) -> false
//	IsArmstrong(0) -> true
//
// Hint: gunakan math.Pow
func IsArmstrong(n int) bool {
	// TODO: implementasi di sini
	if n == 0 {
		return true
	}

	val := strconv.Itoa(n)
	for val1 := 0; val1 <= 9; val1++ {
		var count int
		for _, v2 := range val {
			val2, _ := strconv.Atoi(string(v2))
			fmt.Print(val1, val2, int(math.Pow(float64(val1), float64(val2))))

			count += int(math.Pow(float64(val2), float64(val1)))
		}
		fmt.Println(count)
		if count == n {
			return true
		}
		count = 0
	}

	return false
}

// NextPrime mengembalikan bilangan prima terkecil yang lebih besar dari n.
// Contoh: NextPrime(10) -> 11
//
//	NextPrime(13) -> 17
//	NextPrime(0) -> 2
//	NextPrime(1) -> 2
//	NextPrime(2) -> 3
func NextPrime(n int) int {
	// TODO: implementasi di sini
	i := n
	for true {
		i++
		if len(PrimeFactors(i)) == 1 {
			return i
		}
	}
	return 0
}

// PrimesBetween mengembalikan semua bilangan prima antara a dan b (inklusif).
// Jika a > b, kembalikan slice kosong.
// Contoh: PrimesBetween(10, 20) -> []int{11, 13, 17, 19}
//
//	PrimesBetween(1, 10) -> []int{2, 3, 5, 7}
//	PrimesBetween(20, 10) -> []int{}
//	PrimesBetween(14, 16) -> []int{}
func PrimesBetween(a, b int) []int {
	// TODO: implementasi di sini
	if a > b {
		return []int{}
	}
	var val []int
	for i := a; i <= b; i++ {
		if len(PrimeFactors(i)) == 1 {
			val = append(val, i)
		}
	}
	return val
}

// DigitalRoot menghitung digital root dari bilangan non-negatif.
// Digital root: jumlahkan semua digit secara berulang sampai tersisa satu digit.
// Contoh: DigitalRoot(942) -> 6 (9+4+2=15, 1+5=6)
//
//	DigitalRoot(132189) -> 6 (1+3+2+1+8+9=24, 2+4=6)
//	DigitalRoot(0) -> 0
//	DigitalRoot(5) -> 5
func DigitalRoot(n int) int {
	// TODO: implementasi di sini

	word := strconv.Itoa(n)

	for len(word) != 1 {
		count := 0
		for _, v := range word {
			num, _ := strconv.Atoi(string(v))
			count += num
		}
		word = strconv.Itoa(count)
	}
	num, _ := strconv.Atoi(string(word))
	return num
}
