package belajar

// ==================== DAY 16: SET OPERATIONS ====================
// Topik: Operasi himpunan (set) - union dan symmetric difference.
// Konsep penting: set tidak memiliki elemen duplikat, operasi set mirip matematika.

// Union mengembalikan gabungan dua slice (tanpa duplikat).
// Urutan: elemen dari a dulu (urutkan kemunculan pertama), lalu elemen unik dari b.
// Contoh: Union([]int{1, 2, 3}, []int{3, 4, 5}) -> []int{1, 2, 3, 4, 5}
//
//	Union([]int{1, 1, 2}, []int{2, 3, 3}) -> []int{1, 2, 3}
//	Union([]int{}, []int{1, 2}) -> []int{1, 2}
//	Union([]int{5}, []int{5}) -> []int{5}
//	Union([]int{}, []int{}) -> []int{}
func Union(a, b []int) []int {
	// TODO: implementasi di sini
	return nil
}

// SymmetricDifference mengembalikan elemen yang ada di salah satu slice tapi TIDAK di keduanya.
// Hasil harus unik (tanpa duplikat).
// Urutan: elemen unik dari a dulu, lalu elemen unik dari b.
// Contoh: SymmetricDifference([]int{1, 2, 3}, []int{3, 4, 5}) -> []int{1, 2, 4, 5}
//
//	SymmetricDifference([]int{1, 2}, []int{1, 2}) -> []int{}
//	SymmetricDifference([]int{1, 2, 3}, []int{4, 5, 6}) -> []int{1, 2, 3, 4, 5, 6}
//	SymmetricDifference([]int{}, []int{1, 2}) -> []int{1, 2}
//	SymmetricDifference([]int{1, 1, 2}, []int{2, 3, 3}) -> []int{1, 3}
//
// Hint: elemen yang hanya ada di a + elemen yang hanya ada di b.
func SymmetricDifference(a, b []int) []int {
	// TODO: implementasi di sini
	return nil
}
