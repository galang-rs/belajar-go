package belajar

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Custom error variables - gunakan untuk perbandingan error di test.
var (
	ErrDivisionByZero = errors.New("pembagian dengan nol")
	ErrNegativeNumber = errors.New("bilangan negatif tidak diperbolehkan")
	ErrOutOfRange     = errors.New("indeks di luar jangkauan")
	ErrInvalidInput   = errors.New("input tidak valid")
	ErrEmptySlice     = errors.New("slice kosong")
	ErrInvalidFormat  = errors.New("format tidak valid")
	ErrOutOfBounds    = errors.New("nilai di luar batas")
)

// SafeDivide membagi a dengan b.
// Kembalikan error ErrDivisionByZero jika b == 0.
// Contoh: SafeDivide(10, 2) -> 5.0, nil
//
//	SafeDivide(10, 0) -> 0, ErrDivisionByZero
func SafeDivide(a, b float64) (float64, error) {
	// TODO: implementasi di sini
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	return a / b, nil
}

// SafeSqrt mengembalikan akar kuadrat dari n.
// Kembalikan error ErrNegativeNumber jika n negatif.
// Contoh: SafeSqrt(25) -> 5.0, nil
//
//	SafeSqrt(-4) -> 0, ErrNegativeNumber
//
// Hint: gunakan math.Sqrt
func SafeSqrt(n float64) (float64, error) {
	// TODO: implementasi di sini

	if n < 0 {
		return 0, ErrNegativeNumber
	}
	return math.Sqrt(n), nil
}

// SafeIndex mengembalikan elemen pada index dari slice.
// Kembalikan error ErrOutOfRange jika index di luar jangkauan (negatif atau >= len).
// Contoh: SafeIndex([]int{10, 20, 30}, 1) -> 20, nil
//
//	SafeIndex([]int{10, 20, 30}, 5) -> 0, ErrOutOfRange
//	SafeIndex([]int{10, 20, 30}, -1) -> 0, ErrOutOfRange
func SafeIndex(nums []int, index int) (int, error) {
	// TODO: implementasi di sini
	if len(nums)-1 < index || index < 0 {
		return 0, ErrOutOfRange
	}
	return nums[index], nil
}

// ParsePositiveInt mengonversi string ke integer positif.
// Kembalikan error ErrInvalidInput jika string bukan angka valid.
// Kembalikan error ErrNegativeNumber jika angka negatif.
// Contoh: ParsePositiveInt("42") -> 42, nil
//
//	ParsePositiveInt("-5") -> 0, ErrNegativeNumber
//	ParsePositiveInt("abc") -> 0, ErrInvalidInput
//
// Hint: gunakan strconv.Atoi
func ParsePositiveInt(s string) (int, error) {
	// TODO: implementasi di sini
	if _, err := strconv.Atoi(s); err != nil {
		return 0, ErrInvalidInput
	}
	if v, _ := strconv.Atoi(s); int(v) < 0 {
		return 0, ErrNegativeNumber
	}
	v, _ := strconv.Atoi(s)
	return v, nil
}

// ParseAge menerima string dan mengembalikan usia (int).
// Usia harus berada di rentang 0-150 (inklusif).
// Kembalikan error ErrInvalidInput jika bukan angka valid.
// Kembalikan error ErrOutOfBounds jika di luar rentang 0-150.
// Contoh: ParseAge("25") -> 25, nil
//
//	ParseAge("200") -> 0, ErrOutOfBounds
//	ParseAge("-5") -> 0, ErrOutOfBounds
//	ParseAge("abc") -> 0, ErrInvalidInput
func ParseAge(s string) (int, error) {
	// TODO: implementasi di sini
	// Hint: gunakan strconv.Atoi
	v, err := strconv.Atoi(s)

	if err != nil {
		return 0, ErrInvalidInput
	}
	if v < 0 || v > 150 {
		return 0, ErrOutOfBounds
	}

	return v, nil
}

// ValidateEmail melakukan validasi sederhana format email.
// Email valid harus memenuhi:
//  1. Mengandung tepat satu karakter "@"
//  2. Ada karakter sebelum "@"
//  3. Ada karakter "." setelah "@"
//  4. "." tidak boleh jadi karakter terakhir
//
// Kembalikan error ErrInvalidFormat jika tidak valid, nil jika valid.
// Contoh: ValidateEmail("user@example.com") -> nil
//
//	ValidateEmail("invalid") -> ErrInvalidFormat
//	ValidateEmail("@example.com") -> ErrInvalidFormat
//	ValidateEmail("user@com") -> ErrInvalidFormat
func ValidateEmail(email string) error {
	// TODO: implementasi di sini
	// Hint: gunakan strings.Count, strings.Index, strings.LastIndex

	type data struct {
		indexAt  int
		countAt  int
		IndexDot int
	}

	var val data

	for k, v := range email {
		if string(v) == "@" {
			val.indexAt = k
			val.countAt++
		} else if string(v) == "." && k != len(email)-1 {
			val.IndexDot = k
		}
	}
	fmt.Println(val)

	if val.indexAt == 0 || val.IndexDot == 0 || val.countAt > 1 {
		return ErrInvalidFormat
	}
	if val.indexAt > val.IndexDot {
		return ErrInvalidFormat
	}

	return nil
}

