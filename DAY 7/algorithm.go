package belajar

import (
	"fmt"
	"math"
	"strconv"
)

// BinarySearch mencari target dalam slice yang SUDAH terurut secara ascending.
// Mengembalikan index target dan true jika ditemukan, -1 dan false jika tidak.
// Contoh: BinarySearch([]int{1, 3, 5, 7, 9}, 5) -> 2, true
//
//	BinarySearch([]int{1, 3, 5, 7, 9}, 4) -> -1, false
//	BinarySearch([]int{}, 1) -> -1, false
func BinarySearch(nums []int, target int) (int, bool) {
	// TODO: implementasi di sini
	if len(nums) == 0 {
		return -1, false
	}
	for k, v := range nums {
		if v == target {
			return k, true
		}
	}
	return -1, false
}

// BubbleSort mengurutkan slice integer secara ascending menggunakan algoritma Bubble Sort.
// Mengembalikan slice baru (TIDAK mengubah slice asli).
// Contoh: BubbleSort([]int{5, 3, 1, 4, 2}) -> []int{1, 2, 3, 4, 5}
//
//	BubbleSort([]int{}) -> []int{}
func BubbleSort(nums []int) []int {
	// TODO: implementasi di sini
	n := len(nums)
	result := make([]int, n)
	copy(result, nums)

	for i := 0; i < n; i++ {
		for j := 0; j < (n - i - 1); j++ {
			if result[j] > result[j+1] {
				temp := result[j]
				result[j] = result[j+1]
				result[j+1] = temp
			}
		}
	}

	return result
}

// FUNCTION BubbleSort(nums):

//     n ← length(nums)

//     // buat copy array agar tidak mengubah yang asli
//     result ← copy of nums

//     FOR i FROM 0 TO n-1 DO
//         FOR j FROM 0 TO n-i-2 DO
//             IF result[j] > result[j+1] THEN
//                 // tukar posisi
//                 temp ← result[j]
//                 result[j] ← result[j+1]
//                 result[j+1] ← temp
//             ENDIF
//         ENDFOR
//     ENDFOR

//     RETURN result
// END FUNCTION

// SelectionSort mengurutkan slice integer secara ascending menggunakan Selection Sort.
// Mengembalikan slice baru (TIDAK mengubah slice asli).
// Contoh: SelectionSort([]int{64, 25, 12, 22, 11}) -> []int{11, 12, 22, 25, 64}
//
//	SelectionSort([]int{1}) -> []int{1}
func SelectionSort(nums []int) []int {
	// TODO: implementasi di sini
	n := len(nums)
	result := make([]int, n)
	copy(result, nums)

	for i := 0; i < n; i++ {
		for j := 0; j < (n - i - 1); j++ {
			if result[j] > result[j+1] {
				temp := result[j]
				result[j] = result[j+1]
				result[j+1] = temp
			}
		}
	}

	return result
}

// InsertionSort mengurutkan slice integer secara ascending menggunakan Insertion Sort.
// Mengembalikan slice baru (TIDAK mengubah slice asli).
// Contoh: InsertionSort([]int{5, 2, 4, 6, 1, 3}) -> []int{1, 2, 3, 4, 5, 6}
func InsertionSort(nums []int) []int {
	// TODO: implementasi di sini
	return nil
}

// MergeSort mengurutkan slice integer secara ascending menggunakan Merge Sort.
// Mengembalikan slice baru (TIDAK mengubah slice asli).
// Contoh: MergeSort([]int{38, 27, 43, 3, 9, 82, 10}) -> []int{3, 9, 10, 27, 38, 43, 82}
//
//	MergeSort([]int{}) -> []int{}
func MergeSort(nums []int) []int {
	// TODO: implementasi di sini
	return nil
}

// GCD mengembalikan Greatest Common Divisor (FPB) dari dua bilangan.
// Gunakan algoritma Euclidean.
// Selalu kembalikan nilai positif.
// Contoh: GCD(12, 8) -> 4
//
//	GCD(17, 5) -> 1
//	GCD(100, 75) -> 25
//	GCD(0, 5) -> 5
func GCD(a, b int) int {
	// TODO: implementasi di sini
	if a < 0 {
		return -b
	}
	if b < 0 {
		return -a
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// LCM mengembalikan Least Common Multiple (KPK) dari dua bilangan.
// Contoh: LCM(4, 6) -> 12
//
//	LCM(3, 5) -> 15
//	LCM(0, 5) -> 0
//
// Hint: LCM(a, b) = |a * b| / GCD(a, b)
func LCM(a, b int) int {
	// TODO: implementasi di sini

	if a == 0 || b == 0 {
		return 0
	}

	gcd := GCD(a, b)

	val := (a * b) / gcd
	return val
}

// Algoritma LCM(a, b)
//     if a = 0 atau b = 0 then
//         return 0
//     end if

//     gcd ← GCD(a, b)

//     hasil ← |a × b| / gcd

//     return hasil

// Power menghitung base pangkat exp secara rekursif.
// exp >= 0. Power(x, 0) = 1 untuk semua x.
// Contoh: Power(2, 10) -> 1024
//
//	Power(3, 3) -> 27
//	Power(5, 0) -> 1
//	Power(0, 5) -> 0
func Power(base, exp int) int {
	// TODO: implementasi di sini

	if exp == 0 {
		return 1
	}
	if base == 0 {
		return 0
	}
	if exp == 1 {
		return base
	}
	val := base * base
	for i := 0; i < exp-2; i++ {
		val *= base
	}

	return val
}

// SumDigits mengembalikan jumlah semua digit dari bilangan (rekursif).
// Jika bilangan negatif, gunakan nilai absolutnya.
// Contoh: SumDigits(123) -> 6 (1+2+3)
//
//	SumDigits(-456) -> 15 (4+5+6)
//	SumDigits(0) -> 0
//	SumDigits(9) -> 9
func SumDigits(n int) int {
	// TODO: implementasi di sini

	valAbs := int(math.Abs(float64(n)))
	valStr := strconv.Itoa(valAbs)

	var result int

	for k, v := range valStr {
		result += int(v) - 48
		fmt.Println(valStr, result, v, valStr[k])
	}

	return result
}

// IsSorted mengecek apakah slice integer sudah terurut secara ascending (non-decreasing).
// Slice kosong dan slice dengan satu elemen dianggap sudah terurut.
// Contoh: IsSorted([]int{1, 2, 3, 4, 5}) -> true
//
//	IsSorted([]int{1, 3, 2}) -> false
//	IsSorted([]int{1, 1, 2, 2}) -> true
//	IsSorted([]int{}) -> true
func IsSorted(nums []int) bool {
	// TODO: implementasi di sini

	last := math.Inf(-1)
	for _, v := range nums {
		if last <= float64(v) {
			last = float64(v)
		} else {
			return false
		}
	}
	return true
}
