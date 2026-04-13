package belajar

// ==================== DAY 14: POINTER & DEEP COPY ====================
// Topik: Memahami pointer (alamat memori) dan konsep deep copy vs shallow copy.
// Konsep penting: pointer memungkinkan fungsi memodifikasi data di luar scope-nya.

// SwapValues menukar nilai dari dua variabel integer melalui pointer.
// Setelah dipanggil, nilai *a dan *b harus tertukar.
// Contoh:
//
//	x, y := 10, 20
//	SwapValues(&x, &y)
//	// x sekarang 20, y sekarang 10
//
//	a, b := -5, 100
//	SwapValues(&a, &b)
//	// a sekarang 100, b sekarang -5
//
// Hint: gunakan variabel sementara (temp) untuk menyimpan salah satu nilai.
func SwapValues(a, b *int) {
	// TODO: implementasi di sini
}

// DeepCopyMatrix membuat salinan (deep copy) dari matrix 2D.
// Perubahan pada matrix hasil TIDAK boleh mempengaruhi matrix asli, dan sebaliknya.
// Contoh:
//
//	original := [][]int{{1, 2}, {3, 4}}
//	copied := DeepCopyMatrix(original)
//	copied[0][0] = 99
//	// original[0][0] masih tetap 1 (tidak terpengaruh)
//
//	DeepCopyMatrix([][]int{}) -> [][]int{}
//	DeepCopyMatrix(nil) -> nil
//
// Hint: gunakan make() dan copy() untuk setiap baris.
func DeepCopyMatrix(matrix [][]int) [][]int {
	// TODO: implementasi di sini
	return nil
}
