package belajar

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
	return 0
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
	return ""
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
	return nil
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
	return nil
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
	return 0
}
