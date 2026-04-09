package belajar

// BinarySearch mencari target dalam slice yang SUDAH terurut secara ascending.
// Mengembalikan index target dan true jika ditemukan, -1 dan false jika tidak.
// Contoh: BinarySearch([]int{1, 3, 5, 7, 9}, 5) -> 2, true
//
//	BinarySearch([]int{1, 3, 5, 7, 9}, 4) -> -1, false
//	BinarySearch([]int{}, 1) -> -1, false
func BinarySearch(nums []int, target int) (int, bool) {
	// TODO: implementasi di sini
	return -1, false
}

// BubbleSort mengurutkan slice integer secara ascending menggunakan algoritma Bubble Sort.
// Mengembalikan slice baru (TIDAK mengubah slice asli).
// Contoh: BubbleSort([]int{5, 3, 1, 4, 2}) -> []int{1, 2, 3, 4, 5}
//
//	BubbleSort([]int{}) -> []int{}
func BubbleSort(nums []int) []int {
	// TODO: implementasi di sini
	return nil
}

// SelectionSort mengurutkan slice integer secara ascending menggunakan Selection Sort.
// Mengembalikan slice baru (TIDAK mengubah slice asli).
// Contoh: SelectionSort([]int{64, 25, 12, 22, 11}) -> []int{11, 12, 22, 25, 64}
//
//	SelectionSort([]int{1}) -> []int{1}
func SelectionSort(nums []int) []int {
	// TODO: implementasi di sini
	return nil
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
	return 0
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
	return 0
}

// Power menghitung base pangkat exp secara rekursif.
// exp >= 0. Power(x, 0) = 1 untuk semua x.
// Contoh: Power(2, 10) -> 1024
//
//	Power(3, 3) -> 27
//	Power(5, 0) -> 1
//	Power(0, 5) -> 0
func Power(base, exp int) int {
	// TODO: implementasi di sini
	return 0
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
	return 0
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
	return false
}
