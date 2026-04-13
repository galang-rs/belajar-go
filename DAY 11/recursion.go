package belajar

// ==================== DAY 11: REKURSI LANJUTAN ====================
// Topik: Permutasi dan Power Set menggunakan teknik rekursi + backtracking.
// Konsep penting: cara membangun solusi secara bertahap dan "mundur" (backtrack).

// Permutations mengembalikan semua kemungkinan urutan (permutasi) dari slice integer.
// Setiap permutasi harus memiliki panjang yang sama dengan input.
// Input dijamin tidak ada duplikat.
// Urutan permutasi: ikuti urutan eksplorasi rekursif standar (lihat contoh).
// Contoh: Permutations([]int{1, 2, 3}) -> [][]int{
//
//	{1, 2, 3}, {1, 3, 2},
//	{2, 1, 3}, {2, 3, 1},
//	{3, 1, 2}, {3, 2, 1},
//
// }
//
//	Permutations([]int{1}) -> [][]int{{1}}
//	Permutations([]int{}) -> [][]int{}
//
// Hint: gunakan teknik backtracking dengan slice "used" untuk menandai elemen yang sudah dipakai.
func Permutations(nums []int) [][]int {
	// TODO: implementasi di sini
	return nil
}

// PowerSet mengembalikan semua subset (himpunan bagian) dari slice integer.
// Termasuk subset kosong dan slice itu sendiri.
// Input dijamin tidak ada duplikat.
// Urutan subset: ikuti urutan eksplorasi rekursif standar (lihat contoh).
// Contoh: PowerSet([]int{1, 2, 3}) -> [][]int{
//
//	{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3},
//
// }
//
//	PowerSet([]int{1}) -> [][]int{{}, {1}}
//	PowerSet([]int{}) -> [][]int{{}}
//
// Hint: untuk setiap elemen, pilih "ambil" atau "lewati", lalu rekursi ke elemen berikutnya.
func PowerSet(nums []int) [][]int {
	// TODO: implementasi di sini
	return nil
}