// SafeAverage menghitung rata-rata dari slice float64.
// Kembalikan error ErrEmptySlice jika slice kosong.
// Contoh: SafeAverage([]float64{10, 20, 30}) -> 20.0, nil
//
//	SafeAverage([]float64{}) -> 0, ErrEmptySlice
func SafeAverage(nums []float64) (float64, error) {
	// TODO: implementasi di sini
	if len(nums) == 0 {
		return 0, ErrEmptySlice
	}
	var val float64
	for _, v := range nums {
		val += v
	}
	return val / float64(len(nums)), nil
}

// MinMax mengembalikan nilai minimum dan maksimum dari slice int.
// Kembalikan error ErrEmptySlice jika slice kosong.
// Contoh: MinMax([]int{3, 1, 4, 1, 5}) -> 1, 5, nil
//
//	MinMax([]int{}) -> 0, 0, ErrEmptySlice
func MinMax(nums []int) (int, int, error) {
	// TODO: implementasi di sini
	if len(nums) == 0 {
		return 0, 0, ErrEmptySlice
	}

	max := math.Inf(-1)
	min := math.Inf(1)

	for _, v := range nums {
		if float64(v) > max {
			max = float64(v)
		}
		if float64(v) < min {
			min = float64(v)
		}
	}
	return int(min), int(max), nil
}

// ParseHexColor mengurai string hex color "#RRGGBB" menjadi tiga komponen (r, g, b).
// Setiap komponen bernilai 0-255.
// Kembalikan error ErrInvalidFormat jika format tidak valid.
// Contoh: ParseHexColor("#FF8800") -> 255, 136, 0, nil
//
//	ParseHexColor("#00ff00") -> 0, 255, 0, nil (case-insensitive)
//	ParseHexColor("FF8800") -> 0, 0, 0, ErrInvalidFormat (tanpa #)
//	ParseHexColor("#GGG") -> 0, 0, 0, ErrInvalidFormat
//
// Hint: gunakan strconv.ParseInt dengan base 16
func ParseHexColor(hex string) (int, int, int, error) {
	// TODO: implementasi di sini
	val := strings.Index(hex, "#")
	if val == -1 {
		return 0, 0, 0, ErrInvalidFormat
	}
	if val == 0 && len(hex) == 7 {
		var r int64
		var g int64
		var b int64
		r, _ = strconv.ParseInt(string(hex[1])+string(hex[2]), 16, 64)
		g, _ = strconv.ParseInt(string(hex[3])+string(hex[4]), 16, 64)
		b, _ = strconv.ParseInt(string(hex[5])+string(hex[6]), 16, 64)

		return int(r), int(g), int(b), nil
	}
	return 0, 0, 0, nil
}

// Retry menjalankan fungsi fn maksimal maxAttempts kali.
// Jika fn mengembalikan nil (sukses), langsung kembalikan nil.
// Jika semua percobaan gagal, kembalikan error terakhir.
// Jika maxAttempts <= 0, kembalikan ErrInvalidInput.
// Contoh:
//
//	attempt := 0
//	fn := func() error {
//	    attempt++
//	    if attempt < 3 { return errors.New("gagal") }
//	    return nil
//	}
//	Retry(fn, 5) -> nil (berhasil di percobaan ke-3)
//	Retry(fn, 1) -> error (gagal, hanya 1 percobaan)
func Retry(fn func() error, maxAttempts int) error {
	// TODO: implementasi di sini
	if maxAttempts <= 0 {
		return ErrInvalidInput
	}
	var lastError error
	for i := 0; i < maxAttempts; i++ {
		result := fn()
		if result == nil {
			return nil
		} else {
			lastError = result
		}
	}
	return lastError
}

// FUNCTION Retry(fn, maxAttempts):

//     IF maxAttempts <= 0 THEN
//         RETURN ErrInvalidInput
//     ENDIF

//     lastError ← nil

//     FOR i FROM 1 TO maxAttempts DO
//         result ← fn()

//         IF result == nil THEN
//             RETURN nil   // sukses
//         ELSE
//             lastError ← result
//         ENDIF
//     ENDFOR

//     RETURN lastError   // semua percobaan gagal
// END FUNCTION
