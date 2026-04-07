package belajar

// Reverse membalikkan urutan elemen dalam slice.
// Contoh: nums := []int{1, 2, 3}; Reverse(nums) -> nums menjadi []int{3, 2, 1}
func Reverse(nums []int) {
	// TODO: implementasi di sini
}

// Unique membuang elemen duplikat dan mengembalikan slice baru dengan elemen unik.
// Urutan elemen unik harus sesuai dengan kemunculan pertama di slice asli.
// Contoh: Unique([]int{1, 2, 2, 3, 1}) -> []int{1, 2, 3} 
func Unique(nums []int) []int {
	// TODO: implementasi di sini
	return nil
}

// Intersect mengembalikan irisan dari dua slice (elemen yang ada di kedua slice).
// Tiap angka pada irisan hanya boleh muncul sekali (unik).
// Contoh: Intersect([]int{1, 2, 2, 1}, []int{2, 2, 3}) -> []int{2}
func Intersect(nums1 []int, nums2 []int) []int {
	// TODO: implementasi di sini
	return nil
}

// Contains mengecek apakah nilai target terdapat di dalam slice.
// Contoh: Contains([]int{1, 2, 3}, 2) -> true
//         Contains([]int{1, 2, 3}, 4) -> false
func Contains(nums []int, target int) bool {
	// TODO: implementasi di sini
	return false
}

// GroupByParity mengelompokkan elemen menjadi bilangan genap dan ganjil.
// Mengembalikan slice genap pada return pertama dan slice ganjil pada return kedua.
// Contoh: GroupByParity([]int{1, 2, 3, 4, 5}) -> []int{2, 4}, []int{1, 3, 5}
func GroupByParity(nums []int) ([]int, []int) {
	// TODO: implementasi di sini
	return nil, nil
}

// RemoveAt menghapus elemen pada indeks tertentu dan mengembalikan slice baru.
// Kembalikan bool true jika terjadi error / index tidak valid.
// Contoh: RemoveAt([]int{10, 20, 30}, 1) -> []int{10, 30}, false
//         RemoveAt([]int{10, 20, 30}, 5) -> nil, true
func RemoveAt(nums []int, index int) ([]int, bool) {
	// TODO: implementasi di sini
	return nil, true
}
